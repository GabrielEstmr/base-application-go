package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
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

const _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY = "exceptions.architecture.application.issue-DEFAULT"
const _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE = "Architecture application issue"

func TestCreateEmailFallback_ShouldReturnErrorWhenFindByEventIdFails(t *testing.T) {

	msgId := "msgId"
	emailParams := *new(main_domains.EmailParams)
	args := testCreateEmailArgs{
		ctx:         context.TODO(),
		msgId:       msgId,
		emailParams: emailParams,
	}

	errFind := main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("")

	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindByEventId"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, msgId).
			AddOutput(1, *new(main_domains.Email)).
			AddOutput(2, *errFind)
	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailFields{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: *new(test_mocks.CreateEmailBodySendAndPersistAsSentMock),
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	params := []testCreateEmailInputs{
		{
			name:   "TestCreateEmailFallback_ShouldReturnErrorWhenFindByEventIdFails",
			fields: fields,
			args:   args,
			want:   *new(main_domains.Email),
			want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE),
		},
	}
	createEmailFallback_Execute_RunTests(t, params)
}

func TestCreateEmailFallback_ShouldReturnAppErrorWhenSaveNewEmailFails(t *testing.T) {

	msgId := "msgId"
	emailParams := *new(main_domains.EmailParams)
	args := testCreateEmailArgs{
		ctx:         context.TODO(),
		msgId:       msgId,
		emailParams: emailParams,
	}

	emptyEmail := *new(main_domains.Email)
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindByEventId"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, msgId).
			AddOutput(1, emptyEmail).
			AddOutput(2, nil)

	errSave := main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("")
	emailDatabaseGatewayMethodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)).
			AddOutput(1, *new(main_domains.Email)).
			AddOutput(2, *errSave)

	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailFields{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: *new(test_mocks.CreateEmailBodySendAndPersistAsSentMock),
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	params := []testCreateEmailInputs{
		{
			name:   "TestCreateEmailFallback_ShouldReturnAppErrorWhenSaveNewEmailFails",
			fields: fields,
			args:   args,
			want:   *new(main_domains.Email),
			want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE),
		},
	}
	createEmailFallback_Execute_RunTests(t, params)
}

func TestCreateEmailFallback_ShouldReturnSavedEmailAndErrorWhenProcessEmailFails(t *testing.T) {

	msgId := "msgId"
	emailParams := *new(main_domains.EmailParams)
	args := testCreateEmailArgs{
		ctx:         context.TODO(),
		msgId:       msgId,
		emailParams: emailParams,
	}

	emptyEmail := *new(main_domains.Email)
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindByEventId"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, msgId).
			AddOutput(1, emptyEmail).
			AddOutput(2, nil)

	persistedEmail := *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)
	emailDatabaseGatewayMethodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)).
			AddOutput(1, persistedEmail).
			AddOutput(2, nil)
	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	errApp := main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("")
	createBodyAndPersistAsSentMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createBodyAndPersistAsSentMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, persistedEmail).
			AddOutput(1, *new(main_domains.Email)).
			AddOutput(2, *errApp)
	createBodyAndPersistAsSent := test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createBodyAndPersistAsSentMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailFields{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: createBodyAndPersistAsSent,
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	params := []testCreateEmailInputs{
		{
			name:   "TestCreateEmailFallback_ShouldReturnSavedEmailAndErrorWhenProcessEmailFails",
			fields: fields,
			args:   args,
			want:   persistedEmail,
			want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE),
		},
	}
	createEmailFallback_Execute_RunTests(t, params)
}

func TestCreateEmailFallback_ShouldSaveAndProcessEmail(t *testing.T) {

	msgId := "msgId"
	emailParams := *new(main_domains.EmailParams)
	args := testCreateEmailArgs{
		ctx:         context.TODO(),
		msgId:       msgId,
		emailParams: emailParams,
	}

	emptyEmail := *new(main_domains.Email)
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindByEventId"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, msgId).
			AddOutput(1, emptyEmail).
			AddOutput(2, nil)

	persistedEmail := *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)
	emailDatabaseGatewayMethodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)).
			AddOutput(1, persistedEmail).
			AddOutput(2, nil)
	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	persistedSentEmail := persistedEmail.CloneAsSent()
	createBodyAndPersistAsSentMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createBodyAndPersistAsSentMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, persistedEmail).
			AddOutput(1, persistedSentEmail).
			AddOutput(2, nil)
	createBodyAndPersistAsSent := test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createBodyAndPersistAsSentMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailFields{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: createBodyAndPersistAsSent,
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	params := []testCreateEmailInputs{
		{
			name:   "TestCreateEmailFallback_ShouldSaveAndProcessEmail",
			fields: fields,
			args:   args,
			want:   persistedSentEmail,
			want1:  nil,
		},
	}
	createEmailFallback_Execute_RunTests(t, params)
}

