package main_domains

import main_domains_enums "baseapplicationgo/main/domains/enums"

type PhoneContact struct {
	phoneCountry main_domains_enums.Country
	phoneType    main_domains_enums.PhoneType
	phoneNumber  string
}

func NewPhoneContact(
	phoneCountry main_domains_enums.Country,
	phoneType main_domains_enums.PhoneType,
	phoneNumber string,
) *PhoneContact {
	return &PhoneContact{
		phoneCountry: phoneCountry,
		phoneType:    phoneType,
		phoneNumber:  phoneNumber,
	}
}

func (this PhoneContact) GetPhoneCountry() main_domains_enums.Country {
	return this.phoneCountry
}

func (this PhoneContact) GetPhoneType() main_domains_enums.PhoneType {
	return this.phoneType
}

func (this PhoneContact) GetPhoneNumber() string {
	return this.phoneNumber
}
