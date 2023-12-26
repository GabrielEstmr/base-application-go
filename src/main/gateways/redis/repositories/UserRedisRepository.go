package main_gateways_redis_repositories

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	main_gateways_redis_documents "baseapplicationgo/main/gateways/redis/documents"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserRedisRepository struct {
	redisClient *redis.Client
}

func NewUserRedisRepository() *UserRedisRepository {
	return &UserRedisRepository{redisClient: main_configs_cache.GetRedisClusterBean()}
}

func (this *UserRedisRepository) Save(
	userRedisDocument main_gateways_redis_documents.UserRedisDocument) (
	main_gateways_redis_documents.UserRedisDocument, error) {

	userBytes, err := json.Marshal(userRedisDocument)
	if err != nil {
		return main_gateways_redis_documents.UserRedisDocument{}, err
	}

	for _, value := range userRedisDocument.GetKeys() {
		_, err := this.redisClient.Set(context.TODO(), value, userBytes, time.Hour).Result()
		if err != nil {
			return main_gateways_redis_documents.UserRedisDocument{}, err
		}
	}
	return userRedisDocument, nil
}

func (this *UserRedisRepository) FindById(indicatorId string) (
	main_gateways_redis_documents.UserRedisDocument, error) {

	result, err := this.redisClient.Get(context.TODO(), main_gateways_redis_documents.USER_DOC__ID_NAME_PREFIX+indicatorId).Result()

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

func (this *UserRedisRepository) FindByDocumentNumber(documentNumber string) (
	main_gateways_redis_documents.UserRedisDocument, error) {

	result, err := this.redisClient.Get(context.TODO(), main_gateways_redis_documents.USER_DOC__IDX_DOCUMENT_NUMBER_NAME_PREFIX+documentNumber).Result()

	if err == redis.Nil {
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
