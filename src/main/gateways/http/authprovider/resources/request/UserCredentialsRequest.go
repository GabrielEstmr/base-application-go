package main_gateways_http_authprovider_resources_request

type UserCredentialsRequest struct {
	Temporary bool   `json:"temporary,omitempty"`
	Type      string `json:"type,omitempty"`
	Value     string `json:"value,omitempty"`
}
