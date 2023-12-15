package main_configs_ff_lib_redis

import (
	"baseapplicationgo/main/configs/ff/lib"
	"baseapplicationgo/main/configs/ff/lib/redis/repo"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"errors"
)

type FeaturesRedisMethodsImpl struct {
	repo *main_configs_ff_lib_redis_repo.FeaturesRedisRepo
}

func NewFeaturesRedisMethodsImpl(ffConfigData *main_configs_ff_lib.FfConfigData) *FeaturesRedisMethodsImpl {
	return &FeaturesRedisMethodsImpl{repo: main_configs_ff_lib_redis_repo.NewFeaturesRedisRepo(ffConfigData)}
}

func (this *FeaturesRedisMethodsImpl) getFeature(key string) (main_configs_ff_lib_resources.FeaturesData, error) {

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}

	if byId.IsEmpty() {
		return *new(main_configs_ff_lib_resources.FeaturesData), errors.New("feature doesn't exists")
	}

	return byId.ToDomain(), nil
}

func (this *FeaturesRedisMethodsImpl) IsEnabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsEnabled(), nil
}

func (this *FeaturesRedisMethodsImpl) IsDisabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsDisabled(), nil
}
