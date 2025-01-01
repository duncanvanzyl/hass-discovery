package discovery

import "fmt"

type Vacuum struct {

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

	// The MQTT topic to publish commands to control the vacuum
	// Default: <no value>
	CommandTopic string `json:"command_topic,omitempty"`

	// Information about the device this switch is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// The encoding of the payloads received and published messages. Set to `""` to disable decoding of incoming payload
	// Default: utf-8
	Encoding string `json:"encoding,omitempty"`

	// List of possible fan speeds for the vacuum
	// Default: <no value>
	FanSpeedList []string `json:"fan_speed_list,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The name of the vacuum. Can be set to `null` if only the device name is relevant
	// Default: MQTT Vacuum
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload to send to the `command_topic` to begin a spot cleaning cycle
	// Default: clean_spot
	PayloadCleanSpot string `json:"payload_clean_spot,omitempty"`

	// The payload to send to the `command_topic` to locate the vacuum (typically plays a song)
	// Default: locate
	PayloadLocate string `json:"payload_locate,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload to send to the `command_topic` to pause the vacuum
	// Default: pause
	PayloadPause string `json:"payload_pause,omitempty"`

	// The payload to send to the `command_topic` to tell the vacuum to return to base
	// Default: return_to_base
	PayloadReturnToBase string `json:"payload_return_to_base,omitempty"`

	// The payload to send to the `command_topic` to begin the cleaning cycle
	// Default: start
	PayloadStart string `json:"payload_start,omitempty"`

	// The payload to send to the `command_topic` to stop cleaning
	// Default: stop
	PayloadStop string `json:"payload_stop,omitempty"`

	// Must be `vacuum`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// The MQTT topic to publish custom commands to the vacuum
	// Default: <no value>
	SendCommandTopic string `json:"send_command_topic,omitempty"`

	// The MQTT topic to publish commands to control the vacuum's fan speed
	// Default: <no value>
	SetFanSpeedTopic string `json:"set_fan_speed_topic,omitempty"`

	// The MQTT topic subscribed to receive state messages from the vacuum. Messages received on the `state_topic` must be a valid JSON dictionary, with a mandatory `state` key and optionally `battery_level` and `fan_speed` keys as shown in the [example](#configuration-example)
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// List of features that the vacuum supports (possible values are `start`, `stop`, `pause`, `return_home`, `battery`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)
	// Default: `start`, `stop`, `return_home`, `status`, `battery`, `clean_spot`
	SupportedFeatures []string `json:"supported_features,omitempty"`

	// An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Vacuum
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Vacuum
func (d *Vacuum) AnnounceTopic(prefix string) string {
	topicFormat := "%s/vacuum/%s/config"
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
