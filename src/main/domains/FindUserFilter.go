package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"time"
)

type FindUserFilter struct {
	Ids                   []string
	AccountIds            []string
	AuthProviderIds       []string
	DocumentIds           []string
	UserNames             []string
	FirstNames            []string
	LastNames             []string
	Emails                []string
	EmailsVerified        []bool
	Statuses              []main_domains_enums.UserStatus
	Roles                 []string
	ProviderTypes         []main_domains_enums.AuthProviderType
	StartBirthdayDate     time.Time
	EndBirthdayDate       time.Time
	StartCreatedDate      time.Time
	EndCreatedDate        time.Time
	StartLastModifiedDate time.Time
	EndLastModifiedDate   time.Time
}
