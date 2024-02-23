package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"time"
)

type FindUserEmailVerificationFilter struct {
	Ids                   []string
	UserIds               []string
	VerificationCodes     []string
	Scopes                []string
	Statuses              []main_domains_enums.EmailVerificationStatus
	StartCreatedDate      time.Time
	EndCreatedDate        time.Time
	StartLastModifiedDate time.Time
	EndLastModifiedDate   time.Time
}
