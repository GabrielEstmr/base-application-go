package main_gateways

import main_domains "baseapplicationgo/main/domains"

type UserDatabaseGateway interface {
	Save(indicator main_domains.User) (string, error)
	FindById(id string) (main_domains.User, error)
	//FindByDocumentNumber(documentNumber string) (main_domains.User, error)
}
