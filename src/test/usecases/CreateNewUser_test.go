package main_usecases

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases "baseapplicationgo/main/usecases"
	"errors"
	"log/slog"
	"reflect"
	"testing"
)

type UserDatabaseGatewayMOCK struct {
}

func (this UserDatabaseGatewayMOCK) Save(user main_domains.User) (main_domains.User, error) {
	return main_domains.User{}, errors.New("ERROR")
}

func (this UserDatabaseGatewayMOCK) FindById(id string) (main_domains.User, error) {
	return main_domains.User{}, errors.New("ERROR")
}

func (this UserDatabaseGatewayMOCK) FindByDocumentNumber(documentNumber string) (main_domains.User, error) {
	return main_domains.User{}, errors.New("ERROR")
}

func (this UserDatabaseGatewayMOCK) FindByFilter(filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	return main_domains.Page{}, nil
}

type UserDatabaseGatewayMOCK2 struct {
}

func (this UserDatabaseGatewayMOCK2) Save(user main_domains.User) (main_domains.User, error) {
	return main_domains.User{}, nil
}

func (this UserDatabaseGatewayMOCK2) FindById(id string) (main_domains.User, error) {
	return *new(main_domains.User), nil
}

func (this UserDatabaseGatewayMOCK2) FindByDocumentNumber(documentNumber string) (main_domains.User, error) {
	return *new(main_domains.User), nil
}

func (this UserDatabaseGatewayMOCK2) FindByFilter(filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	return main_domains.Page{}, nil
}

const _TEST_MSG_CREATE_NEW_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue"

func TestCreateNewUser_Execute(t *testing.T) {

	type fields struct {
		userDatabaseGateway main_gateways.UserDatabaseGateway
		apLog               *slog.Logger
	}
	type args struct {
		user main_domains.User
	}
	type testArgs struct {
		name   string
		fields fields
		args   args
		want   main_domains.User
		want1  main_domains_exceptions.ApplicationException
	}
	sit1 := testArgs{
		name: "",
		fields: fields{
			userDatabaseGateway: UserDatabaseGatewayMOCK{},
			apLog:               main_configs_logs.GetLogConfigBean(),
		},
		args: args{
			user: main_domains.User{
				Id: "",
			}},
		want:  *new(main_domains.User),
		want1: main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_TEST_MSG_CREATE_NEW_DOC_ARCH_ISSUE),
	}

	sit2 := testArgs{
		name: "",
		fields: fields{
			userDatabaseGateway: UserDatabaseGatewayMOCK2{},
			apLog:               main_configs_logs.GetLogConfigBean(),
		},
		args: args{
			user: main_domains.User{
				Id: "",
			}},
		want:  *new(main_domains.User),
		want1: nil,
	}
	tests := []testArgs{
		sit1,
		sit2,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateNewUserAllArgs(tt.fields.userDatabaseGateway,
				tt.fields.apLog)
			got, got1 := this.Execute(tt.args.user)
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
	type args struct {
		userDatabaseGateway main_gateways.UserDatabaseGateway
	}
	tests := []struct {
		name string
		args args
		want *main_usecases.CreateNewUser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main_usecases.NewCreateNewUser(tt.args.userDatabaseGateway); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateNewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
