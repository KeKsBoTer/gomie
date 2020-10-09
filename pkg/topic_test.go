package gomie

import (
	"reflect"
	"testing"
)

func TestParseTopic(t *testing.T) {
	tests := []struct {
		topic   string
		wantErr bool
	}{
		{
			topic:   "homie/device123/$homie",
			wantErr: false,
		},
		{
			topic:   "homie/device123/$name",
			wantErr: false,
		},
		{
			topic:   "homie/device123/$state",
			wantErr: false,
		},
		{
			topic:   "homie/device123/$extensions",
			wantErr: false,
		},
		{
			topic:   "homie/device123/$nodes",
			wantErr: false,
		},

		{
			topic:   "homie/device123/mythermostat/$name",
			wantErr: false,
		},
		{
			topic:   "homie/device123/mythermostat/$properties",
			wantErr: false,
		},

		{
			topic:   "homie/device123/mythermostat/temperature",
			wantErr: false,
		},
		{
			topic:   "homie/device123/mythermostat/temperature/$name",
			wantErr: false,
		},
		{
			topic:   "homie/device123/mythermostat/temperature/$unit",
			wantErr: false,
		},
		{
			topic:   "homie/device123/mythermostat/temperature/$datatype",
			wantErr: false,
		},
		// cases that should fail
		{
			topic:   "homie/$broadcast",
			wantErr: true,
		},
		{
			topic:   "homie/$broadcast/alert",
			wantErr: true,
		},
		{
			topic:   "homie",
			wantErr: true,
		},
		{
			topic:   "homie/",
			wantErr: true,
		},
		{
			topic:   "homie/dev?ce/$homie",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.topic, func(t *testing.T) {
			got, err := ParseTopic(tt.topic)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			topicString := got.String()
			if !reflect.DeepEqual(topicString, tt.topic) {
				t.Errorf("ParseTopic().String() = %v, want %v", topicString, tt.topic)
			}
		})
	}
}
