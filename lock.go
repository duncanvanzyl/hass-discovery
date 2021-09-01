package discovery

import "fmt"

type Lock struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to change the lock state
	// Default: <no value>
	CommandTopic string `json:"command_topic"`

	// Information about the device this lock is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works through [MQTT discovery](/docs/mqtt/discovery/) and when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
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

	// The name of the lock
	// Default: MQTT Lock
	Name string `json:"name,omitempty"`

	// Flag that defines if lock works in optimistic mode
	// Default: `true` if no `state_topic` defined, else `false`.
	Optimistic string `json:"optimistic,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload that represents enabled/locked state
	// Default: LOCK
	PayloadLock string `json:"payload_lock,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload that represents disabled/unlocked state
	// Default: UNLOCK
	PayloadUnlock string `json:"payload_unlock,omitempty"`

	// The maximum QoS level of the state topic
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// The value that represents the lock to be in locked stat
	// Default: LOCKED
	StateLocked string `json:"state_locked,omitempty"`

	// The MQTT topic subscribed to receive state updates
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// The value that represents the lock to be in unlocked stat
	// Default: UNLOCKED
	StateUnlocked string `json:"state_unlocked,omitempty"`

	// An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the payload
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Lock
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Lock
func (d *Lock) AnnounceTopic(prefix string) string {
	topicFormat := "%s/lock/%s/config"
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
