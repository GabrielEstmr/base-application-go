package main_usecases_lockers

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

const _MSG_CREATE_NEW_USER_ARCH_ISSUE = "exceptions.architecture.application.issue"
const _MSG_CREATE_NEW_USER_LOCK_SAME_EMAIL = "providers.create.user.lock.issue"

type AtomicLockedCreateNewUser struct {
	lockGateway           main_gateways.DistributedLockGateway
	createNewUser         main_usecases_interfaces.CreateNewUser
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewAtomicLockedCreateNewUser(
	lockGateway main_gateways.DistributedLockGateway,
	createNewUser main_usecases_interfaces.CreateNewUser,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *AtomicLockedCreateNewUser {
	return &AtomicLockedCreateNewUser{
		lockGateway:           lockGateway,
		createNewUser:         createNewUser,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageBeans}
}

func (this *AtomicLockedCreateNewUser) Execute(ctx context.Context, user main_domains.User) (
	main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "CreateNewUser-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating new User with email: %s", user.GetEmail()))

	singleLock := this.lockGateway.GetWithScope(
		span.GetCtx(),
		main_domains_enums.LOCK_SCOPE_USER_MODIFICATION,
		user.GetEmail(),
		90*time.Second)
	errLock := singleLock.Lock()

	singleLockEV := this.lockGateway.GetWithScope(
		span.GetCtx(),
		main_domains_enums.LOCK_SCOPE_USER_VERIFICATION_EMAIL_MODIFICATION,
		user.GetEmail(),
		90*time.Second)
	errLockEV := singleLockEV.Lock()

	if errLock == nil && errLockEV == nil {
		defer this.lockGateway.UnlockAndLogIfError(span.GetCtx(), *singleLock)
		defer this.lockGateway.UnlockAndLogIfError(span.GetCtx(), *singleLockEV)

		wc := writeconcern.Majority()
		rc := readconcern.Snapshot()
		transactionOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

		session, errCreateSession := main_configs_mongo.GetMongoDBClientBean().StartSession()
		if errCreateSession != nil {
			return *new(main_domains.User), main_domains_exceptions.
				NewInternalServerErrorExceptionSglMsg(this.messageUtils.
					GetDefaultLocale(_MSG_CREATE_NEW_USER_ARCH_ISSUE))
		}
		defer session.EndSession(context.Background())

		var userResult main_domains.User
		var errResult main_domains_exceptions.ApplicationException = nil

		errSession := mongo.WithSession(span.GetCtx(), session, func(sessionCtx mongo.SessionContext) error {

			if errSession := session.StartTransaction(transactionOpts); errSession != nil {
				this.logsMonitoringGateway.ERROR(span, errSession.Error())
				errResult = main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(_MSG_CREATE_NEW_USER_ARCH_ISSUE))
				return errSession
			}

			opts := main_domains.NewDatabaseOptions().WithSession(sessionCtx)
			userP, errP := this.createNewUser.Execute(span.GetCtx(), user, opts)
			if errP != nil {
				errResult = errP
				return errP
			}

			userResult = userP

			if errFS := session.CommitTransaction(sessionCtx); errFS != nil {
				this.logsMonitoringGateway.ERROR(span, errFS.Error())
				errResult = main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(_MSG_CREATE_NEW_USER_ARCH_ISSUE))
				return errFS
			}
			return nil
		})

		if errSession != nil {
			this.logsMonitoringGateway.ERROR(span, errSession.Error())
			if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
				return *new(main_domains.User), main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(_MSG_CREATE_NEW_USER_ARCH_ISSUE))
			}
			return *new(main_domains.User), errResult
		}

		return userResult, nil

	} else {
		return *new(main_domains.User),
			main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
				_MSG_CREATE_NEW_USER_LOCK_SAME_EMAIL))
	}

}
