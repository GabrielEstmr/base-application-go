package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"reflect"
	"time"
)

type UserEmailVerification struct {
	id               string
	userId           string
	verificationCode string
	scope            main_domains_enums.UserEmailVerificationScope
	emailParams      EmailParams
	status           main_domains_enums.EmailVerificationStatus
	createdDate      time.Time
	lastModifiedDate time.Time
}

func NewUserEmailVerification(
	userId string,
	verificationCode string,
	scope main_domains_enums.UserEmailVerificationScope,
	emailParams EmailParams,
	status main_domains_enums.EmailVerificationStatus,
) *UserEmailVerification {
	return &UserEmailVerification{
		userId:           userId,
		verificationCode: verificationCode,
		scope:            scope,
		emailParams:      emailParams,
		status:           status,
	}
}

func NewUserEmailVerificationAllArgs(
	id string,
	userId string,
	verificationCode string,
	scope main_domains_enums.UserEmailVerificationScope,
	emailParams EmailParams,
	status main_domains_enums.EmailVerificationStatus,
	createdDate time.Time,
	lastModifiedDate time.Time,
) *UserEmailVerification {
	return &UserEmailVerification{
		id:               id,
		userId:           userId,
		verificationCode: verificationCode,
		scope:            scope,
		emailParams:      emailParams,
		status:           status,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate,
	}
}

func (this UserEmailVerification) GetId() string {
	return this.id
}

func (this UserEmailVerification) GetUserId() string {
	return this.userId
}

func (this UserEmailVerification) GetVerificationCode() string {
	return this.verificationCode
}

func (this UserEmailVerification) GetScope() main_domains_enums.UserEmailVerificationScope {
	return this.scope
}

func (this UserEmailVerification) GetEmailParams() EmailParams {
	return this.emailParams
}

func (this UserEmailVerification) GetStatus() main_domains_enums.EmailVerificationStatus {
	return this.status
}

func (this UserEmailVerification) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this UserEmailVerification) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this UserEmailVerification) CloneAsSent() UserEmailVerification {
	return UserEmailVerification{
		id:               this.id,
		userId:           this.userId,
		verificationCode: this.verificationCode,
		emailParams:      this.emailParams,
		status:           main_domains_enums.EMAIL_VERIFICATION_SENT,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this UserEmailVerification) CloneAsUsed() UserEmailVerification {
	return UserEmailVerification{
		id:               this.id,
		userId:           this.userId,
		verificationCode: this.verificationCode,
		emailParams:      this.emailParams,
		status:           main_domains_enums.EMAIL_VERIFICATION_USED,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this UserEmailVerification) IsEmpty() bool {
	document := *new(UserEmailVerification)
	return reflect.DeepEqual(this, document)
}
