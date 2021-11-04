package filetostring

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-readfile")

const (
	ivFile   = "filename"
	ovOutput = "output"
)

// ReadFileActivity is a stub for your Activity implementation
type ReadFileActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ReadFileActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *ReadFileActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ReadFileActivity) Eval(ctx activity.Context) (done bool, err error) {

	b, err := ioutil.ReadFile(ctx.GetInput(ivFile).(string)) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//str := string(b) // convert content to a 'string'
	str := strings.ReplaceAll(string(b), "\r\n", " ")

	activityLog.Debugf("Data read from file: %s", str)
	ctx.SetOutput(ovOutput, str)

	return true, nil
}
