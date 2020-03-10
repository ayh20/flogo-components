package xplanetelemetry

import (
	"encoding/hex"
	"fmt"

	//"io/ioutil"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

/* func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	aInput := &Input{XmlData: `<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`}
	tc.SetInputObject(aInput)
	done, _ := act.Eval(tc)
	assert.True(t, done)
	aOutput := &Output{}
	err := tc.GetOutputObject(aOutput)
	assert.Nil(t, err)
	assert.Equal(t, "world", aOutput.JsonObject["hello"])
} */

func TestEvalQuality(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	fmt.Println(" ")
	fmt.Println("#######   Testing -  Data ###############################################################")

	//var buf bytes.Buffer

	arr, err1 := hex.DecodeString("8c16459ba25a50e54934e40a0800450000b1f84300008011be79c0a8010ac0a80124bf69bf68009d7e1f444154412a000000001ba1894133339f4100c079c4bc166e3d96133c3db6b0743d0000803f0000803f03000000000000006e39f0368eabf136a99ff13600c079c400000000f50d0b37f50d0b370b000000734034b21cacea31bb7cc73100c079c40000000000c079c400c079c400c079c425000000aebd8344aebd8344000000000000000000000000000000000000000000000000")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}

	tc.SetInput("buffer", arr)
	//err := struc.Pack(&buf, arr)

	//o := &F1Header{}

	//err = struc.Unpack(&buf, o)
	//if err != nil {
	//	fmt.Printf("error code: %v \n", err)
	//	return
	//}

	//fmt.Printf("F1 Header : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rtype := tc.GetOutput("msgtype")
	rdata := tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

}
