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

type testReprocessEmailEventInputs struct {
	name   string
	fields testReprocessEmailEventFields
	args   testReprocessEmailEventArgs
	want1  main_domains.Email
	want2  main_domains_exceptions.ApplicationException
}

type testReprocessEmailEventFields struct {
	emailDatabaseGateway                main_gateways.EmailDatabaseGateway
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent
	lockGateway                         main_gateways.DistributedLockGateway
	logsMonitoringGateway               main_gateways.LogsMonitoringGateway
	spanGateway                         main_gateways.SpanGateway
	messageUtils                        main_utils_messages.ApplicationMessages
}

type testReprocessEmailEventArgs struct {
	ctx context.Context
	id  string
}

func TestReprocessEmailEvent_Execute_ShouldReturnNewResourceNotFoundExceptionWhenEmailHasBeenNotFound(t *testing.T) {
	id := "ID"
	emptyEmail := *new(main_domains.Email)
	errApp := *main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("")
	errWant := *main_domains_exceptions.NewResourceNotFoundExceptionSglMsg("")

	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindById"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, id).
			AddOutput(1, emptyEmail).
			AddOutput(2, errApp)
	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
	})

	fields := testReprocessEmailEventFields{
		emailDatabaseGatewayMock,
		new(test_mocks.CreateEmailBodySendAndPersistAsSentMock),
		new(test_mocks.LockGatewayMock),
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl),
		messageUtilsMock,
	}

	args := testReprocessEmailEventArgs{
		context.Background(),
		id,
	}

	params := []testReprocessEmailEventInputs{
		{
			name:   "TestReprocessEmailEvent_Execute_ShouldReturnNewResourceNotFoundExceptionWhenEmailHasBeenNotFound",
			fields: fields,
			args:   args,
			want1:  emptyEmail,
			want2:  errWant,
		},
	}

	testReprocessEmailEvent_Execute_RunTests(t, params)
}

func TestReprocessEmailEvent_Execute_ShouldReturnNilWhenEmailIsEmpty(t *testing.T) {
	id := "ID"
	emptyEmail := *new(main_domains.Email)

	emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailDatabaseGatewayMethodMocks["FindById"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, id).
			AddOutput(1, emptyEmail).
			AddOutput(2, nil)
	emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
	})

	fields := testReprocessEmailEventFields{
		emailDatabaseGatewayMock,
		new(test_mocks.CreateEmailBodySendAndPersistAsSentMock),
		new(test_mocks.LockGatewayMock),
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl),
		messageUtilsMock,
	}

	args := testReprocessEmailEventArgs{
		context.Background(),
		id,
	}

	params := []testReprocessEmailEventInputs{
		{
			name:   "TestReprocessEmailEvent_Execute_ShouldReturnNilWhenEmailIsEmpty",
			fields: fields,
			args:   args,
			want1:  emptyEmail,
			want2:  nil,
		},
	}

	testReprocessEmailEvent_Execute_RunTests(t, params)
}

func TestReprocessEmailEvent_Execute_ShouldNotReprocessEmailEventEmailAsSentOrError(t *testing.T) {

	statuses := []main_domains_enums.EmailStatus{
		main_domains_enums.EMAIL_STATUS_SENT,
		main_domains_enums.EMAIL_STATUS_ERROR,
	}

	for _, v := range statuses {
		id := "ID"
		email := *main_domains.NewEmail("eventId", *new(main_domains.EmailParams), v)
		emptyEmail := *new(main_domains.Email)

		emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
		emailDatabaseGatewayMethodMocks["FindById"] =
			*test_mocks_support.NewArgsMockSupport().
				AddInput(1, id).
				AddOutput(1, email).
				AddOutput(2, nil)
		emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

		messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
			_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
		})

		fields := testReprocessEmailEventFields{
			emailDatabaseGatewayMock,
			new(test_mocks.CreateEmailBodySendAndPersistAsSentMock),
			new(test_mocks.LockGatewayMock),
			new(test_mocks.LogsMonitoringGatewayMock),
			new(test_mocks.SpanGatewayMockImpl),
			messageUtilsMock,
		}

		args := testReprocessEmailEventArgs{
			context.Background(),
			id,
		}

		params := []testReprocessEmailEventInputs{
			{
				name:   "TestReprocessEmailEvent_Execute_ShouldNotReprocessEmailEventEmailAsSentOrError",
				fields: fields,
				args:   args,
				want1:  emptyEmail,
				want2:  nil,
			},
		}

		testReprocessEmailEvent_Execute_RunTests(t, params)
	}

}

