package main_gateways

import main_domains "baseapplicationgo/main/domains"

type UserDatabaseGateway interface {
	Save(user main_domains.User) (main_domains.User, error)
	FindById(id string) (main_domains.User, error)
	FindByDocumentNumber(documentNumber string) (main_domains.User, error)
	FindByFilter(filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error)
}
