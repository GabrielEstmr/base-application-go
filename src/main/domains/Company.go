package main_domains

import (
	"time"
)

type Company struct {
	id               string
	name             string
	ownerId          string
	country          string
	phoneContacts    []PhoneContact
	createdDate      time.Time
	lastModifiedDate time.Time
}
