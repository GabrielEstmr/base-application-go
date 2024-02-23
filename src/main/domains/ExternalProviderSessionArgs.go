package main_domains

import main_domains_enums "baseapplicationgo/main/domains/enums"

type ExternalProviderSessionArgs struct {
	token    string
	provider main_domains_enums.AuthProviderType
}

func NewExternalProviderSessionArgs(
	token string,
	provider main_domains_enums.AuthProviderType,
) *ExternalProviderSessionArgs {
	return &ExternalProviderSessionArgs{
		token:    token,
		provider: provider,
	}
}

func (this ExternalProviderSessionArgs) GetToken() string {
	return this.token
}

func (this ExternalProviderSessionArgs) GetProvider() main_domains_enums.AuthProviderType {
	return this.provider
}
