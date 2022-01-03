package discovery

//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/alarm_control_panel.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/binary_sensor.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/camera.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/cover.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/device_tracker.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/device_trigger.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/fan.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/humidifier.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/climate.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/light.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/lock.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/number.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/scene.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/select.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/sensor.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/switch.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/tag.mqtt.markdown
//go:generate go run generator/generator.go https://raw.githubusercontent.com/home-assistant/home-assistant.io/current/source/_integrations/vacuum.mqtt.markdown

// Availability is used by mulitple discovery configurations as a list of MQTT topics subscribed to
// receive availability (online/offline) updates. Must not be used together with availability_topic.
// TODO: Make this auto-generated by the generator.
type Availability struct {
	// payload_available string (optional, default: online)
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// payload_not_available string (optional, default: offline)
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// topic string REQUIRED
	// An MQTT topic subscribed to receive availability (online/offline) updates.
	Topic string `json:"topic"`
}

// Device is used by multiple discovery configurations as information about the device this object
// is a part of to tie it into the device registry. Only works through MQTT discovery and when
// unique_id is set. At least one of identifiers or connections must be present to identify the
// device.
// TODO: Make this auto-generated by the generator.
type Device struct {
	// connections list | map (optional)
	// A list of connections of the device to the outside world as a list of tuples [connection_type, connection_identifier]. For example the MAC address of a network interface: 'connections': ['mac', '02:5b:26:a8:dc:12'].
	Connections [][2]string `json:"connections,omitempty"`
	// identifiers list | string (optional)
	// A list of IDs that uniquely identify the device. For example a serial number.
	Identifiers []string `json:"identifiers,omitempty"`
	// manufacturer string (optional)
	// The manufacturer of the device.
	Manufacturer string `json:"manufacturer,omitempty"`
	// model string (optional)
	// The model of the device.
	Model string `json:"model,omitempty"`
	// name string (optional)
	// The name of the device.
	Name string `json:"name,,omitempty"`
	// suggested_area string (optional)
	// Suggest an area if the device isn’t in one yet.
	SuggestedArea string `json:"suggested_area,omitempty"`
	// sw_version string (optional)
	// The firmware version of the device.
	SWVersion string `json:"sw_version,omitempty"`
	// via_device string (optional)
	// Identifier of a device that routes messages between this device and Home Assistant. Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant.
	ViaDevice string `json:"via_device,omitempty"`
}

// Announcer is an interface for things that can announce themselves.
// Intended usage is to use AnnounceTopic to create the topic to announce to home assistant, and
// then use json.Marshal to convert the announcer into a payload.
type Announcer interface {
	AnnounceTopic(string) string
}
