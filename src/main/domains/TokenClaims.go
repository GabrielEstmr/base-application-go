package main_domains

type TokenClaims struct {
	Exp               float64             `json:"exp"`
	Iat               float64             `json:"iat"`
	Jti               string              `json:"jti"`
	Iss               string              `json:"iss"`
	Aud               string              `json:"aud"`
	Sub               string              `json:"sub"`
	Typ               string              `json:"typ"`
	Azp               string              `json:"azp"`
	SessionState      string              `json:"session_state"`
	Acr               string              `json:"acr"`
	AllowedOrigins    []string            `json:"allowed-origins"`
	RealmAccess       TokenRealmAccess    `json:"realm_access"`
	ResourceAccess    TokenResourceAccess `json:"resource_access"`
	Scope             string              `json:"scope"`
	Sid               string              `json:"sid"`
	EmailVerified     bool                `json:"email_verified"`
	Name              string              `json:"name"`
	PreferredUsername string              `json:"preferred_username"`
	GivenName         string              `json:"given_name"`
	FamilyName        string              `json:"family_name"`
	Email             string              `json:"email"`
}

func (this TokenClaims) FromMap(m map[string]interface{}) TokenClaims {
	this.Exp = m["exp"].(float64)
	this.Iat = m["iat"].(float64)
	this.Jti = m["jti"].(string)
	this.Iss = m["iss"].(string)
	this.Aud = m["aud"].(string)
	this.Sub = m["sub"].(string)
	this.Typ = m["typ"].(string)
	this.Azp = m["azp"].(string)
	this.SessionState = m["session_state"].(string)
	this.Acr = m["acr"].(string)

	AllowedOriginsTemp := m["allowed-origins"].([]interface{})
	allowedOrigins := make([]string, 0)
	for _, v := range AllowedOriginsTemp {
		allowedOrigins = append(allowedOrigins, v.(string))
	}

	this.AllowedOrigins = allowedOrigins
	this.RealmAccess = new(TokenRealmAccess).FromMap(m["realm_access"].(map[string]interface{}))
	this.ResourceAccess = new(TokenResourceAccess).FromMap(m["resource_access"].(map[string]interface{}))
	this.Scope = m["scope"].(string)
	this.Sid = m["sid"].(string)
	this.EmailVerified = m["email_verified"].(bool)
	this.Name = m["name"].(string)
	this.PreferredUsername = m["preferred_username"].(string)
	this.GivenName = m["given_name"].(string)
	this.FamilyName = m["family_name"].(string)
	this.Email = m["email"].(string)

	return this
}
