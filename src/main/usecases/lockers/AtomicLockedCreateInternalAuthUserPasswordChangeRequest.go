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

type AtomicLockedCreateInternalAuthUserPasswordChangeRequest struct {
	_MSG_KEY_ARCH_ISSUE                         string
	_MSG_EMAIL_VERIFICATION_LOCK_ISSUE          string
	createInternalAuthUserPasswordChangeRequest main_usecases_interfaces.CreateInternalAuthUserPasswordChangeRequest
	lockGateway                                 main_gateways.DistributedLockGateway
	userDatabaseGateway                         main_gateways.UserDatabaseGateway
	logsMonitoringGateway                       main_gateways.LogsMonitoringGateway
	spanGateway                                 main_gateways.SpanGateway
	messageUtils                                main_utils_messages.ApplicationMessages
}

func NewAtomicLockedCreateInternalAuthUserPasswordChangeRequest(
	createInternalAuthUserPasswordChangeRequest main_usecases_interfaces.CreateInternalAuthUserPasswordChangeRequest,
	lockGateway main_gateways.DistributedLockGateway,
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *AtomicLockedCreateInternalAuthUserPasswordChangeRequest {
	return &AtomicLockedCreateInternalAuthUserPasswordChangeRequest{
		_MSG_KEY_ARCH_ISSUE:                         "exceptions.architecture.application.issue",
		_MSG_EMAIL_VERIFICATION_LOCK_ISSUE:          "providers.modify.user.verification.email.error.lock.issue",
		createInternalAuthUserPasswordChangeRequest: createInternalAuthUserPasswordChangeRequest,
		lockGateway:           lockGateway,
		userDatabaseGateway:   userDatabaseGateway,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageUtils,
	}
}

func (this *AtomicLockedCreateInternalAuthUserPasswordChangeRequest) Execute(
	ctx context.Context,
	userId string,
) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "AtomicLockedCreateInternalAuthUserPasswordChangeRequest-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating password change request for user id: %s", userId))

	persistedUser, errF := this.userDatabaseGateway.FindById(span.GetCtx(), userId, nil)
	if errF != nil {
		this.logsMonitoringGateway.ERROR(span, errF.Error())
		return main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				this._MSG_KEY_ARCH_ISSUE))
	}

	singleLockEV := this.lockGateway.GetWithScope(
		span.GetCtx(),
		main_domains_enums.LOCK_SCOPE_USER_VERIFICATION_EMAIL_MODIFICATION,
		persistedUser.GetEmail(),
		90*time.Second)
	errLockEV := singleLockEV.Lock()

	if errLockEV == nil {
		defer this.lockGateway.UnlockAndLogIfError(span.GetCtx(), *singleLockEV)

		wc := writeconcern.Majority()
		rc := readconcern.Snapshot()
		transactionOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

		session, errCreateSession := main_configs_mongo.GetMongoDBClientBean().StartSession()
		if errCreateSession != nil {
			return main_domains_exceptions.
				NewInternalServerErrorExceptionSglMsg(this.messageUtils.
					GetDefaultLocale(this._MSG_KEY_ARCH_ISSUE))
		}
		defer session.EndSession(context.Background())

		var errResult main_domains_exceptions.ApplicationException = nil

		errSession := mongo.WithSession(span.GetCtx(), session, func(sessionCtx mongo.SessionContext) error {
			if errSession := session.StartTransaction(transactionOpts); errSession != nil {
				this.logsMonitoringGateway.ERROR(span, errSession.Error())
				errResult = main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(this._MSG_KEY_ARCH_ISSUE))
				return errSession
			}

			opts := main_domains.NewDatabaseOptions().WithSession(sessionCtx)
			errUP := this.createInternalAuthUserPasswordChangeRequest.Execute(
				span.GetCtx(),
				userId,
				opts)
			if errUP != nil {
				errResult = errUP
				return errUP
			}

			if errFS := session.CommitTransaction(sessionCtx); errFS != nil {
				this.logsMonitoringGateway.ERROR(span, errFS.Error())
				errResult = main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(this._MSG_KEY_ARCH_ISSUE))
				return errFS
			}
			return nil
		})

		if errSession != nil {
			this.logsMonitoringGateway.ERROR(span, errSession.Error())
			if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
				return main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(this._MSG_KEY_ARCH_ISSUE))
			}
			return errResult
		}

		return nil

	} else {
		return main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
			this._MSG_EMAIL_VERIFICATION_LOCK_ISSUE))
	}
}
