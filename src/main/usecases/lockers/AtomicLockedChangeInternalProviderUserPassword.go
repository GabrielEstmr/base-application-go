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

type AtomicLockedChangeInternalProviderUserPassword struct {
	_MSG_ENABLE_USER_CODE_NOT_FOUND    string
	_MSG_ENABLE_EMAIL_LOCK_ISSUE       string
	changeInternalProviderUserPassword main_usecases_interfaces.ChangeInternalProviderUserPassword
	userDatabaseGateway                main_gateways.UserDatabaseGateway
	lockGateway                        main_gateways.DistributedLockGateway
	logsMonitoringGateway              main_gateways.LogsMonitoringGateway
	spanGateway                        main_gateways.SpanGateway
	messageUtils                       main_utils_messages.ApplicationMessages
}

func NewAtomicLockedChangeInternalProviderUserPassword(
	changeInternalProviderUserPassword main_usecases_interfaces.ChangeInternalProviderUserPassword,
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *AtomicLockedChangeInternalProviderUserPassword {
	return &AtomicLockedChangeInternalProviderUserPassword{
		_MSG_ENABLE_USER_CODE_NOT_FOUND:    "exceptions.architecture.application.issue",
		_MSG_ENABLE_EMAIL_LOCK_ISSUE:       "providers.modify.user.verification.email.error.lock.issue",
		changeInternalProviderUserPassword: changeInternalProviderUserPassword,
		userDatabaseGateway:                userDatabaseGateway,
		lockGateway:                        lockGateway,
		logsMonitoringGateway:              logsMonitoringGateway,
		spanGateway:                        spanGateway,
		messageUtils:                       messageBeans,
	}
}

func (this *AtomicLockedChangeInternalProviderUserPassword) Execute(
	ctx context.Context,
	userId string,
	password string,
	verificationCode string,
) (
	main_domains.User,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(ctx, "EnableExternalProviderUser-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Verifing user. id: %s", userId))

	persistedUser, errF := this.userDatabaseGateway.FindById(span.GetCtx(), userId, nil)
	if errF != nil {
		return *new(main_domains.User), errF
	}

	if persistedUser.IsEmpty() {
		return *new(main_domains.User), main_domains_exceptions.
			NewResourceNotFoundExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_ENABLE_USER_CODE_NOT_FOUND))
	}

	singleLock := this.lockGateway.GetWithScope(
		span.GetCtx(),
		main_domains_enums.LOCK_SCOPE_USER_MODIFICATION,
		persistedUser.GetEmail(),
		90*time.Second)
	errLock := singleLock.Lock()

	singleLockEV := this.lockGateway.GetWithScope(
		span.GetCtx(),
		main_domains_enums.LOCK_SCOPE_USER_VERIFICATION_EMAIL_MODIFICATION,
		persistedUser.GetEmail(),
		90*time.Second)
	errLockEV := singleLockEV.Lock()

	if errLock == nil && errLockEV == nil {
		defer this.lockGateway.UnlockAndLogIfError(span.GetCtx(), *singleLock)
		defer this.lockGateway.UnlockAndLogIfError(span.GetCtx(), *singleLockEV)

		var userResult main_domains.User
		var errResult main_domains_exceptions.ApplicationException = nil

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

		errSession := mongo.WithSession(span.GetCtx(), session, func(sessionCtx mongo.SessionContext) error {
			if errSession := session.StartTransaction(transactionOpts); errSession != nil {
				this.logsMonitoringGateway.ERROR(span, errSession.Error())
				errResult = main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(_MSG_CREATE_NEW_USER_ARCH_ISSUE))
				return errSession
			}

			dbOpts := main_domains.NewDatabaseOptions().WithSession(sessionCtx)

			updatedUser, errE := this.changeInternalProviderUserPassword.Execute(
				span.GetCtx(),
				userId,
				password,
				verificationCode,
				dbOpts)
			if errE != nil {
				errResult = errE
				return errE
			}
			userResult = updatedUser

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
				this._MSG_ENABLE_EMAIL_LOCK_ISSUE))
	}
}
