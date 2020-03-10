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
	fmt.Println("#######   Testing -  Data ###############")

	//var buf bytes.Buffer

	arr, err1 := hex.DecodeString("444154412a030000000000a8c1e4ab30398e9d30390e95303900c079c4b954c1c1d53e4b39d63e4b390800000000000000000000000000000000c079c400c079c400c079c400c079c400c079c40a00000000000000000000800000000000c079c400c079c400c079c400c079c400c079c40b0000003403cf379a00033ca33ba43b00c079c40000000000c079c400c079c400c079c40e0000000000803f0000803f00000080000000000000803f00c079c400c079c400c079c40f0000000012304073ece03d611fc7be00c079c400c079c400c079c400c079c400c079c410000000bb5086b718f50538637bb5b800c079c400c079c400c079c400c079c400c079c411000000fc3f8e3fcc5b233d7772f441e2e2024200c079c400c079c400c079c400c079c414000000e12a2d42752b06c1286f5f405562863e0000803f6d6f5f4000002e42000000c1150000008818f4c6c48feac2f36fb446c02355383179b5b6fdbf92380ad9ef3e5bafa1381a00000000000000000000000000000000000000000000000000000000000000000000002500000037782544ffff4743ffff4743ffff4743ffff4743ffff4743ffff4743ffff47432d000000041b894000000000000000000000000000000000000000000000000000000000840000000d9530395472ad3900c079c42c476f3300c079c400c079c400c079c400c079c4")

	if err1 != nil {
		fmt.Printf("error code: %v \n", err1)
		return
	}

	tc.SetInput("buffer", arr)

	fmt.Println("#######   Call Activity")
	act.Eval(tc)
	fmt.Println("#######   Activity Result")

	rtype := tc.GetOutput("msgtype")
	rdata := tc.GetOutput("data")

	fmt.Printf("Msg Type: %v \n", rtype)
	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("msgtype") == nil {
		t.Fail()
	}

}
