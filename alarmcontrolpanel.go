package discovery

import "fmt"

type AlarmControlPanel struct {

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

	// If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](#configurations-with-remote-code-validation)
	// Default: <no value>
	Code string `json:"code,omitempty"`

	// If true the code is required to arm the alarm. If false the code is not validated
	// Default: true
	CodeArmRequired bool `json:"code_arm_required,omitempty"`

	// If true the code is required to disarm the alarm. If false the code is not validated
	// Default: true
	CodeDisarmRequired bool `json:"code_disarm_required,omitempty"`

	// If true the code is required to trigger the alarm. If false the code is not validated
	// Default: true
	CodeTriggerRequired bool `json:"code_trigger_required,omitempty"`

	// The [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) used for the command payload. Available variables: `action` and `code`
	// Default: action
	CommandTemplate string `json:"command_template,omitempty"`

	// The MQTT topic to publish commands to change the alarm state
	// Default: <no value>
	CommandTopic string `json:"command_topic"`

	// Information about the device this alarm panel is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
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

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The name of the alarm. Can be set to `null` if only the device name is relevant
	// Default: MQTT Alarm
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

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

	// The payload to set armed-vacation mode on your Alarm Panel
	// Default: ARM_VACATION
	PayloadArmVacation string `json:"payload_arm_vacation,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload to disarm your Alarm Panel
	// Default: DISARM
	PayloadDisarm string `json:"payload_disarm,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload to trigger the alarm on your Alarm Panel
	// Default: TRIGGER
	PayloadTrigger string `json:"payload_trigger,omitempty"`

	// Must be `alarm_control_panel`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// The MQTT topic subscribed to receive state updates. A "None" payload resets to an `unknown` state. An empty payload is ignored. Valid state payloads are: `armed_away`, `armed_custom_bypass`, `armed_home`, `armed_night`, `armed_vacation`, `arming`, `disarmed`, `disarming` `pending` and `triggered`
	// Default: <no value>
	StateTopic string `json:"state_topic"`

	// A list of features that the alarm control panel supports. The available list options are `arm_home`, `arm_away`, `arm_night`, `arm_vacation`, `arm_custom_bypass`, and `trigger`
	// Default: [arm_home arm_away arm_night arm_vacation arm_custom_bypass trigger]
	SupportedFeatures string `json:"supported_features,omitempty"`

	// An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable AlarmControlPanel
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the AlarmControlPanel
func (d *AlarmControlPanel) AnnounceTopic(prefix string) string {
	topicFormat := "%s/alarm_control_panel/%s/config"
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
