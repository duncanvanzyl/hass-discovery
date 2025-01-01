package discovery

import "fmt"

type Humidifier struct {

	// A template to render the value received on the `action_topic` with
	// Default: <no value>
	ActionTemplate string `json:"action_template,omitempty"`

	// The MQTT topic to subscribe for changes of the current action. Valid values: `off`, `humidifying`, `drying`, `idle
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

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`
	// Default: <no value>
	CommandTemplate string `json:"command_template,omitempty"`

	// The MQTT topic to publish commands to change the humidifier state
	// Default: <no value>
	CommandTopic string `json:"command_topic"`

	// A template with which the value received on `current_humidity_topic` will be rendered
	// Default: <no value>
	CurrentHumidityTemplate string `json:"current_humidity_template,omitempty"`

	// The MQTT topic on which to listen for the current humidity. A `"None"` value received will reset the current humidity. Empty values (`'''`) will be ignored
	// Default: <no value>
	CurrentHumidityTopic string `json:"current_humidity_topic,omitempty"`

	// Information about the device this humidifier is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// The device class of the MQTT device. Must be either `humidifier`, `dehumidifier` or `null`
	// Default: humidifier
	DeviceClass string `json:"device_class,omitempty"`

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

	// The minimum target humidity percentage that can be set
	// Default: 100
	MaxHumidity string `json:"max_humidity,omitempty"`

	// The maximum target humidity percentage that can be set
	// Default: 0
	MinHumidity string `json:"min_humidity,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `mode_command_topic`
	// Default: <no value>
	ModeCommandTemplate string `json:"mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the `mode` on the humidifier. This attribute ust be configured together with the `modes` attribute
	// Default: <no value>
	ModeCommandTopic string `json:"mode_command_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `mode` state
	// Default: <no value>
	ModeStateTemplate string `json:"mode_state_template,omitempty"`

	// The MQTT topic subscribed to receive the humidifier `mode`
	// Default: <no value>
	ModeStateTopic string `json:"mode_state_topic,omitempty"`

	// List of available modes this humidifier is capable of running at. Common examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`, `auto` and `baby`. These examples offer built-in translations but other custom modes are allowed as well.  This attribute ust be configured together with the `mode_command_topic` attribute
	// Default: []
	Modes string `json:"modes,omitempty"`

	// The name of the humidifier. Can be set to `null` if only the device name is relevant
	// Default: MQTT humidifier
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

	// Flag that defines if humidifier works in optimistic mod
	// Default: `true` if no state topic defined, else `false`.
	Optimistic bool `json:"optimistic,omitempty"`

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

	// A special payload that resets the `target_humidity` state attribute to an `unknown` state when received at the `target_humidity_state_topic`. When received at `current_humidity_topic` it will reset the current humidity state
	// Default: "None"
	PayloadResetHumidity string `json:"payload_reset_humidity,omitempty"`

	// A special payload that resets the `mode` state attribute to an `unknown` state when received at the `mode_state_topic`
	// Default: "None"
	PayloadResetMode string `json:"payload_reset_mode,omitempty"`

	// Must be `humidifier`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: true
	Retain bool `json:"retain,omitempty"`

	// The MQTT topic subscribed to receive state updates. A "None" payload resets to an `unknown` state. An empty payload is ignored. Valid state payloads are `OFF` and `ON`. Custom `OFF` and `ON` values can be set with the `payload_off` and `payload_on` config options
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state
	// Default: <no value>
	StateValueTemplate string `json:"state_value_template,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`
	// Default: <no value>
	TargetHumidityCommandTemplate string `json:"target_humidity_command_template,omitempty"`

	// The MQTT topic to publish commands to change the humidifier target humidity state based on a percentage
	// Default: <no value>
	TargetHumidityCommandTopic string `json:"target_humidity_command_topic"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `target_humidity` state
	// Default: <no value>
	TargetHumidityStateTemplate string `json:"target_humidity_state_template,omitempty"`

	// The MQTT topic subscribed to receive humidifier target humidity
	// Default: <no value>
	TargetHumidityStateTopic string `json:"target_humidity_state_topic,omitempty"`

	// An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Humidifier
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Humidifier
func (d *Humidifier) AnnounceTopic(prefix string) string {
	topicFormat := "%s/humidifier/%s/config"
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
