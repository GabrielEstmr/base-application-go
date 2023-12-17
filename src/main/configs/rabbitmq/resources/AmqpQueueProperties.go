package main_configs_rabbitmq_resources

import "github.com/rabbitmq/amqp091-go"

type AmqpQueueProperties struct {
	queueName        string
	durable          bool
	deleteWhenUnused bool
	exclusive        bool
	nowait           bool
	args             amqp091.Table
}

func NewAmqpQueueParameters(
	queueName string,
	durable bool,
	deleteWhenUnused bool,
	exclusive bool,
	nowait bool,
	args amqp091.Table,
) *AmqpQueueProperties {
	return &AmqpQueueProperties{
		queueName:        queueName,
		durable:          durable,
		deleteWhenUnused: deleteWhenUnused,
		exclusive:        exclusive,
		nowait:           nowait,
		args:             args,
	}
}

func (this *AmqpQueueProperties) GetQueueName() string {
	return this.queueName
}

func (this *AmqpQueueProperties) GetDurable() bool {
	return this.durable
}

func (this *AmqpQueueProperties) GetDeleteWhenUnused() bool {
	return this.deleteWhenUnused
}

func (this *AmqpQueueProperties) GetExclusive() bool {
	return this.exclusive
}

func (this *AmqpQueueProperties) GetNowait() bool {
	return this.nowait
}

func (this *AmqpQueueProperties) GetArgs() amqp091.Table {
	return this.args
}
