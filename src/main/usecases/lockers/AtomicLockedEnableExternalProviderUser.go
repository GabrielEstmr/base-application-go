package main_usecases_lockers

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_domains "baseapplicationgo/main/domains"
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

type AtomicLockedEnableExternalProviderUser struct {
	_MSG_ENABLE_USER_CODE_NOT_FOUND string
	_MSG_ENABLE_EMAIL_LOCK_ISSUE    string
	enableExternalProviderUser      main_usecases_interfaces.EnableExternalProviderUser
	userDatabaseGateway             main_gateways.UserDatabaseGateway
	lockGateway                     main_gateways.DistributedLockGateway
	logsMonitoringGateway           main_gateways.LogsMonitoringGateway
	spanGateway                     main_gateways.SpanGateway
	messageUtils                    main_utils_messages.ApplicationMessages
}

func NewAtomicLockedEnableExternalProviderUser(
	enableExternalProviderUser main_usecases_interfaces.EnableExternalProviderUser,
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *AtomicLockedEnableExternalProviderUser {
	return &AtomicLockedEnableExternalProviderUser{
		_MSG_ENABLE_USER_CODE_NOT_FOUND: "exceptions.architecture.application.issue",
		_MSG_ENABLE_EMAIL_LOCK_ISSUE:    "providers.modify.user.verification.email.error.lock.issue",
		enableExternalProviderUser:      enableExternalProviderUser,
		userDatabaseGateway:             userDatabaseGateway,
		lockGateway:                     lockGateway,
		logsMonitoringGateway:           logsMonitoringGateway,
		spanGateway:                     spanGateway,
		messageUtils:                    messageBeans,
	}
}

func (this *AtomicLockedEnableExternalProviderUser) Execute(
	ctx context.Context,
	userId string,
	args main_domains.EnableExternalUserArgs,
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
		span.GetCtx(), "User Modification-Create-Update", persistedUser.GetEmail(), 90*time.Second)
	errLock := singleLock.Lock()

	if errLock == nil {
		defer this.lockGateway.UnlockAndLogIfError(span.GetCtx(), *singleLock)

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

			updatedUser, errE := this.enableExternalProviderUser.Execute(span.GetCtx(), persistedUser, args, dbOpts)
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
