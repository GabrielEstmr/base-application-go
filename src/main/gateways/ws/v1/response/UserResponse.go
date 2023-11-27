package main_gateways_ws_v1_response

import (
	main_domains "baseapplicationgo/main/domains"
	"time"
)

type UserResponse struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	DocumentNumber   string    `json:"documentNumber"`
	Birthday         time.Time `json:"birthday"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

func NewUserResponse(user main_domains.User) UserResponse {
	return UserResponse{
		Id:               user.Id,
		Name:             user.Name,
		DocumentNumber:   user.DocumentNumber,
		Birthday:         user.Birthday,
		CreatedDate:      user.CreatedDate,
		LastModifiedDate: user.LastModifiedDate,
	}
}

func (this *UserResponse) ToDomain() main_domains.User {
	return main_domains.User{
		Id:               this.Id,
		Name:             this.Name,
		DocumentNumber:   this.DocumentNumber,
		Birthday:         this.Birthday,
		CreatedDate:      this.CreatedDate,
		LastModifiedDate: this.LastModifiedDate,
	}
}
