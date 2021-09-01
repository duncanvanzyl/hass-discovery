package discovery

import "fmt"

type Cover struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// The MQTT topic subscribed to to receive birth and LWT messages from the MQTT cover device. If an `availability` topic is not defined, the cover availability state will always be `available`. If an `availability` topic is defined, the cover availability state will be `unavailable` by default. Must not be used together with `availability`
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to control the cover
	// Default: <no value>
	CommandTopic string `json:"command_topic,omitempty"`

	// Information about the device this cover is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works through [MQTT discovery](/docs/mqtt/discovery/) and when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// Sets the [class of the device](/integrations/cover/), changing the device state and icon that is displayed on the frontend
	// Default: <no value>
	DeviceClass string `json:"device_class,omitempty"`

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

	// The name of the cover
	// Default: MQTT Cover
	Name string `json:"name,omitempty"`

	// Flag that defines if switch works in optimistic mode
	// Default: `false` if state or position topic defined, else `true`.
	Optimistic string `json:"optimistic,omitempty"`

	// The payload that represents the online state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The command payload that closes the cover
	// Default: CLOSE
	PayloadClose string `json:"payload_close,omitempty"`

	// The payload that represents the offline state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The command payload that opens the cover
	// Default: OPEN
	PayloadOpen string `json:"payload_open,omitempty"`

	// The command payload that stops the cover
	// Default: STOP
	PayloadStop string `json:"payload_stop,omitempty"`

	// Number which represents closed position
	// Default: 0
	PositionClosed int `json:"position_closed,omitempty"`

	// Number which represents open position
	// Default: 100
	PositionOpen int `json:"position_open,omitempty"`

	// Defines a [template](/topics/templating/) that can be used to extract the payload for the `position_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function
	// Default: <no value>
	PositionTemplate string `json:"position_template,omitempty"`

	// The MQTT topic subscribed to receive cover position messages
	// Default: <no value>
	PositionTopic string `json:"position_topic,omitempty"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// Defines if published messages should have the retain flag set
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// Defines a [template](/topics/templating/) to define the position to be sent to the `set_position_topic` topic. Incoming position value is available for use in the template `{% raw %}{{ position }}{% endraw %}`. Within the template the following variables are available: `entity_id`, `position`, the target position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function
	// Default: <no value>
	SetPositionTemplate string `json:"set_position_template,omitempty"`

	// The MQTT topic to publish position commands to. You need to set position_topic as well if you want to use position topic. Use template if position topic wants different values than within range `position_closed` - `position_open`. If template is not defined and `position_closed != 100` and `position_open != 0` then proper position value is calculated from percentage position
	// Default: <no value>
	SetPositionTopic string `json:"set_position_topic,omitempty"`

	// The payload that represents the closed state
	// Default: closed
	StateClosed string `json:"state_closed,omitempty"`

	// The payload that represents the closing state
	// Default: closing
	StateClosing string `json:"state_closing,omitempty"`

	// The payload that represents the open state
	// Default: open
	StateOpen string `json:"state_open,omitempty"`

	// The payload that represents the opening state
	// Default: opening
	StateOpening string `json:"state_opening,omitempty"`

	// The payload that represents the stopped state (for covers that do not report `open`/`closed` state)
	// Default: stopped
	StateStopped string `json:"state_stopped,omitempty"`

	// The MQTT topic subscribed to receive cover state messages. State topic can only read (`open`, `opening`, `closed`, `closing` or `stopped`) state
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// The value that will be sent on a `close_cover_tilt` command
	// Default: 0
	TiltClosedValue int `json:"tilt_closed_value,omitempty"`

	// Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_command_topic` topic. Within the template the following variables are available: `entity_id`, `tilt_position`, the target tilt position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function
	// Default: <no value>
	TiltCommandTemplate string `json:"tilt_command_template,omitempty"`

	// The MQTT topic to publish commands to control the cover tilt
	// Default: <no value>
	TiltCommandTopic string `json:"tilt_command_topic,omitempty"`

	// The maximum tilt value
	// Default: 100
	TiltMax int `json:"tilt_max,omitempty"`

	// The minimum tilt value
	// Default: 0
	TiltMin int `json:"tilt_min,omitempty"`

	// The value that will be sent on an `open_cover_tilt` command
	// Default: 100
	TiltOpenedValue int `json:"tilt_opened_value,omitempty"`

	// Flag that determines if tilt works in optimistic mode
	// Default: `true` if `tilt_status_topic` is not defined, else `false`
	TiltOptimistic bool `json:"tilt_optimistic,omitempty"`

	// Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_status_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function
	// Default: <no value>
	TiltStatusTemplate string `json:"tilt_status_template,omitempty"`

	// The MQTT topic subscribed to receive tilt status update values
	// Default: <no value>
	TiltStatusTopic string `json:"tilt_status_topic,omitempty"`

	// An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines a [template](/topics/templating/) that can be used to extract the payload for the `state_topic` topic
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Cover
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Cover
func (d *Cover) AnnounceTopic(prefix string) string {
	topicFormat := "%s/cover/%s/config"
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
