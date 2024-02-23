package main_domains

type AuthProviderUser struct {
	id               string
	createdTimestamp int64
	username         string
	enabled          bool
	totp             bool
	emailVerified    bool
	firstName        string
	lastName         string
	email            string
	notBefore        int
}

func NewAuthProviderUser(
	id string,
	createdTimestamp int64,
	username string,
	enabled bool,
	totp bool,
	emailVerified bool,
	firstName string,
	lastName string,
	email string,
	notBefore int,
) *AuthProviderUser {
	return &AuthProviderUser{
		id:               id,
		createdTimestamp: createdTimestamp,
		username:         username,
		enabled:          enabled,
		totp:             totp,
		emailVerified:    emailVerified,
		firstName:        firstName,
		lastName:         lastName,
		email:            email,
		notBefore:        notBefore,
	}
}

func (this AuthProviderUser) GetId() string {
	return this.id
}

func (this AuthProviderUser) GetCreatedTimestamp() int64 {
	return this.createdTimestamp
}

func (this AuthProviderUser) GetUsername() string {
	return this.username
}

func (this AuthProviderUser) GetEnabled() bool {
	return this.enabled
}

func (this AuthProviderUser) GetTotp() bool {
	return this.totp
}

func (this AuthProviderUser) GetEmailVerified() bool {
	return this.emailVerified
}

func (this AuthProviderUser) GetFirstName() string {
	return this.firstName
}

func (this AuthProviderUser) GetLastName() string {
	return this.lastName
}

func (this AuthProviderUser) GetEmail() string {
	return this.email
}

func (this AuthProviderUser) GetNotBefore() int {
	return this.notBefore
}
