package mqtt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/project-flogo/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestCreate(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	//setup attrs

	fmt.Println("Publishing a flogo test message to topic 'flogo' on broker 'localhost:1883'")

	tc.SetInput("broker", "tcp://127.0.0.1:1883")
	tc.SetInput("id", "flogo_tester")
	tc.SetInput("topic", "flogo")
	tc.SetInput("qos", 0)
	tc.SetInput("enabletls", false)
	tc.SetInput("message", "This is a test message from flogo")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Println("result: ", result)

	if result == nil {
		t.Fail()
	}

}

func TestEvalTLSBosch(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	//setup attrs

	fmt.Println("Publishing a flogo test message to topic 'event' on broker 'mqtt.bosch-iot-hub.com:8883'")

	tc.SetInput("broker", "ssl://mqtt.bosch-iot-hub.com:8883")
	tc.SetInput("id", "basicPubSub")
	tc.SetInput("topic", "event")
	tc.SetInput("user", "little-sensor@tcef56e88b16548f9a4a49cd5b92150af")
	tc.SetInput("password", "plaintextPassword")
	tc.SetInput("qos", 1)
	tc.SetInput("enabletls", true)
	tc.SetInput("certstore", "C:/Users/ahampshi/Documents/BoschIoTStuff/iothub.crt")
	tc.SetInput("message", "This is a test message from flogo")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Println("result: ", result)

	if result == nil {
		t.Fail()
	}

}

func TestEvalTLSAWS(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	//setup attrs

	fmt.Println("Publishing a flogo test message to topic 'topic_1' on broker 'a1ck3umk9w128s-ats.iot.eu-west-1.amazonaws.com:8883'")

	tc.SetInput("broker", "ssl://a1ck3umk9w128s-ats.iot.eu-west-1.amazonaws.com:8883")
	tc.SetInput("id", "testclient")
	tc.SetInput("topic", "topic_1")
	tc.SetInput("qos", 0)
	tc.SetInput("enabletls", true)
	tc.SetInput("thing", "F1FlogoClient-Sam")
	tc.SetInput("certstore", "C:/Users/ahampshi/Documents/F1Demo/AMW_IoT/JavaStarter")
	tc.SetInput("message", "This is a test message from flogo")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Println("result: ", result)

	if result == nil {
		t.Fail()
	}

}

type jsonMessage struct {
	ObuReport struct {
		TruckID   string `json:"truckId"`
		DateTime  string `json:"dateTime"`
		FuelLevel string `json:"fuelLevel"`
		Position  struct {
			Lat   string `json:"lat"`
			Long  string `json:"long"`
			Speed string `json:"speed"`
		} `json:"position"`
	} `json:"obureport"`
}

func TestEvalJSON(t *testing.T) {

	testJSONstring := `{"obureport":{"truckId":"1","dateTime":"3435352424","fuelLevel":"50","position":{"lat":"47.34","long":"23.34","speed":"35"}}}`
	testData := jsonMessage{}
	json.Unmarshal([]byte(testJSONstring), &testData)
	log.Infof("Test data: %v", testData)

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	//setup attrs

	fmt.Println("Publishing a flogo test message to topic 'flogo' on broker 'localhost:1883'")

	tc.SetInput("broker", "tcp://127.0.0.1:1883")
	tc.SetInput("id", "flogo_tester")
	tc.SetInput("topic", "flogo")
	tc.SetInput("qos", 0)
	tc.SetInput("message", testData)

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Println("result: ", result)

	if result == nil {
		t.Fail()
	}

}