func TestCreateEmailFallback_ShouldGetAndProcessAlreadyPersistedStartedEmail(t *testing.T) {

	msgId := "msgId"
	emailParams := *new(main_domains.EmailParams)
	args := testCreateEmailArgs{
		ctx:         context.TODO(),
		msgId:       msgId,
		emailParams: emailParams,
	}

	persistedEmail := *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindByEventId"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, msgId).
			AddOutput(1, persistedEmail).
			AddOutput(2, nil)
	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	persistedSentEmail := persistedEmail.CloneAsSent()
	createBodyAndPersistAsSentMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createBodyAndPersistAsSentMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, persistedEmail).
			AddOutput(1, persistedSentEmail).
			AddOutput(2, nil)
	createBodyAndPersistAsSent := test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createBodyAndPersistAsSentMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailFields{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: createBodyAndPersistAsSent,
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	params := []testCreateEmailInputs{
		{
			name:   "TestCreateEmailFallback_ShouldGetAndProcessAlreadyPersistedStartedEmail",
			fields: fields,
			args:   args,
			want:   persistedSentEmail,
			want1:  nil,
		},
	}
	createEmailFallback_Execute_RunTests(t, params)
}

func TestCreateEmailFallback_ShouldGetAndProcessAlreadyPersistedIntegrationErrorEmail(t *testing.T) {

	msgId := "msgId"
	emailParams := *new(main_domains.EmailParams)
	args := testCreateEmailArgs{
		ctx:         context.TODO(),
		msgId:       msgId,
		emailParams: emailParams,
	}

	persistedEmail := *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_INTEGRATION_ERROR)
	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindByEventId"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, msgId).
			AddOutput(1, persistedEmail).
			AddOutput(2, nil)
	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	persistedSentEmail := persistedEmail.CloneAsSent()
	createBodyAndPersistAsSentMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	createBodyAndPersistAsSentMethodMocks["Execute"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, persistedEmail).
			AddOutput(1, persistedSentEmail).
			AddOutput(2, nil)
	createBodyAndPersistAsSent := test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createBodyAndPersistAsSentMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailFields{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: createBodyAndPersistAsSent,
		logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                        messageUtilsMock,
	}

	params := []testCreateEmailInputs{
		{
			name:   "TestCreateEmailFallback_ShouldGetAndProcessAlreadyPersistedIntegrationErrorEmail",
			fields: fields,
			args:   args,
			want:   persistedSentEmail,
			want1:  nil,
		},
	}
	createEmailFallback_Execute_RunTests(t, params)
}

func TestCreateEmailFallback_ShouldGetAndNotProcessSentOrErrorEmail(t *testing.T) {

	statuses := []main_domains_enums.EmailStatus{
		main_domains_enums.EMAIL_STATUS_SENT,
		main_domains_enums.EMAIL_STATUS_ERROR,
	}

	for _, v := range statuses {
		msgId := "msgId"
		emailParams := *new(main_domains.EmailParams)
		args := testCreateEmailArgs{
			ctx:         context.TODO(),
			msgId:       msgId,
			emailParams: emailParams,
		}

		persistedEmail := *main_domains.NewEmail(msgId, emailParams, v)
		emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
		emailDatabaseGatewayMethodMocks["FindByEventId"] =
			*test_mocks_support.NewArgsMockSupport().
				AddInput(1, msgId).
				AddOutput(1, persistedEmail).
				AddOutput(2, nil)
		emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

		messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
			_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE_VALUE,
		})

		fields := testCreateEmailFields{
			emailDatabaseGateway:                emailDatabaseGateway,
			createEmailBodySendAndPersistAsSent: new(test_mocks.CreateEmailBodySendAndPersistAsSentMock),
			logsMonitoringGateway:               new(test_mocks.LogsMonitoringGatewayMock),
			spanGateway:                         new(test_mocks.SpanGatewayMockImpl),
			messageUtils:                        messageUtilsMock,
		}

		params := []testCreateEmailInputs{
			{
				name:   "TestCreateEmailFallback_ShouldGetAndProcessAlreadyPersistedIntegrationErrorEmail",
				fields: fields,
				args:   args,
				want:   persistedEmail,
				want1:  nil,
			},
		}
		createEmailFallback_Execute_RunTests(t, params)
	}

}

func createEmailFallback_Execute_RunTests(t *testing.T, params []testCreateEmailInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateEmailFallbackAllArgs(
				tt.fields.emailDatabaseGateway,
				tt.fields.createEmailBodySendAndPersistAsSent,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway,
				tt.fields.messageUtils,
			)
			got, got1 := this.Execute(tt.args.ctx, tt.args.msgId, tt.args.emailParams)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
