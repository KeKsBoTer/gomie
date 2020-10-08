package gomie

import "fmt"

type Operation string

const (
	Read Operation = "read"
	Set  Operation = "set"
)

type TopicLevel string

// Topic published with a message
type Topic interface {
	fmt.Stringer

	// Base is the MQTT base topic
	Base() TopicID
}

type baseTopic struct {
	name TopicID
}

func (b baseTopic) String() string {
	return string(b.name)
}

func (b baseTopic) Base() TopicID {
	return b.name
}

type DeviceTopic struct {
	base      baseTopic
	deviceID  TopicID
	attribute AttributeName
	operation Operation
}

func (d DeviceTopic)