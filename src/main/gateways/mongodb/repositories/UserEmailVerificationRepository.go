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
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const _USER_VERIFICATION_EMAIL_COLLECTION_NAME = main_configs_mongo_collections.USER_VERIFICATION_EMAILS_COLLECTION_NAME
const _USER_VERIFICATION_EMAIL_IDX_INDICATOR_MONGO_ID = main_configs_mongo.DEFAULT_BSON_ID_NAME

type UserEmailVerificationRepository struct {
	_USER_ID            string
	_VERIFICATION_CODE  string
	_STATUS             string
	_CREATED_DATE       string
	_LAST_MODIFIED_DATE string
	database            *mongo.Database
	spanGateway         main_gateways.SpanGateway
}

func NewUserEmailVerificationRepository() *UserEmailVerificationRepository {
	return &UserEmailVerificationRepository{
		_USER_ID:            "userId",
		_VERIFICATION_CODE:  "verificationCode",
		_STATUS:             "status",
		_CREATED_DATE:       "createdDate",
		_LAST_MODIFIED_DATE: "lastModifiedDate",
		database:            main_configs_mongo.GetMongoDBDatabaseBean(),
		spanGateway:         main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *UserEmailVerificationRepository) Save(
	ctx context.Context,
	userVerificationEmail main_gateways_mongodb_documents.UserEmailVerificationDocument,
	options main_domains.DatabaseOptions,
) (
	main_gateways_mongodb_documents.UserEmailVerificationDocument, error,
) {
	span := this.spanGateway.Get(ctx, "UserEmailVerificationRepository-Save")
	defer span.End()

	collection := this.database.Collection(_USER_VERIFICATION_EMAIL_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	userVerificationEmail.CreatedDate = now
	userVerificationEmail.LastModifiedDate = now

	fnInsert := func(
		ctx context.Context,
		userVerificationEmail main_gateways_mongodb_documents.UserEmailVerificationDocument,
		options main_domains.DatabaseOptions,
	) (*mongo.InsertOneResult, error) {
		if options == nil {
			return collection.InsertOne(ctx, userVerificationEmail)
		} else {
			return collection.InsertOne(options.GetSession(), userVerificationEmail)
		}
	}

	result, err := fnInsert(span.GetCtx(), userVerificationEmail, options)
	if err != nil {
		return *new(main_gateways_mongodb_documents.UserEmailVerificationDocument), err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	userVerificationEmail.Id = oid
	return userVerificationEmail, nil
}

func (this *UserEmailVerificationRepository) Update(
	ctx context.Context,
	userVerificationEmail main_gateways_mongodb_documents.UserEmailVerificationDocument,
	options main_domains.DatabaseOptions,
) (
	main_gateways_mongodb_documents.UserEmailVerificationDocument, error,
) {
	span := this.spanGateway.Get(ctx, "UserEmailVerificationRepository-Update")
	defer span.End()

	collection := this.database.Collection(_USER_VERIFICATION_EMAIL_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	userVerificationEmail.LastModifiedDate = now

	filter := bson.D{{_USER_VERIFICATION_EMAIL_IDX_INDICATOR_MONGO_ID, userVerificationEmail.Id}}

	update := bson.D{{"$set", bson.D{
		{this._USER_ID, userVerificationEmail.UserId},
		{this._VERIFICATION_CODE, userVerificationEmail.VerificationCode},
		{this._STATUS, userVerificationEmail.Status},
		{this._CREATED_DATE, userVerificationEmail.CreatedDate},
		{this._LAST_MODIFIED_DATE, userVerificationEmail.LastModifiedDate},
	}}}

	fnUpdate := func(
		ctx context.Context,
		filter bson.D,
		update bson.D,
		options main_domains.DatabaseOptions,
	) (*mongo.UpdateResult, error) {
		if options == nil {
			return collection.UpdateOne(ctx, filter, update)
		} else {
			return collection.UpdateOne(options.GetSession(), filter, update)
		}
	}

	_, err := fnUpdate(span.GetCtx(), filter, update, options)
	if err != nil {
		return *new(main_gateways_mongodb_documents.UserEmailVerificationDocument), err
	}
	return userVerificationEmail, nil
}

func (this *UserEmailVerificationRepository) FindById(
	ctx context.Context,
	id string,
	options main_domains.DatabaseOptions,
) (main_gateways_mongodb_documents.UserEmailVerificationDocument, error) {
	span := this.spanGateway.Get(ctx, "UserEmailVerificationRepository-FindById")
	defer span.End()

	collection := this.database.Collection(_USER_VERIFICATION_EMAIL_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.UserEmailVerificationDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.D{{_USER_VERIFICATION_EMAIL_IDX_INDICATOR_MONGO_ID, objectId}}

	fnFindOne := func(
		ctx context.Context,
		options main_domains.DatabaseOptions,
	) error {
		if options == nil {
			return collection.FindOne(span.GetCtx(), filter).Decode(&result)
		}
		return collection.FindOne(options.GetSession(), filter).Decode(&result)
	}

	err2 := fnFindOne(span.GetCtx(), options)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, err2
	}
	return result, nil
}

func (this *UserEmailVerificationRepository) FindByFilter(
	ctx context.Context,
	filter main_domains.FindUserEmailVerificationFilter,
	pageable main_domains.Pageable,
	options main_domains.DatabaseOptions,
) (main_domains.Page, error) {

	span := this.spanGateway.Get(ctx, "UserEmailVerificationRepository-FindByFilter")
	defer span.End()

	collection := this.database.Collection(_USER_VERIFICATION_EMAIL_COLLECTION_NAME)

	if pageable.HasEmptySort() {
		defaultSort := make(map[string]int)
		defaultSort[_USER_VERIFICATION_EMAIL_IDX_INDICATOR_MONGO_ID] = 1
		pageable.SetSort(defaultSort)
	}

	opt := main_gateways_mongodb_utils.NewPageableUtils().GetOptsFromPageable(pageable)

	filterCriterias := bson.D{}

	if len(filter.Ids) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _USER_VERIFICATION_EMAIL_IDX_INDICATOR_MONGO_ID, Value: bson.M{"$in": filter.Ids}})
	}
	if len(filter.UserIds) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._USER_ID, Value: bson.M{"$in": filter.UserIds}})
	}
	if len(filter.VerificationCodes) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._VERIFICATION_CODE, Value: bson.M{"$in": filter.VerificationCodes}})
	}
	if len(filter.Statuses) > 0 {
		status := make([]string, 0)
		for _, v := range filter.Statuses {
			status = append(status, v.Name())
		}
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._STATUS, Value: bson.M{"$in": status}})
	}
	if !filter.StartCreatedDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._CREATED_DATE, Value: bson.M{"$gte": filter.StartCreatedDate}})
	}
	if !filter.EndCreatedDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._CREATED_DATE, Value: bson.M{"$lt": filter.EndCreatedDate}})
	}
	if !filter.StartLastModifiedDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._LAST_MODIFIED_DATE, Value: bson.M{"$gte": filter.StartLastModifiedDate}})
	}
	if !filter.EndLastModifiedDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._LAST_MODIFIED_DATE, Value: bson.M{"$lt": filter.EndLastModifiedDate}})
	}

	var results []main_gateways_mongodb_documents.UserEmailVerificationDocument

	fnFind := func(
		options main_domains.DatabaseOptions,
	) (*mongo.Cursor, error) {
		if options == nil {
			return collection.Find(context.TODO(), filterCriterias, opt)
		}
		return collection.Find(options.GetSession(), filterCriterias, opt)
	}

	cursor, err := fnFind(options)
	if err != nil {
		return *new(main_domains.Page), err
	}
	if err = cursor.All(span.GetCtx(), &results); err != nil {
		return *new(main_domains.Page), err
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}

	fnCount := func(
		options main_domains.DatabaseOptions,
	) (int64, error) {
		if options == nil {
			return collection.CountDocuments(context.TODO(), filterCriterias)
		}
		return collection.CountDocuments(options.GetSession(), filterCriterias)
	}

	numberDocs, err := fnCount(options)
	if err != nil {
		return *new(main_domains.Page), err
	}

	var contents []any
	for _, value := range results {
		contents = append(contents, value)
	}

	return *main_domains.NewPage(contents, pageable.GetPage(), pageable.GetSize(), numberDocs), nil
}
