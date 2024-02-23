package main_gateways_http_authprovider_resources_request

type VerifyUserRequest struct {
	EmailVerified bool `json:"emailVerified,omitempty"`
	Enabled       bool `json:"enabled,omitempty"`
}

func NewVerifyUserRequest(
	emailVerified bool,
	enabled bool,
) *VerifyUserRequest {
	return &VerifyUserRequest{
		EmailVerified: emailVerified,
		Enabled:       enabled,
	}
}
