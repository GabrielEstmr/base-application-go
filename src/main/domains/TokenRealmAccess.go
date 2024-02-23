package main_domains

type TokenRealmAccess struct {
	Roles []string `json:"roles"`
}

func (this TokenRealmAccess) FromMap(m map[string]interface{}) TokenRealmAccess {
	rolesM := m["roles"].([]interface{})
	roles := make([]string, 0)
	for _, v := range rolesM {
		roles = append(roles, v.(string))
	}
	this.Roles = roles
	return this
}
