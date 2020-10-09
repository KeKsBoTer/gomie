package gomie

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

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

type BaseTopic struct {
	name TopicID
}

func (b BaseTopic) String() string {
	return string(b.name)
}

func (b BaseTopic) Base() TopicID {
	return b.name
}

type DeviceTopic struct {
	base      BaseTopic
	DeviceID  TopicID
	Attribute *AttributeName
}

func (t DeviceTopic) String() string {
	topic := path.Join(t.base.String(), string(t.DeviceID))
	if t.Attribute != nil {
		topic = path.Join(topic, t.Attribute.String())
	}
	return topic
}

func (t DeviceTopic) Base() TopicID {
	return t.base.name
}

type NodeTopic struct {
	Device    DeviceTopic
	NodeID    TopicID
	Attribute *AttributeName
}

func (t NodeTopic) String() string {
	topic := path.Join(t.Device.String(), string(t.NodeID))
	if t.Attribute != nil {
		topic = path.Join(topic, t.Attribute.String())
	}
	return topic
}

func (t NodeTopic) Base() TopicID {
	return t.Device.Base()
}

type PropertyTopic struct {
	Node       NodeTopic
	PropertyID TopicID
	Attribute  *AttributeName
	Operation  Operation
}

func (t PropertyTopic) String() string {
	topic := path.Join(t.Node.String(), string(t.PropertyID))
	if t.Attribute != nil {
		topic = path.Join(topic, t.Attribute.String())
	}
	if t.Operation == Set {
		return path.Join(topic, "set")
	}
	return topic
}

func (t PropertyTopic) Base() TopicID {
	return t.Node.Base()
}

type topicPart struct {
	id          TopicID
	isAttribute bool
}

func (p topicPart) String() string {
	id := string(p.id)
	if p.isAttribute {
		return "$" + id
	}
	return id
}

func (p topicPart) error() error {
	return &TopicIDFormatError{topicID: p.String()}
}

func parseTopicParts(parts []string) ([]topicPart, error) {
	topicParts := make([]topicPart, len(parts))
	for i, p := range parts {
		if len(p) == 0 {
			return nil, &TopicIDFormatError{topicID: p}
		}
		isAttribute := p[0] == '$'
		if isAttribute {
			p = p[1:]
		}
		part, err := NewTopicID(p)
		if err != nil {
			return nil, err
		}
		topicParts[i] = topicPart{
			id:          *part,
			isAttribute: isAttribute,
		}
	}
	return topicParts, nil
}

func ParseTopic(topic string) (Topic, error) {
	parts, err := parseTopicParts(strings.Split(topic, "/"))
	if err != nil {
		return nil, err
	}
	/*
		base/device/$attribute
		base/device/node/$attribute
		base/device/node/property
		base/device/node/property/set
		base/device/node/property/$attribute
	*/
	if len(parts) < 3 {
		return nil, errors.New("expected device level topic")
	}
	base := parts[0]
	if base.isAttribute {
		return nil, base.error()
	}
	device := parts[1]
	if device.isAttribute {
		return nil, device.error()
	}
	deviceTopic := DeviceTopic{
		base:     BaseTopic{name: base.id},
		DeviceID: device.id,
	}
	switch len(parts) {
	case 3:
		attr := parts[2]
		if !attr.isAttribute {
			return nil, errors.New("expected attribute at end")
		}
		attrName := AttributeName(attr.id)
		deviceTopic.Attribute = &attrName
		return &deviceTopic, nil
	case 4:
		node := parts[2]
		if node.isAttribute {
			return nil, node.error()
		}
		nodeTopic := NodeTopic{
			Device: deviceTopic,
			NodeID: node.id,
		}
		next := parts[3]
		if next.isAttribute {
			attrName := AttributeName(next.id)
			nodeTopic.Attribute = &attrName
			return &nodeTopic, nil
		}

		return PropertyTopic{
			Node:       nodeTopic,
			PropertyID: next.id,
			Attribute:  nil,
			Operation:  Read,
		}, nil
	case 5:
		node := parts[2]
		if node.isAttribute {
			return nil, node.error()
		}
		nodeTopic := NodeTopic{
			Device: deviceTopic,
			NodeID: node.id,
		}
		property := parts[3]
		if property.isAttribute {
			return nil, property.error()
		}
		propertyTopic := PropertyTopic{
			Node:       nodeTopic,
			PropertyID: property.id,
		}

		if parts[4].isAttribute {
			attrName := AttributeName(parts[4].id)
			propertyTopic.Attribute = &attrName
			propertyTopic.Operation = Read
		} else if parts[3].id == "set" {
			propertyTopic.Operation = Set
		} else {
			return nil, errors.New("expected attribute or 'set' as last part in topic")
		}
		return propertyTopic, nil
	default:
		return nil, errors.New("invalid topic (must consist of 3,4 or 5 parts)")
	}
}
