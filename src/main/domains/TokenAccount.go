package main_domains

type TokenAccount struct {
	Roles []string `json:"roles"`
}

func (this *TokenAccount) FromMap(m map[string]interface{}) *TokenAccount {
	rolesM := m["roles"].([]interface{})
	roles := make([]string, 0)
	for _, v := range rolesM {
		roles = append(roles, v.(string))
	}
	this.Roles = roles
	return this
}
