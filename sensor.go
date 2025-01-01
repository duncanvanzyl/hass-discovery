package discovery

import "fmt"

type Sensor struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`
	// Default: <no value>
	AvailabilityTemplate string `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// Information about the device this sensor is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/device_registry_index/). Only works when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend. The `device_class` can be `null`
	// Default: <no value>
	DeviceClass string `json:"device_class,omitempty"`

	// Flag which defines if the entity should be enabled when first added
	// Default: true
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`

	// The encoding of the payloads received. Set to `""` to disable decoding of incoming payload
	// Default: utf-8
	Encoding string `json:"encoding,omitempty"`

	// The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity. When set, the entity category must be `diagnostic` for sensors
	// Default: <no value>
	EntityCategory string `json:"entity_category,omitempty"`

	// Picture URL for the entity
	// Default: <no value>
	EntityPicture string `json:"entity_picture,omitempty"`

	// If set, it defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires
	// Default: 0
	ExpireAfter int `json:"expire_after,omitempty"`

	// Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history
	// Default: false
	ForceUpdate bool `json:"force_update,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the last_reset. When `last_reset_value_template` is set, the `state_class` option must be `total`. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes
	// Default: <no value>
	LastResetValueTemplate string `json:"last_reset_value_template,omitempty"`

	// The name of the MQTT sensor. Can be set to `null` if only the device name is relevant
	// Default: MQTT Sensor
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

	// List of allowed sensor state value. An empty list is not allowed. The sensor's `device_class` must be set to `enum`. The `options` option cannot be used together with `state_class` or `unit_of_measurement`
	// Default: <no value>
	Options string `json:"options,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// Must be `sensor`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor
	// Default: <no value>
	StateClass string `json:"state_class,omitempty"`

	// The MQTT topic subscribed to receive sensor values. If `device_class`, `state_class`, `unit_of_measurement` or `suggested_display_precision` is set, and a numeric value is expected, an empty value `''` will be ignored and will not update the state, a `'null'` value will set the sensor to an `unknown` state. The `device_class` can be `null`
	// Default: <no value>
	StateTopic string `json:"state_topic"`

	// The number of decimals which should be used in the sensor's state after rounding
	// Default: <no value>
	SuggestedDisplayPrecision int `json:"suggested_display_precision,omitempty"`

	// An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines the units of measurement of the sensor, if any. The `unit_of_measurement` can be `null`
	// Default: <no value>
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value. If the template throws an error, the current state will be used instead
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
