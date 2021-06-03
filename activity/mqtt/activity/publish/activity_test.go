package mqtt

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	_ "github.com/ayh20/flogo-components/activity/mqtt/connection"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

var settingsjson = `{
	"settings": {
		"connection": {
			"name": "myConn",
			"description": "Local MQTT Connection",
			"ref": "github.com/ayh20/flogo-components/activity/mqtt/connection",
			"settings": {
				"name": "myConn",
				"description": "Local MQTT Connection",
				"broker": "tcp://localhost:1883",
				"id": "myid",
				"user": "",
				"password": "",
				"enabletls": false,
				"certstore": "",
				"thing": ""
			}
		},
		"topic": "topic1",
        "qos": 1,
        "jsonpayload": false
	}
}`

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.ToMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func Test_One(t *testing.T) {
	log.RootLogger().Info("****TEST : Executing One start****")
	m := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(settingsjson), &m)

	log.RootLogger().Infof("Input Settings are : %v", m["settings"])
	assert.Nil(t, err1)

	mf := mapper.NewFactory(resolve.GetBasicResolver())

	support.RegisterAlias("connection", "connection", "github.com/ayh20/flogo-components/activity/mqtt/connection")

	//fmt.Println("=======Settings========", m["settings"])
	iCtx := test.NewActivityInitContext(m["settings"], mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("message", "My test message1")

	_, err = act.Eval(tc)

	// Getting outputs
	testOutput := tc.GetOutput("result")
	jsonOutput, _ := json.Marshal(testOutput)
	log.RootLogger().Infof("jsonOutput is : %s", string(jsonOutput))
	log.RootLogger().Info("****TEST : Executing Find One ends****")
	assert.Nil(t, err)
}
func Test_Two(t *testing.T) {
	log.RootLogger().Info("****TEST : Executing Two start****")
	m := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(settingsjson), &m)

	log.RootLogger().Infof("Input Settings are : %v", m["settings"])
	assert.Nil(t, err1)

	mf := mapper.NewFactory(resolve.GetBasicResolver())

	support.RegisterAlias("connection", "connection", "github.com/ayh20/flogo-components/activity/mqtt/connection")

	//fmt.Println("=======Settings========", m["settings"])
	iCtx := test.NewActivityInitContext(m["settings"], mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("message", "My test message2")

	_, err = act.Eval(tc)

	// Getting outputs
	testOutput := tc.GetOutput("result")
	jsonOutput, _ := json.Marshal(testOutput)
	log.RootLogger().Infof("jsonOutput is : %s", string(jsonOutput))
	log.RootLogger().Info("****TEST : Executing Find two ends****")
	assert.Nil(t, err)
}

// func TestCreate(t *testing.T) {
// 	act := NewActivity(getActivityMetadata())
// 	if act == nil {
// 		t.Error("Activity Not Created")
// 		t.Fail()
// 		return
// 	}
// }

// func TestEval(t *testing.T) {

// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()

// 	act := NewActivity(getActivityMetadata())
// 	tc := test.NewTestActivityContext(getActivityMetadata())
// 	//setup attrs

// 	fmt.Println("Publishing a flogo test message to topic 'flogo' on broker 'localhost:1883'")

// 	tc.SetInput("broker", "tcp://127.0.0.1:1883")
// 	tc.SetInput("id", "flogo_tester")
// 	tc.SetInput("topic", "flogo")
// 	tc.SetInput("qos", 0)
// 	tc.SetInput("enabletls", false)
// 	tc.SetInput("message", "This is a test message from flogo")

// 	act.Eval(tc)

// 	//check result attr
// 	result := tc.GetOutput("result")
// 	fmt.Println("result: ", result)

// 	if result == nil {
// 		t.Fail()
// 	}

// }

// func TestEvalTLSBosch(t *testing.T) {

// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()

