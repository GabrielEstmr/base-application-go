package main_usecases_factories_test

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases/factories"
	test_mocks "baseapplicationgo/test/mocks"
	"context"
	"reflect"
	"testing"
)

type testSendEmailGatewayFactoryInputs struct {
	name   string
	fields testSendEmailGatewayFactoryFields
	args   testSendEmailGatewayFactoryArgs
	want   main_gateways.EmailGateway
}

type testSendEmailGatewayFactoryFields struct {
	gmailEmailGatewayImpl main_gateways.EmailGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

type testSendEmailGatewayFactoryArgs struct {
	ctx       context.Context
	emailType main_domains_enums.EmailTemplateType
}

func TestSendEmailGatewayFactory_Get_ShouldReturnGmailEmailGatewayWhenIsWelcomeEmailType(t *testing.T) {
	gmailEmailGatewayImplMock := new(test_mocks.EmailGatewayMock)
	logsMonitoringGatewayMock := new(test_mocks.LogsMonitoringGatewayMock)
	spanGatewayMockImpl := new(test_mocks.SpanGatewayMockImpl)

	fields := testSendEmailGatewayFactoryFields{
		gmailEmailGatewayImplMock,
		logsMonitoringGatewayMock,
		spanGatewayMockImpl,
	}

	args := testSendEmailGatewayFactoryArgs{
		context.Background(),
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
	}

	params := []testSendEmailGatewayFactoryInputs{
		{
			name:   "TestSendEmailGatewayFactory_Get_ShouldReturnGmailEmailGatewayWhenIsWelcomeEmailType",
			fields: fields,
			args:   args,
			want:   gmailEmailGatewayImplMock,
		},
	}

	runTestSendEmailGatewayFactory_Get(t, params)
}

func TestSendEmailGatewayFactory_Get_ShouldReturnGmailEmailGatewayWhenIsOtherEmailTypes(t *testing.T) {
	gmailEmailGatewayImplMock := new(test_mocks.EmailGatewayMock)
	logsMonitoringGatewayMock := new(test_mocks.LogsMonitoringGatewayMock)
	spanGatewayMockImpl := new(test_mocks.SpanGatewayMockImpl)

	fields := testSendEmailGatewayFactoryFields{
		gmailEmailGatewayImplMock,
		logsMonitoringGatewayMock,
		spanGatewayMockImpl,
	}

	args := testSendEmailGatewayFactoryArgs{
		context.Background(),
		main_domains_enums.EMAIL_TYPE_NOTIFICATION_USER,
	}

	params := []testSendEmailGatewayFactoryInputs{
		{
			name:   "TestSendEmailGatewayFactory_Get_ShouldReturnGmailEmailGatewayWhenIsOtherEmailTypes",
			fields: fields,
			args:   args,
			want:   gmailEmailGatewayImplMock,
		},
	}

	runTestSendEmailGatewayFactory_Get(t, params)
}

func runTestSendEmailGatewayFactory_Get(t *testing.T, params []testSendEmailGatewayFactoryInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases_factories.NewSendEmailGatewayFactoryAllArgs(
				tt.fields.gmailEmailGatewayImpl,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway)
			if got := this.Get(tt.args.ctx, tt.args.emailType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
