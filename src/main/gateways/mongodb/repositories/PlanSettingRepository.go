package main_gateways_mongodb_repositories

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_mongo_collections "baseapplicationgo/main/configs/mongodb/collections"
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
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

const _PLAN_SETTINGS_COLLECTION_NAME = main_configs_mongo_collections.PLAN_SETTINGS_COLLECTION_NAME
const _PLAN_SETTINGS_IDX_INDICATOR_MONGO_ID = main_configs_mongo.DEFAULT_BSON_ID_NAME

type PlanSettingRepository struct {
	_PLAN_SETTINGS_COLLECTION_NAME        string
	_PLAN_SETTINGS_IDX_INDICATOR_MONGO_ID string
	_PLAN_TYPE                            string
	_METADATA                             string
	_CREATION_USER_EMAIL                  string
	_START_DATE                           string
	_END_DATE                             string
	_CREATED_DATE                         string
	_LAST_MODIFIED_DATE                   string
	database                              *mongo.Database
	spanGateway                           main_gateways.SpanGateway
}

func NewPlanSettingRepository() *PlanSettingRepository {
	return &PlanSettingRepository{

		_PLAN_TYPE:           "planType",
		_METADATA:            "metadata",
		_CREATION_USER_EMAIL: "creationUserEmail",
		_START_DATE:          "startDate",
		_END_DATE:            "endDate",
		_CREATED_DATE:        "createdDate",
		_LAST_MODIFIED_DATE:  "lastModifiedDate",

		database:    main_configs_mongo.GetMongoDBDatabaseBean(),
		spanGateway: main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *PlanSettingRepository) Save(
	ctx context.Context,
	planSetting main_gateways_mongodb_documents.PlanSettingDocument,
) (
	main_gateways_mongodb_documents.PlanSettingDocument, error,
) {
	span := this.spanGateway.Get(ctx, "PlanSettingRepository-Save")
	defer span.End()

	collection := this.database.Collection(_PLAN_SETTINGS_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	planSetting.CreatedDate = now
	planSetting.LastModifiedDate = now

	result, err := collection.InsertOne(context.TODO(), planSetting)
	if err != nil {
		return *new(main_gateways_mongodb_documents.PlanSettingDocument), err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	planSetting.Id = oid
	return planSetting, nil
}

func (this *PlanSettingRepository) FindById(ctx context.Context, id string) (main_gateways_mongodb_documents.PlanSettingDocument, error) {
	span := this.spanGateway.Get(ctx, "PlanSettingRepository-FindById")
	defer span.End()

	collection := this.database.Collection(_PLAN_SETTINGS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.PlanSettingDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.D{{_PLAN_SETTINGS_IDX_INDICATOR_MONGO_ID, objectId}}
	err2 := collection.FindOne(span.GetCtx(), filter).Decode(result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, err2
	}
	return result, nil
}

func (this *PlanSettingRepository) FindByPlanTypeAndHasEndDate(
	ctx context.Context,
	planType main_domains_enums.PlanType, hasEndDate bool,
) (main_gateways_mongodb_documents.PlanSettingDocument, error) {
	span := this.spanGateway.Get(ctx, "PlanSettingRepository-FindByPlanTypeAndHasEndDate")
	defer span.End()

	collection := this.database.Collection(_PLAN_SETTINGS_COLLECTION_NAME)
	var result main_gateways_mongodb_documents.PlanSettingDocument
	filter := bson.D{
		{this._PLAN_TYPE, planType},
		{this._END_DATE, bson.M{"$lt": hasEndDate}},
	}
	err2 := collection.FindOne(span.GetCtx(), filter).Decode(result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return result, nil
		}
		return result, err2
	}
	return result, nil
}

func (this *PlanSettingRepository) Update(
	ctx context.Context,
	planSetting main_gateways_mongodb_documents.PlanSettingDocument,
) (
	main_gateways_mongodb_documents.PlanSettingDocument, error,
) {
	span := this.spanGateway.Get(ctx, "PlanSettingRepository-Update")
	defer span.End()

	collection := this.database.Collection(_PLAN_SETTINGS_COLLECTION_NAME)
	now := primitive.NewDateTimeFromTime(time.Now())
	planSetting.LastModifiedDate = now

	filter := bson.D{{_PLAN_SETTINGS_IDX_INDICATOR_MONGO_ID, planSetting.Id}}
	update := bson.D{{"$set", bson.D{
		{this._PLAN_TYPE, planSetting.PlanType},
		{this._METADATA, planSetting.Metadata},
		{this._CREATION_USER_EMAIL, planSetting.CreationUserEmail},
		{this._START_DATE, planSetting.StartDate},
		{this._END_DATE, planSetting.EndDate},
		{this._LAST_MODIFIED_DATE, planSetting.LastModifiedDate},
	}}}

	_, err := collection.UpdateOne(span.GetCtx(), filter, update)
	if err != nil {
		return *new(main_gateways_mongodb_documents.PlanSettingDocument), err
	}
	return planSetting, nil
}

func (this *PlanSettingRepository) FindByFilter(
	ctx context.Context, filter main_domains.FindPlanSettingFilter,
	pageable main_domains.Pageable,
) (main_domains.Page, error) {
	span := this.spanGateway.Get(ctx, "PlanSettingRepository-FindByFilter")
	defer span.End()

	collection := this.database.Collection(_PLAN_SETTINGS_COLLECTION_NAME)

	log.Println(len(pageable.GetSort()))
	if pageable.HasEmptySort() {
		defaultSort := make(map[string]int)
		defaultSort[_PLAN_SETTINGS_IDX_INDICATOR_MONGO_ID] = 1
		pageable.SetSort(defaultSort)
	}

	opt := main_gateways_mongodb_utils.NewPageableUtils().GetOptsFromPageable(pageable)

	filterCriterias := bson.D{}

	if len(filter.Ids) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _PLAN_SETTINGS_IDX_INDICATOR_MONGO_ID, Value: bson.M{"$in": filter.Ids}})
	}
	if len(filter.PlanTypes) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: _PLAN_SETTINGS_IDX_INDICATOR_MONGO_ID, Value: bson.M{"$in": filter.PlanTypes}})
	}
	if len(filter.CreationUserEmails) > 0 {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._CREATION_USER_EMAIL, Value: bson.M{"$in": filter.CreationUserEmails}})
	}
	if !filter.StartStartDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._START_DATE, Value: bson.M{"$gte": filter.StartStartDate}})
	}
	if !filter.EndStartDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._START_DATE, Value: bson.M{"$lt": filter.EndStartDate}})
	}
	if !filter.StartEndDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._END_DATE, Value: bson.M{"$gte": filter.StartEndDate}})
	}
	if !filter.EndEndDate.IsZero() {
		filterCriterias = append(filterCriterias,
			bson.E{Key: this._END_DATE, Value: bson.M{"$lt": filter.EndEndDate}})
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
	if len(filter.HasEndDate) > 0 {
		for _, v := range filter.HasEndDate {
			filterCriterias = append(filterCriterias,
				bson.E{Key: this._END_DATE, Value: bson.M{"$exists": v}})
		}
	}

	var results []main_gateways_mongodb_documents.PlanSettingDocument
	cursor, err := collection.Find(context.TODO(), filterCriterias, opt)
	if err = cursor.All(span.GetCtx(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}

	numberDocs, err := collection.CountDocuments(span.GetCtx(), filterCriterias)
	if err != nil {
		return *new(main_domains.Page), err
	}

	var contents []any
	for _, value := range results {
		contents = append(contents, value)
	}

	return *main_domains.NewPage(contents, pageable.GetPage(), pageable.GetSize(), numberDocs), nil
}
