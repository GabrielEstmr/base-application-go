package main_gateways_redis_documents

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_utils "baseapplicationgo/main/utils"
	"fmt"
	"reflect"
	"time"
)

const _USER_COLLECTION_NAME = "User"
const _USER_ID_NAME = "_id"
const _USER_IDX_DOCUMENT_ID_NAME = "_idx_documentId"
const _USER_IDX_USERNAME_NAME = "_idx_userName"
const _USER_IDX_EMAIL = "_idx_email"

const USER_DOC_ID_NAME_PREFIX = _USER_COLLECTION_NAME + _USER_ID_NAME
const USER_DOC_IDX_DOCUMENT_ID_NAME_PREFIX = _USER_COLLECTION_NAME + _USER_IDX_DOCUMENT_ID_NAME
const USER_DOC_IDX_USERNAME_NAME_PREFIX = _USER_COLLECTION_NAME + _USER_IDX_USERNAME_NAME
const USER_DOC_IDX_EMAIL_NAME_PREFIX = _USER_COLLECTION_NAME + _USER_IDX_EMAIL

type UserDocument struct {
	Id               string                 `json:"_id,omitempty"`
	AccountId        string                 `json:"accountId,omitempty"`
	AuthProviderId   string                 `json:"authProviderId,omitempty"`
	DocumentId       string                 `json:"documentId,omitempty"`
	UserName         string                 `json:"userName,omitempty"`
	FirstName        string                 `json:"firstName,omitempty"`
	LastName         string                 `json:"lastName,omitempty"`
	Email            string                 `json:"email,omitempty"`
	EmailVerified    bool                   `json:"emailVerified,omitempty"`
	Status           string                 `json:"status,omitempty"`
	Roles            []string               `json:"roles,omitempty"`
	Birthday         time.Time              `json:"birthday,omitempty"`
	PhoneContacts    []PhoneContactDocument `json:"phoneContacts,omitempty"`
	ProviderType     string                 `json:"providerType,omitempty"`
	CreatedDate      time.Time              `json:"createdDate,omitempty"`
	LastModifiedDate time.Time              `json:"lastModifiedDate,omitempty"`
}

func NewUserDocument(user main_domains.User) UserDocument {
	contacts := make([]PhoneContactDocument, 0)
	for _, v := range user.GetPhoneContacts() {
		contacts = append(contacts, *NewPhoneContactDocument(v))
	}
	return UserDocument{
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

func (this UserDocument) GetKeys() map[string]string {
	keys := make(map[string]string)
	if !new(main_utils.StringUtils).IsEmpty(this.Id) {
		keys[_USER_ID_NAME] = fmt.Sprintf(USER_DOC_ID_NAME_PREFIX+"%s", this.Id)
	}
	if !new(main_utils.StringUtils).IsEmpty(this.DocumentId) {
		keys[_USER_IDX_DOCUMENT_ID_NAME] = fmt.Sprintf(USER_DOC_IDX_DOCUMENT_ID_NAME_PREFIX+"%s", this.DocumentId)
	}
	if !new(main_utils.StringUtils).IsEmpty(this.UserName) {
		keys[_USER_IDX_USERNAME_NAME] = fmt.Sprintf(USER_DOC_IDX_USERNAME_NAME_PREFIX+"%s", this.UserName)
	}
	if !new(main_utils.StringUtils).IsEmpty(this.Email) {
		keys[_USER_IDX_EMAIL] = fmt.Sprintf(USER_DOC_IDX_EMAIL_NAME_PREFIX+"%s", this.Email)
	}
	return keys
}

func (this UserDocument) IsEmpty() bool {
	document := *new(UserDocument)
	return reflect.DeepEqual(this, document)
}

func (this UserDocument) ToDomain() main_domains.User {
	if this.IsEmpty() {
		return *new(main_domains.User)
	}
	contacts := make([]main_domains.PhoneContact, 0)
	for _, v := range this.PhoneContacts {
		contacts = append(contacts, v.ToDomain())
	}
	return *main_domains.NewUserNoPassword(
		this.Id,
		this.AccountId,
		this.AuthProviderId,
		this.DocumentId,
		this.UserName,
		this.FirstName,
		this.LastName,
		this.Email,
		this.EmailVerified,
		new(main_domains_enums.UserStatus).FromValue(this.Status),
		this.Roles,
		this.Birthday,
		contacts,
		new(main_domains_enums.AuthProviderType).FromValue(this.ProviderType),
		this.CreatedDate,
		this.LastModifiedDate,
	)
}
