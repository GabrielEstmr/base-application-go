package main_gateways_mongodb_repositories

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_mongo_collections "baseapplicationgo/main/configs/mongodb/collections"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const _TRANSACTION_COLLECTION_NAME = main_configs_mongo_collections.TRANSACTIONS_COLLECTION_NAME

const _TRANSACTION_REPO_ACCOUNT = "accountId"

type TransactionRepository struct {
	database    *mongo.Database
	spanGateway main_gateways.SpanGateway
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		database:    main_configs_mongo.GetMongoDBDatabaseBean(),
		spanGateway: main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *TransactionRepository) Save(
	ctx context.Context,
	transaction main_gateways_mongodb_documents.TransactionDocument,
) (main_gateways_mongodb_documents.TransactionDocument, error) {

	span := this.spanGateway.Get(ctx, "TransactionRepository-Save")
	defer span.End()

	collection := this.database.Collection(_TRANSACTION_COLLECTION_NAME)
	indexModel := mongo.IndexModel{
		Keys: bson.D{{_TRANSACTION_REPO_ACCOUNT, 1}},
	}
	_, err := this.database.Collection(_TRANSACTION_COLLECTION_NAME).Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	now := primitive.NewDateTimeFromTime(time.Now())
	transaction.CreatedDate = now
	transaction.LastModifiedDate = now

	result, err := collection.InsertOne(context.TODO(), transaction)
	if err != nil {
		return main_gateways_mongodb_documents.TransactionDocument{}, err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	transaction.Id = oid
	return transaction, nil
}
