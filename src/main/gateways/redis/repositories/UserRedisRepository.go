package main_gateways_redis_repositories

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_redis_documents "baseapplicationgo/main/gateways/redis/documents"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserRedisRepository struct {
	redisClient *redis.Client
	spanGateway main_gateways.SpanGateway
}

func NewUserRedisRepository() *UserRedisRepository {
	return &UserRedisRepository{
		main_configs_cache.GetRedisClusterBean(),
		main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *UserRedisRepository) Save(ctx context.Context,
	userRedisDocument main_gateways_redis_documents.UserRedisDocument) (
	main_gateways_redis_documents.UserRedisDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRedisRepository-Save")
	defer span.End()

	userBytes, err := json.Marshal(userRedisDocument)
	if err != nil {
		return main_gateways_redis_documents.UserRedisDocument{}, err
	}

	for _, value := range userRedisDocument.GetKeys() {
		_, err := this.redisClient.Set(span.GetCtx(), value, userBytes, time.Hour).Result()
		if err != nil {
			return main_gateways_redis_documents.UserRedisDocument{}, err
		}
	}
	return userRedisDocument, nil
}

func (this *UserRedisRepository) FindById(ctx context.Context, indicatorId string) (
	main_gateways_redis_documents.UserRedisDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRedisRepository-FindById")
	defer span.End()

	result, err := this.redisClient.Get(span.GetCtx(), main_gateways_redis_documents.USER_DOC__ID_NAME_PREFIX+indicatorId).Result()

	if errors.Is(err, redis.Nil) {
		return main_gateways_redis_documents.UserRedisDocument{}, nil
	}

	if err != nil {
		return main_gateways_redis_documents.UserRedisDocument{}, err
	}

	var cachedIndicatorDocument main_gateways_redis_documents.UserRedisDocument
	if err = json.Unmarshal([]byte(result), &cachedIndicatorDocument); err != nil {
		return main_gateways_redis_documents.UserRedisDocument{}, err
	}

	return cachedIndicatorDocument, nil
}

func (this *UserRedisRepository) FindByDocumentNumber(ctx context.Context, documentNumber string) (
	main_gateways_redis_documents.UserRedisDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRedisRepository-FindByDocumentNumber")
	defer span.End()

	result, err := this.redisClient.Get(span.GetCtx(), main_gateways_redis_documents.USER_DOC__IDX_DOCUMENT_NUMBER_NAME_PREFIX+documentNumber).Result()

	if errors.Is(err, redis.Nil) {
		return main_gateways_redis_documents.UserRedisDocument{}, nil
	}

	if err != nil {
		return main_gateways_redis_documents.UserRedisDocument{}, err
	}

	var cachedIndicatorDocument main_gateways_redis_documents.UserRedisDocument
	if err = json.Unmarshal([]byte(result), &cachedIndicatorDocument); err != nil {
		return main_gateways_redis_documents.UserRedisDocument{}, err
	}

	return cachedIndicatorDocument, nil
}
