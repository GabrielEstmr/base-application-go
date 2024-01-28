package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"time"
)

type FindEmailFilter struct {
	ids                   []string
	statuses              []main_domains_enums.EmailStatus
	emailTemplateTypes    []string
	appOwners             []string
	startCreatedDate      time.Time
	endCreatedDate        time.Time
	startLastModifiedDate time.Time
	endLastModifiedDate   time.Time
}

func NewFindEmailFilter(
	ids []string,
	statuses []main_domains_enums.EmailStatus,
	emailTemplateTypes []string,
	appOwners []string,
	startCreatedDate time.Time,
	endCreatedDate time.Time,
	startLastModifiedDate time.Time,
	endLastModifiedDate time.Time,
) *FindEmailFilter {
	return &FindEmailFilter{
		ids:                   ids,
		statuses:              statuses,
		emailTemplateTypes:    emailTemplateTypes,
		appOwners:             appOwners,
		startCreatedDate:      startCreatedDate,
		endCreatedDate:        endCreatedDate,
		startLastModifiedDate: startLastModifiedDate,
		endLastModifiedDate:   endLastModifiedDate,
	}
}

func (this FindEmailFilter) GetIds() []string {
	return this.ids
}

func (this FindEmailFilter) GetStatuses() []main_domains_enums.EmailStatus {
	return this.statuses
}

func (this FindEmailFilter) GetEmailTemplateTypes() []string {
	return this.emailTemplateTypes
}

func (this FindEmailFilter) GetAppOwners() []string {
	return this.appOwners
}

func (this FindEmailFilter) GetStartCreatedDate() time.Time {
	return this.startCreatedDate
}

func (this FindEmailFilter) GetEndCreatedDate() time.Time {
	return this.endCreatedDate
}

func (this FindEmailFilter) GetStartLastModifiedDate() time.Time {
	return this.startLastModifiedDate
}

func (this FindEmailFilter) GetEndLastModifiedDate() time.Time {
	return this.endLastModifiedDate
}

func (this FindEmailFilter) WithIds(ids []string) FindEmailFilter {
	this.ids = ids
	return this
}

func (this FindEmailFilter) WithStatuses(statuses []main_domains_enums.EmailStatus) FindEmailFilter {
	this.statuses = statuses
	return this
}
