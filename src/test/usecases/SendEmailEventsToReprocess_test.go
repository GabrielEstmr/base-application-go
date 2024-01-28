package main_usecases

import (
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	test_mocks "baseapplicationgo/test/mocks"
	test_mocks_gateways "baseapplicationgo/test/mocks/gateways"
	"context"
	"testing"
)

type testSendEmailEventsToReprocessInputs struct {
	name   string
	fields testSendEmailEventsToReprocessFields
	args   testSendEmailEventsToReprocessArgs
}

type testSendEmailEventsToReprocessFields struct {
	reprocessEmailEventProducerGateway main_gateways.ReprocessEmailEventProducerGateway
	logsMonitoringGateway              main_gateways.LogsMonitoringGateway
	spanGateway                        main_gateways.SpanGateway
	messageUtils                       main_utils_messages.ApplicationMessages
}
type testSendEmailEventsToReprocessArgs struct {
	ctx context.Context
	ids []string
}

func TestSendEmailEventsToReprocess_Execute(t *testing.T) {

	messageUtilsMock := *main_utils_messages.NewApplicationMessagesAllArgs(map[string]string{
		_MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_KEY: _MSG_FIND_EMAIL_BY_FILTER_ARCH_ISSUE_VALUE,
	})

	fields := testSendEmailEventsToReprocessFields{
		reprocessEmailEventProducerGateway: new(test_mocks_gateways.ReprocessEmailEventProducerGateway),
		logsMonitoringGateway:              new(test_mocks.LogsMonitoringGatewayMock),
		spanGateway:                        new(test_mocks.SpanGatewayMockImpl),
		messageUtils:                       messageUtilsMock,
	}

	args := testSendEmailEventsToReprocessArgs{
		ctx: context.Background(),
		ids: []string{"1"},
	}
	params := []testSendEmailEventsToReprocessInputs{
		{
			name:   "",
			fields: fields,
			args:   args,
		},
	}

	testSendEmailEventsToReprocess_RunTests(t, params)
}

func testSendEmailEventsToReprocess_RunTests(t *testing.T, params []testSendEmailEventsToReprocessInputs) {
	for _, tt := range params {
		t.Run(tt.name, func(t *testing.T) {
			this := main_usecases.NewSendEmailEventsToReprocessAllArgs(
				tt.fields.reprocessEmailEventProducerGateway,
				tt.fields.logsMonitoringGateway,
				tt.fields.spanGateway,
				tt.fields.messageUtils,
			)
			this.Execute(tt.args.ctx, tt.args.ids)
		})
	}
}
