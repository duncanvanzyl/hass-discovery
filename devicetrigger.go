package discovery

import "fmt"

type DeviceTrigger struct {

	// The type of automation, must be 'trigger'
	// Default: <no value>
	AutomationType string `json:"automation_type"`

	// Information about the device this device trigger is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device"`

	// Optional payload to match the payload being sent over the topic
	// Default: <no value>
	Payload string `json:"payload,omitempty"`

	// Must be `device_automation`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button
	// Default: <no value>
	Subtype string `json:"subtype"`

	// The MQTT topic subscribed to receive trigger events
	// Default: <no value>
	Topic string `json:"topic"`

	// The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1
	// Default: <no value>
	Type string `json:"type"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable DeviceTrigger
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the DeviceTrigger
func (d *DeviceTrigger) AnnounceTopic(prefix string) string {
	topicFormat := "%s/device_trigger/%s/config"
	objectID := hash(d)

	return fmt.Sprintf(topicFormat, prefix, objectID)
}
