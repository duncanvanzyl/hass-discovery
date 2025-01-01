package discovery

import "fmt"

type DeviceTracker struct {

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

	// Information about the device this device tracker is a part of that ties it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary message containing device tracker attributes. This topic can be used to set the location of the device tracker under the following conditions:
	// - If the attributes in the JSON message include `longitude`, `latitude`, and `gps_accuracy` (optional).
	//  - If the device tracker is within a configured [zone](/integrations/zone/).
	//
	// If these conditions are met, it is not required to configure `state_topic`.
	//
	//  Be aware that any location message received at `state_topic`  overrides the location received via `json_attributes_topic` until a message configured with `payload_reset` is received at `state_topic`. For a more generic usage example of the `json_attributes_topic`, refer to the [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The name of the MQTT device_tracker
	// Default: <no value>
	Name string `json:"name,omitempty"`

	// Used instead of `name` for automatic generation of `entity_id
	// Default: <no value>
	ObjectId string `json:"object_id,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload value that represents the 'home' state for the device
	// Default: home
	PayloadHome string `json:"payload_home,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload value that represents the 'not_home' state for the device
	// Default: not_home
	PayloadNotHome string `json:"payload_not_home,omitempty"`

	// The payload value that will have the device's location automatically derived from Home Assistant's zones
	// Default: "None"
	PayloadReset string `json:"payload_reset,omitempty"`

	// Must be `device_tracker`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)
	// Default: <no value>
	Platform string `json:"platform"`

	// The maximum QoS level to be used when receiving and publishing messages
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`
	// Default: <no value>
	SourceType string `json:"source_type,omitempty"`

	// The MQTT topic subscribed to receive device tracker state changes. The states defined in `state_topic` override the location states defined by the `json_attributes_topic`. This state override is turned inactive if the `state_topic` receives a message containing `payload_reset`. The `state_topic` can only be omitted if `json_attributes_topic` is used. An empty payload is ignored. Valid payloads are `not_home`, `home` or any other custom location or zone name. Payloads for `not_home`, `home` can be overridden with the `payload_not_home`and `payload_home` config options
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// An ID that uniquely identifies this device_tracker. If two device_trackers have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a device tracker state
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable DeviceTracker
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the DeviceTracker
func (d *DeviceTracker) AnnounceTopic(prefix string) string {
	topicFormat := "%s/device_tracker/%s/config"
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
