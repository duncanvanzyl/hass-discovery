package discovery

import "fmt"

type Lock struct {

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

	// A regular expression to validate a supplied code when it is set during the action to `open`, `lock` or `unlock` the MQTT lock
	// Default: <no value>
	CodeFormat string `json:"code_format,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`. The lock command template accepts the parameters `value` and `code`. The `value` parameter will contain the configured value for either `payload_open`, `payload_lock` or `payload_unlock`. The `code` parameter is set during the action to `open`, `lock` or `unlock` the MQTT lock and will be set `None` if no code was passed
	// Default: <no value>
	CommandTemplate string `json:"command_template,omitempty"`

	// The MQTT topic to publish commands to change the lock state
	// Default: <no value>
	CommandTopic string `json:"command_topic"`

	// Information about the device this lock is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
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

	// The name of the lock. Can be set to `null` if only the device name is relevant
	// Default: MQTT Lock
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

	// Flag that defines if lock works in optimistic mode
	// Default: `true` if no `state_topic` defined, else `false`.
	Optimistic bool `json:"optimistic,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload sent to the lock to lock it
	// Default: LOCK
	PayloadLock string `json:"payload_lock,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload sent to the lock to open it
	// Default: <no value>
	PayloadOpen string `json:"payload_open,omitempty"`

	// A special payload that resets the state to `unknown` when received on the `state_topic`
	// Default: "None"
	PayloadReset string `json:"payload_reset,omitempty"`

	// The payload sent to the lock to unlock it
	// Default: UNLOCK
	PayloadUnlock string `json:"payload_unlock,omitempty"`

	// Must be `lock`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// The payload sent to `state_topic` by the lock when it's jammed
	// Default: JAMMED
	StateJammed string `json:"state_jammed,omitempty"`

	// The payload sent to `state_topic` by the lock when it's locked
	// Default: LOCKED
	StateLocked string `json:"state_locked,omitempty"`

	// The payload sent to `state_topic` by the lock when it's locking
	// Default: LOCKING
	StateLocking string `json:"state_locking,omitempty"`

	// The MQTT topic subscribed to receive state updates. It accepts states configured with `state_jammed`, `state_locked`, `state_unlocked`, `state_locking` or `state_unlocking`. A "None" payload resets to an `unknown` state. An empty payload is ignored
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// The payload sent to `state_topic` by the lock when it's unlocked
	// Default: UNLOCKED
	StateUnlocked string `json:"state_unlocked,omitempty"`

	// The payload sent to `state_topic` by the lock when it's unlocking
	// Default: UNLOCKING
	StateUnlocking string `json:"state_unlocking,omitempty"`

	// An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a state value from the payload
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
