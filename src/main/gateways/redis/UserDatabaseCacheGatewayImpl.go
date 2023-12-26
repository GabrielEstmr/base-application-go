package main_gateways_redis

import (
	main_domains "baseapplicationgo/main/domains"
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

func (this *UserDatabaseCacheGatewayImpl) Save(ctx context.Context, user main_domains.User) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-Save")
	defer span.End()

	userDocument := main_gateways_redis_documents.NewUserRedisDocument(user)
	userDocument, err := this.userRedisRepository.Save(span.GetCtx(), userDocument)
	if err != nil {
		return userDocument.ToDomain(), err
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) FindById(ctx context.Context, id string) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-FindById")
	defer span.End()

	userDocument, err := this.userRedisRepository.FindById(span.GetCtx(), id)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseCacheGatewayImpl) FindByDocumentNumber(ctx context.Context, documentNumber string) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "UserDatabaseCacheGatewayImpl-FindByDocumentNumber")
	defer span.End()

	userDocument, err := this.userRedisRepository.FindByDocumentNumber(span.GetCtx(), documentNumber)
	return userDocument.ToDomain(), err
}
