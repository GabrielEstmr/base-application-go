package main_gateways_http_authprovider_resources_response

import main_domains "baseapplicationgo/main/domains"

type UserResponse struct {
	ID               string                 `json:"id"`
	CreatedTimestamp int64                  `json:"createdTimestamp"`
	Username         string                 `json:"username"`
	Enabled          bool                   `json:"enabled"`
	Totp             bool                   `json:"totp"`
	EmailVerified    bool                   `json:"emailVerified"`
	FirstName        string                 `json:"firstName"`
	LastName         string                 `json:"lastName"`
	Email            string                 `json:"email"`
	Attributes       UserAttributesResponse `json:"attributes"`
	NotBefore        int                    `json:"notBefore"`
	Access           UserAccessResponse     `json:"access"`
}

func (this UserResponse) ToDomain() main_domains.AuthProviderUser {
	return *main_domains.NewAuthProviderUser(
		this.ID,
		this.CreatedTimestamp,
		this.Username,
		this.Enabled,
		this.Totp,
		this.EmailVerified,
		this.FirstName,
		this.LastName,
		this.Email,
		this.NotBefore,
	)
}
