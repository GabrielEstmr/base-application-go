package main_configs_rabbitmq_resources

import "github.com/rabbitmq/amqp091-go"

type AmqpExchangeProperties struct {
	exchangeName string
	kind         string
	durable      bool
	autoDelete   bool
	internal     bool
	nowait       bool
	args         amqp091.Table
}

func NewAmqpExchangeParameters(
	exchangeName string,
	kind string,
	durable bool,
	autoDelete bool,
	internal bool,
	nowait bool,
	args amqp091.Table,
) *AmqpExchangeProperties {
	return &AmqpExchangeProperties{
		exchangeName: exchangeName,
		kind:         kind,
		durable:      durable,
		autoDelete:   autoDelete,
		internal:     internal,
		nowait:       nowait,
		args:         args,
	}
}

func (this *AmqpExchangeProperties) GetExchangeName() string {
	return this.exchangeName
}

func (this *AmqpExchangeProperties) GetKind() string {
	return this.kind
}

func (this *AmqpExchangeProperties) GetDurable() bool {
	return this.durable
}

func (this *AmqpExchangeProperties) GetAutoDelete() bool {
	return this.autoDelete
}

func (this *AmqpExchangeProperties) GetInternal() bool {
	return this.internal
}

func (this *AmqpExchangeProperties) GetNowait() bool {
	return this.nowait
}

func (this *AmqpExchangeProperties) GetArgs() amqp091.Table {
	return this.args
}
