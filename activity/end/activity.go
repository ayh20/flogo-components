package end

import (
	"os"

	"github.com/project-flogo/core/activity"
)

type Activity struct {
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {

	ctx.Logger().Info("In New activity")

	act := &Activity{}
	return act, nil
}

// Eval implements api.Activity.Eval - Logs the Message
//func (a *f1telemetry) Eval(context activity.Context) (done bool, err error) {
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	os.Exit(0)

	return true, nil
}
