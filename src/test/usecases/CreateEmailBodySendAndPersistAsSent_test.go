package main_usecases_test

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
	main_usecases_factories_interfaces "baseapplicationgo/main/usecases/factories/interfaces"
	test_mocks "baseapplicationgo/test/mocks"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
	"testing"
)

type testCreateEmailBodySendAndPersistAsSentInputs struct {
	name   string
	fields testCreateEmailBodySendAndPersistAsSentFields
	args   testCreateEmailBodySendAndPersistAsSentArgs
	want   main_domains.Email
	want1  main_domains_exceptions.ApplicationException
}

type testCreateEmailBodySendAndPersistAsSentFields struct {
	emailDatabaseGateway    main_gateways.EmailDatabaseGateway
	sendEmailGatewayFactory main_usecases_factories_interfaces.SendEmailGatewayFactory
	createEmailBodyFactory  main_usecases_factories_interfaces.CreateEmailBodyFactory
	logsMonitoringGateway   main_gateways.LogsMonitoringGateway
	spanGateway             main_gateways.SpanGateway
}

type testCreateEmailBodySendAndPersistAsSentArgs struct {
	ctx   context.Context
	email main_domains.Email
}

func TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionAndUpdateEmailAsErrorWhenEmailParamsIsEmpty(t *testing.T) {

	emailEmptyParams := *main_domains.NewEmail("", *new(main_domains.EmailParams), main_domains_enums.EMAIL_STATUS_STARTED)
	emptyEmail := *new(main_domains.Email)
	emailUpdatedAsError := emptyEmail.CloneAsError("empty email params")
	errApp := main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("empty email params")

	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Update"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, emailUpdatedAsError).
			AddOutput(1, emailUpdatedAsError).
			AddOutput(2, nil)

	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	fields := testCreateEmailBodySendAndPersistAsSentFields{
		emailDatabaseGatewayMock,
		new(test_mocks.SendEmailGatewayFactoryMock),
		new(test_mocks.CreateEmailBodyFactoryMock),
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl)}

	args := testCreateEmailBodySendAndPersistAsSentArgs{
		context.Background(),
		emailEmptyParams,
	}

	params := []testCreateEmailBodySendAndPersistAsSentInputs{
		{
			"TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionAndUpdateEmailAsErrorWhenEmailParamsIsEmpty",
			fields,
			args,
			emailUpdatedAsError,
			errApp,
		},
	}

	createEmailBodySendAndPersistAsSent_Execute_RunTests(t, params)
}

func TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionAndReturnEmptyEmailWhenUpdatedFails(t *testing.T) {

	emailEmptyParams := *main_domains.NewEmail("", *new(main_domains.EmailParams), main_domains_enums.EMAIL_STATUS_STARTED)
	emptyEmail := *new(main_domains.Email)
	emailUpdatedAsError := emptyEmail.CloneAsError("empty email params")
	errUpdateEmail := *main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("")
	errApp := main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("empty email params")

	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Update"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, emailUpdatedAsError).
			AddOutput(1, emailUpdatedAsError).
			AddOutput(2, errUpdateEmail)

	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	fields := testCreateEmailBodySendAndPersistAsSentFields{
		emailDatabaseGatewayMock,
		new(test_mocks.SendEmailGatewayFactoryMock),
		new(test_mocks.CreateEmailBodyFactoryMock),
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl)}

	args := testCreateEmailBodySendAndPersistAsSentArgs{
		context.Background(),
		emailEmptyParams,
	}

	params := []testCreateEmailBodySendAndPersistAsSentInputs{
		{
			"TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionAndReturnEmptyEmailWhenUpdatedFails",
			fields,
			args,
			emailEmptyParams,
			errApp,
		},
	}

	createEmailBodySendAndPersistAsSent_Execute_RunTests(t, params)
}

func TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenCreateEmailBodyFails(t *testing.T) {

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
		"",
		"",
		[]string{""},
		"",
		nil,
	)
	email := *main_domains.NewEmail("", emailParams, main_domains_enums.EMAIL_STATUS_STARTED)

	body := []byte("Email")
	errApp := main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("empty email params")
	createEmailBodyMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams()).
			AddOutput(1, body).
			AddOutput(2, *errApp)
	createEmailBody := test_mocks.NewCreateEmailBodyMock(createEmailBodyMethodMocks)

	createEmailBodyFactoryMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyFactoryMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, createEmailBody)
	createEmailBodyFac := test_mocks.NewCreateEmailBodyFactoryMock(createEmailBodyFactoryMethodMocks)

	emailUpdatesAsError := email.CloneAsError("empty email params")
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Update"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, emailUpdatesAsError).
			AddOutput(1, emailUpdatesAsError).
			AddOutput(2, nil)
	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	fields := testCreateEmailBodySendAndPersistAsSentFields{
		emailDatabaseGatewayMock,
		new(test_mocks.SendEmailGatewayFactoryMock),
		createEmailBodyFac,
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl)}

	args := testCreateEmailBodySendAndPersistAsSentArgs{
		context.Background(),
		email,
	}

	params := []testCreateEmailBodySendAndPersistAsSentInputs{
		{
			"TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenCreateEmailBodyFails",
			fields,
			args,
			emailUpdatesAsError,
			errApp,
		},
	}

	createEmailBodySendAndPersistAsSent_Execute_RunTests(t, params)
}

func TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenSendAndUpdateEmailFails(t *testing.T) {

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
		"",
		"",
		[]string{""},
		"",
		nil,
	)
	email := *main_domains.NewEmail("", emailParams, main_domains_enums.EMAIL_STATUS_STARTED)

	body := []byte("Email")
	createEmailBodyMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams()).
			AddOutput(1, body).
			AddOutput(2, nil)
	createEmailBody := test_mocks.NewCreateEmailBodyMock(createEmailBodyMethodMocks)

	createEmailBodyFactoryMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyFactoryMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, createEmailBody)
	createEmailBodyFac := test_mocks.NewCreateEmailBodyFactoryMock(createEmailBodyFactoryMethodMocks)

	errSend := main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("errSend")
	emailGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayMethodMocks["SendMail"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetTo()).
			AddInput(2, body).
			AddOutput(2, *errSend)
	emailGateway := test_mocks.NewEmailGatewayMock(emailGatewayMethodMocks)

	emailGatewayFacMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayFacMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, emailGateway)
	emailGatewayFactory := test_mocks.NewSendEmailGatewayFactoryMock(emailGatewayFacMethodMocks)

	errUpdate := main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("errUpdate")
	emptyEmail := *new(main_domains.Email)
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Update"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.CloneAsIntegrationError(errSend.Error())).
			AddOutput(1, emptyEmail).
			AddOutput(2, *errUpdate)
	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	fields := testCreateEmailBodySendAndPersistAsSentFields{
		emailDatabaseGatewayMock,
		emailGatewayFactory,
		createEmailBodyFac,
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl)}

	args := testCreateEmailBodySendAndPersistAsSentArgs{
		context.Background(),
		email,
	}

	params := []testCreateEmailBodySendAndPersistAsSentInputs{
		{
			"TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenSendAndUpdateEmailFails",
			fields,
			args,
			email,
			main_domains_exceptions.
				NewInternalServerErrorExceptionSglMsg(errSend.Error()),
		},
	}

	createEmailBodySendAndPersistAsSent_Execute_RunTests(t, params)
}

func TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenSendEmailFails(t *testing.T) {

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
		"",
		"",
		[]string{""},
		"",
		nil,
	)
	email := *main_domains.NewEmail("", emailParams, main_domains_enums.EMAIL_STATUS_STARTED)

	body := []byte("Email")
	createEmailBodyMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams()).
			AddOutput(1, body).
			AddOutput(2, nil)
	createEmailBody := test_mocks.NewCreateEmailBodyMock(createEmailBodyMethodMocks)

	createEmailBodyFactoryMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyFactoryMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, createEmailBody)
	createEmailBodyFac := test_mocks.NewCreateEmailBodyFactoryMock(createEmailBodyFactoryMethodMocks)

	errSend := main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("errSend")
	emailGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayMethodMocks["SendMail"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetTo()).
			AddInput(2, body).
			AddOutput(2, *errSend)
	emailGateway := test_mocks.NewEmailGatewayMock(emailGatewayMethodMocks)

	emailGatewayFacMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayFacMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, emailGateway)
	emailGatewayFactory := test_mocks.NewSendEmailGatewayFactoryMock(emailGatewayFacMethodMocks)

	updatedEmailAsIntErr := email.CloneAsIntegrationError(errSend.Error())
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Update"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, updatedEmailAsIntErr).
			AddOutput(1, updatedEmailAsIntErr).
			AddOutput(2, nil)
	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	fields := testCreateEmailBodySendAndPersistAsSentFields{
		emailDatabaseGatewayMock,
		emailGatewayFactory,
		createEmailBodyFac,
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl)}

	args := testCreateEmailBodySendAndPersistAsSentArgs{
		context.Background(),
		email,
	}

	params := []testCreateEmailBodySendAndPersistAsSentInputs{
		{
			"TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenSendEmailFails",
			fields,
			args,
			updatedEmailAsIntErr,
			errSend,
		},
	}

	createEmailBodySendAndPersistAsSent_Execute_RunTests(t, params)
}

func TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenUpdatedEmailAsSentFails(t *testing.T) {

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
		"",
		"",
		[]string{""},
		"",
		nil,
	)
	email := *main_domains.NewEmail("", emailParams, main_domains_enums.EMAIL_STATUS_STARTED)

	body := []byte("Email")
	createEmailBodyMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams()).
			AddOutput(1, body).
			AddOutput(2, nil)
	createEmailBody := test_mocks.NewCreateEmailBodyMock(createEmailBodyMethodMocks)

	createEmailBodyFactoryMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyFactoryMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, createEmailBody)
	createEmailBodyFac := test_mocks.NewCreateEmailBodyFactoryMock(createEmailBodyFactoryMethodMocks)

	emailGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayMethodMocks["SendMail"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetTo()).
			AddInput(2, body).
			AddOutput(2, nil)
	emailGateway := test_mocks.NewEmailGatewayMock(emailGatewayMethodMocks)

	emailGatewayFacMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayFacMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, emailGateway)
	emailGatewayFactory := test_mocks.NewSendEmailGatewayFactoryMock(emailGatewayFacMethodMocks)

	updatedEmailAsSent := email.CloneAsSent()
	errUpdateSent := main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg("errUpdate")
	emptyEmail := *new(main_domains.Email)
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Update"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, updatedEmailAsSent).
			AddOutput(1, emptyEmail).
			AddOutput(2, *errUpdateSent)
	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	fields := testCreateEmailBodySendAndPersistAsSentFields{
		emailDatabaseGatewayMock,
		emailGatewayFactory,
		createEmailBodyFac,
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl)}

	args := testCreateEmailBodySendAndPersistAsSentArgs{
		context.Background(),
		email,
	}

	params := []testCreateEmailBodySendAndPersistAsSentInputs{
		{
			"TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldReturnInternalServerErrorExceptionWhenUpdatedEmailAsSentFails",
			fields,
			args,
			email,
			main_domains_exceptions.
				NewInternalServerErrorExceptionSglMsg(errUpdateSent.Error()),
		},
	}

	createEmailBodySendAndPersistAsSent_Execute_RunTests(t, params)
}

func TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldCreateEmailBodySendAndPersistAsSent(t *testing.T) {

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
		"",
		"",
		[]string{""},
		"",
		nil,
	)
	email := *main_domains.NewEmail("", emailParams, main_domains_enums.EMAIL_STATUS_STARTED)

	body := []byte("Email")
	createEmailBodyMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams()).
			AddOutput(1, body).
			AddOutput(2, nil)
	createEmailBody := test_mocks.NewCreateEmailBodyMock(createEmailBodyMethodMocks)

	createEmailBodyFactoryMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createEmailBodyFactoryMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, createEmailBody)
	createEmailBodyFac := test_mocks.NewCreateEmailBodyFactoryMock(createEmailBodyFactoryMethodMocks)

	emailGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayMethodMocks["SendMail"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetTo()).
			AddInput(2, body).
			AddOutput(2, nil)
	emailGateway := test_mocks.NewEmailGatewayMock(emailGatewayMethodMocks)

	emailGatewayFacMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayFacMethodMocks["Get"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email.GetEmailParams().GetEmailTemplateType()).
			AddOutput(1, emailGateway)
	emailGatewayFactory := test_mocks.NewSendEmailGatewayFactoryMock(emailGatewayFacMethodMocks)

	updatedEmailAsSent := email.CloneAsSent()
	updatedEmailAsSentPersisted := email.CloneAsSent()
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Update"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, updatedEmailAsSent).
			AddOutput(1, updatedEmailAsSentPersisted).
			AddOutput(2, nil)
	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	fields := testCreateEmailBodySendAndPersistAsSentFields{
		emailDatabaseGatewayMock,
		emailGatewayFactory,
		createEmailBodyFac,
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl)}

	args := testCreateEmailBodySendAndPersistAsSentArgs{
		context.Background(),
		email,
	}

	params := []testCreateEmailBodySendAndPersistAsSentInputs{
		{
			"TestCreateEmailBodySendAndPersistAsSent_Execute_ShouldCreateEmailBodySendAndPersistAsSent",
			fields,
			args,
			updatedEmailAsSentPersisted,
			nil,
		},
	}

	createEmailBodySendAndPersistAsSent_Execute_RunTests(t, params)
}

func createEmailBodySendAndPersistAsSent_Execute_RunTests(t *testing.T, params []testCreateEmailBodySendAndPersistAsSentInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateEmailBodySendAndPersistAsSentAllArgs(
				tt.fields.emailDatabaseGateway,
				tt.fields.sendEmailGatewayFactory,
				tt.fields.createEmailBodyFactory,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway,
			)
			got, got1 := this.Execute(tt.args.ctx, tt.args.email)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
