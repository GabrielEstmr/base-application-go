package main_configs_ff_lib_mongo_repo

import (
	"baseapplicationgo/main/configs/ff/lib"
	main_configs_ff_lib_mongo_documents "baseapplicationgo/main/configs/ff/lib/mongo/documents"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const _KEY = "_id"

type FeaturesMongoRepo struct {
	ffConfigData *main_configs_ff_lib.FfConfigData
}

func NewFeaturesMongoRepo(ffConfigData *main_configs_ff_lib.FfConfigData) *FeaturesMongoRepo {
	return &FeaturesMongoRepo{ffConfigData: ffConfigData}
}

func (this *FeaturesMongoRepo) Save(
	feature main_configs_ff_lib_mongo_documents.FeaturesDataDocument,
) (main_configs_ff_lib_mongo_documents.FeaturesDataDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesDbName())

	result, err := collection.InsertOne(context.TODO(), feature)
	if err != nil {
		return *new(main_configs_ff_lib_mongo_documents.FeaturesDataDocument), err
	}

	key, _ := result.InsertedID.(string)
	feature.Key = key
	return feature, nil
}

func (this *FeaturesMongoRepo) Update(
	feature main_configs_ff_lib_mongo_documents.FeaturesDataDocument,
) (main_configs_ff_lib_mongo_documents.FeaturesDataDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesDbName())
	filter := bson.D{{_KEY, feature.Key}}
	update := bson.D{{"$set", bson.D{{"defaultValue", feature.DefaultValue}}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return *new(main_configs_ff_lib_mongo_documents.FeaturesDataDocument), err
	}
	return feature, nil
}

func (this *FeaturesMongoRepo) FindById(id string) (*main_configs_ff_lib_mongo_documents.FeaturesDataDocument, error) {
	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesDbName())
	var result main_configs_ff_lib_mongo_documents.FeaturesDataDocument
	filter := bson.D{{_KEY, id}}
	err2 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return &result, nil
		}
		return &result, err2
	}
	return &result, nil
}
