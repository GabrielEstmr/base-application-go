package main_domains

import "time"

type Account struct {
	Id               string
	UserId           string
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
