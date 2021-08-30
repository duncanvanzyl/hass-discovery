package discovery

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

	// The maximum QoS level to be used when receiving messages
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
}
