package f1telemetry2019proto

import (
	"github.com/project-flogo/core/activity"
	"google.golang.org/protobuf/proto"

	"bytes"
	"fmt"
	"time"

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

var nsMid = time.Now().UnixNano()

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

	nsMid = time.Now().UnixNano() - time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).UnixNano()

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
	var iCurrentPlayer int32 = int32(unpHeader.PlayerCarIndex)

	output := &Output{}
	//output.Data = ""
	output.MsgType = int(unpHeader.PacketID)

	//outputHeader := fmt.Sprintf("%v,%v,%g,%v,%v.%v", unpHeader.PacketID, unpHeader.SessionUID, unpHeader.SessionTime, unpHeader.PlayerCarIndex, unpHeader.GameMajorVersion, unpHeader.GameMinorVersion)

	td := &TelemetryData{
		FeedGUID:    "ygyugfyawfgawyilgfayilrg",
		FeedName:    "eSportTelemData",
		StreamId:    "eSportAdaptotFeed1",
		StreamType:  StreamType_STREAM_TYPE_LIVE,
		FeedType:    DataFeedType_DATA_FEED_TYPE_TELEMETRY,
		Source:      "eSportAdaptor",
		Frequency:   60,
		Quality:     90,
		Format:      fmt.Sprintf("%v %v %v %v", unpHeader.PacketFormat, unpHeader.GameMajorVersion, unpHeader.GameMinorVersion, unpHeader.PacketVersion),
		SessionGUID: fmt.Sprintf("%v", unpHeader.SessionUID),
		EpochNano:   nsMid,
		Identifier:  fmt.Sprintf("%v", unpHeader.FrameIdentifier),
	}

	switch unpHeader.PacketID {
	case 0: //Motion
		// Unpack the 20 item car motion array
		unpMotion := &F1CarMotion{}
		unpMotionExtra := &F1CarMotionExtra{}

		// First task is to create the data for the "current driver"
		// we have two indexes ... one for the "drivers" car iDPDriver and one for the rest
		var iDPDriver int32 = 900
		var iDP int32 = 2000

		// this is the index used for a loop iteration
		var iDPlocal int32 = 0

		// loop through the buffer unpacking each motion packet
		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpMotion)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarMotion %v", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("Car Array unpacked: %v\n%+v\n", i, unpMotion)

			if i == int(iCurrentPlayer) {
				iDPlocal = iDPDriver
			} else {
				iDPlocal = iDP
			}

			td.DataPoints = append(td.DataPoints,
				setDataPoint(iDPlocal, float64(unpMotion.X)),
				setDataPoint(iDPlocal+1, float64(unpMotion.Y)),
				setDataPoint(iDPlocal+2, float64(unpMotion.Z)),
				setDataPoint(iDPlocal+3, float64(unpMotion.Xv)),
				setDataPoint(iDPlocal+4, float64(unpMotion.Yv)),
				setDataPoint(iDPlocal+5, float64(unpMotion.Zv)),
				setDataPoint(iDPlocal+6, float64(unpMotion.Xd)),
				setDataPoint(iDPlocal+7, float64(unpMotion.Yd)),
				setDataPoint(iDPlocal+8, float64(unpMotion.Zd)),
				setDataPoint(iDPlocal+9, float64(unpMotion.Xr)),
				setDataPoint(iDPlocal+10, float64(unpMotion.Yr)),
				setDataPoint(iDPlocal+11, float64(unpMotion.Zr)),
				setDataPoint(iDPlocal+12, float64(unpMotion.Gforcelat)),
				setDataPoint(iDPlocal+13, float64(unpMotion.Gforcelon)),
				setDataPoint(iDPlocal+14, float64(unpMotion.Gforcevert)),
				setDataPoint(iDPlocal+15, float64(unpMotion.Yaw)),
				setDataPoint(iDPlocal+16, float64(unpMotion.Pitch)),
				setDataPoint(iDPlocal+17, float64(unpMotion.Roll)),
			)
			iDP += 30

		}

		// unpack the trailing extra data for the player
		err = struc.Unpack(buf, unpMotionExtra)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1CarMotionExtra ", err.Error())
			return false, err
		}

		// Format the datapoints
		td.DataPoints = append(td.DataPoints,
			setDataPoint(iDPDriver+18, float64(unpMotionExtra.SuspPosRL)),
			setDataPoint(iDPDriver+19, float64(unpMotionExtra.SuspPosRR)),
			setDataPoint(iDPDriver+20, float64(unpMotionExtra.SuspPosFL)),
			setDataPoint(iDPDriver+21, float64(unpMotionExtra.SuspPosFR)),
			setDataPoint(iDPDriver+22, float64(unpMotionExtra.SuspVelRL)),
			setDataPoint(iDPDriver+23, float64(unpMotionExtra.SuspVelRR)),
			setDataPoint(iDPDriver+24, float64(unpMotionExtra.SuspVelFL)),
			setDataPoint(iDPDriver+25, float64(unpMotionExtra.SuspVelFR)),
			setDataPoint(iDPDriver+26, float64(unpMotionExtra.SuspAccelerationRL)),
			setDataPoint(iDPDriver+27, float64(unpMotionExtra.SuspAccelerationRR)),
			setDataPoint(iDPDriver+28, float64(unpMotionExtra.SuspAccelerationFL)),
			setDataPoint(iDPDriver+29, float64(unpMotionExtra.SuspAccelerationFR)),
			setDataPoint(iDPDriver+30, float64(unpMotionExtra.WheelspeedRL)),
			setDataPoint(iDPDriver+31, float64(unpMotionExtra.WheelspeedRR)),
			setDataPoint(iDPDriver+32, float64(unpMotionExtra.WheelspeedFL)),
			setDataPoint(iDPDriver+33, float64(unpMotionExtra.WheelspeedFR)),
			setDataPoint(iDPDriver+34, float64(unpMotionExtra.WheelslipRL)),
			setDataPoint(iDPDriver+35, float64(unpMotionExtra.WheelslipRR)),
			setDataPoint(iDPDriver+36, float64(unpMotionExtra.WheelslipFL)),
			setDataPoint(iDPDriver+37, float64(unpMotionExtra.WheelslipFR)),
			setDataPoint(iDPDriver+38, float64(unpMotionExtra.XLocalVelocity)),
			setDataPoint(iDPDriver+39, float64(unpMotionExtra.YLocalVelocity)),
			setDataPoint(iDPDriver+40, float64(unpMotionExtra.ZLocalVelocity)),
			setDataPoint(iDPDriver+41, float64(unpMotionExtra.Angvelx)),
			setDataPoint(iDPDriver+42, float64(unpMotionExtra.Angvely)),
			setDataPoint(iDPDriver+43, float64(unpMotionExtra.Angvelz)),
			setDataPoint(iDPDriver+44, float64(unpMotionExtra.AngAccX)),
			setDataPoint(iDPDriver+45, float64(unpMotionExtra.AngAccY)),
			setDataPoint(iDPDriver+46, float64(unpMotionExtra.AngAccZ)),
			setDataPoint(iDPDriver+47, float64(unpMotionExtra.FrontWheelsAngle)),
		)

		// put the data into the return message
		output.Data, err = proto.Marshal(td)

	case 1: //Session
		unpSession := &F1Session{}

		err = struc.Unpack(buf, unpSession)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1Session ", err.Error())
			return false, err
		}

		// Add datapoints
		td.DataPoints = append(td.DataPoints,
			setDataPoint(948, float64(unpSession.Weather)),
			setDataPoint(949, float64(unpSession.TrackTemperature)),
			setDataPoint(950, float64(unpSession.AirTemperature)),
			setDataPoint(951, float64(unpSession.SessionTimeLeft)),
			setDataPoint(952, float64(unpSession.SessionDuration)),
			setDataPoint(953, float64(unpSession.GamePaused)),
			setDataPoint(954, float64(unpSession.ZoneFlag1)),
			setDataPoint(955, float64(unpSession.ZoneFlag2)),
			setDataPoint(956, float64(unpSession.ZoneFlag3)),
			setDataPoint(957, float64(unpSession.ZoneFlag4)),
			setDataPoint(958, float64(unpSession.ZoneFlag5)),
			setDataPoint(959, float64(unpSession.ZoneFlag6)),
			setDataPoint(960, float64(unpSession.ZoneFlag7)),
			setDataPoint(961, float64(unpSession.ZoneFlag8)),
			setDataPoint(962, float64(unpSession.ZoneFlag9)),
			setDataPoint(963, float64(unpSession.ZoneFlag10)),
			setDataPoint(964, float64(unpSession.ZoneFlag11)),
			setDataPoint(965, float64(unpSession.ZoneFlag12)),
			setDataPoint(966, float64(unpSession.ZoneFlag13)),
			setDataPoint(967, float64(unpSession.ZoneFlag14)),
			setDataPoint(968, float64(unpSession.ZoneFlag15)),
			setDataPoint(969, float64(unpSession.ZoneFlag16)),
			setDataPoint(970, float64(unpSession.ZoneFlag17)),
			setDataPoint(971, float64(unpSession.ZoneFlag18)),
			setDataPoint(972, float64(unpSession.ZoneFlag19)),
			setDataPoint(973, float64(unpSession.ZoneFlag20)),
			setDataPoint(974, float64(unpSession.ZoneFlag21)),
			setDataPoint(975, float64(unpSession.SafetyCarStatus)),
			setDataPoint(976, float64(unpSession.IsSpectating)),
			setDataPoint(977, float64(unpSession.SpectatorCarIndex)),
		)
		output.Data, err = proto.Marshal(td)

		//create sessiondata object
		sd := &SessionData{
			FeedGUID:    "ygyugfyawfgawyilgfayilrg",
			FeedName:    "eSportSessData",
			StreamId:    "eSportAdaptotFeed1",
			StreamType:  StreamType_STREAM_TYPE_LIVE,
			Source:      "eSportAdaptor",
			Quality:     90,
			SessionGUID: fmt.Sprintf("%v", unpHeader.SessionUID),
			EpochNano:   nsMid,
			Identifier:  fmt.Sprintf("%v", unpHeader.FrameIdentifier),
		}
		sd.Details = append(sd.Details,
			setNameValue("ZoneStart1", fmt.Sprintf("%f", unpSession.ZoneStart1)),
			setNameValue("ZoneStart2", fmt.Sprintf("%f", unpSession.ZoneStart2)),
			setNameValue("ZoneStart3", fmt.Sprintf("%f", unpSession.ZoneStart3)),
			setNameValue("ZoneStart4", fmt.Sprintf("%f", unpSession.ZoneStart4)),
			setNameValue("ZoneStart5", fmt.Sprintf("%f", unpSession.ZoneStart5)),
			setNameValue("ZoneStart6", fmt.Sprintf("%f", unpSession.ZoneStart6)),
			setNameValue("ZoneStart7", fmt.Sprintf("%f", unpSession.ZoneStart7)),
			setNameValue("ZoneStart8", fmt.Sprintf("%f", unpSession.ZoneStart8)),
			setNameValue("ZoneStart9", fmt.Sprintf("%f", unpSession.ZoneStart9)),
			setNameValue("ZoneStart10", fmt.Sprintf("%f", unpSession.ZoneStart10)),
			setNameValue("ZoneStart11", fmt.Sprintf("%f", unpSession.ZoneStart11)),
			setNameValue("ZoneStart12", fmt.Sprintf("%f", unpSession.ZoneStart12)),
			setNameValue("ZoneStart13", fmt.Sprintf("%f", unpSession.ZoneStart13)),
			setNameValue("ZoneStart14", fmt.Sprintf("%f", unpSession.ZoneStart14)),
			setNameValue("ZoneStart15", fmt.Sprintf("%f", unpSession.ZoneStart15)),
			setNameValue("ZoneStart16", fmt.Sprintf("%f", unpSession.ZoneStart16)),
			setNameValue("ZoneStart17", fmt.Sprintf("%f", unpSession.ZoneStart17)),
			setNameValue("ZoneStart18", fmt.Sprintf("%f", unpSession.ZoneStart18)),
			setNameValue("ZoneStart19", fmt.Sprintf("%f", unpSession.ZoneStart19)),
			setNameValue("ZoneStart20", fmt.Sprintf("%f", unpSession.ZoneStart20)),
			setNameValue("ZoneStart21", fmt.Sprintf("%f", unpSession.ZoneStart21)),
			setNameValue("NetworkGame", fmt.Sprintf("%v", unpSession.NetworkGame)),
			setNameValue("SliProNativeSupport", fmt.Sprintf("%v", unpSession.SliProNativeSupport)),
			setNameValue("NumMarshalZones", fmt.Sprintf("%v", unpSession.NumMarshalZones)),
			setNameValue("PitSpeedLimit ", fmt.Sprintf("%v", unpSession.PitSpeedLimit)),
			setNameValue("Formula", setFormula(unpSession.Formula)),
			setNameValue("TrackID", setTrackName(unpSession.TrackID)),
			setNameValue("SessionType", setSessionType(unpSession.SessionType)),
			setNameValue("TrackLength", fmt.Sprintf("%v", unpSession.TrackLength)),
			setNameValue("TotalLaps", fmt.Sprintf("%v", unpSession.TotalLaps)),
		)

		output.AuxData, err = proto.Marshal(sd)

	case 2: //Lap Data

		// Bypass processing ...
		return false, nil

		/* 		// Unpack the 20 item lap data array
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
		   		output.Data = outputHeader + arraystring */

	case 3: //Event

		// Bypass processing ...
		return false, nil

		/* unpEvent := &F1Event{}
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
		*/
	case 4: //Participants

		// Bypass processing ...
		return false, nil

		/* 		unpParticipant := &F1Participant{}
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
		   		output.Data = outputHeader + fmt.Sprintf("|%v", unpParticipant.NumActiveCars) + arraystring */

	case 5: //Car Setups

		// Bypass processing ...
		return false, nil

		/* 		unpCarSetupData := &F1SetupData{}

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
		*/
	case 6: //Car Telemetery
		unpCarTelemetry := &F1CarTelemetryData{}
		unpCarTelemetryExtra := &F1CarTelemetryDataExtra{}

		//arraystring := ""

		// First task is to create the data for the "current driver"
		// we have two indexes ... one for the "drivers" car iDPDriver and one for the rest
		var iDPDriver int32 = 980
		var iDP int32 = 3000

		// this is the index used for a loop iteration
		var iDPlocal int32 = 0

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpCarTelemetry)
			if err != nil {
				ctx.Logger().Debugf("Unpack Fail: F1CarTelemetry ", err.Error())
				return false, err
			}
			ctx.Logger().Debugf("Car Array unpacked: %v\n%+v\n", i, unpCarTelemetry)

			if i == int(iCurrentPlayer) {
				iDPlocal = iDPDriver
			} else {
				iDPlocal = iDP
			}

			td.DataPoints = append(td.DataPoints,
				setDataPoint(iDPlocal, float64(unpCarTelemetry.Speed)),
				setDataPoint(iDPlocal+1, float64(unpCarTelemetry.Throttle)),
				setDataPoint(iDPlocal+2, float64(unpCarTelemetry.Steer)),
				setDataPoint(iDPlocal+3, float64(unpCarTelemetry.Brake)),
				setDataPoint(iDPlocal+4, float64(unpCarTelemetry.Clutch)),
				setDataPoint(iDPlocal+5, float64(unpCarTelemetry.Gear)),
				setDataPoint(iDPlocal+6, float64(unpCarTelemetry.EngineRPM)),
				setDataPoint(iDPlocal+7, float64(unpCarTelemetry.Drs)),
				setDataPoint(iDPlocal+8, float64(unpCarTelemetry.RevLightsPercent)),
				setDataPoint(iDPlocal+9, float64(unpCarTelemetry.BrakesTemperatureRL)),
				setDataPoint(iDPlocal+10, float64(unpCarTelemetry.BrakesTemperatureRR)),
				setDataPoint(iDPlocal+11, float64(unpCarTelemetry.BrakesTemperatureFL)),
				setDataPoint(iDPlocal+12, float64(unpCarTelemetry.BrakesTemperatureFR)),
				setDataPoint(iDPlocal+13, float64(unpCarTelemetry.TyresSurfaceTemperatureRL)),
				setDataPoint(iDPlocal+14, float64(unpCarTelemetry.TyresSurfaceTemperatureRR)),
				setDataPoint(iDPlocal+15, float64(unpCarTelemetry.TyresSurfaceTemperatureFL)),
				setDataPoint(iDPlocal+16, float64(unpCarTelemetry.TyresSurfaceTemperatureFR)),
				setDataPoint(iDPlocal+17, float64(unpCarTelemetry.TyresInnerTemperatureRL)),
				setDataPoint(iDPlocal+18, float64(unpCarTelemetry.TyresInnerTemperatureRR)),
				setDataPoint(iDPlocal+19, float64(unpCarTelemetry.TyresInnerTemperatureFL)),
				setDataPoint(iDPlocal+20, float64(unpCarTelemetry.TyresInnerTemperatureFR)),
				setDataPoint(iDPlocal+21, float64(unpCarTelemetry.EngineTemperature)),
				setDataPoint(iDPlocal+22, float64(unpCarTelemetry.TyresPressureRL)),
				setDataPoint(iDPlocal+23, float64(unpCarTelemetry.TyresPressureRR)),
				setDataPoint(iDPlocal+24, float64(unpCarTelemetry.TyresPressureFL)),
				setDataPoint(iDPlocal+25, float64(unpCarTelemetry.TyresPressureFR)),
				setDataPoint(iDPlocal+26, float64(unpCarTelemetry.SurfaceTypeRL)),
				setDataPoint(iDPlocal+27, float64(unpCarTelemetry.SurfaceTypeRR)),
				setDataPoint(iDPlocal+28, float64(unpCarTelemetry.SurfaceTypeFL)),
				setDataPoint(iDPlocal+29, float64(unpCarTelemetry.SurfaceTypeFR)),
			)
			iDP += 35

		}

		err = struc.Unpack(buf, unpCarTelemetryExtra)
		if err != nil {
			ctx.Logger().Debugf("Unpack Fail: F1CarTelemetryExtra ", err.Error())
			return false, err
		}

		// Format the datapoints
		td.DataPoints = append(td.DataPoints,
			setDataPoint(iDPDriver+30, float64(unpCarTelemetryExtra.ButtonStatus)),
		)

		// Send all fields
		output.Data, err = proto.Marshal(td)

	case 7: //Car Status
		// Bypass processing ...
		return false, nil
		/* 		unpCarStatus := &F1CarStatus{}
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
		   		output.Data = outputHeader + arraystring */

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

