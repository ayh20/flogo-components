package awsgetsecret

import (
	"fmt"
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

func TestGetStringSecret(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

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
