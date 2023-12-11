package usecases

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_configs_apm_logs_resources "baseapplicationgo/main/configs/apm/logs/resources"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
	"context"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"testing"
)

type UserDatabaseGatewayMock struct {
}

func NewUserDatabaseGatewayMock() *UserDatabaseGatewayMock {
	return &UserDatabaseGatewayMock{}
}

func (this *UserDatabaseGatewayMock) Save(user main_domains.User) (main_domains.User, error) {
	user.Id = "123"
	return user, nil
}

func (this *UserDatabaseGatewayMock) FindById(id string) (main_domains.User, error) {
	return main_domains.User{}, nil
}

func (this *UserDatabaseGatewayMock) FindByDocumentNumber(documentNumber string) (main_domains.User, error) {
	return main_domains.User{}, nil
}

func (this *UserDatabaseGatewayMock) FindByFilter(filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	return main_domains.Page{}, nil
}

type LogsGatewayImplMock struct {
	logConfig main_configs_apm_logs_resources.LogProviderConfig
}

func (this LogsGatewayImplMock) DEBUG(
	span trace.Span,
	msg string,
	args ...any,
) {

}

func (this LogsGatewayImplMock) WARN(
	span trace.Span,
	msg string,
	args ...any,
) {

}

func (this LogsGatewayImplMock) INFO(
	span trace.Span,
	msg string,
	args ...any,
) {

}

func (this LogsGatewayImplMock) ERROR(
	span trace.Span,
	msg string,
	args ...any,
) {

}

func TestCreateNewUser_Execute(t *testing.T) {
	type fields struct {
		userDatabaseGateway main_gateways.UserDatabaseGateway
		logLoki             main_configs_apm_logs_impl.LogsGateway
	}
	type args struct {
		ctx  context.Context
		user main_domains.User
	}

	type TestInputs struct {
		name   string
		fields fields
		args   args
		want   main_domains.User
		want1  main_domains_exceptions.ApplicationException
	}

	userDatabaseGatewayMock := NewUserDatabaseGatewayMock()
	userResponse, _ := userDatabaseGatewayMock.Save(main_domains.User{})
	imput1 := TestInputs{
		name: "Test1",
		fields: fields{
			userDatabaseGateway: userDatabaseGatewayMock,
			logLoki:             LogsGatewayImplMock{},
		},
		args: args{
			ctx:  context.Background(),
			user: main_domains.User{},
		},
		want:  userResponse,
		want1: nil,
	}

	tests := []TestInputs{imput1}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateNewUserAllArgs(tt.fields.userDatabaseGateway,
				tt.fields.logLoki)

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
