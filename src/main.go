package main

import (
	main_configs "baseapplicationgo/main/configs"
	main_configs_apm "baseapplicationgo/main/configs/apm"
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_router "baseapplicationgo/main/configs/router"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_gateways_rabbitmq_subscribers "baseapplicationgo/main/gateways/rabbitmq/subscribers"
	main_gateways_ws "baseapplicationgo/main/gateways/ws"
	main_gateways_ws_beans "baseapplicationgo/main/gateways/ws/beans"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"log/slog"
	"net/http"
)

const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"

func init() {
	main_configs.InitConfigBeans()
	go main_gateways_rabbitmq_subscribers.SubscribeListeners()
}

// TODO: verify beans Nao pode ter bean que usam o mesmo input pois mesmos apontamentos em memoria irão ser usados em classes diferentes
// OLHAR CreateNewUserBean como exemplo
func main() {

	ctx := context.Background()
	main_configs_apm.InitiateApmConfig(&ctx)
	defer main_configs.TerminateConfigBeans(&ctx)
	applicationPort := main_configs_yml.GetYmlValueByName(main_configs_yml.ApplicationPort)
	routes := main_gateways_ws_beans.GetRoutesBean()
	router := main_gateways_ws.ConfigRoutes(main_configs_router.GetRouterBean(), *routes)
	router.Handle("/metrics", promhttp.Handler())
	log.Printf(MSG_LISTENER, applicationPort)

	err2 := http.ListenAndServe(":"+applicationPort, router)
	if err2 != nil {
		main_configs_error.FailOnError(err2, MSG_APPLICATION_FAILED)
	}
}

func test1() {
	wc := writeconcern.Majority()
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	slog.Info("_MSG_MONGO_BEAN_INITIALIZING")
	databaseUri := main_configs_yml.GetYmlValueByName(main_configs_yml.MongoDBURI)
	client, err0 := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))
	if err0 != nil {
		log.Fatalf(err0.Error())
	}

	episodesCollection := main_configs_mongo.GetMongoDBDatabaseBean().Collection("COLLECTION_NAME")

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			return err
		}
		result, err := episodesCollection.InsertOne(
			sessionContext,
			"Document",
		)
		if err != nil {
			return err
		}
		fmt.Println(result.InsertedID)
		result, err = episodesCollection.InsertOne(
			sessionContext,
			"Document",
		)
		if err != nil {
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			return err
		}
		fmt.Println(result.InsertedID)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
}

//func test() {
//
//	wc := writeconcern.Majority()
//
//	rc := readconcern.Snapshot()
//	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
//
//	slog.Info("_MSG_MONGO_BEAN_INITIALIZING")
//	databaseUri := main_configs_yml.GetYmlValueByName(main_configs_yml.MongoDBURI)
//	client, err0 := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))
//	if err0 != nil {
//		log.Fatalf(err0.Error())
//	}
//
//	episodesCollection := main_configs_mongo.GetMongoDBClient().Collection("_EMAILS_COLLECTION_NAME")
//
//	session, err1 := client.StartSession()
//	if err1 != nil {
//		panic(err1)
//	}
//	defer session.EndSession(context.Background())
//
//	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
//		result, err := episodesCollection.InsertOne(
//			sessionContext,
//			"uahsuahs",
//		)
//		if err != nil {
//			return nil, err
//		}
//		result, err = episodesCollection.InsertOne(
//			sessionContext,
//			"uahsuahs",
//		)
//		if err != nil {
//			return nil, err
//		}
//		return result, err
//	}
//
//	_, errT := session.WithTransaction(context.Background(), callback, txnOpts)
//	if errT != nil {
//		panic(errT)
//	}
//}
//
////
////
////type TestArraysGeneric struct {
////	a []any
////}
////
////func NewTestArraysGeneric(a []string) *TestArraysGeneric {
////	inputs := make([]any, 0)
////	for _, v := range a {
////		inputs = append(inputs, v)
////	}
////	return &TestArraysGeneric{a: inputs}
////}
////
////func toString(a any) string {
////	return a.(string)
////}
////
////type GenericFunctionVoid func(inp ...any)
////
////type GenericFunction func(inp AnyInput) AnyOutput
////
////type AnyInput any
////
////type AnyOutput any
////
////func (this TestArraysGeneric) toAnyInput(inp any) AnyInput {
////	return inp.(AnyInput)
////}
////
////func (this TestArraysGeneric) toAnyOutput(out any) AnyOutput {
////	return out.(AnyOutput)
////}
////
////func (this TestArraysGeneric) toFuncMap(gf func(inp any) any) GenericFunction {
////
////	input := this.toAnyInput(gf)
////	this.toAnyOutput()
////
////	return func gfunc(this.toAnyInput(inp) )
////}
////
////func (this TestArraysGeneric) forEach(gf GenericFunctionVoid) {
////	for _, v := range this.a {
////		gf(v)
////	}
////}
////
////func (this TestArraysGeneric) mapC(gf GenericFunction) []any {
////	result := make([]any, 0)
////	for _, v := range this.a {
////		result = append(result, gf(v))
////	}
////	return result
////}
////
////func (this TestArraysGeneric) mapA(gf func(inp any) any) []any {
////	result := make([]any, 0)
////	for _, v := range this.a {
////		result = append(result, gf(v))
////	}
////	return result
////}
////
////func (this TestArraysGeneric) mapB(gf func(inp string) main_domains_enums.UserStatus) []any {
////	result := make([]any, 0)
////	for _, v := range this.a {
////		result = append(result, gf(v))
////	}
////	return result
////}
