package main_gateways_mongodb_repositories

import (
	configsMongo "baseapplicationgo/main/configs/mongodb"
	gatewaysMongodbDocuments "baseapplicationgo/main/gateways/mongodb/documents"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ACCOUNTS_COLLECTION_NAME = "accounts"
const IDX_INDICATOR_MONGO_ID = "_id"

type AccountRepository struct {
	database *mongo.Database
}

func NewIndicatorRepository() *AccountRepository {
	return &AccountRepository{database: configsMongo.GetMongoDbBean()}
}

func (thisRepository *AccountRepository) FindById(
	id string) (gatewaysMongodbDocuments.AccountDocument, error) {

	collection := thisRepository.database.Collection(ACCOUNTS_COLLECTION_NAME)
	var result gatewaysMongodbDocuments.AccountDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.D{{IDX_INDICATOR_MONGO_ID, objectId}}
	err2 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil && err2 != mongo.ErrNoDocuments {
		return result, err2
	}
	return result, nil
}

func (thisRepository *AccountRepository) Save(
	accountDocument gatewaysMongodbDocuments.AccountDocument) (string, error) {

	collection := thisRepository.database.Collection(ACCOUNTS_COLLECTION_NAME)
	result, err := collection.InsertOne(context.TODO(), accountDocument)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(result.InsertedID), nil
}
