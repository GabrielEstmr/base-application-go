package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"reflect"
	"time"
)

type User struct {
	id               string
	accountId        string
	authProviderId   string
	documentId       string
	userName         string
	password         string
	firstName        string
	lastName         string
	email            string
	emailVerified    bool
	status           main_domains_enums.UserStatus
	roles            []string
	birthday         time.Time
	phoneContacts    []PhoneContact
	providerType     main_domains_enums.AuthProviderType
	createdDate      time.Time
	lastModifiedDate time.Time
}

func NewUserAsCreated(
	documentId string,
	userName string,
	password string,
	firstName string,
	lastName string,
	email string,
	birthday time.Time,
	phoneContacts []PhoneContact,
	providerType main_domains_enums.AuthProviderType,
) *User {
	return &User{
		documentId:    documentId,
		userName:      userName,
		password:      password,
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		status:        main_domains_enums.USER_CREATED,
		birthday:      birthday,
		phoneContacts: phoneContacts,
		providerType:  providerType,
	}
}

func NewUserAsCreatedFromProvider(
	tokenClaims TokenClaims,
	providerType main_domains_enums.AuthProviderType,
) *User {
	return &User{
		authProviderId: tokenClaims.Sub,
		documentId:     tokenClaims.Sub,
		userName:       tokenClaims.Sub,
		firstName:      tokenClaims.GivenName,
		lastName:       tokenClaims.FamilyName,
		email:          tokenClaims.Email,
		status:         main_domains_enums.USER_CREATED,
		emailVerified:  false,
		providerType:   providerType,
	}
}

func NewUserAllArgs(
	id string,
	accountId string,
	authProviderId string,
	documentId string,
	userName string,
	password string,
	firstName string,
	lastName string,
	email string,
	emailVerified bool,
	status main_domains_enums.UserStatus,
	roles []string,
	birthday time.Time,
	phoneContacts []PhoneContact,
	createdDate time.Time,
	lastModifiedDate time.Time,
	providerType main_domains_enums.AuthProviderType,
) *User {
	return &User{
		id:               id,
		accountId:        accountId,
		authProviderId:   authProviderId,
		documentId:       documentId,
		userName:         userName,
		password:         password,
		firstName:        firstName,
		lastName:         lastName,
		email:            email,
		emailVerified:    emailVerified,
		status:           status,
		roles:            roles,
		birthday:         birthday,
		phoneContacts:    phoneContacts,
		providerType:     providerType,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate,
	}
}

func NewUserNoPassword(
	id string,
	accountId string,
	authProviderId string,
	documentId string,
	userName string,
	firstName string,
	lastName string,
	email string,
	emailVerified bool,
	status main_domains_enums.UserStatus,
	roles []string,
	birthday time.Time,
	phoneContacts []PhoneContact,
	providerType main_domains_enums.AuthProviderType,
	createdDate time.Time,
	lastModifiedDate time.Time) *User {
	return &User{
		id:               id,
		accountId:        accountId,
		authProviderId:   authProviderId,
		documentId:       documentId,
		userName:         userName,
		firstName:        firstName,
		lastName:         lastName,
		email:            email,
		emailVerified:    emailVerified,
		status:           status,
		roles:            roles,
		birthday:         birthday,
		phoneContacts:    phoneContacts,
		providerType:     providerType,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate}
}

func (this User) GetId() string {
	return this.id
}

func (this User) GetAccountId() string {
	return this.accountId
}

func (this User) GetAuthProviderId() string {
	return this.authProviderId
}

func (this User) GetDocumentId() string {
	return this.documentId
}

func (this User) GetUserName() string {
	return this.userName
}

func (this User) GetPassword() string {
	return this.password
}

func (this User) GetFirstName() string {
	return this.firstName
}

func (this User) GetLastName() string {
	return this.lastName
}

func (this User) GetEmail() string {
	return this.email
}

func (this User) GetEmailVerified() bool {
	return this.emailVerified
}

func (this User) GetStatus() main_domains_enums.UserStatus {
	return this.status
}

func (this User) GetRoles() []string {
	return this.roles
}

func (this User) GetBirthday() time.Time {
	return this.birthday
}

func (this User) GetPhoneContacts() []PhoneContact {
	return this.phoneContacts
}

func (this User) GetProviderType() main_domains_enums.AuthProviderType {
	return this.providerType
}

func (this User) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this User) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this User) IsEmpty() bool {
	document := *new(User)
	return reflect.DeepEqual(this, document)
}

func (this User) IsKeycloakProvider() bool {
	return this.providerType == main_domains_enums.AUTH_PROVIDER_KEYCLOAK
}

func (this User) IsExternalAuthProvider() bool {
	return !this.IsKeycloakProvider()
}

func (this User) IsInternalAuthProvider() bool {
	return !this.IsExternalAuthProvider()
}

func (this User) IsStatusCreated() bool {
	return main_domains_enums.USER_CREATED == this.status
}

func (this User) IsStatusEnabled() bool {
	return main_domains_enums.USER_ENABLED == this.status
}

func (this User) CloneWithNewAuthProviderId(authProviderId string) User {
	return User{
		id:               this.id,
		accountId:        this.accountId,
		authProviderId:   authProviderId,
		documentId:       this.documentId,
		userName:         this.userName,
		password:         this.password,
		firstName:        this.firstName,
		lastName:         this.lastName,
		email:            this.email,
		emailVerified:    this.emailVerified,
		status:           this.status,
		roles:            this.roles,
		birthday:         this.birthday,
		phoneContacts:    this.phoneContacts,
		providerType:     this.providerType,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this User) CloneEnabledAsVerified() User {
	return User{
		id:               this.id,
		accountId:        this.accountId,
		authProviderId:   this.authProviderId,
		documentId:       this.documentId,
		userName:         this.userName,
		password:         this.password,
		firstName:        this.firstName,
		lastName:         this.lastName,
		email:            this.email,
		emailVerified:    true,
		status:           main_domains_enums.USER_ENABLED,
		roles:            this.roles,
		birthday:         this.birthday,
		phoneContacts:    this.phoneContacts,
		providerType:     this.providerType,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this User) CloneEnabledAsVerifiedAndAttData(args EnableExternalUserArgs) User {
	return User{
		id:               this.id,
		accountId:        this.accountId,
		authProviderId:   this.authProviderId,
		documentId:       args.GetDocumentId(),
		userName:         args.GetUserName(),
		password:         this.password,
		firstName:        this.firstName,
		lastName:         this.lastName,
		email:            this.email,
		emailVerified:    true,
		status:           main_domains_enums.USER_ENABLED,
		roles:            this.roles,
		birthday:         args.GetBirthday(),
		phoneContacts:    args.GetPhoneContacts(),
		providerType:     this.providerType,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this User) CloneWithNewPassword(password string) User {
	return User{
		id:               this.id,
		accountId:        this.accountId,
		authProviderId:   this.authProviderId,
		documentId:       this.documentId,
		userName:         this.userName,
		password:         password,
		firstName:        this.firstName,
		lastName:         this.lastName,
		email:            this.email,
		emailVerified:    this.emailVerified,
		status:           this.status,
		roles:            this.roles,
		birthday:         this.birthday,
		phoneContacts:    this.phoneContacts,
		providerType:     this.providerType,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}
