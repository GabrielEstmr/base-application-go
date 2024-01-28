package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"reflect"
	"time"
)

type Email struct {
	id               string
	eventId          string
	emailParams      EmailParams
	status           main_domains_enums.EmailStatus
	errorMsg         string
	createdDate      time.Time
	lastModifiedDate time.Time
}

func NewEmail(
	eventId string,
	emailParams EmailParams,
	status main_domains_enums.EmailStatus,
) *Email {
	return &Email{
		eventId:     eventId,
		emailParams: emailParams,
		status:      status,
	}
}

func NewEmailAllArgs(
	id string,
	eventId string,
	emailParams EmailParams,
	status main_domains_enums.EmailStatus,
	errorMsg string,
	createdDate time.Time,
	lastModifiedDate time.Time,
) *Email {
	return &Email{
		id:               id,
		eventId:          eventId,
		emailParams:      emailParams,
		status:           status,
		errorMsg:         errorMsg,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate,
	}
}

func (this Email) GetId() string {
	return this.id
}

func (this Email) GetEventId() string {
	return this.eventId
}

func (this Email) GetEmailParams() EmailParams {
	return this.emailParams
}

func (this Email) GetStatus() main_domains_enums.EmailStatus {
	return this.status
}

func (this Email) GetErrorMsg() string {
	return this.errorMsg
}

func (this Email) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this Email) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this Email) CloneAsSent() Email {
	return Email{
		id:               this.id,
		eventId:          this.eventId,
		emailParams:      this.emailParams,
		status:           main_domains_enums.EMAIL_STATUS_SENT,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this Email) CloneAsIntegrationError(errorMsg string) Email {
	return Email{
		id:               this.id,
		eventId:          this.eventId,
		emailParams:      this.emailParams,
		status:           main_domains_enums.EMAIL_STATUS_INTEGRATION_ERROR,
		errorMsg:         errorMsg,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this Email) CloneAsError(errorMsg string) Email {
	return Email{
		id:               this.id,
		eventId:          this.eventId,
		emailParams:      this.emailParams,
		status:           main_domains_enums.EMAIL_STATUS_ERROR,
		errorMsg:         errorMsg,
		createdDate:      this.createdDate,
		lastModifiedDate: this.lastModifiedDate,
	}
}

func (this Email) IsEmpty() bool {
	document := *new(Email)
	return reflect.DeepEqual(this, document)
}

func (this Email) IsError() bool {
	return this.status == main_domains_enums.EMAIL_STATUS_ERROR
}

func (this Email) IsIntegrationError() bool {
	return this.status == main_domains_enums.EMAIL_STATUS_INTEGRATION_ERROR
}

func (this Email) IsStarted() bool {
	return this.status == main_domains_enums.EMAIL_STATUS_STARTED
}

func (this Email) IsAbleToReprocess() bool {
	return this.IsStarted() || this.IsIntegrationError()
}
