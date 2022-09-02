package parseparquet

import (
	"encoding/json"
	"fmt"

	// do not use, old version
	"github.com/project-flogo/core/activity"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

// Activity is a Parquet parser
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

// ActivityLog is the default logger for the Log Activity
//var activityLog = logger.GetLogger("activity-flogo-parseparquet")

// const (
// 	ivinitRow     = "initRow"
// 	ivmaxRows     = "maxRows"
// 	ivparquetFile = "parquetFile"
// 	ovOutput      = "output"
// )

type RowType struct {
	FirstName string
	LastName  string
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	// Get the runtime values
	ctx.Logger().Debug("Starting")

	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	initRow := in.InitRow
	maxRows := in.MaxRows
	parquetFile := in.ParquetFile

	ctx.Logger().Debugf("Processing file: %s, [%s-%s] ", parquetFile, initRow, maxRows)

	fr, error := local.NewLocalFileReader(parquetFile)
	if error != nil {
		return false, fmt.Errorf("error opening the specified file: %v", error)
	}

	pr, error := reader.NewParquetReader(fr, nil, 4)
	if error != nil {
		return false, fmt.Errorf("error reading the specified file: %v", error)
	}

	num := int(pr.GetNumRows())

	res, error := pr.ReadByNumber(num)
	if error != nil {
		ctx.Logger().Debugf("Read Fail ", error.Error())
		return false, error
	}

	jsonBs, error := json.Marshal(res)
	if error != nil {
		ctx.Logger().Debugf("Marshal Fail ", error.Error())
		return false, error
	}

	fmt.Println("JSON:")
	fmt.Println(string(jsonBs))

	pr.ReadStop()
	fr.Close()

	output := &Output{}

	output.Output = string(jsonBs)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
