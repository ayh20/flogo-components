package End

import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-end")

const ()

// EndActivity is a stub for your Activity implementation
type EndActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &EndActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *EndActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *EndActivity) Eval(ctx activity.Context) (done bool, err error) {
	os.Exit(0)

	return true, nil
}
