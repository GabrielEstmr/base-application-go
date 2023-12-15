package main_configs_ff_lib_mongo

import (
	"baseapplicationgo/main/configs/ff/lib"
	main_configs_ff_lib_mongorepo_documents "baseapplicationgo/main/configs/ff/lib/mongo/documents"
	main_configs_ff_lib_mongo_repo "baseapplicationgo/main/configs/ff/lib/mongo/repo"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
)

type RegisterMethodsMongoImpl struct {
	repo *main_configs_ff_lib_mongo_repo.FeaturesMongoRepo
}

func NewRegisterMethodsMongoImpl(ffConfigData *main_configs_ff_lib.FfConfigData) *RegisterMethodsMongoImpl {
	return &RegisterMethodsMongoImpl{repo: main_configs_ff_lib_mongo_repo.NewFeaturesMongoRepo(ffConfigData)}
}

func (this *RegisterMethodsMongoImpl) getFeature(key string) (main_configs_ff_lib_resources.FeaturesData, error) {

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(main_configs_ff_lib_resources.FeaturesData), err
	}

	if byId.IsEmpty() {
		return *new(main_configs_ff_lib_resources.FeaturesData), nil
	}

	return byId.ToDomain(), nil
}

func (this *RegisterMethodsMongoImpl) RegisterFeatures(features main_configs_ff_lib_resources.Features) error {
	for k, v := range features {
		feature, err := this.getFeature(k)
		if err != nil {
			return err
		}
		if feature.IsEmpty() {
			_, err2 := this.repo.Save(main_configs_ff_lib_mongorepo_documents.NewFeaturesDataDocument(v))
			if err2 != nil {
				return err2
			}
		}
	}
	return nil
}
