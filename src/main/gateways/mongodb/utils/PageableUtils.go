package main_gateways_mongodb_utils

import (
	main_domains "baseapplicationgo/main/domains"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PageableUtils struct {
}

func NewPageableUtils() *PageableUtils {
	return &PageableUtils{}
}

func (this *PageableUtils) GetOptsFromPageable(pageable main_domains.Pageable) *options.FindOptions {
	return options.Find().SetSkip(pageable.GetPage()).SetLimit(
		pageable.GetSize()).SetSort(
		buildBsonSortFromMap(pageable.GetSort()))
}

func buildBsonSortFromMap(metadata map[string]int) bson.D {
	var tmp bson.D
	for k, v := range metadata {
		tmp = append(tmp, bson.E{
			Key:   k,
			Value: v,
		})

	}
	return tmp
}
