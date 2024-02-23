package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"time"
)

type Account struct {
	id                string
	status            main_domains_enums.AccountStatus
	ownerId           string
	companies         []Company
	subscriptionDate  time.Time
	creationUserEmail string
	createdDate       time.Time
	lastModifiedDate  time.Time
}
