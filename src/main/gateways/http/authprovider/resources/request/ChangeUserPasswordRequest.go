package main_gateways_http_authprovider_resources_request

import main_domains "baseapplicationgo/main/domains"

type ChangeUserPasswordRequest struct {
	Credentials []UserCredentialsRequest `json:"credentials,omitempty"`
}

func NewChangeUserPasswordRequest(
	user main_domains.User,
) *ChangeUserPasswordRequest {
	return &ChangeUserPasswordRequest{
		Credentials: []UserCredentialsRequest{{
			Temporary: false,
			Type:      "password",
			Value:     user.GetPassword(),
		}},
	}
}
