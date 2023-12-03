package main_domains

import "time"

type FindUserFilter struct {
	name                  []string
	documentNumber        []string
	birthday              []time.Time
	startCreatedDate      []time.Time
	endCreatedDate        []time.Time
	startLastModifiedDate []time.Time
	endLastModifiedDate   []time.Time
	page                  Page
}

func NewFindUserFilter(
	name []string,
	documentNumber []string,
	birthday []time.Time,
	startCreatedDate []time.Time,
	endCreatedDate []time.Time,
	startLastModifiedDate []time.Time,
	endLastModifiedDate []time.Time,
	page Page,
) *FindUserFilter {
	return &FindUserFilter{
		name:                  name,
		documentNumber:        documentNumber,
		birthday:              birthday,
		startCreatedDate:      startCreatedDate,
		endCreatedDate:        endCreatedDate,
		startLastModifiedDate: startLastModifiedDate,
		endLastModifiedDate:   endLastModifiedDate,
		page:                  page}
}

func (f FindUserFilter) getName() []string {
	return f.name
}

func (f FindUserFilter) getDocumentNumber() []string {
	return f.documentNumber
}

func (f FindUserFilter) getBirthday() []time.Time {
	return f.birthday
}

func (f FindUserFilter) getStartCreatedDate() []time.Time {
	return f.startCreatedDate
}

func (f FindUserFilter) getEndCreatedDate() []time.Time {
	return f.endCreatedDate
}

func (f FindUserFilter) getStartLastModifiedDate() []time.Time {
	return f.startLastModifiedDate
}

func (f FindUserFilter) getEndLastModifiedDate() []time.Time {
	return f.endLastModifiedDate
}

func (f FindUserFilter) getPage() Page {
	return f.page
}