func TestReprocessEmailEvent_Execute_ShouldReturnInternalServerErrorExceptionWhenCreateEmailBodyFails(t *testing.T) {

	statuses := []main_domains_enums.EmailStatus{
		main_domains_enums.EMAIL_STATUS_STARTED,
		main_domains_enums.EMAIL_STATUS_INTEGRATION_ERROR,
	}

	for _, v := range statuses {
		id := "ID"
		email := *main_domains.NewEmail("eventId", *new(main_domains.EmailParams), v)

		emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
		emailDatabaseGatewayMethodMocks["FindById"] =
			*test_mocks_support.NewArgsMockSupport().
				AddInput(1, id).
				AddOutput(1, email).
				AddOutput(2, nil)
		emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

		messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
			_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
		})

		errApp := *main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("")
		emptyEmail := *new(main_domains.Email)

		createEmailBodyMockMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
		createEmailBodyMockMethodMocks["Execute"] =
			*test_mocks_support.NewArgsMockSupport().
				AddInput(1, email).
				AddOutput(1, emptyEmail).
				AddOutput(2, errApp)
		createEmailBodyMock := test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createEmailBodyMockMethodMocks)

		fields := testReprocessEmailEventFields{
			emailDatabaseGatewayMock,
			createEmailBodyMock,
			new(test_mocks.LockGatewayMock),
			new(test_mocks.LogsMonitoringGatewayMock),
			new(test_mocks.SpanGatewayMockImpl),
			messageUtilsMock,
		}

		args := testReprocessEmailEventArgs{
			context.Background(),
			id,
		}

		params := []testReprocessEmailEventInputs{
			{
				name:   "TestReprocessEmailEvent_Execute_ShouldReturnInternalServerErrorExceptionWhenCreateEmailBodyFails",
				fields: fields,
				args:   args,
				want1:  *new(main_domains.Email),
				want2:  errApp,
			},
		}

		testReprocessEmailEvent_Execute_RunTests(t, params)
	}

}

func TestReprocessEmailEvent_Execute_ShouldReprocessEmailEventWhenEmailIsAsStartedOrIntegrationError(t *testing.T) {

	statuses := []main_domains_enums.EmailStatus{
		main_domains_enums.EMAIL_STATUS_STARTED,
		main_domains_enums.EMAIL_STATUS_INTEGRATION_ERROR,
	}

	for _, v := range statuses {
		id := "ID"
		email := *main_domains.NewEmail("eventId", *new(main_domains.EmailParams), v)
		updatedEmail := *main_domains.NewEmail("eventId", *new(main_domains.EmailParams), main_domains_enums.EMAIL_STATUS_SENT)

		emailDatabaseGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
		emailDatabaseGatewayMethodMocks["FindById"] =
			*test_mocks_support.NewArgsMockSupport().
				AddInput(1, id).
				AddOutput(1, email).
				AddOutput(2, nil)
		emailDatabaseGatewayMock := test_mocks.NewEmailDatabaseGatewayMock(emailDatabaseGatewayMethodMocks)

		messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
			_MSG_CREATE_NEW_EMAIL_ARCH_ISSUE: _MSG_CREATE_NEW_EMAIL_ARCH_ISSUE_VALUE,
		})

		createEmailBodyMockMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
		createEmailBodyMockMethodMocks["Execute"] =
			*test_mocks_support.NewArgsMockSupport().
				AddInput(1, email).
				AddOutput(1, updatedEmail).
				AddOutput(2, nil)
		createEmailBodyMock := test_mocks.NewCreateEmailBodySendAndPersistAsSentMock(createEmailBodyMockMethodMocks)

		fields := testReprocessEmailEventFields{
			emailDatabaseGatewayMock,
			createEmailBodyMock,
			new(test_mocks.LockGatewayMock),
			new(test_mocks.LogsMonitoringGatewayMock),
			new(test_mocks.SpanGatewayMockImpl),
			messageUtilsMock,
		}

		args := testReprocessEmailEventArgs{
			context.Background(),
			id,
		}

		params := []testReprocessEmailEventInputs{
			{
				name:   "TestReprocessEmailEvent_Execute_ShouldReprocessEmailEventWhenEmailIsAsStartedOrIntegrationError",
				fields: fields,
				args:   args,
				want1:  updatedEmail,
				want2:  nil,
			},
		}

		testReprocessEmailEvent_Execute_RunTests(t, params)
	}

}

func testReprocessEmailEvent_Execute_RunTests(t *testing.T, params []testReprocessEmailEventInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			provider := main_usecases.NewReprocessEmailEventAllArgs(
				tt.fields.emailDatabaseGateway,
				tt.fields.createEmailBodySendAndPersistAsSent,
				tt.fields.lockGateway,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway,
				tt.fields.messageUtils)
			got, got1 := provider.Execute(tt.args.ctx, tt.args.id)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got, tt.want1)
			}
			if !reflect.DeepEqual(got1, tt.want2) {
				t.Errorf("Execute() got2 = %v, want %v", got1, tt.want2)
			}
		})
	}
}
