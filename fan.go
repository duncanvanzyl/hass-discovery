package discovery

import "fmt"

type Fan struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`
	// Default: <no value>
	AvailabilityTemplate string `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`
	// Default: <no value>
	CommandTemplate string `json:"command_template,omitempty"`

	// The MQTT topic to publish commands to change the fan state
	// Default: <no value>
	CommandTopic string `json:"command_topic"`

	// Information about the device this fan is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `direction_command_topic`
	// Default: <no value>
	DirectionCommandTemplate string `json:"direction_command_template,omitempty"`

	// The MQTT topic to publish commands to change the direction state
	// Default: <no value>
	DirectionCommandTopic string `json:"direction_command_topic,omitempty"`

	// The MQTT topic subscribed to receive direction state updates
	// Default: <no value>
	DirectionStateTopic string `json:"direction_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the direction
	// Default: <no value>
	DirectionValueTemplate string `json:"direction_value_template,omitempty"`

	// Flag which defines if the entity should be enabled when first added
	// Default: true
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`

	// The encoding of the payloads received and published messages. Set to `""` to disable decoding of incoming payload
	// Default: utf-8
	Encoding string `json:"encoding,omitempty"`

	// The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity
	// Default: <no value>
	EntityCategory string `json:"entity_category,omitempty"`

	// Picture URL for the entity
	// Default: <no value>
	EntityPicture string `json:"entity_picture,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The name of the fan. Can be set to `null` if only the device name is relevant
	// Default: MQTT Fan
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

	// Flag that defines if fan works in optimistic mod
	// Default: `true` if no state topic defined, else `false`.
	Optimistic bool `json:"optimistic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `oscillation_command_topic`
	// Default: <no value>
	OscillationCommandTemplate string `json:"oscillation_command_template,omitempty"`

	// The MQTT topic to publish commands to change the oscillation state
	// Default: <no value>
	OscillationCommandTopic string `json:"oscillation_command_topic,omitempty"`

	// The MQTT topic subscribed to receive oscillation state updates
	// Default: <no value>
	OscillationStateTopic string `json:"oscillation_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the oscillation
	// Default: <no value>
	OscillationValueTemplate string `json:"oscillation_value_template,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload that represents the stop state
	// Default: OFF
	PayloadOff string `json:"payload_off,omitempty"`

	// The payload that represents the running state
	// Default: ON
	PayloadOn string `json:"payload_on,omitempty"`

	// The payload that represents the oscillation off state
	// Default: oscillate_off
	PayloadOscillationOff string `json:"payload_oscillation_off,omitempty"`

	// The payload that represents the oscillation on state
	// Default: oscillate_on
	PayloadOscillationOn string `json:"payload_oscillation_on,omitempty"`

	// A special payload that resets the `percentage` state attribute to `unknown` when received at the `percentage_state_topic`
	// Default: "None"
	PayloadResetPercentage string `json:"payload_reset_percentage,omitempty"`

	// A special payload that resets the `preset_mode` state attribute to `unknown` when received at the `preset_mode_state_topic`
	// Default: "None"
	PayloadResetPresetMode string `json:"payload_reset_preset_mode,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `percentage_command_topic`
	// Default: <no value>
	PercentageCommandTemplate string `json:"percentage_command_template,omitempty"`

	// The MQTT topic to publish commands to change the fan speed state based on a percentage
	// Default: <no value>
	PercentageCommandTopic string `json:"percentage_command_topic,omitempty"`

	// The MQTT topic subscribed to receive fan speed based on percentage
	// Default: <no value>
	PercentageStateTopic string `json:"percentage_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `percentage` value from the payload received on `percentage_state_topic`
	// Default: <no value>
	PercentageValueTemplate string `json:"percentage_value_template,omitempty"`

	// Must be `fan`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`
	// Default: <no value>
	PresetModeCommandTemplate string `json:"preset_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the preset mode
	// Default: <no value>
	PresetModeCommandTopic string `json:"preset_mode_command_topic,omitempty"`

	// The MQTT topic subscribed to receive fan speed based on presets
	// Default: <no value>
	PresetModeStateTopic string `json:"preset_mode_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`
	// Default: <no value>
	PresetModeValueTemplate string `json:"preset_mode_value_template,omitempty"`

	// List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`
	// Default: []
	PresetModes string `json:"preset_modes,omitempty"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: true
	Retain bool `json:"retain,omitempty"`

	// The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`
	// Default: 100
	SpeedRangeMax int `json:"speed_range_max,omitempty"`

	// The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`
	// Default: 1
	SpeedRangeMin int `json:"speed_range_min,omitempty"`

	// The MQTT topic subscribed to receive state updates. A "None" payload resets to an `unknown` state. An empty payload is ignored. By default, valid state payloads are `OFF` and `ON`. The accepted payloads can be overridden with the `payload_off` and `payload_on` config options
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state
	// Default: <no value>
	StateValueTemplate string `json:"state_value_template,omitempty"`

	// An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Fan
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Fan
func (d *Fan) AnnounceTopic(prefix string) string {
	topicFormat := "%s/fan/%s/config"
	objectID := ""
	switch {
	case d.UniqueId != "":
		objectID = d.UniqueId
	case d.Name != "":
		objectID = d.Name
	default:
		objectID = hash(d)
	}

	return fmt.Sprintf(topicFormat, prefix, objectID)
}
