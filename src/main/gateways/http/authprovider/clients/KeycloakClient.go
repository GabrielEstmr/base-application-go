package main_gateways_http_authprovider_clients

import (
	main_configs_authprovider "baseapplicationgo/main/configs/authprovider"
	main_configs_authprovider_resource "baseapplicationgo/main/configs/authprovider/resource"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_http_authprovider_resources_request "baseapplicationgo/main/gateways/http/authprovider/resources/request"
	gateways_authprovider_resources "baseapplicationgo/main/gateways/http/commons/response"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_utils "baseapplicationgo/main/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type KeycloakClient struct {
	config                main_configs_authprovider_resource.AuthProviderConfig
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewKeycloakClient() *KeycloakClient {
	return &KeycloakClient{
		config:                *main_configs_authprovider.GetAuthProviderBean(),
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this KeycloakClient) CreateUser(ctx context.Context, userRequest main_gateways_http_authprovider_resources_request.CreateUserRequest,
) (
	gateways_authprovider_resources.HttpResponse, error) {

	span := this.spanGateway.Get(ctx, "KeycloakClient-CreateUser")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("KeycloakClient. CreateUser userEmail %s", userRequest.Email))

	config := clientcredentials.Config{
		ClientID:     this.config.GetClientId(),
		ClientSecret: this.config.GetClientSecret(),
		TokenURL:     this.config.GetBaseUrl() + "/realms/" + this.config.GetRealm() + "/protocol/openid-connect/token",
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	bytesJSON, err := json.Marshal(userRequest)
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	uri := "/admin/realms/" + this.config.GetRealm() + "/users"
	req, err := http.NewRequest(
		http.MethodPost,
		this.config.GetBaseUrl()+uri, bytes.NewBuffer(bytesJSON))
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := this.config.GetClient().Do(req)
	defer resp.Body.Close()
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	var id string
	if resp.StatusCode == http.StatusCreated {
		location, err := resp.Location()
		if err != nil && resp.StatusCode == http.StatusCreated {
			return *new(gateways_authprovider_resources.HttpResponse), err
		}
		id = strings.TrimPrefix(location.Path, uri+"/")
	} else {
		id = ""
	}
	return main_utils.NewHttpExternalClientUtils().BuildHttpResponseWithBody(resp, id)
}

func (this KeycloakClient) CreateOauthExchangeSession(ctx context.Context, token string, provider string) (
	gateways_authprovider_resources.HttpResponse, error) {

	span := this.spanGateway.Get(ctx, "KeycloakClient-CreateOauthExchangeSession")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("KeycloakClient. CreateOauthExchangeSession provider %s", provider))

	q := url.Values{}
	q.Add("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange")
	q.Add("subject_token_type", "urn:ietf:params:oauth:token-type:access_token")
	q.Add("client_id", this.config.GetTokenExchangeClientId())
	q.Add("subject_token", token)
	q.Add("subject_issuer", provider)

	resp, errP := this.config.GetClient().PostForm(this.config.GetBaseUrl()+
		"/realms/"+this.config.GetRealm()+"/protocol/openid-connect/token",
		q)
	defer resp.Body.Close()

	if errP != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errP
	}

	bodyText, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errR
	}
	return main_utils.NewHttpExternalClientUtils().BuildHttpResponseWithBodyBytes(resp, bodyText)
}

func (this KeycloakClient) CreateSession(ctx context.Context, username string, password string) (
	gateways_authprovider_resources.HttpResponse, error) {

	span := this.spanGateway.Get(ctx, "KeycloakClient-CreateSession")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("KeycloakClient. CreateSession username %s", username))

	q := url.Values{}
	q.Add("client_id", this.config.GetTokenExchangeClientId())
	q.Add("client_secret", this.config.GetClientSecret())
	q.Add("grant_type", "password")
	q.Add("username", username)
	q.Add("password", password)

	resp, errP := this.config.GetClient().PostForm(this.config.GetBaseUrl()+
		"/realms/"+this.config.GetRealm()+"/protocol/openid-connect/token",
		q)
	defer resp.Body.Close()

	if errP != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errP
	}

	bodyText, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errR
	}
	return main_utils.NewHttpExternalClientUtils().BuildHttpResponseWithBodyBytes(resp, bodyText)
}

