package mqtt

import (
	"context"
	"fmt"
	"io/ioutil"
	golog "log"
	"time"

	//MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

var jsonTestMetadata = getTestJsonMetadata()
var listentime time.Duration = 10

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
	  "broker": "ssl://mqtt.bosch-iot-hub.com:8883",
	  "id": "flogoEngine",
	  "user": "flogo.tibco.com_my-device-id-4712@t90ab69dfe0e54fb9bb38c9083ebb2936_hub",
	  "password": "password",
	  "store": "",
	  "qos": "1",
	  "keepalive": "30",
	  "cleansess": "false",
		"enabletls": "true",
		"autoreconnect": "true",
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

type TestRunner struct {
}

func (tr *TestRunner) Execute(ctx context.Context, act action.Action, inputs map[string]*data.Attribute) (results map[string]*data.Attribute, err error) {
	golog.Printf("Ran Action: %v", act.Metadata().ID)
	return nil, nil
}

// Run implements action.Runner.Run
func (tr *TestRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {
	golog.Printf("Ran Action: %v", uri)
	return 0, nil, nil
}

func (tr *TestRunner) RunAction(ctx context.Context, act action.Action, options map[string]interface{}) (results map[string]*data.Attribute, err error) {
	golog.Printf("Ran Action: %v", act.Metadata().ID)
	return nil, nil
}

func TestrunTest(config *trigger.Config, expectSucceed bool, testName string, configOnly bool) error {
	golog.Printf("Test %s starting\n", testName)
	defer func() error {
		if r := recover(); r != nil {
			if expectSucceed {
				golog.Printf("Test %s was expected to succeed but did not because: %s", testName, r)
				return fmt.Errorf("%s", r)
			}
		}
		return nil
	}()

	// New  factory
	md := trigger.NewMetadata(jsonTestMetadata)
	f := NewFactory(md)

	//f := &KafkasubFactory{}
	tgr := f.New(config)
	golog.Printf("\t%s trigger created\n", testName)
	//runner := &TestRunner{}
	//tgr.Init(runner)
	golog.Printf("\t%s trigger initialized \n", testName)
	if configOnly {
		golog.Printf("Test %s complete\n", testName)
		return nil
	}
	defer tgr.Stop()
	error := tgr.Start()
	if !expectSucceed {
		if error == nil {
			return fmt.Errorf("Test was expected to fail, but didn't")
		}
		fmt.Printf("Test was expected to fail and did with error: %s", error)
		return nil
	}
	golog.Printf("\t%s listening for messages for %d seconds\n", testName, listentime)
	time.Sleep(time.Second * listentime)
	golog.Printf("Test %s complete\n", testName)
	return nil

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
