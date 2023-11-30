package main_gateways_redis_repositories

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	main_gateways_redis_documents "baseapplicationgo/main/gateways/redis/documents"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
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

	val, err := this.redisClient.Set(context.TODO(), userRedisDocument.Id, userBytes, time.Hour).Result()
	log.Println(val)

	if err != nil {
		return main_gateways_redis_documents.UserRedisDocument{}, err
	}

	return userRedisDocument, nil
}

func (this *UserRedisRepository) FindById(indicatorId string) (
	main_gateways_redis_documents.UserRedisDocument, error) {

	result, err := this.redisClient.Get(context.TODO(), indicatorId).Result()

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
