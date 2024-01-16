package main_configs_ff_lib_mongo

import (
	"baseapplicationgo/main/configs/ff/lib"
	"baseapplicationgo/main/configs/ff/lib/mongo/repo"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"errors"
)

type FeaturesMongoMethodsImpl struct {
	repo *main_configs_ff_lib_mongo_repo.FeaturesMongoRepo
}

func NewFeaturesMongoMethodsImpl(ffConfigData *main_configs_ff_lib.FfConfigData) *FeaturesMongoMethodsImpl {
	return &FeaturesMongoMethodsImpl{repo: main_configs_ff_lib_mongo_repo.NewFeaturesMongoRepo(ffConfigData)}
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
	return feature.IsEnabled(), nil
}

func (this *FeaturesMongoMethodsImpl) IsDisabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsDisabled(), nil
}

func (this *FeaturesMongoMethodsImpl) Enable(key string) (main_configs_ff_lib_resources.FeaturesData, error) {
	featureDoc, err := this.repo.FindById(key)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}
	if featureDoc.IsEmpty() {
		return *new(main_configs_ff_lib_resources.FeaturesData), errors.New("feature not found")
	}

	if featureDoc.IsDisabled() {
		featureDoc.DefaultValue = true
		savedFeatureDoc, err := this.repo.Update(*featureDoc)
		if err != nil {
			return *new(main_configs_ff_lib_resources.FeaturesData), err
		}
		return savedFeatureDoc.ToDomain(), nil
	}
	return featureDoc.ToDomain(), nil
}

func (this *FeaturesMongoMethodsImpl) Disable(key string) (main_configs_ff_lib_resources.FeaturesData, error) {
	featureDoc, err := this.repo.FindById(key)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}
	if featureDoc.IsEmpty() {
		return *new(main_configs_ff_lib_resources.FeaturesData), errors.New("feature not found")
	}

	if featureDoc.IsEnabled() {
		featureDoc.DefaultValue = false
		savedFeatureDoc, err := this.repo.Update(*featureDoc)
		if err != nil {
			return *new(main_configs_ff_lib_resources.FeaturesData), err
		}
		return savedFeatureDoc.ToDomain(), nil
	}
	return featureDoc.ToDomain(), nil
}
