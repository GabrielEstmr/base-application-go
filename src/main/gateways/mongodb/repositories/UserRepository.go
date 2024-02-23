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
	"log"
	"time"
)

const _USERS_COLLECTION_NAME = main_configs_mongo_collections.USERS_COLLECTION_NAME
const _USERS_IDX_INDICATOR_MONGO_ID = main_configs_mongo.DEFAULT_BSON_ID_NAME

type UserRepository struct {
	_ACCOUNT_ID         string
	_AUTH_PROVIDER_ID   string
	_DOCUMENT_ID        string
	_USER_NAME          string
	_FIRST_NAME         string
	_LAST_NAME          string
	_EMAIL              string
	_EMAIL_VERIFIED     string
	_STATUS             string
	_ROLES              string
	_BIRTHDAY           string
	_PHONE_CONTACTS     string
	_PROVIDER_TYPE      string
	_CREATED_DATE       string
	_LAST_MODIFIED_DATE string
	database            *mongo.Database
	spanGateway         main_gateways.SpanGateway
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		_ACCOUNT_ID:         "accountId",
		_AUTH_PROVIDER_ID:   "authProviderId",
		_DOCUMENT_ID:        "documentId",
		_USER_NAME:          "userName",
		_FIRST_NAME:         "firstName",
		_LAST_NAME:          "lastName",
		_EMAIL:              "email",
		_EMAIL_VERIFIED:     "emailVerified",
		_STATUS:             "status",
		_ROLES:              "roles",
		_BIRTHDAY:           "birthday",
		_PHONE_CONTACTS:     "phoneContacts",
		_PROVIDER_TYPE:      "providerType",
		_CREATED_DATE:       "createdDate",
		_LAST_MODIFIED_DATE: "lastModifiedDate",
		database:            main_configs_mongo.GetMongoDBDatabaseBean(),
		spanGateway:         main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *UserRepository) Save(
	ctx context.Context,
	user main_gateways_mongodb_documents.UserDocument,
	options main_domains.DatabaseOptions,
) (main_gateways_mongodb_documents.UserDocument, error) {

	span := this.spanGateway.Get(ctx, "UserRepository-Save")
	defer span.End()

	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	user.CreatedDate = now
	user.LastModifiedDate = now

	optTest := options.GetPropertyByName("session").(mongo.SessionContext)
	fmt.Println(optTest)

	fnInsert := func(
		ctx context.Context,
		user main_gateways_mongodb_documents.UserDocument,
		options main_domains.DatabaseOptions,
	) (*mongo.InsertOneResult, error) {
		if options == nil {
			return collection.InsertOne(ctx, user)
		} else {
			return collection.InsertOne(options.GetSession(), user)
		}
	}

	result, err := fnInsert(ctx, user, options)
	if err != nil {
		return *new(main_gateways_mongodb_documents.UserDocument), err
	}
	oid, _ := result.InsertedID.(primitive.ObjectID)
	user.Id = oid
	return user, nil
}

