// Package readfile implements a file reader for Flogo
package readdir

// Imports
import (
	//"encoding/json"

	"os"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
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
	ctx.Logger().Debug("DirName: ", in.DirName)
	dirname, _ := coerce.ToString(in.DirName)

	// Check if the file exists
	_, err = os.Stat(dirname)
	if err != nil {
		ctx.Logger().Debugf("Error while tryinf to find file: %s", err.Error())
		return false, err
	}

	// Read the file
	osdirdata, err := os.ReadDir(dirname)
	if err != nil {
		ctx.Logger().Debugf("Error while reading file: %s\n", err.Error())
		return false, err
	}
	type Dirdata struct {
		Name      string
		Entrytype string
	}
	type Filedata struct {
		File []Dirdata
	}

	dirs := []Dirdata{}
	files := Filedata{dirs}

	for _, file := range osdirdata {
		filetype := "File"
		if file.IsDir() {
			filetype = "Dir"
		}
		dir := Dirdata{Name: file.Name(), Entrytype: filetype}
		files.File = append(files.File, dir)
	}

	//myjson, _ := json.Marshal(files)
	//fmt.Println(string(myjson))

	output.Result = files

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}
	return true, nil
}
