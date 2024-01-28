package main_usecases_test

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	test_mocks "baseapplicationgo/test/mocks"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
	"testing"
)

type testCreateEmailInputs struct {
	name   string
	fields testCreateEmailFields
	args   testCreateEmailArgs
	want   main_domains.Email
	want1  main_domains_exceptions.ApplicationException
}

type testCreateEmailFields struct {
	emailDatabaseGateway                main_gateways.EmailDatabaseGateway
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent
	logsMonitoringGateway               main_gateways.LogsMonitoringGateway
	spanGateway                         main_gateways.SpanGateway
	messageUtils                        main_utils_messages.ApplicationMessages
}

type testCreateEmailArgs struct {
	ctx         context.Context
	msgId       string
	emailParams main_domains.EmailParams
}

const _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE = "exceptions.architecture.application.issue"
const _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE = "Architecture application issue"

func TestCreateEmail_Execute_ShouldLogAndReturnArchErrorWHenErrorOnSave(t *testing.T) {

	emailParams := *new(main_domains.EmailParams)

	email := *main_domains.NewEmail(
		"eventId",
		emailParams,
		main_domains_enums.EMAIL_STATUS_STARTED,
	)

	err := *main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("err")

	methodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	methodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email).
			AddOutput(1, email).
			AddOutput(2, err)

	var emailDatabaseGateway main_gateways.EmailDatabaseGateway = test_mocks.NewEmailDatabaseGatewayMock(methodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
	})

	providerFields := testCreateEmailFields{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: new(test_mocks.CreateEmailBodySendAndPersistAsSentMock),
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	args := testCreateEmailArgs{
		ctx:         context.Background(),
		msgId:       "eventId",
		emailParams: emailParams,
	}

	tests := []testCreateEmailInputs{
		{
			name:   "TestCreateEmail_Execute_ShouldLogAndReturnArchErrorWHenErrorOnSave",
			fields: providerFields,
			args:   args,
			want:   main_domains.Email{},
			want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(""),
		},
	}

	runExecuteTestCases(t, tests)
}

func TestCreateEmail_Execute_ShouldLogAndReturnArchErrorWhenCreateEmailBodySendAndPersistAsSentReturnsError(t *testing.T) {

	emailParams := *new(main_domains.EmailParams)

	email := *main_domains.NewEmail(
		"eventId",
		emailParams,
		main_domains_enums.EMAIL_STATUS_STARTED,
	)

	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email).
			AddOutput(1, email).
			AddOutput(2, nil)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
	})

	err := *main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("err")
	createBodyAndSentMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createBodyAndSentMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email).
			AddOutput(1, email).
			AddOutput(2, err)

	providerFields := testCreateEmailFields{
		emailDatabaseGateway:                test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks),
		createEmailBodySendAndPersistAsSent: test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createBodyAndSentMethodMocks),
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	args := testCreateEmailArgs{
		ctx:         context.Background(),
		msgId:       "eventId",
		emailParams: emailParams,
	}

	tests := []testCreateEmailInputs{
		{
			name:   "TestCreateEmail_Execute_ShouldLogAndReturnArchErrorWhenCreateEmailBodySendAndPersistAsSentReturnsError",
			fields: providerFields,
			args:   args,
			want:   main_domains.Email{},
			want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(""),
		},
	}

	runExecuteTestCases(t, tests)
}

func TestCreateEmail_Execute_ShouldCreateEmail(t *testing.T) {

	emailParams := *new(main_domains.EmailParams)

	email := *main_domains.NewEmail(
		"eventId",
		emailParams,
		main_domains_enums.EMAIL_STATUS_STARTED,
	)

	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email).
			AddOutput(1, email).
			AddOutput(2, nil)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
	})

	createBodyAndSentMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createBodyAndSentMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, email).
			AddOutput(1, email).
			AddOutput(2, nil)

	providerFields := testCreateEmailFields{
		emailDatabaseGateway:                test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks),
		createEmailBodySendAndPersistAsSent: test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createBodyAndSentMethodMocks),
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	args := testCreateEmailArgs{
		ctx:         context.Background(),
		msgId:       "eventId",
		emailParams: emailParams,
	}

	tests := []testCreateEmailInputs{
		{
			name:   "TestCreateEmail_Execute_ShouldLogAndReturnArchErrorWhenCreateEmailBodySendAndPersistAsSentReturnsError",
			fields: providerFields,
			args:   args,
			want:   email,
			want1:  nil,
		},
	}

	runExecuteTestCases(t, tests)
}

func runExecuteTestCases(t *testing.T, params []testCreateEmailInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			provider := main_usecases.NewCreateEmailAllArgs(
				tt.fields.emailDatabaseGateway,
				tt.fields.createEmailBodySendAndPersistAsSent,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway,
				tt.fields.messageUtils,
			)
			got, got1 := provider.Execute(tt.args.ctx, tt.args.msgId, tt.args.emailParams)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
