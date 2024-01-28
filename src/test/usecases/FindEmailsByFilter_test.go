package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	test_mocks "baseapplicationgo/test/mocks"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
	"testing"
)

const _MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_KEY = "exceptions.architecture.application.issue-DEFAULT"
const _MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_VALUE = "Architecture application issue"

type testFindEmailsByFilterInputs struct {
	name   string
	fields testFindEmailsByFilterFields
	args   testFindEmailsByFilterArgs
	want   main_domains.Page
	want1  main_domains_exceptions.ApplicationException
}

type testFindEmailsByFilterFields struct {
	emailDatabaseGateway  main_gateways.EmailDatabaseGateway
	messageUtils          main_utils_messages.ApplicationMessages
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

type testFindEmailsByFilterArgs struct {
	ctx      context.Context
	filter   main_domains.FindEmailFilter
	pageable main_domains.Pageable
}

func TestFindEmailsByFilter_ShouldReturnInternalServerErrorExceptionWhenGatewayFails(t *testing.T) {

	filter := main_domains.FindEmailFilter{}.
		WithIds([]string{"1", "2"}).
		WithStatuses([]main_domains_enums.EmailStatus{main_domains_enums.EMAIL_STATUS_STARTED})
	pageable := *main_domains.NewPageable(0, 20, nil)
	errApp := *main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("Gateway Err")

	emailGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayMethodMocks["FindByFilter"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, filter).
			AddInput(2, pageable).
			AddOutput(1, *new(main_domains.Page)).
			AddOutput(2, errApp)
	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailGatewayMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_KEY: _MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_VALUE,
	})

	fields := testFindEmailsByFilterFields{
		emailDatabaseGateway:  emailDatabaseGateway,
		messageUtils:          messageUtilsMock,
		spanGateway:           new(test_mocks.SpanGatewayMockImpl),
		logsMonitoringGateway: new(test_mocks.LogsMonitoringGatewayMock),
	}

	args := testFindEmailsByFilterArgs{
		context.Background(),
		filter,
		pageable,
	}

	params := []testFindEmailsByFilterInputs{
		{
			name:   "TestFindEmailsByFilter_ShouldReturnInternalServerErrorExceptionWhenGatewayFails",
			fields: fields,
			args:   args,
			want:   *new(main_domains.Page),
			want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_VALUE),
		},
	}

	testFindEmailsByFilter_RunTests(t, params)
}

func TestFindEmailsByFilter_ShouldReturnPageOfEmails(t *testing.T) {

	filter := main_domains.FindEmailFilter{}.
		WithIds([]string{"1", "2"}).
		WithStatuses([]main_domains_enums.EmailStatus{main_domains_enums.EMAIL_STATUS_STARTED})
	pageable := *main_domains.NewPageable(0, 20, nil)

	email := *main_domains.NewEmail("", *new(main_domains.EmailParams), main_domains_enums.EMAIL_STATUS_STARTED)

	page := *main_domains.NewPage(
		[]any{email},
		0,
		20,
		200,
	)

	emailGatewayMethodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	emailGatewayMethodMocks["FindByFilter"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, filter).
			AddInput(2, pageable).
			AddOutput(1, page).
			AddOutput(2, nil)
	emailDatabaseGateway := test_mocks.NewEmailDatabaseGatewayMock(emailGatewayMethodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_KEY: _MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_VALUE,
	})

	fields := testFindEmailsByFilterFields{
		emailDatabaseGateway:  emailDatabaseGateway,
		messageUtils:          messageUtilsMock,
		spanGateway:           new(test_mocks.SpanGatewayMockImpl),
		logsMonitoringGateway: new(test_mocks.LogsMonitoringGatewayMock),
	}

	args := testFindEmailsByFilterArgs{
		context.Background(),
		filter,
		pageable,
	}

	params := []testFindEmailsByFilterInputs{
		{
			name:   "TestFindEmailsByFilter_ShouldReturnInternalServerErrorExceptionWhenGatewayFails",
			fields: fields,
			args:   args,
			want:   page,
			want1:  nil,
		},
	}

	testFindEmailsByFilter_RunTests(t, params)
}

func testFindEmailsByFilter_RunTests(t *testing.T, params []testFindEmailsByFilterInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewFindEmailsByFilter(
				tt.fields.emailDatabaseGateway,
				tt.fields.messageUtils,
				tt.fields.spanGateway,
				tt.fields.logsMonitoringGateway,
			)
			got, got1 := this.Execute(tt.args.ctx, tt.args.filter, tt.args.pageable)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
