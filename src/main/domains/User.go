package main_domains

import "time"

type User struct {
	Id               string
	Name             string
	DocumentNumber   string
	Birthday         time.Time
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
