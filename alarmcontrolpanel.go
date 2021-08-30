package discovery

type AlarmControlPanel struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// If defined, specifies a code to enable or disable the alarm in the frontend
	// Default: <no value>
	Code string `json:"code,omitempty"`

	// If true the code is required to arm the alarm. If false the code is not validated
	// Default: true
	CodeArmRequired bool `json:"code_arm_required,omitempty"`

	// If true the code is required to disarm the alarm. If false the code is not validated
	// Default: true
	CodeDisarmRequired bool `json:"code_disarm_required,omitempty"`

	// The [template](/docs/configuration/templating/#processing-incoming-data) used for the command payload. Available variables: `action` and `code`
	// Default: action
	CommandTemplate string `json:"command_template,omitempty"`

	// The MQTT topic to publish commands to change the alarm state
	// Default: <no value>
	CommandTopic string `json:"command_topic"`

	// Information about the device this alarm panel is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works through [MQTT discovery](/docs/mqtt/discovery/) and when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// Flag which defines if the entity should be enabled when first added
	// Default: true
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The name of the alarm
	// Default: MQTT Alarm
	Name string `json:"name,omitempty"`

	// The payload to set armed-away mode on your Alarm Panel
	// Default: ARM_AWAY
	PayloadArmAway string `json:"payload_arm_away,omitempty"`

	// The payload to set armed-custom-bypass mode on your Alarm Panel
	// Default: ARM_CUSTOM_BYPASS
	PayloadArmCustomBypass string `json:"payload_arm_custom_bypass,omitempty"`

	// The payload to set armed-home mode on your Alarm Panel
	// Default: ARM_HOME
	PayloadArmHome string `json:"payload_arm_home,omitempty"`

	// The payload to set armed-night mode on your Alarm Panel
	// Default: ARM_NIGHT
	PayloadArmNight string `json:"payload_arm_night,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload to disarm your Alarm Panel
	// Default: DISARM
	PayloadDisarm string `json:"payload_disarm,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The maximum QoS level of the state topic
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// The MQTT topic subscribed to receive state updates
	// Default: <no value>
	StateTopic string `json:"state_topic"`

	// An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}
