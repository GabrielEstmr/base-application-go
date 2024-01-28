package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	test_mocks "baseapplicationgo/test/mocks"
	"bytes"
	"context"
	"fmt"
	"reflect"
	"testing"
	"text/template"
)

const _MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_KEY = "exceptions.architecture.application.issue-DEFAULT"
const _MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_VALUE = "Architecture application issue"

const _PATH_WELCOME_TEMPLATE_EMAIL = "../../zresources/emailtemplates/welcome-email-template.html"

type testCreateEmailBodyForWelcomeEmailInputs struct {
	name   string
	fields testCreateEmailBodyForWelcomeEmailFields
	args   testCreateEmailBodyForWelcomeEmailArgs
	want   []byte
	want1  main_domains_exceptions.ApplicationException
}
type testCreateEmailBodyForWelcomeEmailFields struct {
	templatePath          string
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

type testCreateEmailBodyForWelcomeEmailArgs struct {
	ctx         context.Context
	emailParams main_domains.EmailParams
}

func TestCreateEmailBodyForWelcomeEmail_ShouldReturnInternalServerErrorExceptionWhenParseFileFails(t *testing.T) {

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailBodyForWelcomeEmailFields{
		templatePath:          "./invalid-path",
		logsMonitoringGateway: new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:           new(test_mocks.SpanGatewayMockImpl),
		messageUtils:          messageUtilsMock,
	}

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
		"",
		"",
		[]string{""},
		"",
		nil,
	)

	args := testCreateEmailBodyForWelcomeEmailArgs{
		ctx:         context.Background(),
		emailParams: emailParams,
	}

	params := []testCreateEmailBodyForWelcomeEmailInputs{
		{
			name:   "TestCreateEmailBodyForWelcomeEmail_ShouldReturnInternalServerErrorExceptionWhenParseFileFails",
			fields: fields,
			args:   args,
			want:   nil,
			want1:  main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_VALUE),
		},
	}
	createEmailBodyForWelcomeEmail_RunTests(t, params)
}

func TestCreateEmailBodyForWelcomeEmail_CreateEmailBodyForWelcomeEmail(t *testing.T) {

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_VALUE,
	})

	fields := testCreateEmailBodyForWelcomeEmailFields{
		templatePath:          _PATH_WELCOME_TEMPLATE_EMAIL,
		logsMonitoringGateway: new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:           new(test_mocks.SpanGatewayMockImpl),
		messageUtils:          messageUtilsMock,
	}

	emailParams := *main_domains.NewEmailParams(
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
		"rm-exemple",
		"123",
		[]string{"gabriel@gmail.com"},
		"Email's Subject",
		map[string]string{
			"Name":    "Name",
			"Message": "Message",
		},
	)

	args := testCreateEmailBodyForWelcomeEmailArgs{
		ctx:         context.Background(),
		emailParams: emailParams,
	}

	params := []testCreateEmailBodyForWelcomeEmailInputs{
		{
			name:   "TestCreateEmailBodyForWelcomeEmail_CreateEmailBodyForWelcomeEmail",
			fields: fields,
			args:   args,
			want:   getWelComeTemplateBody(emailParams),
			want1:  nil,
		},
	}
	createEmailBodyForWelcomeEmail_RunTests(t, params)
}

func createEmailBodyForWelcomeEmail_RunTests(t *testing.T, params []testCreateEmailBodyForWelcomeEmailInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewCreateEmailBodyForWelcomeEmailAllArgs(
				tt.fields.templatePath,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway,
				tt.fields.messageUtils,
			)
			got, got1 := this.Execute(tt.args.ctx, tt.args.emailParams)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func getWelComeTemplateBody(emailParams main_domains.EmailParams) []byte {
	template, _ := template.ParseFiles(_PATH_WELCOME_TEMPLATE_EMAIL)
	var body bytes.Buffer
	subject := "Subject:" + emailParams.GetSubject() + "\n%s\n\n"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf(subject, mimeHeaders)))

	emailVariables := struct {
		Name    string
		Message string
	}{
		Name:    emailParams.GetBodyParams()["Name"],
		Message: emailParams.GetBodyParams()["Message"],
	}
	template.Execute(&body, emailVariables)

	return body.Bytes()
}

func TestNewCreateEmailBodyForWelcomeEmail(t *testing.T) {
	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_KEY: _MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_VALUE,
	})

	type args struct {
		logsMonitoringGateway main_gateways.LogsMonitoringGateway
		spanGateway           main_gateways.SpanGateway
		messageUtils          main_utils_messages.ApplicationMessages
	}
	tests := []struct {
		name string
		args args
		want *main_usecases.CreateEmailBodyForWelcomeEmail
	}{
		{
			name: "TestNewCreateEmailBodyForWelcomeEmail",
			args: args{
				messageUtils:          messageUtilsMock,
				logsMonitoringGateway: new(test_mocks.LogsMonitoringGatewayMock),
				spanGateway:           new(test_mocks.SpanGatewayMockImpl),
			},
			want: main_usecases.NewCreateEmailBodyForWelcomeEmailAllArgs(
				"./zresources/emailtemplates/welcome-email-template.html",
				new(test_mocks.LogsMonitoringGatewayMock),
				new(test_mocks.SpanGatewayMockImpl),
				messageUtilsMock,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main_usecases.NewCreateEmailBodyForWelcomeEmail(tt.args.logsMonitoringGateway, tt.args.spanGateway, tt.args.messageUtils); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateEmailBodyForWelcomeEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
