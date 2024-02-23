package main_domains

import "time"

type EnableExternalUserArgs struct {
	documentId    string
	userName      string
	birthday      time.Time
	phoneContacts []PhoneContact
}

func NewEnableExternalUserArgs(
	documentId string,
	userName string,
	birthday time.Time,
	phoneContacts []PhoneContact) *EnableExternalUserArgs {
	return &EnableExternalUserArgs{
		documentId:    documentId,
		userName:      userName,
		birthday:      birthday,
		phoneContacts: phoneContacts}
}

func (this EnableExternalUserArgs) GetDocumentId() string {
	return this.documentId
}

func (this EnableExternalUserArgs) GetUserName() string {
	return this.userName
}

func (this EnableExternalUserArgs) GetBirthday() time.Time {
	return this.birthday
}

func (this EnableExternalUserArgs) GetPhoneContacts() []PhoneContact {
	return this.phoneContacts
}
