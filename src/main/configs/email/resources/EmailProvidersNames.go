package main_configs_email_resources

const (
	GMAIL = "GMAIL"
)

func GetEmailProviderValues() []string {
	return []string{GMAIL}
}

func GetEmailProviderMapValues() map[string]string {
	return map[string]string{
		GMAIL: GMAIL,
	}
}
