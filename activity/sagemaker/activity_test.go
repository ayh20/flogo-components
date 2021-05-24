// Package amazons3 uploads or downloads files from Amazon Simple Storage Service (S3)
package sagemaker

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

func TestEvalDownload(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "us-west-2")
	tc.SetInput("assumeRole", true)
	tc.SetInput("roleARN", "arn:aws:iam::624719220700:role/TIBCO/Administrator")
	tc.SetInput("roleSessionName", "xxxx")
	tc.SetInput("endpointname", "xgboost-2021-05-21-15-23-14-953")
	//tc.SetInput("body", `{ "instances": [ { "start": "2018-03-13 00:00:00", "target": [100, 12] } ] }`)
	//tc.SetInput("body", "31,1,999,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,0,1,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,0,1,0")
	tc.SetInput("body", "57,2,999,1,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,1,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,1,0,0")
	//tc.SetInput("contenttype", "application/json")
	tc.SetInput("contenttype", "text/csv")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("Result is: [%s]", result)
}
