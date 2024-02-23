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
	docTtl      time.Duration
	redisClient *redis.Client
	spanGateway main_gateways.SpanGateway
}

func NewUserRedisRepository() *UserRedisRepository {
	return &UserRedisRepository{
		docTtl:      time.Hour,
		redisClient: main_configs_cache.GetRedisClusterBean(),
		spanGateway: main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *UserRedisRepository) Save(ctx context.Context,
	userRedisDocument main_gateways_redis_documents.UserDocument) (
	main_gateways_redis_documents.UserDocument, error) {

	span := this.spanGateway.Get(ctx, "UserRedisRepository-Save")
	defer span.End()

	userBytes, errJ := json.Marshal(userRedisDocument)
	if errJ != nil {
		return *new(main_gateways_redis_documents.UserDocument), errJ
	}

	for _, value := range userRedisDocument.GetKeys() {
		_, err := this.redisClient.Set(context.TODO(), value, userBytes, this.docTtl).Result()
		if err != nil {
			return *new(main_gateways_redis_documents.UserDocument), err
		}
	}
	return userRedisDocument, nil
}

func (this *UserRedisRepository) Update(ctx context.Context,
	userRedisDocument main_gateways_redis_documents.UserDocument) (
	main_gateways_redis_documents.UserDocument, error) {

	span := this.spanGateway.Get(ctx, "UserRedisRepository-Update")
	defer span.End()

	return this.Save(span.GetCtx(), userRedisDocument)
}

func (this *UserRedisRepository) FindById(ctx context.Context, indicatorId string) (

	main_gateways_redis_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRedisRepository-FindById")
	defer span.End()

	result, err := this.redisClient.Get(context.TODO(), main_gateways_redis_documents.USER_DOC_ID_NAME_PREFIX+indicatorId).Result()

	if errors.Is(err, redis.Nil) {
		return *new(main_gateways_redis_documents.UserDocument), nil
	}

	if err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	var cachedIndicatorDocument main_gateways_redis_documents.UserDocument
	if err = json.Unmarshal([]byte(result), &cachedIndicatorDocument); err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	return cachedIndicatorDocument, nil
}

func (this *UserRedisRepository) FindByDocumentId(ctx context.Context, documentNumber string) (
	main_gateways_redis_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRedisRepository-FindByDocumentNumber")
	defer span.End()

	result, err := this.redisClient.Get(context.TODO(),
		main_gateways_redis_documents.USER_DOC_IDX_DOCUMENT_ID_NAME_PREFIX+documentNumber).Result()

	if errors.Is(err, redis.Nil) {
		return *new(main_gateways_redis_documents.UserDocument), nil
	}

	if err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	var cachedIndicatorDocument main_gateways_redis_documents.UserDocument
	if err = json.Unmarshal([]byte(result), &cachedIndicatorDocument); err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	return cachedIndicatorDocument, nil
}

func (this *UserRedisRepository) FindByUserName(ctx context.Context, userName string) (
	main_gateways_redis_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRedisRepository-FindByDocumentNumber")
	defer span.End()

	result, err := this.redisClient.Get(context.TODO(),
		main_gateways_redis_documents.USER_DOC_IDX_USERNAME_NAME_PREFIX+userName).Result()

	if errors.Is(err, redis.Nil) {
		return *new(main_gateways_redis_documents.UserDocument), nil
	}

	if err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	var cachedIndicatorDocument main_gateways_redis_documents.UserDocument
	if err = json.Unmarshal([]byte(result), &cachedIndicatorDocument); err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	return cachedIndicatorDocument, nil
}

func (this *UserRedisRepository) FindByEmail(ctx context.Context, email string) (
	main_gateways_redis_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRedisRepository-FindByDocumentNumber")
	defer span.End()

	result, err := this.redisClient.Get(context.TODO(),
		main_gateways_redis_documents.USER_DOC_IDX_EMAIL_NAME_PREFIX+email).Result()

	if errors.Is(err, redis.Nil) {
		return *new(main_gateways_redis_documents.UserDocument), nil
	}

	if err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	var cachedIndicatorDocument main_gateways_redis_documents.UserDocument
	if err = json.Unmarshal([]byte(result), &cachedIndicatorDocument); err != nil {
		return *new(main_gateways_redis_documents.UserDocument), err
	}

	return cachedIndicatorDocument, nil
}