func (this KeycloakClient) RefreshSession(ctx context.Context, refreshToken string) (
	gateways_authprovider_resources.HttpResponse, error) {

	span := this.spanGateway.Get(ctx, "KeycloakClient-RefreshSession")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "KeycloakClient. RefreshSession username")

	q := url.Values{}
	q.Add("client_id", this.config.GetTokenExchangeClientId())
	q.Add("client_secret", this.config.GetClientSecret())
	q.Add("grant_type", "refresh_token")
	q.Add("refresh_token", refreshToken)

	resp, errP := this.config.GetClient().PostForm(this.config.GetBaseUrl()+
		"/realms/"+this.config.GetRealm()+"/protocol/openid-connect/token",
		q)
	defer resp.Body.Close()

	if errP != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errP
	}

	bodyText, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errR
	}
	return main_utils.NewHttpExternalClientUtils().BuildHttpResponseWithBodyBytes(resp, bodyText)
}

func (this KeycloakClient) EndSession(ctx context.Context, refreshToken string) (
	gateways_authprovider_resources.HttpResponse, error) {

	span := this.spanGateway.Get(ctx, "KeycloakClient-EndSession")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "KeycloakClient. EndSession")

	q := url.Values{}
	q.Add("client_id", this.config.GetTokenExchangeClientId())
	q.Add("client_secret", this.config.GetClientSecret())
	q.Add("grant_type", "refresh_token")
	q.Add("refresh_token", refreshToken)

	resp, errP := this.config.GetClient().PostForm(this.config.GetBaseUrl()+
		"/realms/"+this.config.GetRealm()+"/protocol/openid-connect/token",
		q)
	defer resp.Body.Close()

	if errP != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errP
	}

	bodyText, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errR
	}
	return main_utils.NewHttpExternalClientUtils().BuildHttpResponseWithBodyBytes(resp, bodyText)
}

func (this KeycloakClient) UpdateUser(ctx context.Context, id string, userRequest any) (gateways_authprovider_resources.HttpResponse, error) {

	span := this.spanGateway.Get(ctx, "KeycloakClient-UpdateUser")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("KeycloakClient. UpdateUser id %s", id))

	config := clientcredentials.Config{
		ClientID:     this.config.GetClientId(),
		ClientSecret: this.config.GetClientSecret(),
		TokenURL:     this.config.GetBaseUrl() + "/realms/" + this.config.GetRealm() + "/protocol/openid-connect/token",
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	bytesJSON, err := json.Marshal(userRequest)
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	uri := "/admin/realms/" + this.config.GetRealm() + "/users/" + id
	req, err := http.NewRequest(
		http.MethodPut,
		this.config.GetBaseUrl()+uri, bytes.NewBuffer(bytesJSON))
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := this.config.GetClient().Do(req)
	defer resp.Body.Close()
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}
	return main_utils.NewHttpExternalClientUtils().BuildHttpResponse(resp)
}

func (this KeycloakClient) GetUserByUserName(ctx context.Context, username string) (gateways_authprovider_resources.HttpResponse, error) {

	span := this.spanGateway.Get(ctx, "KeycloakClient-GetUserByUserName")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("KeycloakClient. GetUserByUserName username %s", username))

	config := clientcredentials.Config{
		ClientID:     this.config.GetClientId(),
		ClientSecret: this.config.GetClientSecret(),
		TokenURL:     this.config.GetBaseUrl() + "/realms/" + this.config.GetRealm() + "/protocol/openid-connect/token",
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		this.config.GetBaseUrl()+"/admin/realms/"+this.config.GetRealm()+"/users", nil)
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	q := req.URL.Query()
	q.Add("username", username)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err := this.config.GetClient().Do(req)
	defer resp.Body.Close()
	if err != nil {
		return *new(gateways_authprovider_resources.HttpResponse), err
	}

	return main_utils.NewHttpExternalClientUtils().BuildHttpResponse(resp)
}
