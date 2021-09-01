package discovery

import "fmt"

type Vacuum struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// Defines a [template](/topics/templating/) to define the battery level of the vacuum. This is required if `battery_level_topic` is set
	// Default: <no value>
	BatteryLevelTemplate string `json:"battery_level_template,omitempty"`

	// The MQTT topic subscribed to receive battery level values from the vacuum
	// Default: <no value>
	BatteryLevelTopic string `json:"battery_level_topic,omitempty"`

	// Defines a [template](/topics/templating/) to define the charging state of the vacuum. This is required if `charging_topic` is set
	// Default: <no value>
	ChargingTemplate string `json:"charging_template,omitempty"`

	// The MQTT topic subscribed to receive charging state values from the vacuum
	// Default: <no value>
	ChargingTopic string `json:"charging_topic,omitempty"`

	// Defines a [template](/topics/templating/) to define the cleaning state of the vacuum. This is required if `cleaning_topic` is set
	// Default: <no value>
	CleaningTemplate string `json:"cleaning_template,omitempty"`

	// The MQTT topic subscribed to receive cleaning state values from the vacuum
	// Default: <no value>
	CleaningTopic string `json:"cleaning_topic,omitempty"`

	// The MQTT topic to publish commands to control the vacuum
	// Default: <no value>
	CommandTopic string `json:"command_topic,omitempty"`

	// Defines a [template](/topics/templating/) to define the docked state of the vacuum. This is required if `docked_topic` is set
	// Default: <no value>
	DockedTemplate string `json:"docked_template,omitempty"`

	// The MQTT topic subscribed to receive docked state values from the vacuum
	// Default: <no value>
	DockedTopic string `json:"docked_topic,omitempty"`

	// Flag which defines if the entity should be enabled when first added
	// Default: true
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`

	// Defines a [template](/topics/templating/) to define potential error messages emitted by the vacuum. This is required if `error_topic` is set
	// Default: <no value>
	ErrorTemplate string `json:"error_template,omitempty"`

	// The MQTT topic subscribed to receive error messages from the vacuum
	// Default: <no value>
	ErrorTopic string `json:"error_topic,omitempty"`

	// List of possible fan speeds for the vacuum
	// Default: <no value>
	FanSpeedList []string `json:"fan_speed_list,omitempty"`

	// Defines a [template](/topics/templating/) to define the fan speed of the vacuum. This is required if `fan_speed_topic` is set
	// Default: <no value>
	FanSpeedTemplate string `json:"fan_speed_template,omitempty"`

	// The MQTT topic subscribed to receive fan speed values from the vacuum
	// Default: <no value>
	FanSpeedTopic string `json:"fan_speed_topic,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The name of the vacuum
	// Default: MQTT Vacuum
	Name string `json:"name,omitempty"`

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

	// The payload to send to the `command_topic` to tell the vacuum to return to base
	// Default: return_to_base
	PayloadReturnToBase string `json:"payload_return_to_base,omitempty"`

	// The payload to send to the `command_topic` to start or pause the vacuum
	// Default: start_pause
	PayloadStartPause string `json:"payload_start_pause,omitempty"`

	// The payload to send to the `command_topic` to stop the vacuum
	// Default: stop
	PayloadStop string `json:"payload_stop,omitempty"`

	// The payload to send to the `command_topic` to turn the vacuum off
	// Default: turn_off
	PayloadTurnOff string `json:"payload_turn_off,omitempty"`

	// The payload to send to the `command_topic` to begin the cleaning cycle
	// Default: turn_on
	PayloadTurnOn string `json:"payload_turn_on,omitempty"`

	// The maximum QoS level of the state topic
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// The schema to use. Must be `legacy` or omitted to select the legacy schema
	// Default: legacy
	Schema string `json:"schema,omitempty"`

	// The MQTT topic to publish custom commands to the vacuum
	// Default: <no value>
	SendCommandTopic string `json:"send_command_topic,omitempty"`

	// The MQTT topic to publish commands to control the vacuum's fan speed
	// Default: <no value>
	SetFanSpeedTopic string `json:"set_fan_speed_topic,omitempty"`

	// List of features that the vacuum supports (possible values are `turn_on`, `turn_off`, `pause`, `stop`, `return_home`, `battery`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)
	// Default: `turn_on`, `turn_off`, `stop`, `return_home`, `status`, `battery`, `clean_spot`
	SupportedFeatures []string `json:"supported_features,omitempty"`

	// An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception
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
