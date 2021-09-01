package discovery

import "fmt"

type Sensor struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// Information about the device this sensor is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works through [MQTT discovery](/docs/mqtt/discovery/) and when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend
	// Default: None
	DeviceClass string `json:"device_class,omitempty"`

	// Flag which defines if the entity should be enabled when first added
	// Default: true
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`

	// Defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`
	// Default: 0
	ExpireAfter int `json:"expire_after,omitempty"`

	// Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history
	// Default: false
	ForceUpdate bool `json:"force_update,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The MQTT topic subscribed to receive timestamps for when an accumulating sensor such as an energy meter was reset. If the sensor never resets, set `last_reset_topic` to same as `state_topic` and set the `last_reset_value_template` to a constant valid timstamp, for example UNIX epoch 0: `1970-01-01T00:00:00+00:00`
	// Default: <no value>
	LastResetTopic string `json:"last_reset_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the last_reset. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes
	// Default: <no value>
	LastResetValueTemplate string `json:"last_reset_value_template,omitempty"`

	// The name of the MQTT sensor
	// Default: MQTT Sensor
	Name string `json:"name,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The maximum QoS level of the state topic
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor
	// Default: None
	StateClass string `json:"state_class,omitempty"`

	// The MQTT topic subscribed to receive sensor values
	// Default: <no value>
	StateTopic string `json:"state_topic"`

	// An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines the units of measurement of the sensor, if any
	// Default: <no value>
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Sensor
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Sensor
func (d *Sensor) AnnounceTopic(prefix string) string {
	topicFormat := "%s/sensor/%s/config"
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
