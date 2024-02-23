package main_configs_yml

type YmlToken string

const (
	ApplicationPort                               YmlToken = "ApplicationPort"
	ApplicationLogLevel                           YmlToken = "ApplicationLogLevel"
	MongoDBURI                                    YmlToken = "MongoDBURI"
	MongoDBDatabaseName                           YmlToken = "MongoDBDatabaseName"
	RabbitMQURI                                   YmlToken = "RabbitMQURI"
	IntegrationAuthProviderUrl                    YmlToken = "IntegrationAuthProviderUrl"
	IntegrationAuthProviderTimeoutInMilliseconds  YmlToken = "IntegrationAuthProviderTimeoutInMilliseconds"
	IntegrationAuthProviderClientId               YmlToken = "IntegrationAuthProviderClientId"
	IntegrationAuthProviderClientSecret           YmlToken = "IntegrationAuthProviderClientSecret"
	IntegrationAuthProviderRealm                  YmlToken = "IntegrationAuthProviderRealm"
	IntegrationAuthProviderClientIdTokenExchange  YmlToken = "IntegrationAuthProviderClientIdTokenExchange"
	IntegrationAuthProviderTokenPublicSecret      YmlToken = "IntegrationAuthProviderTokenPublicSecret"
	RedisHosts                                    YmlToken = "RedisHosts"
	ApmServerName                                 YmlToken = "ApmServerName"
	ApmServerOtlpCollectorGrpcHost                YmlToken = "ApmServerOtlpCollectorGrpcHost"
	ApmServerLokiCollectorHttpHost                YmlToken = "ApmServerLokiCollectorHttpHost"
	ApmServerLokiCollectorHttpTimeoutMilliseconds YmlToken = "ApmServerLokiCollectorHttpTimeoutMilliseconds"
	EmailGmailCredentialsEmail                    YmlToken = "EmailGmailCredentialsEmail"
	EmailGmailCredentialsPassword                 YmlToken = "EmailGmailCredentialsPassword"
)

var ymlTokenEnum = map[YmlToken]YmlToken{
	ApplicationPort:            ApplicationPort,
	ApplicationLogLevel:        ApplicationLogLevel,
	MongoDBURI:                 MongoDBURI,
	MongoDBDatabaseName:        MongoDBDatabaseName,
	RabbitMQURI:                RabbitMQURI,
	IntegrationAuthProviderUrl: IntegrationAuthProviderUrl,
	IntegrationAuthProviderTimeoutInMilliseconds:  IntegrationAuthProviderTimeoutInMilliseconds,
	IntegrationAuthProviderClientId:               IntegrationAuthProviderClientId,
	IntegrationAuthProviderClientSecret:           IntegrationAuthProviderClientSecret,
	IntegrationAuthProviderRealm:                  IntegrationAuthProviderRealm,
	IntegrationAuthProviderClientIdTokenExchange:  IntegrationAuthProviderClientIdTokenExchange,
	IntegrationAuthProviderTokenPublicSecret:      IntegrationAuthProviderTokenPublicSecret,
	RedisHosts:                                    RedisHosts,
	ApmServerName:                                 ApmServerName,
	ApmServerOtlpCollectorGrpcHost:                ApmServerOtlpCollectorGrpcHost,
	ApmServerLokiCollectorHttpHost:                ApmServerLokiCollectorHttpHost,
	ApmServerLokiCollectorHttpTimeoutMilliseconds: ApmServerLokiCollectorHttpTimeoutMilliseconds,
	EmailGmailCredentialsEmail:                    EmailGmailCredentialsEmail,
	EmailGmailCredentialsPassword:                 EmailGmailCredentialsPassword,
}

var ymlTokenEnumFromNames = map[string]YmlToken{
	"ApplicationPort":            ApplicationPort,
	"ApplicationLogLevel":        ApplicationLogLevel,
	"MongoDBURI":                 MongoDBURI,
	"MongoDBDatabaseName":        MongoDBDatabaseName,
	"RabbitMQURI":                RabbitMQURI,
	"IntegrationAuthProviderUrl": IntegrationAuthProviderUrl,
	"IntegrationAuthProviderTimeoutInMilliseconds": IntegrationAuthProviderTimeoutInMilliseconds,
	"IntegrationAuthProviderClientId":              IntegrationAuthProviderClientId,
	"IntegrationAuthProviderClientSecret":          IntegrationAuthProviderClientSecret,
	"IntegrationAuthProviderRealm":                 IntegrationAuthProviderRealm,
	"IntegrationAuthProviderClientIdTokenExchange": IntegrationAuthProviderClientIdTokenExchange,
	"IntegrationAuthProviderTokenPublicSecret":     IntegrationAuthProviderTokenPublicSecret,
	"RedisHosts":                                    RedisHosts,
	"ApmServerName":                                 ApmServerName,
	"ApmServerOtlpCollectorGrpcHost":                ApmServerOtlpCollectorGrpcHost,
	"ApmServerLokiCollectorHttpHost":                ApmServerLokiCollectorHttpHost,
	"ApmServerLokiCollectorHttpTimeoutMilliseconds": ApmServerLokiCollectorHttpTimeoutMilliseconds,
	"EmailGmailCredentialsEmail":                    EmailGmailCredentialsEmail,
	"EmailGmailCredentialsPassword":                 EmailGmailCredentialsPassword,
}

var ymlTokenDescriptionEnum = map[YmlToken]string{
	ApplicationPort:            "Application.Port",
	ApplicationLogLevel:        "Application.log.level",
	MongoDBURI:                 "MongoDB.URI",
	MongoDBDatabaseName:        "MongoDB.DatabaseName",
	RabbitMQURI:                "RabbitMQ.URI",
	IntegrationAuthProviderUrl: "Integration.authProvider.url",
	IntegrationAuthProviderTimeoutInMilliseconds:  "Integration.authProvider.timeout-in-milliseconds",
	IntegrationAuthProviderClientId:               "Integration.authProvider.client.id",
	IntegrationAuthProviderClientSecret:           "Integration.authProvider.client.secret",
	IntegrationAuthProviderRealm:                  "Integration.authProvider.realm",
	IntegrationAuthProviderClientIdTokenExchange:  "Integration.authProvider.client.id.token.exchange",
	IntegrationAuthProviderTokenPublicSecret:      "Integration.authProvider.token.public.secret",
	RedisHosts:                                    "Redis.hosts",
	ApmServerName:                                 "Apm.server.name",
	ApmServerOtlpCollectorGrpcHost:                "Apm.server.otlp.collector.grpc.host",
	ApmServerLokiCollectorHttpHost:                "Apm.server.loki.collector.http.host",
	ApmServerLokiCollectorHttpTimeoutMilliseconds: "Apm.server.loki.collector.http.timeout.milliseconds",
	EmailGmailCredentialsEmail:                    "Email.gmail.credentials.email",
	EmailGmailCredentialsPassword:                 "Email.gmail.credentials.password",
}

func (this YmlToken) Exists() bool {
	_, exists := ymlTokenEnum[this]
	return exists
}

func (this YmlToken) FromValue(value string) YmlToken {
	valueMap, exists := ymlTokenEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this YmlToken) Name() string {
	valueMap, exists := ymlTokenEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}

func (this YmlToken) GetToken() string {
	valueMap, exists := ymlTokenDescriptionEnum[this]
	if exists {
		return valueMap
	}
	return ""
}

func (this YmlToken) Values() []YmlToken {
	values := make([]YmlToken, 0)
	for _, v := range ymlTokenEnum {
		values = append(values, v)
	}
	return values
}
