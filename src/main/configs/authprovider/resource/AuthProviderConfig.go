package main_configs_authprovider_resource

import (
	"net/http"
	"time"
)

type AuthProviderConfig struct {
	client                http.Client
	baseUrl               string
	clientId              string
	clientSecret          string
	realm                 string
	tokenExchangeClientId string
}

func NewAuthProviderConfig(
	baseUrl string,
	timeoutMilliseconds int64,
	clientId string,
	clientSecret string,
	realm string,
	tokenExchangeClientId string,
) *AuthProviderConfig {
	return &AuthProviderConfig{
		client:                buildHttpClient(timeoutMilliseconds),
		baseUrl:               baseUrl,
		clientId:              clientId,
		clientSecret:          clientSecret,
		realm:                 realm,
		tokenExchangeClientId: tokenExchangeClientId,
	}
}

func buildHttpClient(timeoutMilliseconds int64) http.Client {
	return http.Client{
		Timeout: time.Duration(timeoutMilliseconds) * time.Millisecond,
	}
}

func (this AuthProviderConfig) GetClient() *http.Client {
	return &this.client
}

func (this AuthProviderConfig) GetBaseUrl() string {
	return this.baseUrl
}

func (this AuthProviderConfig) GetClientId() string {
	return this.clientId
}

func (this AuthProviderConfig) GetClientSecret() string {
	return this.clientSecret
}

func (this AuthProviderConfig) GetRealm() string {
	return this.realm
}

func (this AuthProviderConfig) GetTokenExchangeClientId() string {
	return this.tokenExchangeClientId
}
