package main_gateways_http_authprovider_resources_request

import main_domains "baseapplicationgo/main/domains"

type CreateUserRequest struct {
	Attributes    UserAttributesRequest    `json:"attributes,omitempty"`
	Credentials   []UserCredentialsRequest `json:"credentials,omitempty"`
	UserName      string                   `json:"username,omitempty"`
	FirstName     string                   `json:"firstName,omitempty"`
	LastName      string                   `json:"lastName,omitempty"`
	Email         string                   `json:"email,omitempty"`
	EmailVerified bool                     `json:"emailVerified,omitempty"`
	Enabled       bool                     `json:"enabled,omitempty"`
}

func NewCreateUserRequest(
	user main_domains.User,
) *CreateUserRequest {
	return &CreateUserRequest{
		Attributes: UserAttributesRequest{
			AttributeKey: "test_value",
		},
		Credentials: []UserCredentialsRequest{{
			Temporary: false,
			Type:      "password",
			Value:     user.GetPassword(),
		}},
		UserName:      user.GetEmail(),
		FirstName:     user.GetFirstName(),
		LastName:      user.GetLastName(),
		Email:         user.GetEmail(),
		EmailVerified: false,
		Enabled:       false,
	}
}
