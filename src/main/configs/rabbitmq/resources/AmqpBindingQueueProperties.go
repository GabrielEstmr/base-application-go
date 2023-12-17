package main_configs_rabbitmq_resources

type AmqpBindingQueueProperties struct {
	bindingId        string
	queueName        string
	routingKey       string
	exchange         string
	durable          bool
	deleteWhenUnused bool
	exclusive        bool
	nowait           bool
}

func NewAmqpBindingQueueProperties(
	queueName string,
	routingKey string,
	exchange string,
	durable bool,
	deleteWhenUnused bool,
	exclusive bool,
	nowait bool) *AmqpBindingQueueProperties {
	return &AmqpBindingQueueProperties{
		bindingId:        exchange + "_" + routingKey + "_" + queueName,
		queueName:        queueName,
		routingKey:       routingKey,
		exchange:         exchange,
		durable:          durable,
		deleteWhenUnused: deleteWhenUnused,
		exclusive:        exclusive,
		nowait:           nowait,
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

func (this *AmqpBindingQueueProperties) GetDurable() bool {
	return this.durable
}

func (this *AmqpBindingQueueProperties) GetDeleteWhenUnused() bool {
	return this.deleteWhenUnused
}

func (this *AmqpBindingQueueProperties) GetExclusive() bool {
	return this.exclusive
}

func (this *AmqpBindingQueueProperties) GetNowait() bool {
	return this.nowait
}
