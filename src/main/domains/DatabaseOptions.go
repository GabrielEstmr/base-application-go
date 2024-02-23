package main_domains

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseOptions map[string]interface{}

func NewDatabaseOptions() *DatabaseOptions {
	test := make(DatabaseOptions)
	return &test
}

func (this DatabaseOptions) WithSession(sessionCtx mongo.SessionContext) DatabaseOptions {
	this["session"] = sessionCtx
	return this
}

func (this DatabaseOptions) GetPropertyByName(key string) interface{} {
	return this[key]
}

func (this DatabaseOptions) GetSession() mongo.SessionContext {
	return this["session"].(mongo.SessionContext)
}
