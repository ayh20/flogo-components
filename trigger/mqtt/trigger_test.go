package mqtt

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	//MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

var jsonTestMetadata = getTestJsonMetadata()

func getTestJsonMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

const testConfig string = `{
  "name": "flogo-mqtt",
  "settings": {
    "topic": "flogo/#",
    "broker": "tcp://127.0.0.1:1883",
    "id": "flogoEngine",
    "user": "",
    "password": "",
    "store": "",
    "qos": "0",
    "cleansess": "false"
  },
  "handlers": [
    {
      "actionId": "device_info",
      "settings": {
        "topic": "test_start"
      }
    }
  ]
}`

const testConfig2 string = `{
	"name": "flogo-mqtt",
	"settings": {
	  "topic": "flogo/#",
	  "broker": "ssl://mqtt.bosch-iot-hub.com:8883",
	  "id": "flogoEngine",
	  "user": "little-sensor@tcef56e88b16548f9a4a49cd5b92150af",
	  "password": "plaintextPassword",
	  "store": "",
	  "qos": "1",
	  "keepalive": "30",
	  "cleansess": "false",
	  "enabletls": "true",
	  "certstore": "C:/Users/ahampshi/Documents/BoschIoTStuff/iothub.crt"

	},
	"handlers": [
	  {
		"actionId": "device_info",
		"settings": {
		  "topic": "event"
		}
	  }
	]
  }`

func TestInit(t *testing.T) {

	// New  factory
	md := trigger.NewMetadata(jsonTestMetadata)
	f := NewFactory(md)

	// New Trigger
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), config)
	f.New(&config)

	json.Unmarshal([]byte(testConfig2), config)
	f.New(&config)

}

/*
// TODO Fix this test
func TestEndpoint(t *testing.T) {

	// New  factory
	md := trigger.NewMetadata(jsonMetadata)
	f := NewFactory(md)

	// New Trigger
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), &config)
	tgr := f.New(&config)

	runner := &TestRunner{}

	tgr.Init(runner)

	tgr.Start()
	defer tgr.Stop()

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("flogo_test")
	opts.SetUsername("")
	opts.SetPassword("")
	opts.SetCleanSession(false)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Debug("---- doing first publish ----")

	token := client.Publish("test_start", 0, false, "Test message payload!")
	token.Wait()

	duration2 := time.Duration(2)*time.Second
	time.Sleep(duration2)

	log.Debug("---- doing second publish ----")

	token = client.Publish("test_start", 0, false, "Test message payload!")
	token.Wait()

	duration5 := time.Duration(5)*time.Second
	time.Sleep(duration5)

	client.Disconnect(250)
	log.Debug("Sample Publisher Disconnected")
}
*/
