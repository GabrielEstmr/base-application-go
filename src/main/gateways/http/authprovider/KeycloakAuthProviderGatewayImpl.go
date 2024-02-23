package main_gateways_http_authprovider

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_http_authprovider_clients "baseapplicationgo/main/gateways/http/authprovider/clients"
	main_gateways_http_authprovider_resources_request "baseapplicationgo/main/gateways/http/authprovider/resources/request"
	main_gateways_http_authprovider_resources_response "baseapplicationgo/main/gateways/http/authprovider/resources/response"
	main_gateways_http_commons_decoders "baseapplicationgo/main/gateways/http/commons/decoders"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"encoding/json"
	"fmt"
)

type KeycloakAuthProviderGatewayImpl struct {
	_MSG_INVALID_CREDENTIALS string
	defaultHttpErrorDecoder  main_gateways_http_commons_decoders.HttpErrorDecoder
	keycloakClient           main_gateways_http_authprovider_clients.KeycloakClient
	logsMonitoringGateway    main_gateways.LogsMonitoringGateway
	spanGateway              main_gateways.SpanGateway
	messageUtils             main_utils_messages.ApplicationMessages
}

func NewKeycloakAuthProviderGatewayImpl(
	keycloakClient main_gateways_http_authprovider_clients.KeycloakClient,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *KeycloakAuthProviderGatewayImpl {
	return &KeycloakAuthProviderGatewayImpl{
		_MSG_INVALID_CREDENTIALS: "providers.create.session.invalid.credentials",
		defaultHttpErrorDecoder:  main_gateways_http_commons_decoders.NewDefaultHttpErrorDecoder(),
		keycloakClient:           keycloakClient,
		logsMonitoringGateway:    logsMonitoringGateway,
		spanGateway:              spanGateway,
		messageUtils:             messageBeans,
	}
}

func (this *KeycloakAuthProviderGatewayImpl) CreateUser(
	ctx context.Context,
	user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-CreateUser")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new User on authProvider documentNumberId: %s, email: %s, userName: %s",
			user.GetDocumentId(),
			user.GetEmail(),
			user.GetUserName(),
		))

	resp, err := this.keycloakClient.CreateUser(span.GetCtx(), *main_gateways_http_authprovider_resources_request.NewCreateUserRequest(user))
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return *new(main_domains.User), errApp
	}

	fmt.Println(string(resp.GetResponseBody()))
	updatedUser := user.CloneWithNewAuthProviderId(string(resp.GetResponseBody()))

	return updatedUser, nil
}

func (this *KeycloakAuthProviderGatewayImpl) CreateOauthExchangeSession(
	ctx context.Context, args main_domains.ExternalProviderSessionArgs,
) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-CreateOauthExchangeSession")
	defer span.End()

	resp, err := this.keycloakClient.CreateOauthExchangeSession(span.GetCtx(), args.GetToken(), args.GetProvider().GetProviderDescription())
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return *new(main_domains.SessionCredentials), errApp
	}

	var sessionResponse main_gateways_http_authprovider_resources_response.SessionCredentialsResponse
	errJSON := json.Unmarshal(resp.GetResponseBody(), &sessionResponse)
	if errJSON != nil {
		return *new(main_domains.SessionCredentials),
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errJSON.Error())
	}

	return sessionResponse.ToDomain(), nil
}

func (this *KeycloakAuthProviderGatewayImpl) CreateSession(
	ctx context.Context, username string, password string,
) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-CreateSession")
	defer span.End()

	resp, err := this.keycloakClient.CreateSession(span.GetCtx(), username, password)
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return *new(main_domains.SessionCredentials), errApp
	}

	var sessionResponse main_gateways_http_authprovider_resources_response.SessionCredentialsResponse
	errJSON := json.Unmarshal(resp.GetResponseBody(), &sessionResponse)
	if errJSON != nil {
		return *new(main_domains.SessionCredentials),
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errJSON.Error())
	}

	return sessionResponse.ToDomain(), nil
}

func (this *KeycloakAuthProviderGatewayImpl) RefreshSession(
	ctx context.Context, refreshToken string,
) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-CreateSession")
	defer span.End()

	resp, err := this.keycloakClient.RefreshSession(span.GetCtx(), refreshToken)
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return *new(main_domains.SessionCredentials), errApp
	}

	var sessionResponse main_gateways_http_authprovider_resources_response.SessionCredentialsResponse
	errJSON := json.Unmarshal(resp.GetResponseBody(), &sessionResponse)
	if errJSON != nil {
		return *new(main_domains.SessionCredentials),
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errJSON.Error())
	}

	return sessionResponse.ToDomain(), nil
}

func (this *KeycloakAuthProviderGatewayImpl) EndSession(
	ctx context.Context, refreshToken string,
) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-CreateSession")
	defer span.End()

	resp, err := this.keycloakClient.EndSession(span.GetCtx(), refreshToken)
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return errApp
	}

	var sessionResponse main_gateways_http_authprovider_resources_response.SessionCredentialsResponse
	errJSON := json.Unmarshal(resp.GetResponseBody(), &sessionResponse)
	if errJSON != nil {
		return main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errJSON.Error())
	}

	return nil
}

func (this *KeycloakAuthProviderGatewayImpl) ChangeUserStatusAndEmailVerification(
	ctx context.Context,
	user main_domains.User,
	enabled bool,
) (main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-ChangeUserStatusAndEmailVerification")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("ChangeUserStatusAndEmailVerification User by id: %s", user.GetAuthProviderId()))

	resp, err := this.keycloakClient.UpdateUser(span.GetCtx(),
		user.GetAuthProviderId(),
		*main_gateways_http_authprovider_resources_request.NewVerifyUserRequest(user.GetEmailVerified(), enabled))
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return *new(main_domains.User), errApp
	}

	return user, nil
}

func (this *KeycloakAuthProviderGatewayImpl) ChangeUsersPassword(
	ctx context.Context,
	user main_domains.User,
) (main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-ChangeUsersPassword")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("ChangeUsersPassword User by id: %s", user.GetAuthProviderId()))

	resp, err := this.keycloakClient.UpdateUser(span.GetCtx(),
		user.GetAuthProviderId(),
		*main_gateways_http_authprovider_resources_request.NewChangeUserPasswordRequest(user))
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return *new(main_domains.User), errApp
	}

	return user, nil
}

func (this *KeycloakAuthProviderGatewayImpl) GetUsers(
	ctx context.Context,
	email string) ([]main_domains.AuthProviderUser, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "KeycloakAuthProviderGatewayImpl-GetUserByEmail")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("Finding User by email: %s", email))

	resp, err := this.keycloakClient.GetUserByUserName(span.GetCtx(), email)
	errApp := this.defaultHttpErrorDecoder.DecodeErrors(span, resp, err)
	if errApp != nil {
		return *new([]main_domains.AuthProviderUser), errApp
	}

	var userResponse []main_gateways_http_authprovider_resources_response.UserResponse
	errJSON := json.Unmarshal(resp.GetResponseBody(), &userResponse)
	if errJSON != nil {
		return *new([]main_domains.AuthProviderUser), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errJSON.Error())
	}

	users := make([]main_domains.AuthProviderUser, 0)
	for _, v := range userResponse {
		users = append(users, v.ToDomain())
	}

	return users, nil
}