// 	act := NewActivity(getActivityMetadata())
// 	tc := test.NewTestActivityContext(getActivityMetadata())
// 	//setup attrs

// 	fmt.Println("Publishing a flogo test message to topic 'event' on broker 'mqtt.bosch-iot-hub.com:8883'")

// 	tc.SetInput("broker", "ssl://mqtt.bosch-iot-hub.com:8883")
// 	tc.SetInput("id", "basicPubSub")
// 	tc.SetInput("topic", "event")
// 	tc.SetInput("user", "little-sensor@tcef56e88b16548f9a4a49cd5b92150af")
// 	tc.SetInput("password", "plaintextPassword")
// 	tc.SetInput("qos", 1)
// 	tc.SetInput("enabletls", true)
// 	tc.SetInput("certstore", "C:/Users/ahampshi/Documents/BoschIoTStuff/iothub.crt")
// 	tc.SetInput("message", "This is a test message from flogo")

// 	act.Eval(tc)

// 	//check result attr
// 	result := tc.GetOutput("result")
// 	fmt.Println("result: ", result)

// 	if result == nil {
// 		t.Fail()
// 	}

// }

// func TestEvalTLSAWS(t *testing.T) {

// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()

// 	act := NewActivity(getActivityMetadata())
// 	tc := test.NewTestActivityContext(getActivityMetadata())
// 	//setup attrs

// 	fmt.Println("Publishing a flogo test message to topic 'topic_1' on broker 'a1ck3umk9w128s-ats.iot.eu-west-1.amazonaws.com:8883'")

// 	tc.SetInput("broker", "ssl://a1ck3umk9w128s-ats.iot.eu-west-1.amazonaws.com:8883")
// 	tc.SetInput("id", "testclient")
// 	tc.SetInput("topic", "topic_1")
// 	tc.SetInput("qos", 0)
// 	tc.SetInput("enabletls", true)
// 	tc.SetInput("thing", "F1FlogoClient-Sam")
// 	tc.SetInput("certstore", "C:/Users/ahampshi/Documents/F1Demo/AMW_IoT/JavaStarter")
// 	tc.SetInput("message", "This is a test message from flogo")

// 	act.Eval(tc)

// 	//check result attr
// 	result := tc.GetOutput("result")
// 	fmt.Println("result: ", result)

// 	if result == nil {
// 		t.Fail()
// 	}

// }

// type jsonMessage struct {
// 	ObuReport struct {
// 		TruckID   string `json:"truckId"`
// 		DateTime  string `json:"dateTime"`
// 		FuelLevel string `json:"fuelLevel"`
// 		Position  struct {
// 			Lat   string `json:"lat"`
// 			Long  string `json:"long"`
// 			Speed string `json:"speed"`
// 		} `json:"position"`
// 	} `json:"obureport"`
// }

// func TestEvalJSON(t *testing.T) {

// 	testJSONstring := `{"obureport":{"truckId":"1","dateTime":"3435352424","fuelLevel":"50","position":{"lat":"47.34","long":"23.34","speed":"35"}}}`
// 	testData := jsonMessage{}
// 	json.Unmarshal([]byte(testJSONstring), &testData)
// 	log.Infof("Test data: %v", testData)

// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()

// 	act := NewActivity(getActivityMetadata())
// 	tc := test.NewTestActivityContext(getActivityMetadata())
// 	//setup attrs

// 	fmt.Println("Publishing a flogo test message to topic 'flogo' on broker 'localhost:1883'")

// 	tc.SetInput("broker", "tcp://127.0.0.1:1883")
// 	tc.SetInput("id", "flogo_tester")
// 	tc.SetInput("topic", "flogo")
// 	tc.SetInput("qos", 0)
// 	tc.SetInput("message", testData)

// 	act.Eval(tc)

// 	//check result attr
// 	result := tc.GetOutput("result")
// 	fmt.Println("result: ", result)

// 	if result == nil {
// 		t.Fail()
// 	}

// }
