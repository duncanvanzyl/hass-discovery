# hass-discovery

Provides the structures nessecary to create
[MQTT Discovery](https://www.home-assistant.io/docs/mqtt/discovery/) messages for use
with [Home Assistant](https://www.home-assistant.io/).

## Usage

Create an entry for the type of device you want to announce. Then publish to a valid
[topic](https://www.home-assistant.io/docs/mqtt/discovery/#discovery-topic). A valid
topic can be generated from the device entry with `device.AnnounceTopic(prefix)`. The
default prefix in `homeassistant`, but is configurable your `configuration.yaml` file.

Example:

```go
func PublishBinarySensor(cli mqtt.Client) error {
  uid := "someuniqueidentifier"

  s := BinarySensor{
    StateTopic:    "some/sensor",
    Name:          "ON/OFF Sensor",
    DeviceClass:   "safety",
    UniqueId:      uid,
    ExpireAfter:   60 * 60 * 12,
    ValueTemplate: "{{ value_json.state }}",
    PayloadOn:     "on",
    PayloadOff:    "off",
    Device: &Device{
      Identifiers:   []string{uid},
      Manufacturer:  "Super Widgets Inc.",
      Model:         "Widget 1.0",
      Name:          "Widget",
      SuggestedArea: "outside",
      SWVersion:     "v0.0.3",
    },
  }

  bs, err := json.Marshal(s)
  if err != nil {
    return fmt.Errorf("could not marshal Binary Sensor: %v", err)
  }

  announceTopic := s.AnnounceTopic("homeassistant")
  cli.Publish(announceTopic, 0, true, bs)

  return nil
}
```

## Generation

The structs are created directly from the
[documentation](https://www.home-assistant.io/docs/mqtt/discovery/) for the mqtt devices
that support discovery.  
Recreate the structs with `go generate ./...`  

### Sources

It's much easier to generate the structs from the yaml used to create the documentation
than it is to scrape <https://www.home-assistant.io/integrations/#search/mqtt> . So the
actual scraping is done on markdown found on
[github.com](https://github.com/home-assistant/home-assistant.io/tree/current/source/_integrations).

### Issues

- Device and Availability structs are hard coded and defined in
- [hassdiscovery.go](./hassdiscovery.go).
