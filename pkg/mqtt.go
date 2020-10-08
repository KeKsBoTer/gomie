package gomie

type Message interface {
	Duplicate() bool
	Qos() byte
	Retained() bool
	Topic() string
	MessageID() uint16
	Payload() []byte
	Ack()
}

type MessageHandler func(Message)

type MQTTClient interface {
	Subscribe(topic string, message MessageHandler) error
	Unsubscribe(topic string) error
	Publish(topic string, qos byte, retained bool, payload string) error
}
