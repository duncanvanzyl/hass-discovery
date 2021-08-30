package discovery

import (
	"encoding/json"
	"testing"
)

func TestBinarySensor(t *testing.T) {
	s := BinarySensor{
		StateTopic: "test/test",
	}

	bs, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("could not Binary Sensor: %v", err)
	}

	t.Fatalf("%s", bs)
}

func TestBiggerSensor(t *testing.T) {
	uid := "someuniqueidentifier"

	s := BinarySensor{
		StateTopic:    "some/sensor",
		Name:          "ON/OFF Sensor",
		DeviceClass:   "safety",
		UniqueId:      uid,
		ExpireAfter:   60 * 60 * 12,
		ValueTemplate: "{{ value_json.state }}",
		PayloadOn:     "on",
		PayloadOff:    "off",
		Device: &Device{
			Identifiers:   []string{uid},
			Manufacturer:  "Super Widgets Inc.",
			Model:         "Widget 1.0",
			Name:          "Widget",
			SuggestedArea: "outside",
			SWVersion:     "v0.0.3",
		},
	}

	bs, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("could not Binary Sensor: %v", err)
	}

	t.Fatalf("%s", bs)
}
