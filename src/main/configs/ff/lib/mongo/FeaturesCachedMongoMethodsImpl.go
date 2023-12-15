package main_configs_ff_lib_mongo

import (
	"baseapplicationgo/main/configs/ff/lib"
	main_configs_ff_lib_mongo_repo "baseapplicationgo/main/configs/ff/lib/mongo/repo"
	main_configs_ff_lib_redis_documents "baseapplicationgo/main/configs/ff/lib/redis/documents"
	main_configs_ff_lib_redis_repo "baseapplicationgo/main/configs/ff/lib/redis/repo"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"errors"
)

type FeaturesCachedMongoMethodsImpl struct {
	repo      *main_configs_ff_lib_mongo_repo.FeaturesMongoRepo
	repoCache *main_configs_ff_lib_redis_repo.FeaturesRedisRepo
}

func NewFeaturesCachedMongoMethodsImpl(ffConfigData *main_configs_ff_lib.FfConfigData) *FeaturesCachedMongoMethodsImpl {
	return &FeaturesCachedMongoMethodsImpl{
		repo:      main_configs_ff_lib_mongo_repo.NewFeaturesMongoRepo(ffConfigData),
		repoCache: main_configs_ff_lib_redis_repo.NewFeaturesRedisRepo(ffConfigData),
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
	return feature.IsEnabled(), nil
}

func (this *FeaturesCachedMongoMethodsImpl) IsDisabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsDisabled(), nil
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
