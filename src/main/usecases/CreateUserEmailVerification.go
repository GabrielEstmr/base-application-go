package main_usecases

import (
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type CreateUserEmailVerification struct {
	generateEmailVerificationCode main_usecases_interfaces.GenerateEmailVerificationCode
	lockGateway                   main_gateways.DistributedLockGateway
	logsMonitoringGateway         main_gateways.LogsMonitoringGateway
	spanGateway                   main_gateways.SpanGateway
	messageUtils                  main_utils_messages.ApplicationMessages
}

func NewCreateUserEmailVerification(
	generateEmailVerificationCode main_usecases_interfaces.GenerateEmailVerificationCode,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *CreateUserEmailVerification {
	return &CreateUserEmailVerification{
		generateEmailVerificationCode: generateEmailVerificationCode,
		lockGateway:                   lockGateway,
		logsMonitoringGateway:         logsMonitoringGateway,
		spanGateway:                   spanGateway,
		messageUtils:                  messageUtils,
	}
}

func (this *CreateUserEmailVerification) Execute(
	ctx context.Context,
	user main_domains.User,
	scope main_domains_enums.UserEmailVerificationScope,
) main_domains.UserEmailVerification {

	span := this.spanGateway.Get(ctx, "CreateUserVerificationEmail-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating UserVerificationEmail for user id: %s", user.GetId()))

	verificationCode := this.generateEmailVerificationCode.Execute()

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_VERIFICATION_USER,
		main_configs_yml.GetYmlValueByName(main_configs_yml.ApmServerName),
		user.GetId(),
		[]string{user.GetEmail()},
		"Verification Email",
		map[string]string{
			"Name":    "PEPETA TEST",
			"Message": verificationCode,
		},
	)

	return *main_domains.NewUserEmailVerification(
		user.GetId(),
		verificationCode,
		scope,
		emailParams,
		main_domains_enums.EMAIL_VERIFICATION_CREATED,
	)
}
