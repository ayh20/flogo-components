package log

import (
	"strconv"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-compare")

const (
	ivInput1      = "input1"
	ivInput2      = "input2"
	ivDataType    = "datatype"
	ivCompareMode = "comparemode"

	ovResult = "result"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// Compare is an Activity that is used to compare two values. The input are two strings plus the
// origin datatype and a compare mode ... ie "=" or ">"
// inputs : {input1, input2, datatype, comparemode}
// outputs: result (bool)
type Compare struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &Compare{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *Compare) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Compare) Eval(context activity.Context) (done bool, err error) {

	// Get the runtime values
	input1, _ := context.GetInput(ivInput1).(string)
	input2, _ := context.GetInput(ivInput2).(string)
	datatype, _ := context.GetInput(ivDataType).(string)
	comparemode, _ := context.GetInput(ivCompareMode).(string)

	res := false

	//  perform the compare based on data type ... look for better way of doing this
	switch datatype {
	case "int":
		int1, _ := strconv.ParseInt(input1, 0, 64)
		int2, _ := strconv.ParseInt(input2, 0, 64)
		res = compareInt(int1, int2, comparemode)
	case "uint":
		int1, _ := strconv.ParseUint(input1, 0, 64)
		int2, _ := strconv.ParseUint(input2, 0, 64)
		res = compareUint(int1, int2, comparemode)
	case "float":
		int1, _ := strconv.ParseFloat(input1, 64)
		int2, _ := strconv.ParseFloat(input2, 64)
		res = compareFloat(int1, int2, comparemode)
	}

	activityLog.Info(strings.Join([]string{"Returning result", strconv.FormatBool(res)}, " "))

	context.SetOutput(ovResult, res)

	return true, nil
}

func compareInt(num1 int64, num2 int64, compare string) bool {

	activityLog.Info("Compare Int")

	switch compare {
	case "=", "==":
		if num1 == num2 {
			return true
		}
	case "!=":
		if num1 != num2 {
			return true
		}
	case ">":
		if num1 > num2 {
			return true
		}
	case "<":
		if num1 < num2 {
			return true
		}
	case ">=":
		if num1 >= num2 {
			return true
		}
	case "<=":
		if num1 <= num2 {
			return true
		}
	default:
		return false

	}
	return false
}
func compareUint(num1, num2 uint64, compare string) bool {

	activityLog.Info("Compare Uint")

	switch compare {
	case "=", "==":
		if num1 == num2 {
			return true
		}
	case "!=":
		if num1 != num2 {
			return true
		}
	case ">":
		if num1 > num2 {
			return true
		}
	case "<":
		if num1 < num2 {
			return true
		}
	case ">=":
		if num1 >= num2 {
			return true
		}
	case "<=":
		if num1 <= num2 {
			return true
		}
	default:
		return false

	}
	return false
}
func compareFloat(num1, num2 float64, compare string) bool {

	activityLog.Info("Compare float64")

	switch compare {
	case "=", "==":
		if num1 == num2 {
			return true
		}
	case "!=":
		if num1 != num2 {
			return true
		}
	case ">":
		if num1 > num2 {
			return true
		}
	case "<":
		if num1 < num2 {
			return true
		}
	case ">=":
		if num1 >= num2 {
			return true
		}
	case "<=":
		if num1 <= num2 {
			return true
		}
	default:
		return false

	}
	return false
}
