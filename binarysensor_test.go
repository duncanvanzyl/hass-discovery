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
