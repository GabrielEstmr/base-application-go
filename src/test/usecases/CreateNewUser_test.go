package main_usecases_test

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	test_mocks "baseapplicationgo/test/mocks"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

const _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS = "providers.create.user.user.with.given.document.already.exists-DEFAULT"
const _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE = "Document with the given document already exists"
const _MSG_CREATE_NEW_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue-DEFAULT"
const _MSG_CREATE_NEW_DOC_ARCH_ISSUE_VALUE = "Architecture application issue"

type testCreateNewUserInputs struct {
	name   string
	fields testCreateNewUserFields
	args   testCreateNewUserArgs
	want   main_domains.User
	want1  main_domains_exceptions.ApplicationException
}

type testCreateNewUserFields struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	logLoki             main_gateways.LogsMonitoringGateway
	spanGateway         main_gateways.SpanGateway
	messageUtils        main_utils_messages.ApplicationMessages
}

type testCreateNewUserArgs struct {
	ctx  context.Context
	user main_domains.User
}

func TestCreateNewUser_Execute_ShouldSaveNewUser(t *testing.T) {

	user := main_domains.User{
		Name:           "Name",
		DocumentNumber: "DocumentNumber",
		Birthday:       time.Now(),
	}

	userResponse := main_domains.User{
		Id:               "",
		Name:             "Name",
		DocumentNumber:   "DocumentNumber",
		Birthday:         time.Now(),
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}

	methodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	methodMocks["FindByDocumentNumber"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, user.DocumentNumber).
			AddOutput(1, main_domains.User{}).
			AddOutput(2, nil)
	methodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, user).
			AddOutput(1, userResponse).
			AddOutput(2, nil)

	mock := test_mocks.NewUserDatabaseGatewayMock(methodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS: _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE,
	})

	testCreateNewUserFields := testCreateNewUserFields{
		mock,
		new(test_mocks.LogsMonitoringGatewayMock),
		new(test_mocks.SpanGatewayMockImpl),
		messageUtilsMock,
	}

	args := testCreateNewUserArgs{
		context.Background(), user,
	}

	parameters := []testCreateNewUserInputs{testCreateNewUserInputs{
		name:   "testCreateNewUser_Execute_ShouldSaveNewUser",
		fields: testCreateNewUserFields,
		args:   args,
		want:   userResponse,
		want1:  nil,
	}}
	runTestCreateNewUser_Execute(t, parameters)
}

func TestCreateNewUser_Execute_ShouldReturnInternalServerErrorWhenFindByDocNumberReturnsError(t *testing.T) {

	user := main_domains.User{
		Name:           "Name",
		DocumentNumber: "DocumentNumber",
		Birthday:       time.Now(),
	}

	methodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	methodMocks["FindByDocumentNumber"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, user.DocumentNumber).
			AddOutput(1, *new(main_domains.User)).
			AddOutput(2, errors.New("generic Error"))

	mock := test_mocks.NewUserDatabaseGatewayMock(methodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_DOC_ARCH_ISSUE: _MSG_CREATE_NEW_DOC_ARCH_ISSUE_VALUE,
	})

	testCreateNewUserFields := testCreateNewUserFields{
		mock, new(test_mocks.LogsMonitoringGatewayMock), new(test_mocks.SpanGatewayMockImpl), messageUtilsMock,
	}

	args := testCreateNewUserArgs{
		context.Background(), user,
	}

	parameters := []testCreateNewUserInputs{testCreateNewUserInputs{
		name:   "testCreateNewUser_Execute_ShouldReturnInternalServerErrorWhenFindByDocNumberReturnsError",
		fields: testCreateNewUserFields,
		args:   args,
		want:   *new(main_domains.User),
		want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_NEW_DOC_ARCH_ISSUE_VALUE),
	}}
	runTestCreateNewUser_Execute(t, parameters)
}

