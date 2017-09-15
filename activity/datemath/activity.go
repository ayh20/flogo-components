package datemath

import (
	"strconv"
	"strings"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-datemath")

const (
	ivDate   = "date"
	ivAmount = "amount"
	ivUnit   = "unit"
	ivMode   = "function"

	ovResult = "result"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// Datemath is an Activity that is used to add or subtract a time internal to/from a date.
// The input are two strings plus the
// origin datatype and a compare mode ... ie "=" or ">"
// inputs : {input1, input2, datatype, comparemode}
// outputs: result (bool)
type Datemath struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &Datemath{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *Datemath) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Datemath) Eval(context activity.Context) (done bool, err error) {

	// Get the runtime values
	date, _ := context.GetInput(ivDate).(string)
	amount, _ := context.GetInput(ivAmount).(string)
	unit, _ := context.GetInput(ivUnit).(string)
	mode, _ := context.GetInput(ivMode).(string)

	dt1, err := time.Parse("2006-01-02T15:04:05-07:00", date)
	if err != nil {
		activityLog.Error("Input date format", err.Error())
		panic(err)
	}

	amt, _ := strconv.ParseInt(amount, 0, 16)

	if mode != "Add" {
		amt = -amt
	}

	activityLog.Info("amt = ", amt)

	var timeres time.Time

	switch unit {
	case "Day":
		{
			amt = amt * 24
			timeres = dt1.Add(time.Hour * time.Duration(amt))
		}
	case "Hour":
		{
			timeres = dt1.Add(time.Hour * time.Duration(amt))
		}
	case "Min":
		{
			timeres = dt1.Add(time.Minute * time.Duration(amt))
		}
	case "Sec":
		{
			timeres = dt1.Add(time.Second * time.Duration(amt))
		}
	}

	res := timeres.Format("2006-01-02T15:04:05-07:00")

	activityLog.Debug(strings.Join([]string{"Returning result", res}, " "))

	context.SetOutput(ovResult, res)

	return true, nil
}
