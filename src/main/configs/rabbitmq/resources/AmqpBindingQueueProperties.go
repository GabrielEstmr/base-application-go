package main_configs_rabbitmq_resources

import "github.com/rabbitmq/amqp091-go"

type AmqpBindingQueueProperties struct {
	bindingId  string
	queueName  string
	routingKey string
	exchange   string
	nowait     bool
	args       amqp091.Table
}

func NewAmqpBindingQueueProperties(
	queueName string,
	routingKey string,
	exchange string,
	nowait bool,
	args amqp091.Table,
) *AmqpBindingQueueProperties {
	return &AmqpBindingQueueProperties{
		bindingId:  exchange + "_" + routingKey + "_" + queueName,
		queueName:  queueName,
		routingKey: routingKey,
		exchange:   exchange,
		nowait:     nowait,
		args:       args,
	}
}

func (this *AmqpBindingQueueProperties) GetBindingId() string {
	return this.bindingId
}

func (this *AmqpBindingQueueProperties) GetQueueName() string {
	return this.queueName
}

func (this *AmqpBindingQueueProperties) GetRoutingKey() string {
	return this.routingKey
}

func (this *AmqpBindingQueueProperties) GetExchange() string {
	return this.exchange
}

func (this *AmqpBindingQueueProperties) GetNowait() bool {
	return this.nowait
}

func (this *AmqpBindingQueueProperties) GetArgs() amqp091.Table {
	return this.args
}
