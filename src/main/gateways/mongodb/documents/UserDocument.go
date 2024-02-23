package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type UserDocument struct {
	Id               primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountId        string                 `json:"accountId,omitempty" bson:"accountId,omitempty"`
	AuthProviderId   string                 `json:"authProviderId,omitempty" bson:"authProviderId,omitempty"`
	DocumentId       string                 `json:"documentId,omitempty" bson:"documentId,omitempty"`
	UserName         string                 `json:"userName,omitempty" bson:"userName,omitempty"`
	FirstName        string                 `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName         string                 `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email            string                 `json:"email,omitempty" bson:"email,omitempty"`
	EmailVerified    bool                   `json:"emailVerified" bson:"emailVerified"`
	Status           string                 `json:"status,omitempty" bson:"status,omitempty"`
	Roles            []string               `json:"roles,omitempty" bson:"roles,omitempty"`
	Birthday         primitive.DateTime     `json:"birthday,omitempty" bson:"birthday"`
	PhoneContacts    []PhoneContactDocument `json:"phoneContacts,omitempty" bson:"phoneContacts,omitempty"`
	ProviderType     string                 `json:"providerType,omitempty" bson:"providerType,omitempty"`
	CreatedDate      primitive.DateTime     `json:"createdDate,omitempty" bson:"createdDate"`
	LastModifiedDate primitive.DateTime     `json:"lastModifiedDate,omitempty" bson:"lastModifiedDate"`
}

func NewUserDocument(user main_domains.User) UserDocument {
	oId, _ := primitive.ObjectIDFromHex(user.GetId())
	contacts := make([]PhoneContactDocument, 0)
	for _, v := range user.GetPhoneContacts() {
		contacts = append(contacts, *NewPhoneContactDocument(v))
	}
	return UserDocument{
		Id:               oId,
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
		Birthday:         primitive.NewDateTimeFromTime(user.GetBirthday()),
		PhoneContacts:    contacts,
		ProviderType:     user.GetProviderType().Name(),
		CreatedDate:      primitive.NewDateTimeFromTime(user.GetCreatedDate()),
		LastModifiedDate: primitive.NewDateTimeFromTime(user.GetLastModifiedDate()),
	}
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
		this.Id.Hex(),
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
		this.Birthday.Time(),
		contacts,
		new(main_domains_enums.AuthProviderType).FromValue(this.ProviderType),
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}
