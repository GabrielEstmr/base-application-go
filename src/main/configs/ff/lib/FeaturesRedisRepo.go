package main_configs_ff_lib

import (
	main_configs_ff_lib_redis_documents "baseapplicationgo/main/configs/ff/lib/redis/documents"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type FeaturesRedisRepo struct {
	ffConfigData *FfConfigData
}

func NewFeaturesRedisRepo(ffConfigData *FfConfigData) *FeaturesRedisRepo {
	return &FeaturesRedisRepo{ffConfigData}
}

func (this *FeaturesRedisRepo) Save(
	feature main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument) (
	main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument, error) {

	featureBytes, err := json.Marshal(feature)
	if err != nil {
		return *new(main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument), err
	}

	this.ffConfigData.GetCacheClient().Set(context.TODO(), feature.Key, featureBytes, time.Hour).Result()
	return feature, nil
}

func (this *FeaturesRedisRepo) FindById(key string) (
	main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument, error) {

	result, err := this.ffConfigData.GetCacheClient().
		Get(context.TODO(), this.ffConfigData.GetCachingPrefix()+"_"+key).Result()

	if errors.Is(err, redis.Nil) {
		return *new(main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument), nil
	}

	if err != nil {
		return *new(main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument), err
	}

	var feature main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument
	if err = json.Unmarshal([]byte(result), &feature); err != nil {
		return *new(main_configs_ff_lib_redis_documents.FeaturesDataRedisDocument), err
	}

	return feature, nil
}