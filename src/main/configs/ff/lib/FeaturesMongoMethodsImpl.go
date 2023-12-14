package main_configs_ff_lib

import (
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"errors"
)

type FeaturesMongoMethodsImpl struct {
	repo *FeaturesMongoRepo
}

func NewFeaturesMongoMethodsImpl(ffConfigData *FfConfigData) *FeaturesMongoMethodsImpl {
	return &FeaturesMongoMethodsImpl{repo: NewFeaturesMongoRepo(ffConfigData)}
}

func (this *FeaturesMongoMethodsImpl) getFeature(key string) (main_configs_ff_lib_resources.FeaturesData, error) {

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}

	if byId.IsEmpty() {
		return *new(main_configs_ff_lib_resources.FeaturesData), errors.New("feature doesn't exists")
	}

	return byId.ToDomain(), nil
}

func (this *FeaturesMongoMethodsImpl) IsEnabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.GetDefaultValue() == true, nil
}

func (this *FeaturesMongoMethodsImpl) IsDisabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.GetDefaultValue() == false, nil
}
