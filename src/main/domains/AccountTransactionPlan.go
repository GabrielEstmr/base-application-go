package main_domains

import (
	"time"
)

type AccountTransactionPlan struct {
	id                string
	accountId         string
	planSetting       PlanSetting
	creationUserId    string
	creationUserEmail string
	startDate         time.Time
	endDate           time.Time
	createdDate       time.Time
	lastModifiedDate  time.Time
}
