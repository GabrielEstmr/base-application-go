package main_usecases

import (
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type BuildTokenClaim struct {
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewBuildTokenClaim(
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *BuildTokenClaim {
	return &BuildTokenClaim{
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageBeans,
	}
}

func (this *BuildTokenClaim) Execute(ctx context.Context, accessToken string) (main_domains.TokenClaims, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "BuildTokenClaim-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "BuildTokenClaim")

	CLIENT_PUBLIC_SECRET := main_configs_yml.GetYmlValueByName(main_configs_yml.IntegrationAuthProviderTokenPublicSecret)

	public_key := "-----BEGIN PUBLIC KEY-----\n" + CLIENT_PUBLIC_SECRET + "\n-----END PUBLIC KEY-----"
	pkey, errP := jwt.ParseRSAPublicKeyFromPEM([]byte(public_key))
	if errP != nil {
		return *new(main_domains.TokenClaims), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errP.Error())
	}

	tokenP, errParse := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return pkey, nil
	})
	if errParse != nil {
		return *new(main_domains.TokenClaims), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errParse.Error())
	}

	test := tokenP.Claims.(jwt.MapClaims)
	var result map[string]interface{} = test
	testS := new(main_domains.TokenClaims).FromMap(result)

	return testS, nil
}
