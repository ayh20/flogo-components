// Package readfile implements a file reader for Flogo
package readdir

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

func TestEvalReadNonExistingFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	// Set required attributes
	tc.SetInput("dirname", `C:\tmpx`)

	// Execute the activity
	_, err := act.Eval(tc)

	assert.Error(t, err, "")
}

func TestEvalReadExistingFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	// Set required attributes
	tc.SetInput("dirname", "C:\\tmp")

	// Execute the activity
	_, err := act.Eval(tc)

	assert.NoError(t, err, "")

	//check result attr
	aOutput := &Output{}
	err = tc.GetOutputObject(aOutput)
	assert.Nil(t, err)
	result := aOutput.Result
	fmt.Println(result)

	//assert.Contains(t, result, "This is some data in a file to read ...")

}
