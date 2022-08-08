package f1telemetry2021

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

	// Test for valid 2021 data..
	if unpHeader.PacketFormat != 2021 {
		ctx.Logger().Warnf("F1 Data: Unsupported packet format %v %v.%v %v %v", unpHeader.PacketFormat, unpHeader.GameMajorVersion, unpHeader.GameMinorVersion, unpHeader.PacketID, unpHeader.PacketVersion)

		return false, fmt.Errorf("F1 Data: Unsupported packet format %v %v.%v %v %v", unpHeader.PacketFormat, unpHeader.GameMajorVersion, unpHeader.GameMinorVersion, unpHeader.PacketID, unpHeader.PacketVersion)
	}

	output := &Output{}
	output.Data = ""
	output.MsgType = int(unpHeader.PacketID)

	outputHeader := fmt.Sprintf("%v,%v,%v,%g,%v,%v.%v.%v", unpHeader.PacketID, unpHeader.SessionUID, in.Source, unpHeader.SessionTime, unpHeader.PlayerCarIndex, unpHeader.PacketFormat, unpHeader.GameMajorVersion, unpHeader.GameMinorVersion)

	switch unpHeader.PacketID {
	case 0: //Motion
		// Unpack the 20 item car motion array
		// Note - Output array is:  Timestamp + array of car CSV data seprated by a "|"

		unpMotion := &F1CarMotion{}
		unpMotionExtra := &F1CarMotionExtra{}

		arraystring := ""

		for i := 0; i <= 21; i++ {
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
			ctx.Logger().Errorf("Unpack Fail: F1Session ", err.Error())
			return false, err
		}

		if unpSession.TrackID == 0 || unpSession.TrackID == 4 || unpSession.TrackID == 14 {
			unpSession.TrackID += 50
		}

		// Send all fields
		output.Data = outputHeader + "|" + getStrings(unpSession)

	case 2: //Lap Data
		// Unpack the 20 item lap data array
		// Note - Output array is:  Timestamp + array of car CSV data seprated by a "|"

		unpLapdata := &F1LapData{}

		arraystring := ""

		for i := 0; i <= 21; i++ {
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

		//  NOTE: Some events don't have data do we just get the code SSTA, SEND, DRSE, DRSD, LGOT, CHQF
		switch unpEvent.EventString {
		case "FTLP":
			unpEventFL := &F1EventDetailsFastestLap{}
			err = struc.Unpack(buf, unpEventFL)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1EventDetailsFastestLap ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventFL)

		case "RTMT", "TMPT", "RCWN", "DTSV", "SGSV":
			//
			unpEventExtra := &F1EventDetailsExtraIndex{}
			err = struc.Unpack(buf, unpEventExtra)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1EventDetailsExtraIndex ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventExtra)

		case "PENA":
			//
			unpEventPena := &F1EventDetailsPenalty{}
			err = struc.Unpack(buf, unpEventPena)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1EventDetailsPenalty ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventPena)

		case "SPTP":
			//
			unpEventSptp := &F1EventDetailsSpeedTrap{}
			err = struc.Unpack(buf, unpEventSptp)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1EventDetailsSpeedTrap ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventSptp)

		case "BUTN":
			//
			unpEventBut := &F1EventDetailsButtons{}
			err = struc.Unpack(buf, unpEventBut)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1EventDetailsButtons ", err.Error())
				return false, err
			}
			extradata = "," + fmt.Sprintf("%08b", unpEventBut.ButtonStatus) + "," + getStrings(unpEventBut)

		case "STLG":
			//
			unpEventSl := &F1EventDetailStartLights{}
			err = struc.Unpack(buf, unpEventSl)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1EventDetailStartLights ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventSl)

		case "FLBK":
			//
			unpEventFb := &F1EventDetailsFlashback{}
			err = struc.Unpack(buf, unpEventFb)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1EventDetailsFlashback ", err.Error())
				return false, err
			}
			extradata = "," + getStrings(unpEventFb)
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

		cnt := int(unpParticipant.NumActiveCars) - 1
		for i := 0; i <= cnt; i++ {
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

		for i := 0; i <= 21; i++ {
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

		for i := 0; i <= 21; i++ {
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

		for i := 0; i <= 21; i++ {
			err = struc.Unpack(buf, unpCarStatus)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarStatus ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("CarStatus unpacked: %v\n%+v\n", i, unpCarStatus)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpCarStatus)

		}
		output.Data = outputHeader + arraystring

	case 8: //Final Classification
		//ctx.Logger().Error(unpHeader.PacketID)
		unpFinalClassificationData := &F1FinalClassificationData{}
		unpFinalClassificationPacket := &F1FinalClassificationPacket{}

		arraystring := ""

		err = struc.Unpack(buf, unpFinalClassificationPacket)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1FinalClassificationPacket ", err.Error())
			return false, err
		}
		cnt := int(unpFinalClassificationPacket.NumCars) - 1
		for i := 0; i <= cnt; i++ {
			err = struc.Unpack(buf, unpFinalClassificationData)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1FinalClassificationData ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("F1FinalClassificationData unpacked: %v\n%+v\n", i, unpFinalClassificationData)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpFinalClassificationData)
		}

		output.Data = outputHeader + "|" + getStrings(unpFinalClassificationPacket) + arraystring

	case 9: //Lobby Info
		unpF1LobbyInfo := &F1LobbyInfo{}
		unpF1LobbyInfoData := &F1LobbyInfoData{}

		arraystring := ""

		err = struc.Unpack(buf, unpF1LobbyInfo)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1LobbyInfo ", err.Error())
			return false, err
		}
		cnt := int(unpF1LobbyInfo.NumPlayers) - 1
		for i := 0; i <= cnt; i++ {
			err = struc.Unpack(buf, unpF1LobbyInfoData)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1LobbyInfoData ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("F1LobbyInfoData unpacked: %v\n%+v\n", i, unpF1LobbyInfoData)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpF1LobbyInfoData)

		}
		output.Data = outputHeader + "|" + getStrings(unpF1LobbyInfo) + arraystring

	case 10: //Car Damage
		unpF1CarDamageData := &F1CarDamageData{}

		arraystring := ""

		for i := 0; i <= 21; i++ {
			err = struc.Unpack(buf, unpF1CarDamageData)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarDamageData ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("F1CarDamageData unpacked: %v\n%+v\n", i, unpF1CarDamageData)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpF1CarDamageData)

		}
		output.Data = outputHeader + arraystring

	case 11: //Session History

		unpF1SessionHistory := &F1SessionHistory{}
		// unpF1LapHistory := &F1LapHistory{}
		// unpF1TyreStintHistory := &F1TyreStintHistory{}

		err = struc.Unpack(buf, unpF1SessionHistory)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1SessionHistory ", err.Error())
			return false, err
		}

		// Send all fields
		output.Data = outputHeader + "|" + getStrings(unpF1SessionHistory)

		// sessHistory := "|" + getStrings(unpF1SessionHistory)

		// arraystring := ""
		// arraystring2 := ""

		// for i := 0; i <= 99; i++ {
		// 	err = struc.Unpack(buf, unpF1LapHistory)
		// 	if err != nil {
		// 		ctx.Logger().Debugf("Unpack Fail: F1LapHistory ", err.Error())
		// 		return false, err
		// 	}
		// 	ctx.Logger().Debugf("F1LapHistory unpacked: %v\n%+v\n", i, unpF1LapHistory)
		// 	arraystring = arraystring + fmt.Sprintf("|%v,", i) + getStrings(unpF1LapHistory)
		// }

		// for i := 0; i <= 7; i++ {
		// 	err = struc.Unpack(buf, unpF1TyreStintHistory)
		// 	if err != nil {
		// 		ctx.Logger().Debugf("Unpack Fail: F1TyreStintHistory ", err.Error())
		// 		return false, err
		// 	}
		// 	ctx.Logger().Debugf("F1TyreStintHistory unpacked: %v\n%+v\n", i, unpF1TyreStintHistory)
		// 	arraystring2 = arraystring2 + fmt.Sprintf("|%v,", i) + getStrings(unpF1TyreStintHistory)
		// }
		// output.Data = outputHeader + sessHistory + arraystring + arraystring2

	default:
		//fmt.Println("Error")
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
		//fmt.Printf("kind: %v", v.Field(i).Kind())
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			ss[i] = strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 32)
		case reflect.Array:
			ss[i] = strings.Trim(fmt.Sprintf("%v", v.Field(i)), "\x00")
			ss[i] = strings.ReplaceAll(ss[i], "} {", ";")
			ss[i] = strings.ReplaceAll(ss[i], "[{", "")
			ss[i] = strings.ReplaceAll(ss[i], "}]", "")
			ss[i] = strings.ReplaceAll(ss[i], " ", ":")
		default:
			ss[i] = strings.Trim(fmt.Sprintf("%v", v.Field(i)), "\x00")
		}
	}

	// Return the data as a CSV style string
	return strings.Join(ss, ",")
}
