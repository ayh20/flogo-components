package end

import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-end")

const ()

// endActivity is a stub for your Activity implementation
type endActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &endActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *endActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *endActivity) Eval(ctx activity.Context) (done bool, err error) {
	os.Exit(0)

	return true, nil
}
