package main_configs_ff_lib

import (
	main_configs_ff_lib_mongorepo "baseapplicationgo/main/configs/ff/lib/mongorepo"
	main_configs_ff_lib_mongorepo_documents "baseapplicationgo/main/configs/ff/lib/mongorepo/documents"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
)

type RegisterMethodsMongoImpl struct {
	repo *main_configs_ff_lib_mongorepo.FeaturesMongoRepo
}

func NewRegisterMethodsMongoImpl() *RegisterMethodsMongoImpl {
	return &RegisterMethodsMongoImpl{repo: main_configs_ff_lib_mongorepo.NewFeaturesMongoRepo()}
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

func (this *RegisterMethodsMongoImpl) RegisterFeatures(features map[string]main_configs_ff_lib_resources.FeaturesData) error {
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
