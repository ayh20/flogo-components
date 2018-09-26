package aggregate

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

func TestEval(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivFunction, "moving_avg")
	tc.SetInput(ivWindowSize, 2)
	tc.SetInput(ivValue, 2)
	tc.SetInput(ivDataKey, "1")

	act.Eval(tc)

	report := tc.GetOutput(ovReport).(bool)
	result := tc.GetOutput(ovResult)

	if result != 0.0 {
		t.Errorf("Result is %d instead of 0", result)
	}
	if report {
		t.Error("Window should not report after first value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc2 := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc2.SetInput(ivFunction, "moving_avg")
	tc2.SetInput(ivWindowSize, 5)
	tc2.SetInput(ivValue, 3)
	tc2.SetInput(ivDataKey, "1")

	act.Eval(tc2)

	report = tc2.GetOutput(ovReport).(bool)
	result = tc2.GetOutput(ovResult)

	if result != 2.5 {
		t.Errorf("Result is %d instead of 2.5", result)
	}

	if !report {
		t.Error("Window should report after second value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc3 := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc3.SetInput(ivFunction, "moving_avg")
	tc3.SetInput(ivWindowSize, 5)
	tc3.SetInput(ivValue, 3)
	tc3.SetInput(ivDataKey, "1")

	act.Eval(tc3)

	report = tc3.GetOutput(ovReport).(bool)
	result = tc3.GetOutput(ovResult)

	if result != 3.0 {
		t.Errorf("Result is %d instead of 3.0", result)
	}

	if !report {
		t.Error("Window should report after third value")
	}

	fmt.Printf("test %v, %v", report, result)

}

func TestResetEval(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivFunction, "block_avg")
	tc.SetInput(ivWindowSize, 2)
	tc.SetInput(ivValue, 2)

	act.Eval(tc)

	report := tc.GetOutput(ovReport).(bool)
	result := tc.GetOutput(ovResult)

	if result != 0.0 {
		t.Errorf("Result is %d instead of 0", result)
	}
	if report {
		t.Error("Window should not report after first value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc2 := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc2.SetInput(ivFunction, "block_avg")
	tc2.SetInput(ivWindowSize, 2)
	tc2.SetInput(ivValue, 3)

	act.Eval(tc2)

	report = tc2.GetOutput(ovReport).(bool)
	result = tc2.GetOutput(ovResult)

	if result != 2.5 {
		t.Errorf("Result is %d instead of 2.5", result)
	}

	if !report {
		t.Error("Window should report after second value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc3 := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc3.SetInput(ivFunction, "block_avg")
	tc3.SetInput(ivWindowSize, 2)
	tc3.SetInput(ivValue, 3)

	act.Eval(tc3)

	report = tc3.GetOutput(ovReport).(bool)
	result = tc3.GetOutput(ovResult)

	if report {
		t.Error("Window should not report after third value")
	}

	if result != 0.0 {
		t.Errorf("Result is %d instead of 0.0", result)
	}

	fmt.Printf("test %v, %v", report, result)

}

func TestVaryingData(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivFunction, "block_avg")
	tc.SetInput(ivWindowSize, 2)
	tc.SetInput(ivValue, 2)
	tc.SetInput(ivDataKey, "1")

	act.Eval(tc)

	report := tc.GetOutput(ovReport).(bool)
	result := tc.GetOutput(ovResult)

	if result != 0.0 {
		t.Errorf("Result is %d instead of 0", result)
	}
	if report {
		t.Error("Window should not report after first value for key 1")
	}

	fmt.Printf("test %v, %v", report, result)

	tca := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tca.SetInput(ivFunction, "block_avg")
	tca.SetInput(ivWindowSize, 2)
	tca.SetInput(ivValue, 2)
	tca.SetInput(ivDataKey, "2")

	act.Eval(tca)

	report = tca.GetOutput(ovReport).(bool)
	result = tca.GetOutput(ovResult)

	if result != 0.0 {
		t.Errorf("Result is %d instead of 0", result)
	}
	if report {
		t.Error("Window should not report after first value for key 2")
	}

	fmt.Printf("test %v, %v", report, result)

	tc2 := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc2.SetInput(ivFunction, "block_avg")
	tc2.SetInput(ivWindowSize, 2)
	tc2.SetInput(ivValue, 3)
	tc2.SetInput(ivDataKey, "1")

	act.Eval(tc2)

	report = tc2.GetOutput(ovReport).(bool)
	result = tc2.GetOutput(ovResult)

	if result != 2.5 {
		t.Errorf("Result is %d instead of 2.5", result)
	}

	if !report {
		t.Error("Window should report after second value for key 1")
	}

	fmt.Printf("test %v, %v", report, result)

	tc3 := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc3.SetInput(ivFunction, "block_avg")
	tc3.SetInput(ivWindowSize, 2)
	tc3.SetInput(ivValue, 3)
	tc3.SetInput(ivDataKey, "1")

	act.Eval(tc3)

	report = tc3.GetOutput(ovReport).(bool)
	result = tc3.GetOutput(ovResult)

	if report {
		t.Error("Window should not report after third value for key 1")
	}

	if result != 0.0 {
		t.Errorf("Result is %d instead of 0.0", result)
	}

	fmt.Printf("test %v, %v", report, result)

	tca2 := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tca2.SetInput(ivFunction, "block_avg")
	tca2.SetInput(ivWindowSize, 2)
	tca2.SetInput(ivValue, 3)
	tca2.SetInput(ivDataKey, "2")

	act.Eval(tca2)

	report = tca2.GetOutput(ovReport).(bool)
	result = tca2.GetOutput(ovResult)

	if result != 2.5 {
		t.Errorf("Result is %d instead of 2.5", result)
	}

	if !report {
		t.Error("Window should report after second value for key 2")
	}

	fmt.Printf("test %v, %v", report, result)

}
