package awsgetsecret

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
	tc.SetInput("secretARN", "arn:aws:secretsmanager:us-west-2:624719220700:secret:ProdTCI-unhgur")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("secret")
	fmt.Printf("Result is: [%s]", result)
}
