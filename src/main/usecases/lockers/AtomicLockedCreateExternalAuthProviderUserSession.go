package main_usecases_lockers

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

type AtomicLockedCreateExternalAuthProviderUserSession struct {
	_MSG_TOKEN_MUST_NOT_BE_EMPTY          string
	_MSG_ARCH_ISSUE                       string
	_MSG_CONFLICT_CREATE_USER_ISSUE       string
	createExternalAuthProviderUserSession main_usecases_interfaces.CreateExternalAuthProviderUserSession
	authProvider                          main_gateways.AuthProviderGateway
	buildTokenClaim                       main_usecases_interfaces.BuildTokenClaim
	lockGateway                           main_gateways.DistributedLockGateway
	logsMonitoringGateway                 main_gateways.LogsMonitoringGateway
	spanGateway                           main_gateways.SpanGateway
	messageUtils                          main_utils_messages.ApplicationMessages
}

func NewAtomicLockedCreateExternalAuthProviderUserSession(
	createExternalAuthProviderUserSession main_usecases_interfaces.CreateExternalAuthProviderUserSession,
	authProvider main_gateways.AuthProviderGateway,
	buildTokenClaim main_usecases_interfaces.BuildTokenClaim,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *AtomicLockedCreateExternalAuthProviderUserSession {
	return &AtomicLockedCreateExternalAuthProviderUserSession{
		_MSG_TOKEN_MUST_NOT_BE_EMPTY:          "providers.create.session.by.identity.provider.token.empty",
		_MSG_ARCH_ISSUE:                       "exceptions.architecture.application.issue",
		_MSG_CONFLICT_CREATE_USER_ISSUE:       "providers.create.user.lock.issue",
		createExternalAuthProviderUserSession: createExternalAuthProviderUserSession,
		authProvider:                          authProvider,
		buildTokenClaim:                       buildTokenClaim,
		lockGateway:                           lockGateway,
		logsMonitoringGateway:                 logsMonitoringGateway,
		spanGateway:                           spanGateway,
		messageUtils:                          messageBeans,
	}
}

func (this *AtomicLockedCreateExternalAuthProviderUserSession) Execute(
	ctx context.Context,
	args main_domains.ExternalProviderSessionArgs,
) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "AtomicLockedCreateExternalAuthProviderUserSession-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("AtomicLockedCreateExternalAuthProviderUserSession. provider: %s", args.GetProvider()))

	if main_utils.NewStringUtils().IsEmpty(args.GetToken()) {
		return *new(main_domains.SessionCredentials),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					this._MSG_TOKEN_MUST_NOT_BE_EMPTY))
	}

	sessionCredentials, errS := this.authProvider.CreateOauthExchangeSession(span.GetCtx(), args)
	if errS != nil {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(errS.Error()))
		return *new(main_domains.SessionCredentials), errS
	}

	tokenClaims, errT := this.buildTokenClaim.Execute(span.GetCtx(), sessionCredentials.GetAccessToken())
	if errT != nil {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(errT.Error()))
		return *new(main_domains.SessionCredentials), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				this._MSG_ARCH_ISSUE))
	}

	singleLock := this.lockGateway.GetWithScope(
		span.GetCtx(),
		main_domains_enums.LOCK_SCOPE_USER_MODIFICATION,
		tokenClaims.Email,
		90*time.Second)
	errLock := singleLock.Lock()

	if errLock == nil {
		defer this.lockGateway.UnlockAndLogIfError(span.GetCtx(), *singleLock)

		wc := writeconcern.Majority()
		rc := readconcern.Snapshot()
		transactionOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

		session, errCreateSession := main_configs_mongo.GetMongoDBClientBean().StartSession()
		if errCreateSession != nil {
			return *new(main_domains.SessionCredentials), main_domains_exceptions.
				NewInternalServerErrorExceptionSglMsg(this.messageUtils.
					GetDefaultLocale(_MSG_CREATE_NEW_USER_ARCH_ISSUE))
		}
		defer session.EndSession(context.Background())

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
			errP := this.createExternalAuthProviderUserSession.Execute(span.GetCtx(), args, tokenClaims, opts)
			if errP != nil {
				errResult = errP
				return errP
			}

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
				return *new(main_domains.SessionCredentials), main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(_MSG_CREATE_NEW_USER_ARCH_ISSUE))
			}
			return *new(main_domains.SessionCredentials), errResult
		}

		return sessionCredentials, nil

	} else {
		return *new(main_domains.SessionCredentials),
			main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
				this._MSG_CONFLICT_CREATE_USER_ISSUE))
	}

}
