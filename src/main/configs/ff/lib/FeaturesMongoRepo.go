package main_configs_ff_lib

import (
	main_configs_ff_lib_mongorepo_documents "baseapplicationgo/main/configs/ff/lib/mongorepo/documents"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const _KEY = "_id"

type FeaturesMongoRepo struct {
	ffConfigData *FfConfigData
}

func NewFeaturesMongoRepo(ffConfigData *FfConfigData) *FeaturesMongoRepo {
	return &FeaturesMongoRepo{ffConfigData: ffConfigData}
}

func (this *FeaturesMongoRepo) Save(
	feature main_configs_ff_lib_mongorepo_documents.FeaturesDataDocument,
) (main_configs_ff_lib_mongorepo_documents.FeaturesDataDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())

	result, err := collection.InsertOne(context.TODO(), feature)
	if err != nil {
		return *new(main_configs_ff_lib_mongorepo_documents.FeaturesDataDocument), err
	}

	key, _ := result.InsertedID.(string)
	feature.Key = key
	return feature, nil
}

func (this *FeaturesMongoRepo) FindById(id string) (*main_configs_ff_lib_mongorepo_documents.FeaturesDataDocument, error) {
	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	var result main_configs_ff_lib_mongorepo_documents.FeaturesDataDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &result, nil
	}
	filter := bson.D{{_KEY, objectId}}
	err2 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return &result, nil
		}
		return &result, err2
	}
	return &result, nil
}
