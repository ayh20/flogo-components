package tcmtcipublisher

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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

func TestIt(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing ")
	//test1
	tc.SetInput("message", `{"_dest":"FlightData","hex":"48433a","flight":"KZR941","lat": 51.522354,"lon":-0.031771,"altitude":4225,"track":236,"speed":148,"messages":76 }`)
	tc.SetInput("key", "773b42654dec94de14b659a9c1f01c69")
	tc.SetInput("url", "wss://eu.messaging.cloud.tibco.com/tcm/01BKABHMAZKJDJA12PWHDR6WAP/channel")
	tc.SetInput("channel", "myChannel")
	act.Eval(tc)

	res := tc.GetOutput("result").(string)

	fmt.Println("Result: ", res)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

}
