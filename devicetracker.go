package discovery

import "fmt"

type DeviceTracker struct {

	// List of devices with their topic
	// Default: <no value>
	Devices string `json:"devices"`

	// The payload value that represents the 'home' state for the device
	// Default: home
	PayloadHome string `json:"payload_home,omitempty"`

	// The payload value that represents the 'not_home' state for the device
	// Default: not_home
	PayloadNotHome string `json:"payload_not_home,omitempty"`

	// The QoS level of the topic
	// Default: <no value>
	Qos int `json:"qos,omitempty"`

	// Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`
	// Default: <no value>
	SourceType string `json:"source_type,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable DeviceTracker
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the DeviceTracker
func (d *DeviceTracker) AnnounceTopic(prefix string) string {
	topicFormat := "%s/device_tracker/%s/config"
	objectID := hash(d)

	return fmt.Sprintf(topicFormat, prefix, objectID)
}
