package discovery

import "fmt"

type Light struct {

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with `availability_topic`
	// Default: <no value>
	Availability []Availability `json:"availability,omitempty"`

	// When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability
	// Default: latest
	AvailabilityMode string `json:"availability_mode,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`
	// Default: <no value>
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to change the light’s brightness
	// Default: <no value>
	BrightnessCommandTopic string `json:"brightness_command_topic,omitempty"`

	// Defines the maximum brightness value (i.e., 100%) of the MQTT device
	// Default: 255
	BrightnessScale int `json:"brightness_scale,omitempty"`

	// The MQTT topic subscribed to receive brightness state updates
	// Default: <no value>
	BrightnessStateTopic string `json:"brightness_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the brightness value
	// Default: <no value>
	BrightnessValueTemplate string `json:"brightness_value_template,omitempty"`

	// The MQTT topic subscribed to receive color mode updates. If this is not configured, `color_mode` will be automatically set according to the last received valid color or color temperatur
	// Default: <no value>
	ColorModeStateTopic string `json:"color_mode_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the color mode
	// Default: <no value>
	ColorModeValueTemplate string `json:"color_mode_value_template,omitempty"`

	// Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `color_temp_command_topic`. Available variables: `value`
	// Default: <no value>
	ColorTempCommandTemplate string `json:"color_temp_command_template,omitempty"`

	// The MQTT topic to publish commands to change the light’s color temperature state. The color temperature command slider has a range of 153 to 500 mireds (micro reciprocal degrees)
	// Default: <no value>
	ColorTempCommandTopic string `json:"color_temp_command_topic,omitempty"`

	// The MQTT topic subscribed to receive color temperature state updates. If the light also supports setting colors, also define a `white_value_state_topic`.
	// Default: <no value>
	ColorTempStateTopic string `json:"color_temp_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the color temperature value
	// Default: <no value>
	ColorTempValueTemplate string `json:"color_temp_value_template,omitempty"`

	// The MQTT topic to publish commands to change the switch state
	// Default: <no value>
	CommandTopic string `json:"command_topic"`

	// Information about the device this light is a part of to tie it into the [device registry](https://developers.home-assistant.io/docs/en/device_registry_index.html). Only works through [MQTT discovery](/docs/mqtt/discovery/) and when [`unique_id`](#unique_id) is set. At least one of identifiers or connections must be present to identify the device
	// Default: <no value>
	Device *Device `json:"device,omitempty"`

	// The MQTT topic to publish commands to change the light's effect state
	// Default: <no value>
	EffectCommandTopic string `json:"effect_command_topic,omitempty"`

	// The list of effects the light supports
	// Default: <no value>
	EffectList []string `json:"effect_list,omitempty"`

	// The MQTT topic subscribed to receive effect state updates
	// Default: <no value>
	EffectStateTopic string `json:"effect_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the effect value
	// Default: <no value>
	EffectValueTemplate string `json:"effect_value_template,omitempty"`

	// Flag which defines if the entity should be enabled when first added
	// Default: true
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`

	// The MQTT topic to publish commands to change the light's color state in HS format (Hue Saturation). Range for Hue: 0° .. 360°, Range of Saturation: 0..100. Note: Brightness is sent separately in the `brightness_command_topic`
	// Default: <no value>
	HsCommandTopic string `json:"hs_command_topic,omitempty"`

	// The MQTT topic subscribed to receive color state updates in HS format. Note: Brightness is received separately in the `brightness_state_topic`
	// Default: <no value>
	HsStateTopic string `json:"hs_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the HS value
	// Default: <no value>
	HsValueTemplate string `json:"hs_value_template,omitempty"`

	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity
	// Default: <no value>
	Icon string `json:"icon,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation
	// Default: <no value>
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation
	// Default: <no value>
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`

	// The maximum color temperature in mireds
	// Default: <no value>
	MaxMireds int `json:"max_mireds,omitempty"`

	// The minimum color temperature in mireds
	// Default: <no value>
	MinMireds int `json:"min_mireds,omitempty"`

	// The name of the light
	// Default: MQTT Light
	Name string `json:"name,omitempty"`

	// Defines when on the payload_on is sent. Using `last` (the default) will send any style (brightness, color, etc) topics first and then a `payload_on` to the `command_topic`. Using `first` will send the `payload_on` and then any style topics. Using `brightness` will only send brightness commands instead of the `payload_on` to turn the light on
	// Default: <no value>
	OnCommandType string `json:"on_command_type,omitempty"`

	// Flag that defines if switch works in optimistic mode
	// Default: `true` if no state topic defined, else `false`.
	Optimistic bool `json:"optimistic,omitempty"`

	// The payload that represents the available state
	// Default: online
	PayloadAvailable string `json:"payload_available,omitempty"`

	// The payload that represents the unavailable state
	// Default: offline
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`

	// The payload that represents disabled state
	// Default: OFF
	PayloadOff string `json:"payload_off,omitempty"`

	// The payload that represents enabled state
	// Default: ON
	PayloadOn string `json:"payload_on,omitempty"`

	// The maximum QoS level of the state topic
	// Default: 0
	Qos int `json:"qos,omitempty"`

	// If the published message should have the retain flag on or not
	// Default: false
	Retain bool `json:"retain,omitempty"`

	// Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgb_command_topic`. Available variables: `red`, `green` and `blue`
	// Default: <no value>
	RgbCommandTemplate string `json:"rgb_command_template,omitempty"`

	// The MQTT topic to publish commands to change the light's RGB state. Please note that the color value sent by Home Assistant is normalized to full brightness if `brightness_command_topic` is set. Brightness information is in this case sent separately in the `brightness_command_topic`. This will cause a light that expects an absolute color value (including brightness) to flicker
	// Default: <no value>
	RgbCommandTopic string `json:"rgb_command_topic,omitempty"`

	// The MQTT topic subscribed to receive RGB state updates. The expected payload is the RGB values separated by commas, for example, `255,0,127`. Please note that the color value received by Home Assistant is normalized to full brightness. Brightness information is received separately in the `brightness_state_topic`
	// Default: <no value>
	RgbStateTopic string `json:"rgb_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the RGB value
	// Default: <no value>
	RgbValueTemplate string `json:"rgb_value_template,omitempty"`

	// The schema to use. Must be `default` or omitted to select the default schema"
	// Default: default
	Schema string `json:"schema,omitempty"`

	// The MQTT topic subscribed to receive state updates
	// Default: <no value>
	StateTopic string `json:"state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the state value. The template should match the payload `on` and `off` values, so if your light uses `power on` to turn on, your `state_value_template` string should return `power on` when the switch is on. For example if the message is just `on`, your `state_value_template` should be `power {{ value }}`
	// Default: <no value>
	StateValueTemplate string `json:"state_value_template,omitempty"`

	// An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception
	// Default: <no value>
	UniqueId string `json:"unique_id,omitempty"`

	// The MQTT topic to publish commands to change the light to white mode with a given brightness
	// Default: <no value>
	WhiteCommandTopic string `json:"white_command_topic,omitempty"`

	// Defines the maximum white level (i.e., 100%) of the MQTT device
	// Default: 255
	WhiteScale int `json:"white_scale,omitempty"`

	// The MQTT topic to publish commands to change the light's XY state
	// Default: <no value>
	XyCommandTopic string `json:"xy_command_topic,omitempty"`

	// The MQTT topic subscribed to receive XY state updates
	// Default: <no value>
	XyStateTopic string `json:"xy_state_topic,omitempty"`

	// Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the XY value
	// Default: <no value>
	XyValueTemplate string `json:"xy_value_template,omitempty"`
}

// AnnounceTopic returns the topic to announce the discoverable Light
// Topic has the format below:
//   <discovery_prefix>/<component>/<object_id>/config
// 'object_id' is either the UniqueId, the Name, or a hash of the Light
func (d *Light) AnnounceTopic(prefix string) string {
	topicFormat := "%s/light/%s/config"
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
