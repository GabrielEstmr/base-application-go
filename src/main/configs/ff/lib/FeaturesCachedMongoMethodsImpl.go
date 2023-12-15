package main_configs_ff_lib

import (
	main_configs_ff_lib_redis_documents "baseapplicationgo/main/configs/ff/lib/redis/documents"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"errors"
)

type FeaturesCachedMongoMethodsImpl struct {
	repo      *FeaturesMongoRepo
	repoCache *FeaturesRedisRepo
}

func NewFeaturesCachedMongoMethodsImpl(ffConfigData *FfConfigData) *FeaturesCachedMongoMethodsImpl {
	return &FeaturesCachedMongoMethodsImpl{
		repo:      NewFeaturesMongoRepo(ffConfigData),
		repoCache: NewFeaturesRedisRepo(ffConfigData),
	}
}

func (this *FeaturesCachedMongoMethodsImpl) getFeature(key string) (main_configs_ff_lib_resources.FeaturesData, error) {

	byIdCached, err := this.repoCache.FindById(key)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}
	if !byIdCached.IsEmpty() {
		return byIdCached.ToDomain(), nil
	}

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}

	if byId.IsEmpty() {
		return *new(main_configs_ff_lib_resources.FeaturesData), errors.New("feature doesn't exists")
	}

	go this.repoCache.Save(main_configs_ff_lib_redis_documents.NewFeaturesDataRedisDocument(byId.ToDomain()))
	return byId.ToDomain(), nil
}

func (this *FeaturesCachedMongoMethodsImpl) IsEnabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.GetDefaultValue() == true, nil
}

func (this *FeaturesCachedMongoMethodsImpl) IsDisabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.GetDefaultValue() == false, nil
}

func (this *FeaturesCachedMongoMethodsImpl) Enable(key string) (main_configs_ff_lib_resources.FeaturesData, error) {
	featureDoc, errFind := this.repo.FindById(key)
	if errFind != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), errFind
	}

	featureDoc.DefaultValue = true
	_, errCache := this.repoCache.Save(main_configs_ff_lib_redis_documents.NewFeaturesDataRedisDocument(featureDoc.ToDomain()))
	if errCache != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), errCache
	}

	savedFeatureDoc, err := this.repo.Update(*featureDoc)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}
	return savedFeatureDoc.ToDomain(), nil
}

func (this *FeaturesCachedMongoMethodsImpl) Disable(key string) (main_configs_ff_lib_resources.FeaturesData, error) {
	featureDoc, errFind := this.repo.FindById(key)
	if errFind != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), errFind
	}

	featureDoc.DefaultValue = false
	_, errCache := this.repoCache.Save(main_configs_ff_lib_redis_documents.NewFeaturesDataRedisDocument(featureDoc.ToDomain()))
	if errCache != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), errCache
	}

	savedFeatureDoc, err := this.repo.Update(*featureDoc)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}
	return savedFeatureDoc.ToDomain(), nil
}
