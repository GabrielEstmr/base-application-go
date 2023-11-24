package main_gateways

import main_domains "baseapplicationgo/main/domains"

type AccountDatabaseGateway interface {
	Save(indicator main_domains.Account) (main_domains.Account, error)
	FindById(id string) (main_domains.Account, error)
}
