package test_usecases

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_configs_messages_resources "baseapplicationgo/main/configs/messages/resources"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
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

type fields struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	logLoki             main_configs_apm_logs_impl.LogsGateway
}
type args struct {
	ctx  context.Context
	user main_domains.User
}
type testInputs struct {
	name   string
	fields fields
	args   args
	want   main_domains.User
	want1  main_domains_exceptions.ApplicationException
}

func TestCreateNewUser_Execute(t *testing.T) {

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

	imput1 := testInputs{
		name: "Test Save Successfully",
		fields: fields{
			userDatabaseGateway: mock,
			logLoki:             new(test_mocks.LogsGatewayImplMock),
		},
		args: args{
			ctx:  context.Background(),
			user: user,
		},
		want:  userResponse,
		want1: nil,
	}

	tests := []testInputs{imput1}

	runTest(t, tests)
}

func TestCreateNewUser_Execute_Error_DocumentNumber_Already_Exists(t *testing.T) {

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
			AddOutput(1, userResponse).
			AddOutput(2, nil)

	mock := test_mocks.NewUserDatabaseGatewayMock(*test_mocks_support.NewMethodArgsMockSupport(methodName))

	imput1 := testInputs{
		name: "Test Save Successfully",
		fields: fields{
			userDatabaseGateway: mock,
			logLoki:             new(test_mocks.LogsGatewayImplMock),
		},
		args: args{
			ctx:  context.Background(),
			user: user,
		},
		want:  *new(main_domains.User),
		want1: main_domains_exceptions.NewConflictExceptionSglMsg(_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS_VALUE),
	}

	tests := []testInputs{imput1}

	runTest(t, tests)
}

func runTest(t *testing.T, tests []testInputs) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateNewUserAllArgs(tt.fields.userDatabaseGateway,
				tt.fields.logLoki, main_configs_messages_resources.NewApplicationMessages(map[string]string{
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
//		logLoki             main_configs_apm_logs_impl.LogsGateway
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
