package f1telemetry2020

import (
	"github.com/project-flogo/core/activity"

	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/lunixbochs/struc"
)

// Activity is a F1 Telemetery decoder activity
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

	// Get the runtime values
	ctx.Logger().Debug("Starting")

	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	//input, _ := context.GetInput(ivInput).([]byte)
	buf := bytes.NewBuffer(in.Buffer)

	ctx.Logger().Debugf("input : \n %x \n", in.Buffer)

	// Create structs to hold unpacked data
	unpHeader := &F1Header{}

	ctx.Logger().Debug("Unpack Header")

	// Unpack the Header
	err = struc.Unpack(buf, unpHeader)
	if err != nil {
		ctx.Logger().Debug("Unpack Fail: F1Header ", err.Error())
		return false, err
	}

	// dump header
	ctx.Logger().Debugf("struct F1Header : \n %+v \n", unpHeader)

	// Test for valid 2019 data..
	if unpHeader.PacketFormat != 2019 {
		ctx.Logger().Debugf("F1 Data: Unsupported packet format %v", unpHeader.PacketFormat)
		return false, fmt.Errorf("F1 Data: Unsupported packet format %v", unpHeader.PacketFormat)
	}

	output := &Output{}
	output.Data = ""
	output.MsgType = int(unpHeader.PacketID)

	outputHeader := fmt.Sprintf("%v,%v,%g,%v,%v.%v", unpHeader.PacketID, unpHeader.SessionUID, unpHeader.SessionTime, unpHeader.PlayerCarIndex, unpHeader.GameMajorVersion, unpHeader.GameMinorVersion)

	switch unpHeader.PacketID {
	case 0: //Motion
		// Unpack the 20 item car motion array
		// Note - Output array is:  Timestamp + array of car CSV data seprated by a "|"

		unpMotion := &F1CarMotion{}
		unpMotionExtra := &F1CarMotionExtra{}

		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpMotion)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarMotion %v", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("Car Array unpacked: %v\n%+v\n", i, unpMotion)

			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpMotion)

		}

		err = struc.Unpack(buf, unpMotionExtra)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1CarMotionExtra ", err.Error())
			return false, err
		}
		// Send all fields
		output.Data = outputHeader + "|" + getStrings(unpMotionExtra) + arraystring

	case 1: //Session
		unpSession := &F1Session{}

		err = struc.Unpack(buf, unpSession)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1Session ", err.Error())
			return false, err
		}

		// Send all fields
		output.Data = outputHeader + "|" + getStrings(unpSession)

	case 2: //Lap Data
		// Unpack the 20 item lap data array
		// Note - Output array is:  Timestamp + array of car CSV data seprated by a "|"

		unpLapdata := &F1LapData{}

		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpLapdata)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1LapData ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("LapData unpacked: %v\n%+v\n", i, unpLapdata)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpLapdata)
		}
		output.Data = outputHeader + arraystring

	case 3: //Event
		unpEvent := &F1Event{}
		extradata := ""

		err = struc.Unpack(buf, unpEvent)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1Event ", err.Error())
			return false, err
		}

		switch unpEvent.EventString {
		case "FTLP":
			unpEventFL := &F1EventDetailsFastestLap{}
			err = struc.Unpack(buf, unpEventFL)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1LapData ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventFL)

		case "RTMT", "TMPT", "RCWN":
			//
			unpEventExtra := &F1EventDetailsExtraIndex{}
			err = struc.Unpack(buf, unpEventExtra)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1LapData ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventExtra)
		}

		output.Data = outputHeader + "|" + unpEvent.EventString + extradata

	case 4: //Participants
		unpParticipant := &F1Participant{}
		unpParticpantData := &F1ParticipantData{}

		err = struc.Unpack(buf, unpParticipant)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1Participant ", err.Error())
			return false, err
		}
		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpParticpantData)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1ParticipantData ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("F1Participant unpacked: %v\n%+v\n", i, unpParticpantData)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpParticpantData)

		}
		output.Data = outputHeader + fmt.Sprintf("|%v", unpParticipant.NumActiveCars) + arraystring

	case 5: //Car Setups
		unpCarSetupData := &F1SetupData{}

		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpCarSetupData)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarSetupData ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("F1CarSetupData unpacked: %v\n%+v\n", i, unpCarSetupData)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpCarSetupData)
		}
		output.Data = outputHeader + arraystring

	case 6: //Car Telemetery
		unpCarTelemetry := &F1CarTelemetryData{}
		unpCarTelemetryExtra := &F1CarTelemetryDataExtra{}

		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpCarTelemetry)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarTelemetry ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("Car Array unpacked: %v\n%+v\n", i, unpCarTelemetry)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpCarTelemetry)

		}

		err = struc.Unpack(buf, unpCarTelemetryExtra)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1CarTelemetryExtra ", err.Error())
			return false, err
		}
		// Send all fields
		output.Data = outputHeader + "|" + getStrings(unpCarTelemetryExtra) + arraystring

	case 7: //Car Status
		unpCarStatus := &F1CarStatus{}
		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpCarStatus)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarStatus ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("CarStatus unpacked: %v\n%+v\n", i, unpCarStatus)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpCarStatus)

		}
		output.Data = outputHeader + arraystring

	default:
		fmt.Println("Error")
		return false, fmt.Errorf("F1 Data: Undefined packet ID %v", unpHeader.PacketID)
	}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getStrings(iface interface{}) string {

	//  The function is passed a pointer to a struct. We use Indirect to get the actual value of the passed struct
	v := reflect.Indirect(reflect.ValueOf(iface))

	// Create a slice that is the correct size for the struct
	ss := make([]string, v.NumField())

	// run through each field and based on it's type format the value as a string
	for i := range ss {
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			ss[i] = strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 32)
		default:
			ss[i] = strings.Trim(fmt.Sprintf("%v", v.Field(i)), "\x00")
		}
	}

	// Return the data as a CSV style string
	return strings.Join(ss, ",")
}