func TestCreateNewUser_Execute_ShouldReturnConflictExceptionWhenExistsDocWithSameDocNumber(t *testing.T) {

	user := main_domains.User{
		Name:           "Name",
		DocumentNumber: "DocumentNumber",
		Birthday:       time.Now(),
	}

	userResponse := main_domains.User{
		Id:               "",
		Name:             "Name",
		DocumentNumber:   "DocumentNumber",
		Birthday:         time.Now(),
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}

	methodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	methodMocks["FindByDocumentNumber"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, user.DocumentNumber).
			AddOutput(1, userResponse).
			AddOutput(2, nil)

	mock := test_mocks.NewUserDatabaseGatewayMock(methodMocks)

	messageUtils := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS: _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE,
	})

	testCreateNewUserFields := testCreateNewUserFields{
		mock, new(test_mocks.LogsMonitoringGatewayMock), new(test_mocks.SpanGatewayMockImpl), messageUtils,
	}

	args := testCreateNewUserArgs{
		context.Background(), user,
	}

	parameters := []testCreateNewUserInputs{testCreateNewUserInputs{
		name:   "TestCreateNewUser_Execute_ShouldReturnConflictExceptionWhenExistsDocWithSameDocNumber",
		fields: testCreateNewUserFields,
		args:   args,
		want:   *new(main_domains.User),
		want1:  main_domains_exceptions.NewConflictExceptionSglMsg(_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE),
	}}
	runTestCreateNewUser_Execute(t, parameters)
}

func TestCreateNewUser_Execute_ShouldReturnInternalServerErrorWhenSaveReturnsError(t *testing.T) {

	user := main_domains.User{
		Name:           "Name",
		DocumentNumber: "DocumentNumber",
		Birthday:       time.Now(),
	}

	methodMocks := make(map[string]test_mocks_support.ArgsMockSupport)
	methodMocks["FindByDocumentNumber"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, user.DocumentNumber).
			AddOutput(1, *new(main_domains.User)).
			AddOutput(2, nil)
	methodMocks["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, user).
			AddOutput(1, *new(main_domains.User)).
			AddOutput(2, errors.New("save error"))

	mock := test_mocks.NewUserDatabaseGatewayMock(methodMocks)

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_DOC_ARCH_ISSUE: _MSG_CREATE_NEW_DOC_ARCH_ISSUE_VALUE,
	})

	testCreateNewUserFields := testCreateNewUserFields{
		mock, new(test_mocks.LogsMonitoringGatewayMock), new(test_mocks.SpanGatewayMockImpl), messageUtilsMock,
	}

	args := testCreateNewUserArgs{
		context.Background(), user,
	}

	parameters := []testCreateNewUserInputs{testCreateNewUserInputs{
		name:   "testCreateNewUser_Execute_ShouldReturnInternalServerErrorWhenFindByDocNumberReturnsError",
		fields: testCreateNewUserFields,
		args:   args,
		want:   *new(main_domains.User),
		want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_NEW_DOC_ARCH_ISSUE_VALUE),
	}}
	runTestCreateNewUser_Execute(t, parameters)
}

func runTestCreateNewUser_Execute(t *testing.T, parameters []testCreateNewUserInputs) {
	for _, tt := range parameters {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateNewUserAllArgs(tt.fields.userDatabaseGateway,
				tt.fields.logLoki, tt.fields.spanGateway, tt.fields.messageUtils)
			got, got1 := this.Execute(tt.args.ctx, tt.args.user)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewCreateNewUser(t *testing.T) {
	type testsArgs struct {
		name                  string
		userDatabaseGateway   main_gateways.UserDatabaseGateway
		logsMonitoringGateway main_gateways.LogsMonitoringGateway
		spanGateway           main_gateways.SpanGateway
		want                  *main_usecases.CreateNewUser
	}

	tests := []testsArgs{
		testsArgs{
			"TestNewCreateNewUser",
			new(test_mocks.UserDatabaseGatewayMock),
			new(test_mocks.LogsMonitoringGatewayMock),
			new(test_mocks.SpanGatewayMockImpl),
			main_usecases.NewCreateNewUserAllArgs(
				new(test_mocks.UserDatabaseGatewayMock),
				new(test_mocks.LogsMonitoringGatewayMock),
				new(test_mocks.SpanGatewayMockImpl),
				*main_utils_messages.NewApplicationMessages()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main_usecases.NewCreateNewUser(
				tt.userDatabaseGateway, tt.logsMonitoringGateway, tt.spanGateway); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateNewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
