package gomie

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type withID interface {
	ID() TopicID
}

type withName interface {
	Name() string
}

type canUpdate interface {
	UpdateValue(key AttributeName, value string)
}

type withAttributes interface {
	Attribute(key AttributeName) string
	Attributes() []AttributeName
}

type PayloadType string

const (
	Integer PayloadType = "integer"
	Float   PayloadType = "float"
	Boolean PayloadType = "boolean"
	String  PayloadType = "string"
	Enum    PayloadType = "enum"
	Color   PayloadType = "color"
)

type TopicID string

var topicRegex = regexp.MustCompile("^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$")

func NewTopicID(topic string) (*TopicID, error) {

	if topicRegex.MatchString(topic) {
		t := TopicID(topic)
		return &t, nil
	}
	return nil, &TopicIDFormatError{topicID: topic}
}

type AttributeName TopicID

func (a AttributeName) String() string {
	return "$" + string(a)
}

type Version struct {
	major int
	minor int
	build int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.build)
}

func ParseHomieVersion(v string) (*Version, error) {
	vOrig := v
	if v[0] == 'v' {
		v = v[1:]
	}
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return nil, &MalformattedVersionError{version: vOrig}
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, &MalformattedVersionError{version: vOrig}
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, &MalformattedVersionError{version: vOrig}
	}
	build, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, &MalformattedVersionError{version: vOrig}
	}

	return &Version{
		major: major,
		minor: minor,
		build: build,
	}, nil
}

type State string

const (
	Init         State = "init"
	Ready        State = "ready"
	Disconnected State = "disconnected"
	Sleeping     State = "sleeping"
	Lost         State = "lost"
	Alert        State = "alert"
)

type Nodes []TopicID

func (n Nodes) String() string {
	result := ""
	for i, v := range n {
		result += string(v)
		if i != len(n)-1 {
			result += ","
		}
	}
	return result
}

type Extentions []string

func (n Extentions) String() string {
	return strings.Join(n, ",")
}

type Properties []TopicID

func (p Properties) String() string {
	result := ""
	for i, v := range p {
		result += string(v)
		if i != len(p)-1 {
			result += ","
		}
	}
	return result
}
