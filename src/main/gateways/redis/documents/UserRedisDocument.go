package main_gateways_redis_documents

import (
	main_domains "baseapplicationgo/main/domains"
	"time"
)

// TODO ver depois: https://stackoverflow.com/questions/11126793/json-and-dealing-with-unexported-fields
type UserRedisDocument struct {
	Id               string    `json:"_id,omitempty"`
	Name             string    `json:"name,omitempty"`
	DocumentNumber   string    `json:"documentNumber,omitempty"`
	Birthday         time.Time `json:"birthday,omitempty"`
	CreatedDate      time.Time `json:"createdDate,omitempty"`
	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`
}

func NewUserRedisDocument(user main_domains.User) UserRedisDocument {
	return UserRedisDocument{
		Id:               user.Id,
		Name:             user.Name,
		DocumentNumber:   user.DocumentNumber,
		Birthday:         user.Birthday,
		CreatedDate:      user.CreatedDate,
		LastModifiedDate: user.LastModifiedDate,
	}
}

func (this *UserRedisDocument) IsEmpty() bool {
	return *this == UserRedisDocument{}
}

func (this *UserRedisDocument) ToDomain() main_domains.User {
	if (*this == UserRedisDocument{}) {
		return main_domains.User{}
	}
	return main_domains.User{
		Id:               this.Id,
		Name:             this.Name,
		DocumentNumber:   this.DocumentNumber,
		Birthday:         this.Birthday,
		CreatedDate:      this.CreatedDate,
		LastModifiedDate: this.LastModifiedDate,
	}
}
