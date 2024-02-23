package main_domains

type SessionCredentials struct {
	accessToken  string
	idToken      string
	expiresIn    int
	tokenType    string
	scope        string
	refreshToken string
}

func NewSessionCredentials(
	accessToken string,
	idToken string,
	expiresIn int,
	tokenType string,
	scope string,
	refreshToken string,
) *SessionCredentials {
	return &SessionCredentials{
		accessToken:  accessToken,
		idToken:      idToken,
		expiresIn:    expiresIn,
		tokenType:    tokenType,
		scope:        scope,
		refreshToken: refreshToken,
	}
}

func (this SessionCredentials) GetAccessToken() string {
	return this.accessToken
}

func (this SessionCredentials) GetIdToken() string {
	return this.idToken
}

func (this SessionCredentials) GetExpiresIn() int {
	return this.expiresIn
}

func (this SessionCredentials) GetTokenType() string {
	return this.tokenType
}

func (this SessionCredentials) GetScope() string {
	return this.scope
}

func (this SessionCredentials) GetRefreshToken() string {
	return this.refreshToken
}