func setDataPoint(parm int32, value float64) *DataPoint {
	samp := &Sample{TimestampNano: nsMid, Value: value}
	dp := &DataPoint{ParameterId: parm}
	dp.Samples = append(dp.Samples, samp)
	return dp
}
func setNameValue(name string, value string) *NameValue {
	nv := &NameValue{Name: name, Value: value}
	return nv
}
func setFormula(id uint8) string {

	// Convert to String // Formular, 0 = F1 modern, 1 = F1 classic, 2 - F2, 3 = F1 Generic
	switch id {
	case 0:
		return "F1 Modern"
	case 1:
		return "F1 Classic"
	case 2:
		return "F1 Generic"
	}
	return "unknown"
}

func setSessionType(id uint8) string {
	//conver to string based on these conditions :   0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2, 12 = Time Trial
	switch id {
	case 1:
		return "P1"
	case 2:
		return "P2"
	case 3:
		return "P3"
	case 4:
		return "Short P"
	case 5:
		return "Q1"
	case 6:
		return "Q2"
	case 7:
		return "Q3"
	case 8:
		return "Short Q"
	case 9:
		return "OSQ"
	case 10:
		return "R"
	case 11:
		return "R2"
	case 12:
		return "Time Trial"

	}
	return "unknown"
}

func setTrackName(id int8) string {

	switch id {
	case 0:
		return "Melbourne"
	case 1:
		return "Paul Ricard"
	case 2:
		return "Shanghai"
	case 3:
		return "Sakhir (Bahrain)"
	case 4:
		return "Catalunya"
	case 5:
		return "Monaco"
	case 6:
		return "Montreal"
	case 7:
		return "Silverstone"
	case 8:
		return "Hockenheim"
	case 9:
		return "Hungaroring"
	case 10:
		return "Spa"
	case 11:
		return "Monza"
	case 12:
		return "Singapore"
	case 13:
		return "Suzuka"
	case 14:
		return "Abu Dhabi"
	case 15:
		return "Texas"
	case 16:
		return "Brazil"
	case 17:
		return "Austria"
	case 18:
		return "Sochi"
	case 19:
		return "Mexico"
	case 20:
		return "Baku (Azerbaijan)"
	case 21:
		return "Sakhir Short"
	case 22:
		return "Silverstone Short"
	case 23:
		return "Texas Short"
	case 24:
		return "Suzuka Short"
	}
	return "unknown"
}
