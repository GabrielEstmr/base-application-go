package main_domains

type TokenResourceAccess struct {
	Account TokenAccount `json:"account"`
}

func (this TokenResourceAccess) FromMap(m map[string]interface{}) TokenResourceAccess {
	accountM := m["account"].(map[string]interface{})
	this.Account = *new(TokenAccount).FromMap(accountM)
	return this
}
