package discovery

import "fmt"

type Climate struct {

	// A template to render the value received on the `action_topic` with
	// Default: <no value>
	ActionTemplate string `json:"action_template,omitempty"`

	// The MQTT topic to subscribe for changes of the current action. If this is set, the climate graph uses the value received as data source. A "None" payload resets the current action state. An empty payload is ignored. Valid action values: `off`, `heating`, `cooling`, `drying`, `idle`, `fan`
	// Default: <no value>
	ActionTopic string `json:"action_topic,omitempty"`

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

	// A template with which the value received on `current_humidity_topic` will be rendered
	// Default: <no value>
	CurrentHumidityTemplate string `json:"current_humidity_template,omitempty"`

	// The MQTT topic on which to listen for the current humidity. A `"None"` value received will reset the current humidity. Empty values (`'''`) will be ignored
	// Default: <no value>
	CurrentHumidityTopic string `json:"current_humidity_topic,omitempty"`

	// A template with which the value received on `current_temperature_topic` will be rendered
	// Default: <no value>
	CurrentTemperatureTemplate string `json:"current_temperature_template,omitempty"`

	// The MQTT topic on which to listen for the current temperature. A `"None"` value received will reset the current temperature. Empty values (`'''`) will be ignored
	// Default: <no value>
	CurrentTemperatureTopic string `json:"current_temperature_topic,omitempty"`

	// Information about the device this HVAC device is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

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

	// A template to render the value sent to the `fan_mode_command_topic` with
	// Default: <no value>
	FanModeCommandTemplate string `json:"fan_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the fan mode
	// Default: <no value>
	FanModeCommandTopic string `json:"fan_mode_command_topic,omitempty"`

	// A template to render the value received on the `fan_mode_state_topic` with
	// Default: <no value>
	FanModeStateTemplate string `json:"fan_mode_state_template,omitempty"`

	// The MQTT topic to subscribe for changes of the HVAC fan mode. If this is not set, the fan mode works in optimistic mode (see below). A "None" payload resets the fan mode state. An empty payload is ignored
	// Default: <no value>
	FanModeStateTopic string `json:"fan_mode_state_topic,omitempty"`

	// A list of supported fan modes
	// Default: [auto low medium high]
	FanModes string `json:"fan_modes,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Set the initial target temperature. The default value depends on the temperature unit and will be 21° or 69.8°F
	// Default: <no value>
	Initial string `json:"initial,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The minimum target humidity percentage that can be set
	// Default: 99
	MaxHumidity string `json:"max_humidity,omitempty"`

	// Maximum set point available. The default value depends on the temperature unit, and will be 35°C or 95°F
	// Default: <no value>
	MaxTemp string `json:"max_temp,omitempty"`

	// The maximum target humidity percentage that can be set
	// Default: 30
	MinHumidity string `json:"min_humidity,omitempty"`

	// Minimum set point available. The default value depends on the temperature unit, and will be 7°C or 44.6°F
	// Default: <no value>
	MinTemp string `json:"min_temp,omitempty"`

	// A template to render the value sent to the `mode_command_topic` with
	// Default: <no value>
	ModeCommandTemplate string `json:"mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the HVAC operation mode
	// Default: <no value>
	ModeCommandTopic string `json:"mode_command_topic,omitempty"`

	// A template to render the value received on the `mode_state_topic` with
	// Default: <no value>
	ModeStateTemplate string `json:"mode_state_template,omitempty"`

	// The MQTT topic to subscribe for changes of the HVAC operation mode. If this is not set, the operation mode works in optimistic mode (see below). A "None" payload resets to an `unknown` state. An empty payload is ignored
	// Default: <no value>
	ModeStateTopic string `json:"mode_state_topic,omitempty"`

	// A list of supported modes. Needs to be a subset of the default values
	// Default: [auto off cool heat dry fan_only]
	Modes string `json:"modes,omitempty"`

	// The name of the HVAC. Can be set to `null` if only the device name is relevant
	// Default: MQTT HVAC
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

	// Flag that defines if the climate works in optimistic mod
	// Default: `true` if no state topic defined, else `false`.
	Optimistic bool `json:"optimistic,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload sent to turn off the device
	// Default: OFF
	PayloadOff string `json:"payload_off,omitempty"`

	// The payload sent to turn the device on
	// Default: ON
	PayloadOn string `json:"payload_on,omitempty"`

	// A template to render the value sent to the `power_command_topic` with. The `value` parameter is the payload set for `payload_on` or `payload_off`
	// Default: <no value>
	PowerCommandTemplate string `json:"power_command_template,omitempty"`

	// The MQTT topic to publish commands to change the HVAC power state. Sends the payload configured with `payload_on` if the climate is turned on via the `climate.turn_on`, or the payload configured with `payload_off` if the climate is turned off via the `climate.turn_off` action. Note that `optimistic` mode is not supported through `climate.turn_on` and `climate.turn_off` actions. When called, these actions will send a power command to the device but will not optimistically update the state of the climate entity. The climate device should report its state back via `mode_state_topic`
	// Default: <no value>
	PowerCommandTopic string `json:"power_command_topic,omitempty"`

	// The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`
	// Default: 0.1 for Celsius and 1.0 for Fahrenheit.
	Precision string `json:"precision,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`
	// Default: <no value>
	PresetModeCommandTemplate string `json:"preset_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the preset mode
	// Default: <no value>
	PresetModeCommandTopic string `json:"preset_mode_command_topic,omitempty"`

	// The MQTT topic subscribed to receive climate speed based on presets. When preset 'none' is received or `None` the `preset_mode` will be reset
	// Default: <no value>
	PresetModeStateTopic string `json:"preset_mode_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`
	// Default: <no value>
	PresetModeValueTemplate string `json:"preset_mode_value_template,omitempty"`

	// List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`
	// Default: []
	PresetModes string `json:"preset_modes,omitempty"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// Defines if published messages should have the retain flag set
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// A template to render the value sent to the `swing_mode_command_topic` with
	// Default: <no value>
	SwingModeCommandTemplate string `json:"swing_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the swing mode
	// Default: <no value>
	SwingModeCommandTopic string `json:"swing_mode_command_topic,omitempty"`

	// A template to render the value received on the `swing_mode_state_topic` with
	// Default: <no value>
	SwingModeStateTemplate string `json:"swing_mode_state_template,omitempty"`

	// The MQTT topic to subscribe for changes of the HVAC swing mode. If this is not set, the swing mode works in optimistic mode (see below)
	// Default: <no value>
	SwingModeStateTopic string `json:"swing_mode_state_topic,omitempty"`

	// A list of supported swing modes
	// Default: [on off]
	SwingModes string `json:"swing_modes,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`
	// Default: <no value>
	TargetHumidityCommandTemplate string `json:"target_humidity_command_template,omitempty"`

	// The MQTT topic to publish commands to change the target humidity
	// Default: <no value>
	TargetHumidityCommandTopic string `json:"target_humidity_command_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the climate `target_humidity` state
	// Default: <no value>
	TargetHumidityStateTemplate string `json:"target_humidity_state_template,omitempty"`

	// The MQTT topic subscribed to receive the target humidity. If this is not set, the target humidity works in optimistic mode (see below). A `"None"` value received will reset the target humidity. Empty values (`'''`) will be ignored
	// Default: <no value>
	TargetHumidityStateTopic string `json:"target_humidity_state_topic,omitempty"`

	// Step size for temperature set point
	// Default: 1
	TempStep string `json:"temp_step,omitempty"`

	// A template to render the value sent to the `temperature_command_topic` with
	// Default: <no value>
	TemperatureCommandTemplate string `json:"temperature_command_template,omitempty"`

	// The MQTT topic to publish commands to change the target temperature
	// Default: <no value>
	TemperatureCommandTopic string `json:"temperature_command_topic,omitempty"`

	// A template to render the value sent to the `temperature_high_command_topic` with
	// Default: <no value>
	TemperatureHighCommandTemplate string `json:"temperature_high_command_template,omitempty"`

	// The MQTT topic to publish commands to change the high target temperature
	// Default: <no value>
	TemperatureHighCommandTopic string `json:"temperature_high_command_topic,omitempty"`

	// A template to render the value received on the `temperature_high_state_topic` with. A `"None"` value received will reset the temperature high set point. Empty values (`'''`) will be ignored
	// Default: <no value>
	TemperatureHighStateTemplate string `json:"temperature_high_state_template,omitempty"`

	// The MQTT topic to subscribe for changes in the target high temperature. If this is not set, the target high temperature works in optimistic mode (see below)
	// Default: <no value>
	TemperatureHighStateTopic string `json:"temperature_high_state_topic,omitempty"`

	// A template to render the value sent to the `temperature_low_command_topic` with
	// Default: <no value>
	TemperatureLowCommandTemplate string `json:"temperature_low_command_template,omitempty"`

	// The MQTT topic to publish commands to change the target low temperature
	// Default: <no value>
	TemperatureLowCommandTopic string `json:"temperature_low_command_topic,omitempty"`

	// A template to render the value received on the `temperature_low_state_topic` with. A `"None"` value received will reset the temperature low set point. Empty values (`'''`) will be ignored
	// Default: <no value>
	TemperatureLowStateTemplate string `json:"temperature_low_state_template,omitempty"`

	// The MQTT topic to subscribe for changes in the target low temperature. If this is not set, the target low temperature works in optimistic mode (see below)
	// Default: <no value>
	TemperatureLowStateTopic string `json:"temperature_low_state_topic,omitempty"`

	// A template to render the value received on the `temperature_state_topic` with
	// Default: <no value>
	TemperatureStateTemplate string `json:"temperature_state_template,omitempty"`

	// The MQTT topic to subscribe for changes in the target temperature. If this is not set, the target temperature works in optimistic mode (see below). A `"None"` value received will reset the temperature set point. Empty values (`'''`) will be ignored
	// Default: <no value>
	TemperatureStateTopic string `json:"temperature_state_topic,omitempty"`

	// Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit
	// Default: <no value>
	TemperatureUnit string `json:"temperature_unit,omitempty"`

	// An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Default template to render the payloads on *all* `*_state_topic`s with
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Climate
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Climate
func (d *Climate) AnnounceTopic(prefix string) string {
	topicFormat := "%s/climate/%s/config"
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
