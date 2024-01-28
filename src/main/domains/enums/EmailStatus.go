package main_domains_enums

type EmailStatus string

const (
	EMAIL_STATUS_STARTED           EmailStatus = "STARTED"
	EMAIL_STATUS_SENT              EmailStatus = "SENT"
	EMAIL_STATUS_ERROR             EmailStatus = "ERROR"
	EMAIL_STATUS_INTEGRATION_ERROR EmailStatus = "EMAIL_STATUS_INTEGRATION_ERROR"
)

var emailStatusEnum = map[EmailStatus]EmailStatus{
	EMAIL_STATUS_STARTED:           EMAIL_STATUS_STARTED,
	EMAIL_STATUS_SENT:              EMAIL_STATUS_SENT,
	EMAIL_STATUS_ERROR:             EMAIL_STATUS_ERROR,
	EMAIL_STATUS_INTEGRATION_ERROR: EMAIL_STATUS_INTEGRATION_ERROR,
}

func ExistsEmailStatus(value EmailStatus) bool {
	_, exists := emailStatusEnum[value]
	return exists
}

func GetEmailStatusDescription(value EmailStatus) string {
	valueMap, exists := emailStatusEnum[value]
	if exists {
		return string(valueMap)
	}
	return ""
}

func GetEmailStatusFromDescription(description string) EmailStatus {
	valueMap, exists := emailStatusEnum[EmailStatus(description)]
	if exists {
		return valueMap
	}
	return ""
}

func GetEmailStatusesFromDescriptions(descriptions []string) []EmailStatus {
	var statues []EmailStatus
	for _, v := range descriptions {
		statues = append(statues, GetEmailStatusFromDescription(v))
	}
	return statues
}
