package main_gateways_mongodb_repositories

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_mongo_collections "baseapplicationgo/main/configs/mongodb/collections"
	main_domains "baseapplicationgo/main/domains"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const _USERS_COLLECTION_NAME = main_configs_mongo_collections.USERS_COLLECTION_NAME
const _USERS_IDX_INDICATOR_MONGO_ID = main_configs_mongo.DEFAULT_BSON_ID_NAME
const _DOCUMENT_NUMBER = "documentNumber"

type UserRepository struct {
	database *mongo.Database
}

func NewUserRepository() *UserRepository {
	return &UserRepository{database: main_configs_mongo.GetMongoDBClient()}
}

func (this *UserRepository) Save(user main_gateways_mongodb_documents.UserDocument) (main_gateways_mongodb_documents.UserDocument, error) {
	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	indexModel := mongo.IndexModel{
		Keys: bson.D{{_DOCUMENT_NUMBER, 1}},
	}
	name, err := this.database.Collection(_USERS_COLLECTION_NAME).Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)

	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return main_gateways_mongodb_documents.UserDocument{}, err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	user.Id = oid
	return user, nil
}

func (this *UserRepository) FindById(id string) (*main_gateways_mongodb_documents.UserDocument, error) {
	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.UserDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &result, nil
	}
	filter := bson.D{{_USERS_IDX_INDICATOR_MONGO_ID, objectId}}
	err2 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return &result, nil
		}
		return &result, err2
	}
	return &result, nil
}

func (this *UserRepository) FindByDocumentNumber(documentNumber string) (*main_gateways_mongodb_documents.UserDocument, error) {
	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	filter := bson.D{{_DOCUMENT_NUMBER, documentNumber}}
	var result main_gateways_mongodb_documents.UserDocument
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &result, nil
		}
		return &result, err
	}
	return &result, nil
}

func (this *UserRepository) FindByFilter(filter main_domains.FindUserFilter) (*main_domains.Page, error) {
	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	filter := bson.D{{_DOCUMENT_NUMBER, documentNumber}}
	var result main_gateways_mongodb_documents.UserDocument
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &result, nil
		}
		return &result, err
	}
	return &result, nil
}
