package discovery

import "fmt"

type Tag struct {

	// Information about the device this device trigger is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device"`

	// The MQTT topic subscribed to receive tag scanned events
	// Default: <no value>
	Topic string `json:"topic"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a tag ID
	// Default: <no value>
	ValueTemplate string `json:"value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Tag
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Tag
func (d *Tag) AnnounceTopic(prefix string) string {
	topicFormat := "%s/tag/%s/config"
	objectID := hash(d)

	return fmt.Sprintf(topicFormat, prefix, objectID)
}
