package main_usecases_factories_test

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases/factories"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	test_mocks "baseapplicationgo/test/mocks"
	"context"
	"reflect"
	"testing"
)

type testCreateEmailBodyFactoryInputs struct {
	name   string
	fields testCreateEmailBodyFactoryFields
	args   testCreateEmailBodyFactoryArgs
	want   main_usecases_interfaces.CreateEmailBody
}

type testCreateEmailBodyFactoryFields struct {
	createWelcomeEmailTemplateBody main_usecases_interfaces.CreateEmailBody
	logsMonitoringGateway          main_gateways.LogsMonitoringGateway
	spanGateway                    main_gateways.SpanGateway
}

type testCreateEmailBodyFactoryArgs struct {
	ctx       context.Context
	emailType main_domains_enums.EmailTemplateType
}

func TestCreateEmailBodyFactory_Get_ShouldReturnCreateWelcomeTemplateBodyWhenEmailIsWelcomeTemplateType(t *testing.T) {

	createEmailBodyMock := new(test_mocks.CreateEmailBodyMock)
	logsMonitoringGatewayMock := new(test_mocks.LogsMonitoringGatewayMock)
	spanGatewayMockImpl := new(test_mocks.SpanGatewayMockImpl)

	fields := testCreateEmailBodyFactoryFields{
		createEmailBodyMock,
		logsMonitoringGatewayMock,
		spanGatewayMockImpl,
	}

	args := testCreateEmailBodyFactoryArgs{
		context.Background(),
		main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL,
	}

	params := []testCreateEmailBodyFactoryInputs{
		{
			"TestCreateEmailBodyFactory_Get_ShouldReturnCreateWelcomeTemplateBodyWhenEmailIsWelcomeTemplateType",
			fields,
			args,
			createEmailBodyMock,
		},
	}

	runGetTestCases(t, params)
}

func TestCreateEmailBodyFactory_Get_ShouldReturnCreateWelcomeTemplateBodyWhenEmailIsNotificationUserTemplateType(t *testing.T) {

	createEmailBodyMock := new(test_mocks.CreateEmailBodyMock)
	logsMonitoringGatewayMock := new(test_mocks.LogsMonitoringGatewayMock)
	spanGatewayMockImpl := new(test_mocks.SpanGatewayMockImpl)

	fields := testCreateEmailBodyFactoryFields{
		createEmailBodyMock,
		logsMonitoringGatewayMock,
		spanGatewayMockImpl,
	}

	args := testCreateEmailBodyFactoryArgs{
		context.Background(),
		main_domains_enums.EMAIL_TYPE_NOTIFICATION_USER,
	}

	params := []testCreateEmailBodyFactoryInputs{
		{
			"TestCreateEmailBodyFactory_Get_ShouldReturnCreateWelcomeTemplateBodyWhenEmailIsNotificationUserTemplateType",
			fields,
			args,
			createEmailBodyMock,
		},
	}

	runGetTestCases(t, params)
}

func runGetTestCases(t *testing.T, params []testCreateEmailBodyFactoryInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases_factories.NewCreateEmailBodyFactoryAllArgs(
				tt.fields.createWelcomeEmailTemplateBody,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway,
			)
			if got := this.Get(tt.args.ctx, tt.args.emailType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
