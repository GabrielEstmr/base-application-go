package main_gateways_redis

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_redis_documents "baseapplicationgo/main/gateways/redis/documents"
	main_gateways_redis_repositories "baseapplicationgo/main/gateways/redis/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
)

type UserDatabaseCacheGatewayImpl struct {
	userRedisRepository main_gateways_redis_repositories.UserRedisRepository
	spanGateway         main_gateways.SpanGateway
}

func NewUserDatabaseCacheGatewayImpl(
	redisUserRepo main_gateways_redis_repositories.UserRedisRepository) *UserDatabaseCacheGatewayImpl {
	return &UserDatabaseCacheGatewayImpl{
		redisUserRepo,
		main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *UserDatabaseCacheGatewayImpl) Save(ctx context.Context, user main_domains.User,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-Save")
	defer span.End()

	userDocument := main_gateways_redis_documents.NewUserDocument(user)
	userDocument, err := this.userRedisRepository.Save(span.GetCtx(), userDocument)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) Update(ctx context.Context, user main_domains.User,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-Update")
	defer span.End()

	userDocument := main_gateways_redis_documents.NewUserDocument(user)
	userDocument, err := this.userRedisRepository.Update(span.GetCtx(), userDocument)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) FindById(ctx context.Context, id string,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-FindById")
	defer span.End()

	userDocument, err := this.userRedisRepository.FindById(span.GetCtx(), id)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) FindByDocumentId(ctx context.Context, documentNumber string,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-FindByDocumentId")
	defer span.End()

	userDocument, err := this.userRedisRepository.FindByDocumentId(span.GetCtx(), documentNumber)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) FindByUserName(ctx context.Context, userName string,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-FindByUserName")
	defer span.End()

	userDocument, err := this.userRedisRepository.FindByUserName(span.GetCtx(), userName)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) FindByEmail(ctx context.Context, email string,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-FindByEmail")
	defer span.End()

	userDocument, err := this.userRedisRepository.FindByEmail(span.GetCtx(), email)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) FindByDocumentIdOrUserNameOrEmail(
	ctx context.Context,
	documentId string,
	userName string,
	email string,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-FindByDocumentIdOrUserNameOrEmail")
	defer span.End()

	userDocument, err := this.FindByDocumentId(span.GetCtx(), documentId)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	if !userDocument.IsEmpty() {
		return userDocument, nil
	}

	userDocumentUN, errUN := this.FindByUserName(span.GetCtx(), userName)
	if errUN != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errUN.Error())
	}
	if !userDocumentUN.IsEmpty() {
		return userDocumentUN, nil
	}

	userDocumentE, errE := this.FindByEmail(span.GetCtx(), email)
	if errE != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errE.Error())
	}
	if !userDocumentE.IsEmpty() {
		return userDocumentE, nil
	}
	return *new(main_domains.User), nil
}
