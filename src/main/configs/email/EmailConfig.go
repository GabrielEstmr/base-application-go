package main_configs_email

import (
	main_configs_email_resources "baseapplicationgo/main/configs/email/resources"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"log/slog"
	"sync"
)

const _MSG_EMAIL_BEAN_INITIALIZING = "Initializing email clients"
const _MSG_EMAIL_BEAN_FINISHED = "Email clients successfully initialized"

var once sync.Once
var emailClientsProps *map[string]main_configs_email_resources.EmailClientProps

func GetEmailConfigsBean() *map[string]main_configs_email_resources.EmailClientProps {

	once.Do(func() {
		if emailClientsProps == nil {
			emailClientsProps = getEmailConfigs()
		}
	})
	return emailClientsProps
}

func getEmailConfigs() *map[string]main_configs_email_resources.EmailClientProps {
	slog.Info(_MSG_EMAIL_BEAN_INITIALIZING)

	props := make(map[string]main_configs_email_resources.EmailClientProps)

	gmailProps := getGmailCredentials()
	props[gmailProps.GetProviderName()] = gmailProps

	slog.Info(_MSG_EMAIL_BEAN_FINISHED)
	return &props
}

func getGmailCredentials() main_configs_email_resources.EmailClientProps {
	gmailClientEmail := main_configs_yml.GetYmlValueByName(main_configs_yml.EmailGmailCredentialsEmail)
	gmailClientPassword := main_configs_yml.GetYmlValueByName(main_configs_yml.EmailGmailCredentialsPassword)
	return *main_configs_email_resources.NewEmailClientProps(main_configs_email_resources.GMAIL, gmailClientEmail, gmailClientPassword)
}
