package main_domains

import "time"

type FindUserFilter struct {
	name                  []string
	documentNumber        []string
	birthday              []time.Time
	startCreatedDate      time.Time
	endCreatedDate        time.Time
	startLastModifiedDate time.Time
	endLastModifiedDate   time.Time
}

func NewFindUserFilter(
	name []string,
	documentNumber []string,
	birthday []time.Time,
	startCreatedDate time.Time,
	endCreatedDate time.Time,
	startLastModifiedDate time.Time,
	endLastModifiedDate time.Time,
) *FindUserFilter {
	return &FindUserFilter{
		name:                  name,
		documentNumber:        documentNumber,
		birthday:              birthday,
		startCreatedDate:      startCreatedDate,
		endCreatedDate:        endCreatedDate,
		startLastModifiedDate: startLastModifiedDate,
		endLastModifiedDate:   endLastModifiedDate,
	}
}

func (f FindUserFilter) GetName() []string {
	return f.name
}

func (f FindUserFilter) GetDocumentNumber() []string {
	return f.documentNumber
}

func (f FindUserFilter) GetBirthday() []time.Time {
	return f.birthday
}

func (f FindUserFilter) GetStartCreatedDate() time.Time {
	return f.startCreatedDate
}

func (f FindUserFilter) GetEndCreatedDate() time.Time {
	return f.endCreatedDate
}

func (f FindUserFilter) GetStartLastModifiedDate() time.Time {
	return f.startLastModifiedDate
}

func (f FindUserFilter) GetEndLastModifiedDate() time.Time {
	return f.endLastModifiedDate
}
