package main_gateways_ws_v1_response

import (
	main_domains "baseapplicationgo/main/domains"
	"time"
)

type UserResponse struct {
	Id               string                 `json:"id"`
	AccountId        string                 `json:"accountId"`
	AuthProviderId   string                 `json:"authProviderId"`
	DocumentId       string                 `json:"documentId"`
	UserName         string                 `json:"userName"`
	FirstName        string                 `json:"firstName"`
	LastName         string                 `json:"lastName"`
	Email            string                 `json:"email"`
	EmailVerified    bool                   `json:"emailVerified"`
	Status           string                 `json:"status"`
	Roles            []string               `json:"roles"`
	Birthday         time.Time              `json:"birthday"`
	PhoneContacts    []PhoneContactResponse `json:"phoneContacts"`
	ProviderType     string                 `json:"providerType"`
	CreatedDate      time.Time              `json:"createdDate"`
	LastModifiedDate time.Time              `json:"lastModifiedDate"`
}

func NewUserResponse(user main_domains.User) UserResponse {
	contacts := make([]PhoneContactResponse, 0)
	for _, v := range user.GetPhoneContacts() {
		contacts = append(contacts, *NewPhoneContactResponse(v))
	}
	return UserResponse{
		Id:               user.GetId(),
		AccountId:        user.GetAccountId(),
		AuthProviderId:   user.GetAuthProviderId(),
		DocumentId:       user.GetDocumentId(),
		UserName:         user.GetUserName(),
		FirstName:        user.GetFirstName(),
		LastName:         user.GetLastName(),
		Email:            user.GetEmail(),
		EmailVerified:    user.GetEmailVerified(),
		Status:           user.GetStatus().Name(),
		Roles:            user.GetRoles(),
		Birthday:         user.GetBirthday(),
		PhoneContacts:    contacts,
		ProviderType:     user.GetProviderType().Name(),
		CreatedDate:      user.GetCreatedDate(),
		LastModifiedDate: user.GetLastModifiedDate(),
	}
}
