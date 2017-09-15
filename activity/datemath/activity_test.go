package datemath

import (
	"fmt"
	"io/ioutil"
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

func TestEvalSubtract(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing Subtract")
	//test1
	tc.SetInput("date", "2017-09-14T09:09:09+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Day")
	tc.SetInput("function", "subtract")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res := tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

	//test1
	tc.SetInput("date", "2017-09-14T09:09:09+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Hour")
	tc.SetInput("function", "subtract")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res = tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

	//test1
	tc.SetInput("date", "2017-09-14T09:09:09+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Min")
	tc.SetInput("function", "subtract")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res = tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

	//test1
	tc.SetInput("date", "2017-09-14T09:09:09+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Sec")
	tc.SetInput("function", "subtract")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res = tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

}

func TestEvalAddy(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing Subtract")
	//test1
	tc.SetInput("date", "2017-09-24T09:09:09+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Day")
	tc.SetInput("function", "Add")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res := tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

	//test1
	tc.SetInput("date", "2017-09-14T14:09:09+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Hour")
	tc.SetInput("function", "Add")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res = tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

	//test1
	tc.SetInput("date", "2017-09-14T09:55:09+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Min")
	tc.SetInput("function", "Add")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res = tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

	//test1
	tc.SetInput("date", "2017-09-14T09:09:55+00:00")
	tc.SetInput("amount", "10")
	tc.SetInput("unit", "Sec")
	tc.SetInput("function", "Add")
	act.Eval(tc)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
	res = tc.GetOutput("result")
	fmt.Println("Start Date:", tc.GetInput("date"), " ", tc.GetInput("function"), " ", tc.GetInput("amount"), " ", tc.GetInput("unit"), " Result", res)

}
