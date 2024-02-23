package main_gateways_redis_documents

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"reflect"
)

type PhoneContactDocument struct {
	PhoneCountry string `json:"phoneCountry,omitempty" bson:"phoneCountry,omitempty"`
	PhoneType    string `json:"phoneType,omitempty" bson:"phoneType,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
}

func NewPhoneContactDocument(
	phoneContact main_domains.PhoneContact,
) *PhoneContactDocument {
	return &PhoneContactDocument{
		PhoneCountry: phoneContact.GetPhoneCountry().Name(),
		PhoneType:    phoneContact.GetPhoneType().Name(),
		PhoneNumber:  phoneContact.GetPhoneNumber(),
	}
}

func (this PhoneContactDocument) IsEmpty() bool {
	document := *new(PhoneContactDocument)
	return reflect.DeepEqual(this, document)
}

func (this PhoneContactDocument) ToDomain() main_domains.PhoneContact {
	if this.IsEmpty() {
		return *new(main_domains.PhoneContact)
	}
	return *main_domains.NewPhoneContact(
		new(main_domains_enums.Country).FromValue(this.PhoneCountry),
		new(main_domains_enums.PhoneType).FromValue(this.PhoneType),
		this.PhoneNumber,
	)
}
