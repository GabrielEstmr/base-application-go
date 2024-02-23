package main_gateways_http_authprovider_resources_response

import main_domains "baseapplicationgo/main/domains"

type SessionCredentialsResponse struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func (this SessionCredentialsResponse) ToDomain() main_domains.SessionCredentials {
	return *main_domains.NewSessionCredentials(
		this.AccessToken,
		this.IdToken,
		this.ExpiresIn,
		this.TokenType,
		this.Scope,
		this.RefreshToken,
	)
}
