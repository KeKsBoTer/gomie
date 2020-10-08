package gomie

import "fmt"

type TopicIDFormatError struct {
	topicID string
}

func (e *TopicIDFormatError) Error() string {
	return fmt.Sprintf("'%s' is not a valid topic ID", e.topicID)
}

type MalformattedVersionError struct {
	version string
}

func (e *MalformattedVersionError) Error() string {
	return fmt.Sprintf("'%s' is not a valid homie version", e.version)
}
