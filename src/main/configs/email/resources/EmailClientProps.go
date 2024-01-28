package main_configs_email_resources

type EmailClientProps struct {
	providerName string
	email        string
	password     string
}

func NewEmailClientProps(
	providerName string,
	email string,
	password string,
) *EmailClientProps {
	return &EmailClientProps{
		providerName: providerName,
		email:        email,
		password:     password,
	}
}

func (e *EmailClientProps) GetProviderName() string {
	return e.providerName
}

func (e *EmailClientProps) GetEmail() string {
	return e.email
}

func (e *EmailClientProps) GetPassword() string {
	return e.password
}
