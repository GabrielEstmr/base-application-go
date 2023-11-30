package main_gateways

import main_domains "baseapplicationgo/main/domains"

type UserDatabaseCacheGateway interface {
	Save(user main_domains.User) (main_domains.User, error)
	FindById(id string) (main_domains.User, error)
	//FindByDocumentNumber(documentNumber string) (main_domains.User, error)
}
