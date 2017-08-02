package compare

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func TestEvalQuality(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing Equals")
	//test1
	tc.SetInput("input1", "1")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "=")
	tc.SetInput("datatype", "int")
	act.Eval(tc)

	res := tc.GetOutput("result").(bool)
	msg := strconv.FormatBool(res)
	fmt.Println("1 = 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "=")
	tc.SetInput("datatype", "int")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2 = 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test3
	tc.SetInput("input1", "1")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "!=")
	tc.SetInput("datatype", "int")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("1 != 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "<")
	tc.SetInput("datatype", "int")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2 != 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
}
func TestEvalGT(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing Greater Than")
	//test1
	tc.SetInput("input1", "1")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", ">")
	act.Eval(tc)

	res := tc.GetOutput("result").(bool)
	msg := strconv.FormatBool(res)
	fmt.Println("1 > 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", ">")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2 > 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test3
	tc.SetInput("input1", "3")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", ">")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("3 > 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	fmt.Println("#######   Testing Greater Than or Equals")
	//test1
	tc.SetInput("input1", "1")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", ">=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("1 >= 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", ">=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2 >= 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test3
	tc.SetInput("input1", "3")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", ">=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("3 >= 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
}
func TestEvalLT(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing Less Than")
	//test1
	tc.SetInput("input1", "1")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "<")
	act.Eval(tc)

	res := tc.GetOutput("result").(bool)
	msg := strconv.FormatBool(res)
	fmt.Println("1 < 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "<")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2 < 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test3
	tc.SetInput("input1", "3")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "<")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("3 < 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	fmt.Println("#######   Testing Less Than or Equals")
	//test1
	tc.SetInput("input1", "1")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "<=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("1 <= 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "<=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2 <= 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test3
	tc.SetInput("input1", "3")
	tc.SetInput("input2", "2")
	tc.SetInput("comparemode", "<=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("3 <= 2: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
}
func TestEvalLTDecimal(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing Less Than")
	//test1
	tc.SetInput("input1", "1.234")
	tc.SetInput("input2", "2.345")
	tc.SetInput("comparemode", "<")
	act.Eval(tc)

	res := tc.GetOutput("result").(bool)
	msg := strconv.FormatBool(res)
	fmt.Println("1.234 < 2.345: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2.345")
	tc.SetInput("input2", "2.345")
	tc.SetInput("comparemode", "<")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2.345 < 2.345: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test3
	tc.SetInput("input1", "3.456")
	tc.SetInput("input2", "2.345")
	tc.SetInput("comparemode", "<")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("3.456 < 2.345: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	fmt.Println("#######   Testing Less Than or Equals")
	//test1
	tc.SetInput("input1", "1.234")
	tc.SetInput("input2", "2.345")
	tc.SetInput("comparemode", "<=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("1.234 <= 2.345: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test2
	tc.SetInput("input1", "2.345")
	tc.SetInput("input2", "2.345")
	tc.SetInput("comparemode", "<=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("2.345 <= 2.345: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}

	//test3
	tc.SetInput("input1", "3.456")
	tc.SetInput("input2", "2.345")
	tc.SetInput("comparemode", "<=")
	act.Eval(tc)

	res = tc.GetOutput("result").(bool)
	msg = strconv.FormatBool(res)
	fmt.Println("3.456 <= 2.345: ", msg)

	if tc.GetOutput("result") == nil {
		t.Fail()
	}
}
