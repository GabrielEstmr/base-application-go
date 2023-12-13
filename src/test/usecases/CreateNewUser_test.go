package test_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	test_mocks "baseapplicationgo/test/mocks"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
	"testing"
	"time"
)

const _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS = "providers.create.user.user.with.given.document.already.exists-DEFAULT"
const _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE = "Document with the given document already exists"
const _MSG_CREATE_NEW_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue"
const _MSG_CREATE_NEW_DOC_ARCH_ISSUE_VALUE = "Architecture application issue"

type testInputs struct {
	name   string
	fields struct {
		userDatabaseGateway main_gateways.UserDatabaseGateway
		logLoki             main_gateways.LogsMonitoringGateway
		messageBeans        main_utils_messages.ApplicationMessages
	}
	args struct {
		ctx  context.Context
		user main_domains.User
	}
	want  main_domains.User
	want1 main_domains_exceptions.ApplicationException
}

func testCreateNewUser_Execute_ShouldSaveNewUser() testInputs {

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

	methodName := make(map[string]test_mocks_support.ArgsMockSupport)
	methodName["FindByDocumentNumber"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, "DocumentNumber").
			AddOutput(1, main_domains.User{}).
			AddOutput(2, nil)
	methodName["Save"] =
		*test_mocks_support.NewArgsMockSupport().
			AddInput(1, user).
			AddOutput(1, userResponse).
			AddOutput(2, nil)

	mock := test_mocks.NewUserDatabaseGatewayMock(*test_mocks_support.NewMethodArgsMockSupport(methodName))

	appMsg := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS: _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE,
	})

	return testInputs{
		name: "",
		fields: struct {
			userDatabaseGateway main_gateways.UserDatabaseGateway
			logLoki             main_gateways.LogsMonitoringGateway
			messageBeans        main_utils_messages.ApplicationMessages
		}{userDatabaseGateway: mock, logLoki: new(test_mocks.LogsMonitoringGateway), messageBeans: appMsg},
		args: struct {
			ctx  context.Context
			user main_domains.User
		}{ctx: context.Background(), user: user},
		want:  userResponse,
		want1: nil,
	}
}

func TestCreateNewUser_Execute(t *testing.T) {

	tests := []testInputs{
		testCreateNewUser_Execute_ShouldSaveNewUser(),
	}

	// ==============

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateNewUserAllArgs(tt.fields.userDatabaseGateway,
				tt.fields.logLoki, *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
					_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS: _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE,
				}))
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

//func TestNewCreateNewUser(t *testing.T) {
//	type args struct {
//		userDatabaseGateway main_gateways.UserDatabaseGateway
//	}
//	tests := []struct {
//		name string
//		args args
//		want *main_usecases.CreateNewUser
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := main_usecases.NewCreateNewUser(tt.args.userDatabaseGateway); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewCreateNewUser() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNewCreateNewUserAllArgs(t *testing.T) {
//	type args struct {
//		userDatabaseGateway main_gateways.UserDatabaseGateway
//		logLoki             main_configs_apm_logs_impl.LogsMethods
//	}
//	tests := []struct {
//		name string
//		args args
//		want *main_usecases.CreateNewUser
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := main_usecases.NewCreateNewUserAllArgs(tt.args.userDatabaseGateway, tt.args.logLoki); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewCreateNewUserAllArgs() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
