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

const USERS_COLLECTION_NAME = "users"
const USERS_IDX_INDICATOR_MONGO_ID = "_id"

type UserRepository struct {
	database *mongo.Database
}

func NewUserRepository() *UserRepository {
	return &UserRepository{database: configsMongo.GetMongoDbBean()}
}

func (thisRepository *UserRepository) FindById(id string) (gatewaysMongodbDocuments.UserDocument, error) {

	collection := thisRepository.database.Collection(USERS_COLLECTION_NAME)
	var result gatewaysMongodbDocuments.UserDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.D{{USERS_IDX_INDICATOR_MONGO_ID, objectId}}
	err2 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil && err2 != mongo.ErrNoDocuments {
		return result, err2
	}
	return result, nil
}

func (thisRepository *UserRepository) Save(
	userDocument gatewaysMongodbDocuments.UserDocument) (string, error) {

	collection := thisRepository.database.Collection(USERS_COLLECTION_NAME)
	result, err := collection.InsertOne(context.TODO(), userDocument)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(result.InsertedID), nil
}