func (this *UserRepository) Update(
	ctx context.Context,
	user main_gateways_mongodb_documents.UserDocument,
	options main_domains.DatabaseOptions,
) (
	main_gateways_mongodb_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRepository-Update")
	defer span.End()

	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	user.LastModifiedDate = now

	filter := bson.D{{_EMAILS_IDX_INDICATOR_MONGO_ID, user.Id}}
	update := bson.D{{"$set", bson.D{
		{this._ACCOUNT_ID, user.AccountId},
		{this._AUTH_PROVIDER_ID, user.AuthProviderId},
		{this._DOCUMENT_ID, user.DocumentId},
		{this._USER_NAME, user.UserName},
		{this._FIRST_NAME, user.FirstName},
		{this._LAST_NAME, user.LastName},
		{this._EMAIL, user.Email},
		{this._EMAIL_VERIFIED, user.EmailVerified},
		{this._STATUS, user.Status},
		{this._ROLES, user.Roles},
		{this._BIRTHDAY, user.Birthday},
		{this._PHONE_CONTACTS, user.PhoneContacts},
		{this._PROVIDER_TYPE, user.ProviderType},
		{this._CREATED_DATE, user.CreatedDate},
		{this._LAST_MODIFIED_DATE, user.LastModifiedDate},
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
		return *new(main_gateways_mongodb_documents.UserDocument), err
	}
	return user, nil
}

func (this *UserRepository) FindById(
	ctx context.Context,
	id string,
	options main_domains.DatabaseOptions,
) (main_gateways_mongodb_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRepository-FindById")
	defer span.End()

	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.UserDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.D{{_USERS_IDX_INDICATOR_MONGO_ID, objectId}}
	fnFindOne := func(
		ctx context.Context,
		options main_domains.DatabaseOptions,
	) error {
		if options == nil {
			return collection.FindOne(span.GetCtx(), filter).Decode(&result)
		}
		return collection.FindOne(options.GetSession(), filter).Decode(&result)
	}

	errFind := fnFindOne(span.GetCtx(), options)
	if errFind != nil {
		if errors.Is(errFind, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, errFind
	}
	return result, nil
}

func (this *UserRepository) FindByDocumentId(
	ctx context.Context,
	documentId string,
	options main_domains.DatabaseOptions,
) (main_gateways_mongodb_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRepository-FindByDocumentNumber")
	defer span.End()

	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.UserDocument
	filter := bson.D{{this._DOCUMENT_ID, documentId}}
	fnFindOne := func(
		ctx context.Context,
		options main_domains.DatabaseOptions,
	) error {
		if options == nil {
			return collection.FindOne(span.GetCtx(), filter).Decode(&result)
		}
		return collection.FindOne(options.GetPropertyByName("session").(mongo.SessionContext), filter).Decode(&result)
	}

	errFind := fnFindOne(span.GetCtx(), options)
	if errFind != nil {
		if errors.Is(errFind, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, errFind
	}
	return result, nil
}

func (this *UserRepository) FindByUserName(
	ctx context.Context,
	userName string,
	options main_domains.DatabaseOptions,
) (main_gateways_mongodb_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRepository-FindByUserName")
	defer span.End()

	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.UserDocument
	filter := bson.D{{this._USER_NAME, userName}}
	fnFindOne := func(
		ctx context.Context,
		options main_domains.DatabaseOptions,
	) error {
		if options == nil {
			return collection.FindOne(span.GetCtx(), filter).Decode(&result)
		}
		return collection.FindOne(options.GetPropertyByName("session").(mongo.SessionContext), filter).Decode(&result)
	}

	errFind := fnFindOne(span.GetCtx(), options)
	if errFind != nil {
		if errors.Is(errFind, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, errFind
	}
	return result, nil
}

func (this *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
	options main_domains.DatabaseOptions,
) (main_gateways_mongodb_documents.UserDocument, error) {
	span := this.spanGateway.Get(ctx, "UserRepository-FindByEmail")
	defer span.End()

	collection := this.database.Collection(_USERS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.UserDocument
	filter := bson.D{{this._EMAIL, email}}
	fnFindOne := func(
		ctx context.Context,
		options main_domains.DatabaseOptions,
	) error {
		if options == nil {
			return collection.FindOne(span.GetCtx(), filter).Decode(&result)
		}
		return collection.FindOne(options.GetPropertyByName("session").(mongo.SessionContext), filter).Decode(&result)
	}

	errFind := fnFindOne(span.GetCtx(), options)
	if errFind != nil {
		if errors.Is(errFind, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, errFind
	}
	return result, nil
}

func (this *UserRepository) FindByFilter(
	ctx context.Context,
	filter main_domains.FindUserFilter,
	pageable main_domains.Pageable,
	options main_domains.DatabaseOptions,
) (main_domains.Page, error) {
	span := this.spanGateway.Get(ctx, "UserRepository-FindByFilter")
	defer span.End()

	collection := this.database.Collection(_USERS_COLLECTION_NAME)

	log.Println(len(pageable.GetSort()))
	if pageable.HasEmptySort() {
		defaultSort := make(map[string]int)
		defaultSort[_USERS_IDX_INDICATOR_MONGO_ID] = 1
		pageable.SetSort(defaultSort)
	}

	opt := main_gateways_mongodb_utils.NewPageableUtils().GetOptsFromPageable(pageable)

	filterCriterias := bson.D{}

	if len(filter.Ids) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _USERS_IDX_INDICATOR_MONGO_ID, Value: bson.M{"$in": filter.Ids}})
	}
	if len(filter.AccountIds) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._ACCOUNT_ID, Value: bson.M{"$in": filter.AccountIds}})
	}
	if len(filter.AuthProviderIds) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._AUTH_PROVIDER_ID, Value: bson.M{"$in": filter.AuthProviderIds}})
	}
	if len(filter.DocumentIds) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._DOCUMENT_ID, Value: bson.M{"$in": filter.DocumentIds}})
	}
	if len(filter.UserNames) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._USER_NAME, Value: bson.M{"$in": filter.UserNames}})
	}
	if len(filter.FirstNames) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._FIRST_NAME, Value: bson.M{"$in": filter.FirstNames}})
	}
	if len(filter.LastNames) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._LAST_NAME, Value: bson.M{"$in": filter.LastNames}})
	}
	if len(filter.Emails) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._EMAIL, Value: bson.M{"$in": filter.Emails}})
	}
	if len(filter.EmailsVerified) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._EMAIL_VERIFIED, Value: bson.M{"$in": filter.EmailsVerified}})
	}
	if len(filter.Statuses) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._STATUS, Value: bson.M{"$in": filter.Statuses}})
	}
	if len(filter.Roles) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._ROLES, Value: bson.M{"$in": filter.Roles}})
	}
	if len(filter.ProviderTypes) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._PROVIDER_TYPE, Value: bson.M{"$in": filter.ProviderTypes}})
	}
	if !filter.StartBirthdayDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._BIRTHDAY, Value: bson.M{"$gte": filter.StartBirthdayDate}})
	}
	if !filter.EndBirthdayDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._BIRTHDAY, Value: bson.M{"$lt": filter.EndBirthdayDate}})
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

	fnFind := func(
		options main_domains.DatabaseOptions,
	) (*mongo.Cursor, error) {
		if options == nil {
			return collection.Find(context.TODO(), filterCriterias, opt)
		}
		return collection.Find(options.GetSession(), filterCriterias, opt)
	}

	var results []main_gateways_mongodb_documents.UserDocument
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
