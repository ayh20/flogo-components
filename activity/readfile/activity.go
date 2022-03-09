// Package readfile implements a file reader for Flogo
package readfile

// Imports
import (
	"io/ioutil"
	"os"

	"github.com/project-flogo/core/activity"
)

type Activity struct {
	//metadata *activity.Metadata
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

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	ctx.Logger().Debug("In Eval")
	in := &Input{}
	output := &Output{}

	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}
	ctx.Logger().Debug("FileName: ", in.Filename)

	// Check if the file exists
	_, err = os.Stat(in.Filename)
	if err != nil {
		ctx.Logger().Debugf("Error while tryinf to find file: %s", err.Error())
		return false, err
	}

	// Read the file
	fileBytes, err := ioutil.ReadFile(in.Filename) //os.ReadFile(in.Filename)
	if err != nil {
		ctx.Logger().Debugf("Error while reading file: %s\n", err.Error())
		return false, err
	}

	output.Result = string(fileBytes)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}
	return true, nil
}
