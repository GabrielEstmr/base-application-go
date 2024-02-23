package main_gateways_mongodb_repositories

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_mongo_collections "baseapplicationgo/main/configs/mongodb/collections"
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_utils "baseapplicationgo/main/gateways/mongodb/utils"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const _EMAILS_COLLECTION_NAME = main_configs_mongo_collections.EMAILS_COLLECTION_NAME
const _EMAILS_IDX_INDICATOR_MONGO_ID = main_configs_mongo.DEFAULT_BSON_ID_NAME

const _EMAIL_REPO_EVENT_ID = "eventId"
const _EMAIL_REPO_STATUS = "status"
const _EMAIL_REPO_EMAIL_PARAMS_EMAIL_TYPE = "emailParams.emailTemplateType"
const _EMAIL_REPO_EMAIL_APP_OWNER = "emailParams.appOwner"
const _EMAIL_REPO_ERROR_MSG = "errorMsg"
const _EMAIL_REPO_CREATED_DATE = "createdDate"
const _EMAIL_REPO_LAST_MODIFIED_DATE = "lastModifiedDate"

type EmailRepository struct {
	database    *mongo.Database
	spanGateway main_gateways.SpanGateway
}

func NewEmailRepository() *EmailRepository {
	return &EmailRepository{
		database:    main_configs_mongo.GetMongoDBDatabaseBean(),
		spanGateway: main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *EmailRepository) Save(ctx context.Context, email main_gateways_mongodb_documents.EmailDocument) (main_gateways_mongodb_documents.EmailDocument, error) {
	span := this.spanGateway.Get(ctx, "EmailRepository-Save")
	defer span.End()

	collection := this.database.Collection(_EMAILS_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	email.CreatedDate = now
	email.LastModifiedDate = now

	result, err := collection.InsertOne(span.GetCtx(), email)
	if err != nil {
		return main_gateways_mongodb_documents.EmailDocument{}, err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	email.Id = oid
	return email, nil
}

func (this *EmailRepository) FindById(ctx context.Context, id string) (main_gateways_mongodb_documents.EmailDocument, error) {
	span := this.spanGateway.Get(ctx, "EmailRepository-FindById")
	defer span.End()

	collection := this.database.Collection(_EMAILS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.EmailDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, nil
	}
	filter := bson.D{{_EMAILS_IDX_INDICATOR_MONGO_ID, objectId}}
	err2 := collection.FindOne(span.GetCtx(), filter).Decode(&result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, err2
	}
	return result, nil
}

func (this *EmailRepository) FindByEventId(ctx context.Context, eventId string) (main_gateways_mongodb_documents.EmailDocument, error) {
	span := this.spanGateway.Get(ctx, "EmailRepository-FindById")
	defer span.End()

	collection := this.database.Collection(_EMAILS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.EmailDocument

	filter := bson.D{{_EMAIL_REPO_EVENT_ID, eventId}}
	err2 := collection.FindOne(span.GetCtx(), filter).Decode(result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, err2
	}
	return result, nil
}

func (this *EmailRepository) Update(ctx context.Context, email main_gateways_mongodb_documents.EmailDocument) (main_gateways_mongodb_documents.EmailDocument, error) {
	span := this.spanGateway.Get(ctx, "EmailRepository-Update")
	defer span.End()

	collection := this.database.Collection(_EMAILS_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	email.LastModifiedDate = now

	filter := bson.D{{_EMAILS_IDX_INDICATOR_MONGO_ID, email.Id}}
	update := bson.D{{"$set", bson.D{
		{_EMAIL_REPO_STATUS, email.Status},
		{_EMAIL_REPO_ERROR_MSG, email.ErrorMsg},
		{_EMAIL_REPO_LAST_MODIFIED_DATE, email.LastModifiedDate},
	}}}

	_, err := collection.UpdateOne(span.GetCtx(), filter, update)
	if err != nil {
		return main_gateways_mongodb_documents.EmailDocument{}, err
	}
	return email, nil
}

func (this *EmailRepository) FindByFilter(ctx context.Context, filter main_domains.FindEmailFilter,
	pageable main_domains.Pageable) (main_domains.Page, error) {
	span := this.spanGateway.Get(ctx, "UserRepository-FindByFilter")
	defer span.End()

	collection := this.database.Collection(_EMAILS_COLLECTION_NAME)

	log.Println(len(pageable.GetSort()))
	if pageable.HasEmptySort() {
		defaultSort := make(map[string]int)
		defaultSort[_EMAILS_IDX_INDICATOR_MONGO_ID] = 1
		pageable.SetSort(defaultSort)
	}

	opt := main_gateways_mongodb_utils.NewPageableUtils().GetOptsFromPageable(pageable)

	filterCriterias := bson.D{}

	if len(filter.GetIds()) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAILS_IDX_INDICATOR_MONGO_ID, Value: bson.M{"$in": filter.GetIds()}})
	}
	if len(filter.GetStatuses()) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAIL_REPO_STATUS, Value: bson.M{"$in": filter.GetStatuses()}})
	}
	if len(filter.GetEmailTemplateTypes()) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAIL_REPO_EMAIL_PARAMS_EMAIL_TYPE, Value: bson.M{"$in": filter.GetEmailTemplateTypes()}})
	}
	if len(filter.GetAppOwners()) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAIL_REPO_EMAIL_APP_OWNER, Value: bson.M{"$in": filter.GetAppOwners()}})
	}
	if !filter.GetStartCreatedDate().IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAIL_REPO_CREATED_DATE, Value: bson.M{"$gte": filter.GetStartCreatedDate()}})
	}
	if !filter.GetEndCreatedDate().IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAIL_REPO_CREATED_DATE, Value: bson.M{"$lt": filter.GetEndCreatedDate()}})
	}
	if !filter.GetStartLastModifiedDate().IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAIL_REPO_LAST_MODIFIED_DATE, Value: bson.M{"$gte": filter.GetStartLastModifiedDate()}})
	}
	if !filter.GetEndLastModifiedDate().IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _EMAIL_REPO_LAST_MODIFIED_DATE, Value: bson.M{"$lt": filter.GetEndLastModifiedDate()}})
	}

	var results []main_gateways_mongodb_documents.EmailDocument
	cursor, errF := collection.Find(span.GetCtx(), filterCriterias, opt)
	if errF != nil {
		return *new(main_domains.Page), errF
	}
	if errC := cursor.All(span.GetCtx(), &results); errC != nil {
		return *new(main_domains.Page), errC
	}
	for _, result := range results {
		_, errM := json.Marshal(result)
		if errM != nil {
			return *new(main_domains.Page), errM
		}
	}

	numberDocs, errC := collection.CountDocuments(span.GetCtx(), filterCriterias)
	if errC != nil {
		return *new(main_domains.Page), errC
	}

	var contents []any
	for _, value := range results {
		contents = append(contents, value)
	}

	return *main_domains.NewPage(contents, pageable.GetPage(), pageable.GetSize(), numberDocs), nil
}
