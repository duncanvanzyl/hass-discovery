package discovery

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
