package main_configs_rabbitmq_resources

import "github.com/rabbitmq/amqp091-go"

type AmqpConsumerProperties struct {
	queueName   string
	consumerTag string
	autoAck     bool
	exclusive   bool
	noLocal     bool
	noWait      bool
	args        amqp091.Table
}

func NewAmqpConsumerProperties(
	queueName string,
	consumerTag string,
	autoAck bool,
	exclusive bool,
	noLocal bool,
	noWait bool,
	args amqp091.Table,
) *AmqpConsumerProperties {
	return &AmqpConsumerProperties{
		queueName:   queueName,
		consumerTag: consumerTag,
		autoAck:     autoAck,
		exclusive:   exclusive,
		noLocal:     noLocal,
		noWait:      noWait,
		args:        args,
	}
}

func (this *AmqpConsumerProperties) GetQueueName() string {
	return this.queueName
}

func (this *AmqpConsumerProperties) GetConsumerTag() string {
	return this.consumerTag
}

func (this *AmqpConsumerProperties) GetAutoAck() bool {
	return this.autoAck
}

func (this *AmqpConsumerProperties) GetExclusive() bool {
	return this.exclusive
}

func (this *AmqpConsumerProperties) GetNoLocal() bool {
	return this.noLocal
}

func (this *AmqpConsumerProperties) GetNoWait() bool {
	return this.noWait
}

func (this *AmqpConsumerProperties) GetArgs() amqp091.Table {
	return this.args
}
