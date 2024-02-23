package main_gateways_ws_v1_response

import main_domains "baseapplicationgo/main/domains"

type SessionCredentialsResponse struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func NewSessionCredentialsResponse(
	sessionCredentials main_domains.SessionCredentials,
) *SessionCredentialsResponse {
	return &SessionCredentialsResponse{
		AccessToken:  sessionCredentials.GetAccessToken(),
		IdToken:      sessionCredentials.GetIdToken(),
		ExpiresIn:    sessionCredentials.GetExpiresIn(),
		TokenType:    sessionCredentials.GetTokenType(),
		Scope:        sessionCredentials.GetScope(),
		RefreshToken: sessionCredentials.GetRefreshToken(),
	}
}
