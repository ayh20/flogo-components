package f1telemetry2018

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/lunixbochs/struc"
	//"gopkg.in/restruct.v1"
)

// activityLog is the default logger for the Log Activity
var log = logger.GetLogger("activity-f1telemetry2018")

const (
	ivInput = "buffer"

	ovOutputData = "data"
	ovOutputType = "msgtype"
)

// F1Header - Struct for the unpacking of the UDP Header
type F1Header struct {
	PacketFormat    uint16  `struc:"uint16,little"`  // 2018
	PacketVersion   uint8   `struc:"uint8,little"`   // Version of this packet type, all start from 1
	PacketID        uint8   `struc:"uint8,little"`   // Identifier for the packet type, see below
	SessionUID      uint64  `struc:"uint64,little"`  // Unique identifier for the session
	SessionTime     float32 `struc:"float32,little"` // Session timestamp
	FrameIdentifier uint    `struc:"uint,little"`    // Identifier for the frame the data was retrieved on
	PlayerCarIndex  uint8   `struc:"uint8,little"`   // Index of player's car in the array
}

// F1CarMotion - Struct for the unpacking of the UDP Motion format
type F1CarMotion struct {
	X          float32 `struc:"float32,little"` // World space position X
	Y          float32 `struc:"float32,little"` // World space position Y
	Z          float32 `struc:"float32,little"` // World space position Z
	Xv         float32 `struc:"float32,little"` // Velocity in world space X
	Yv         float32 `struc:"float32,little"` // Velocity in world space Y
	Zv         float32 `struc:"float32,little"` // Velocity in world space Z
	Xd         int16   `struc:"int16,little"`   // World space forward direction X
	Yd         int16   `struc:"int16,little"`   // World space forward direction Y
	Zd         int16   `struc:"int16,little"`   // World space forward direction Z
	Xr         int16   `struc:"int16,little"`   // World space right direction X
	Yr         int16   `struc:"int16,little"`   // World space right direction Y
	Zr         int16   `struc:"int16,little"`   // World space right direction Z
	Gforcelat  float32 `struc:"float32,little"` // Lateral G-Force component
	Gforcelon  float32 `struc:"float32,little"` // Longitudinal G-Force component
	Gforcevert float32 `struc:"float32,little"` // Vertical G-Force component
	Yaw        float32 `struc:"float32,little"` // Yaw angle in radians
	Pitch      float32 `struc:"float32,little"` // Pitch angle in radians
	Roll       float32 `struc:"float32,little"` // Roll angle in radians
}

// F1CarMotionExtra - Struct for the unpacking of the UDP data format
type F1CarMotionExtra struct {
	SuspPosRL          float32 `struc:"float32,little"` // Suspension position RL, RR, FL, FR  F array
	SuspPosRR          float32 `struc:"float32,little"`
	SuspPosFL          float32 `struc:"float32,little"`
	SuspPosFR          float32 `struc:"float32,little"`
	SuspVelRL          float32 `struc:"float32,little"` // Suspension velocity RL, RR, FL, FR  F array
	SuspVelRR          float32 `struc:"float32,little"`
	SuspVelFL          float32 `struc:"float32,little"`
	SuspVelFR          float32 `struc:"float32,little"`
	SuspAccelerationRL float32 `struc:"float32,little"` // RL, RR, FL, FR
	SuspAccelerationRR float32 `struc:"float32,little"`
	SuspAccelerationFL float32 `struc:"float32,little"`
	SuspAccelerationFR float32 `struc:"float32,little"`
	WheelspeedRL       float32 `struc:"float32,little"` // Wheel Speed RL, RR, FL, FR  F array
	WheelspeedRR       float32 `struc:"float32,little"`
	WheelspeedFL       float32 `struc:"float32,little"`
	WheelspeedFR       float32 `struc:"float32,little"`
	WheelslipRL        float32 `struc:"float32,little"` // Wheel Speed RL, RR, FL, FR  F array
	WheelslipRR        float32 `struc:"float32,little"`
	WheelslipFL        float32 `struc:"float32,little"`
	WheelslipFR        float32 `struc:"float32,little"`
	XLocalVelocity     float32 `struc:"float32,little"` // Velocity in local space
	YLocalVelocity     float32 `struc:"float32,little"` // Velocity in local space
	ZLocalVelocity     float32 `struc:"float32,little"` // Velocity in local space
	Angvelx            float32 `struc:"float32,little"` // angular velocity x-component
	Angvely            float32 `struc:"float32,little"` // angular velocity y-component
	Angvelz            float32 `struc:"float32,little"` // angular velocity z-component
	AngAccX            float32 `struc:"float32,little"` // Angular acceleration x-component
	AngAccY            float32 `struc:"float32,little"` // Angular acceleration y-component
	AngAccZ            float32 `struc:"float32,little"` // Angular acceleration z-component
	FrontWheelsAngle   float32 `struc:"float32,little"` // Current front wheels angle in radians
}

// F1Session - Struct for the unpacking of the UDP data format
type F1Session struct {
	Weather             uint8   `struc:"uint8,little"`   // Weather - 0 = clear, 1 = light cloud, 2 = overcast  3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature    int8    `struc:"int8,little"`    // Track temp. in degrees celsius
	AirTemperature      int8    `struc:"int8,little"`    // Air temp. in degrees celsius
	TotalLaps           uint8   `struc:"uint8,little"`   // Total number of laps in this race
	TrackLength         uint16  `struc:"uint16,little"`  // Track length in metres
	SessionType         uint8   `struc:"uint8,little"`   // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ,  10 = R, 11 = R2, 12 = Time Trial
	TrackID             int8    `struc:"int8,little"`    // -1 for unknown, 0-21 for tracks, see appendix
	Era                 uint8   `struc:"uint8,little"`   // Era, 0 = modern, 1 = classic
	SessionTimeLeft     uint16  `struc:"uint16,little"`  // Time left in session in seconds
	SessionDuration     uint16  `struc:"uint16,little"`  // Session duration in seconds
	PitSpeedLimit       uint8   `struc:"uint8,little"`   // Pit speed limit in kilometres per hour
	GamePaused          uint8   `struc:"uint8,little"`   // Whether the game is paused
	IsSpectating        uint8   `struc:"uint8,little"`   // Whether the player is spectating
	SpectatorCarIndex   uint8   `struc:"uint8,little"`   // Index of the car being spectated
	SliProNativeSupport uint8   `struc:"uint8,little"`   // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones     uint8   `struc:"uint8,little"`   // Number of marshal zones to follow
	ZoneStart1          float32 `struc:"float32,little"` // Fraction (0..1) of way through the lap the marshal Zone starts
	ZoneFlag1           int8    `struc:"int8,little"`    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ZoneStart2          float32 `struc:"float32,little"` // Note defining as separate fields for the moment to make the code simpler
	ZoneFlag2           int8    `struc:"int8,little"`    // review soon....
	ZoneStart3          float32 `struc:"float32,little"`
	ZoneFlag3           int8    `struc:"int8,little"`
	ZoneStart4          float32 `struc:"float32,little"`
	ZoneFlag4           int8    `struc:"int8,little"`
	ZoneStart5          float32 `struc:"float32,little"`
	ZoneFlag5           int8    `struc:"int8,little"`
	ZoneStart6          float32 `struc:"float32,little"`
	ZoneFlag6           int8    `struc:"int8,little"`
	ZoneStart7          float32 `struc:"float32,little"`
	ZoneFlag7           int8    `struc:"int8,little"`
	ZoneStart8          float32 `struc:"float32,little"`
	ZoneFlag8           int8    `struc:"int8,little"`
	ZoneStart9          float32 `struc:"float32,little"`
	ZoneFlag9           int8    `struc:"int8,little"`
	ZoneStart10         float32 `struc:"float32,little"`
	ZoneFlag10          int8    `struc:"int8,little"`
	ZoneStart11         float32 `struc:"float32,little"`
	ZoneFlag11          int8    `struc:"int8,little"`
	ZoneStart12         float32 `struc:"float32,little"`
	ZoneFlag12          int8    `struc:"int8,little"`
	ZoneStart13         float32 `struc:"float32,little"`
	ZoneFlag13          int8    `struc:"int8,little"`
	ZoneStart14         float32 `struc:"float32,little"`
	ZoneFlag14          int8    `struc:"int8,little"`
	ZoneStart15         float32 `struc:"float32,little"`
	ZoneFlag15          int8    `struc:"int8,little"`
	ZoneStart16         float32 `struc:"float32,little"`
	ZoneFlag16          int8    `struc:"int8,little"`
	ZoneStart17         float32 `struc:"float32,little"`
	ZoneFlag17          int8    `struc:"int8,little"`
	ZoneStart18         float32 `struc:"float32,little"`
	ZoneFlag18          int8    `struc:"int8,little"`
	ZoneStart19         float32 `struc:"float32,little"`
	ZoneFlag19          int8    `struc:"int8,little"`
	ZoneStart20         float32 `struc:"float32,little"`
	ZoneFlag20          int8    `struc:"int8,little"`
	ZoneStart21         float32 `struc:"float32,little"`
	ZoneFlag21          int8    `struc:"int8,little"`
	SafetyCarStatus     uint8   `struc:"uint8,little"` // 0 = no safety car, 1 = full safety car 2 = virtual safety car
	NetworkGame         uint8   `struc:"uint8,little"` // 0 = offline, 1 = online
}

// F1LapData - Struct for the unpacking of the UDP data format
type F1LapData struct {
	LastLapTime       float32 `struc:"float32,little"` // Last lap time in seconds
	CurrentLapTime    float32 `struc:"float32,little"` // Current time around the lap in seconds
	BestLapTime       float32 `struc:"float32,little"` // Best lap time of the session in seconds
	Sector1Time       float32 `struc:"float32,little"` // Sector 1 time in seconds
	Sector2Time       float32 `struc:"float32,little"` // Sector 2 time in seconds
	LapDistance       float32 `struc:"float32,little"` // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance     float32 `struc:"float32,little"` // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	SafetyCarDelta    float32 `struc:"float32,little"` // Delta in seconds for safety car
	CarPosition       uint8   `struc:"uint8,little"`   // Car race position
	CurrentLapNum     uint8   `struc:"uint8,little"`   // Current lap number
	PitStatus         uint8   `struc:"uint8,little"`   // 0 = none, 1 = pitting, 2 = in pit area
	Sector            uint8   `struc:"uint8,little"`   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid uint8   `struc:"uint8,little"`   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties         uint8   `struc:"uint8,little"`   // Accumulated time penalties in seconds to be added
	GridPosition      uint8   `struc:"uint8,little"`   // Grid position the vehicle started the race in
	DriverStatus      uint8   `struc:"uint8,little"`   // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus      uint8   `struc:"uint8,little"`   // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = disqualified, 5 = not classified, 6 = retired
}

// F1Event - Struct for the unpacking of the UDP data format
type F1Event struct {
	EventString string `struc:"[4]byte,little"` // Event string code
}

// F1Participant - Struct for the unpacking of the UDP data format
type F1Participant struct {
	NumCars uint8 `struc:"uint8,little"` // Number of cars in the data
}

// F1ParticipantData - Struct for the unpacking of the UDP data format
type F1ParticipantData struct {
	AiControlled uint8  `struc:"uint8,little"`    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID     uint8  `struc:"uint8,little"`    // Driver id
	TeamID       uint8  `struc:"uint8,little"`    // Team id
	RaceNumber   uint8  `struc:"uint8,little"`    // Race number of the car
	Nationality  uint8  `struc:"uint8,little"`    // Nationality of the drive
	Name         string `struc:"[48]byte,little"` // Name of participant in UTF-8 format – null terminated.  Will be truncated with … (U+2026) if too long
}

// F1SetupData - Struct for the unpacking of the UDP data format
type F1SetupData struct {
	FrontWing             uint8   `struc:"uint8,little"`   // Front wing aero
	RearWing              uint8   `struc:"uint8,little"`   // Rear wing aero
	OnThrottle            uint8   `struc:"uint8,little"`   // Differential adjustment on throttle (percentage)
	OffThrottle           uint8   `struc:"uint8,little"`   // Differential adjustment off throttle (percentage)
	FrontCamber           float32 `struc:"float32,little"` // Front camber angle (suspension geometry)
	RearCamber            float32 `struc:"float32,little"` // Rear camber angle (suspension geometry)
	FrontToe              float32 `struc:"float32,little"` // Front toe angle (suspension geometry)
	RearToe               float32 `struc:"float32,little"` // Rear toe angle (suspension geometry)
	FrontSuspension       uint8   `struc:"uint8,little"`   // Front suspension
	RearSuspension        uint8   `struc:"uint8,little"`   // Rear suspension
	FrontAntiRollBar      uint8   `struc:"uint8,little"`   // Front anti-roll bar
	RearAntiRollBar       uint8   `struc:"uint8,little"`   // Front anti-roll bar
	FrontSuspensionHeight uint8   `struc:"uint8,little"`   // Front ride height
	RearSuspensionHeight  uint8   `struc:"uint8,little"`   // Rear ride height
	BrakePressure         uint8   `struc:"uint8,little"`   // Brake pressure (percentage)
	BrakeBias             uint8   `struc:"uint8,little"`   // Brake bias (percentage)
	FrontTyrePressure     float32 `struc:"float32,little"` // Front tyre pressure (PSI)
	RearTyrePressure      float32 `struc:"float32,little"` // Rear tyre pressure (PSI)
	Ballast               uint8   `struc:"uint8,little"`   // Ballast
	FuelLoad              float32 `struc:"float32,little"` // Fuel load
}

// F1CarTelemetryData - Struct for the unpacking of the UDP data format
type F1CarTelemetryData struct {
	Speed                     uint16  `struc:"uint16,little"`  // Speed of car in kilometres per hour
	Throttle                  uint8   `struc:"uint8,little"`   // Amount of throttle applied (0 to 100)
	Steer                     int8    `struc:"int8,little"`    // Steering (-100 (full lock left) to 100 (full lock right))
	Brake                     uint8   `struc:"uint8,little"`   // Amount of brake applied (0 to 100)
	Clutch                    uint8   `struc:"uint8,little"`   // Amount of clutch applied (0 to 100)
	Gear                      uint8   `struc:"uint8,little"`   // Gear selected (1-8, N=0, R=-1)
	EngineRPM                 uint16  `struc:"uint16,little"`  // Engine RPM
	Drs                       uint8   `struc:"uint8,little"`   // 0 = off, 1 = on
	RevLightsPercent          uint8   `struc:"uint8,little"`   // Rev lights indicator (percentage)
	BrakesTemperatureRL       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureRR       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureFL       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	BrakesTemperatureFR       uint16  `struc:"uint16,little"`  // Brakes temperature (celsius)
	TyresSurfaceTemperatureRL uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureRR uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureFL uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresSurfaceTemperatureFR uint16  `struc:"uint16,little"`  // Tyres surface temperature (celsius)
	TyresInnerTemperatureRL   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
	TyresInnerTemperatureRR   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
	TyresInnerTemperatureFL   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
	TyresInnerTemperatureFR   uint16  `struc:"uint16,little"`  // Tyres inner temperature (celsius)
	EngineTemperature         uint16  `struc:"uint16,little"`  // Engine temperature (celsius)
	TyresPressureRL           float32 `struc:"float32,little"` // Tyres pressure (PSI)
	TyresPressureRR           float32 `struc:"float32,little"` // Tyres pressure (PSI)
	TyresPressureFL           float32 `struc:"float32,little"` // Tyres pressure (PSI)
	TyresPressureFR           float32 `struc:"float32,little"` // Tyres pressure (PSI)
}

// F1CarTelemetryDataExtra - Struct for the unpacking of the UDP data format
type F1CarTelemetryDataExtra struct {
	ButtonStatus uint32 `struc:"uint32,little"` // Bit flags specifying which buttons are being pressed currently
}

// F1CarStatus - Struct for the unpacking of the UDP data format
type F1CarStatus struct {
	TractionControl         uint8   `struc:"uint8,little"`   // 0 (off) - 2 (high)
	AntiLockBrakes          uint8   `struc:"uint8,little"`   // 0 (off) - 1 (on)
	FuelMix                 uint8   `struc:"uint8,little"`   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias          uint8   `struc:"uint8,little"`   // Front brake bias (percentage)
	PitLimiterStatus        uint8   `struc:"uint8,little"`   // Pit limiter status - 0 = off, 1 = on
	FuelInTank              float32 `struc:"float32,little"` // Current fuel mass
	FuelCapacity            float32 `struc:"float32,little"` // Fuel capacity
	MaxRPM                  uint16  `struc:"uint16,little"`  // Cars max RPM, point of rev limiter
	IdleRPM                 uint16  `struc:"uint16,little"`  // Cars idle RPM
	MaxGears                uint8   `struc:"uint8,little"`   // Maximum number of gears
	DrsAllowed              uint8   `struc:"uint8,little"`   // 0 = not allowed, 1 = allowed, -1 = unknown
	TyresWearRL             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	TyresWearRR             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	TyresWearFL             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	TyresWearFR             uint8   `struc:"uint8,little"`   // Tyre wear percentage
	TyreCompound            uint8   `struc:"uint8,little"`   // Modern - 0 = hyper soft, 1 = ultra soft, 2 = super soft, 3 = soft, 4 = medium, 5 = hard,  6 = super hard, 7 = inter, 8 = wet // Classic - 0-6 = dry, 7-8 = wet
	TyresDamageRL           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageRR           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageFL           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	TyresDamageFR           uint8   `struc:"uint8,little"`   // Tyre damage (percentage)
	FrontLeftWingDamage     uint8   `struc:"uint8,little"`   // Front left wing damage (percentage)
	FrontRightWingDamage    uint8   `struc:"uint8,little"`   // Front right wing damage (percentage)
	RearWingDamage          uint8   `struc:"uint8,little"`   // Rear wing damage (percentage)
	EngineDamage            uint8   `struc:"uint8,little"`   // Engine damage (percentage)
	GearBoxDamage           uint8   `struc:"uint8,little"`   // Gear box damage (percentage)
	ExhaustDamage           uint8   `struc:"uint8,little"`   // Exhaust damage (percentage)
	VehicleFiaFlags         uint8   `struc:"uint8,little"`   // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ErsStoreEnergy          float32 `struc:"float32,little"` // ERS energy store in Joules
	ErsDeployMode           uint8   `struc:"uint8,little"`   // ERS deployment mode, 0 = none, 1 = low, 2 = medium, 3 = high, 4 = overtake, 5 = hotlap
	ErsHarvestedThisLapMGUK float32 `struc:"float32,little"` // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32 `struc:"float32,little"` // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap      float32 `struc:"float32,little"` // ERS energy deployed this lap
}

func init() {
	log.SetLogLevel(logger.InfoLevel)
	//log.SetLogLevel(logger.DebugLevel)
}

// f1telemetry is an Activity that takes in data from a byte stream and interprets it as data from F1 2017
//
// inputs : {buffer} (byte array) RAW UDP data
// outputs: {data} (string) CSV data
type f1telemetry struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &f1telemetry{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *f1telemetry) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *f1telemetry) Eval(context activity.Context) (done bool, err error) {

	// Get the runtime values
	log.Debug("Starting")

	input, _ := context.GetInput(ivInput).([]byte)
	buf := bytes.NewBuffer(input)

	log.Debugf("input : \n %x \n", input)

	// Create structs to hold unpacked data
	unpHeader := &F1Header{}
	unpMotion := &F1CarMotion{}
	unpMotionExtra := &F1CarMotionExtra{}
	unpSession := &F1Session{}
	unpLapdata := &F1LapData{}
	unpEvent := &F1Event{}
	unpParticipant := &F1Participant{}
	unpParticpantData := &F1ParticipantData{}
	unpCarSetupData := &F1SetupData{}
	unpCarTelemetry := &F1CarTelemetryData{}
	unpCarTelemetryExtra := &F1CarTelemetryDataExtra{}
	unpCarStatus := &F1CarStatus{}

	log.Debug("Unpack Header")

	// Unpack the Header
	err = struc.Unpack(buf, unpHeader)
	if err != nil {
		log.Error("Unpack Fail: F1Header ", err.Error())
		return false, err
	}

	// dump header
	log.Debug("print")
	log.Debugf("struct F1Header : \n %+v \n", unpHeader)

	// Test for valid 2018 data..
	if unpHeader.PacketFormat != 2018 {
		log.Error("F1 Data: Unsupported packet format ", unpHeader.PacketFormat)
		return false, fmt.Errorf("F1 Data: Unsupported packet format %v", unpHeader.PacketFormat)
	}

	context.SetOutput(ovOutputData, "")
	context.SetOutput(ovOutputType, int(unpHeader.PacketID))

	outputHeader := fmt.Sprintf("%v,%v,%g,%v", unpHeader.PacketID, unpHeader.SessionUID, unpHeader.SessionTime, unpHeader.PlayerCarIndex)

	switch unpHeader.PacketID {
	case 0: //Motion
		fmt.Println("zero")

		// Unpack the 20 item car motion array
		// Note - Output array is:
		// Timestamp + array of car CSV data seprated by a "|"

		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpMotion)
			if err != nil {
				log.Error("Unpack Fail: F1CarMotion ", err.Error())
				return false, err
			}
			log.Debugf("Car Array unpacked: %v\n%+v\n", i, unpMotion)
			arrayfields := unpMotion.valueStrings()
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + strings.Join(arrayfields, ",")

		}

		err = struc.Unpack(buf, unpMotionExtra)
		if err != nil {
			log.Error("Unpack Fail: F1CarMotionExtra ", err.Error())
			return false, err
		}
		// Send all fields
		fieldsstring := "|" + strings.Join(unpMotionExtra.valueStrings(), ",")
		context.SetOutput(ovOutputData, outputHeader+fieldsstring+arraystring)

	case 1: //Session
		fmt.Println("one")
		err = struc.Unpack(buf, unpSession)
		if err != nil {
			log.Error("Unpack Fail: F1Session ", err.Error())
			return false, err
		}
		fields := unpSession.valueStrings()
		fieldsstring := "|" + strings.Join(fields, ",")
		// Send all fields
		context.SetOutput(ovOutputData, outputHeader+fieldsstring)

	case 2: //Lap Data
		fmt.Println("two")
		// Unpack the 20 item lap data array
		// Note - Output array is:
		// Timestamp + array of car CSV data seprated by a "|"

		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpLapdata)
			if err != nil {
				log.Error("Unpack Fail: F1LapData ", err.Error())
				return false, err
			}
			log.Debugf("LapData unpacked: %v\n%+v\n", i, unpLapdata)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + strings.Join(unpLapdata.valueStrings(), ",")

		}
		context.SetOutput(ovOutputData, outputHeader+arraystring)
	case 3: //Event
		fmt.Println("three")
		err = struc.Unpack(buf, unpEvent)
		if err != nil {
			log.Error("Unpack Fail: F1Event ", err.Error())
			return false, err
		}

		context.SetOutput(ovOutputData, outputHeader+"|"+unpEvent.EventString)

	case 4: //Participants
		fmt.Println("four")
		err = struc.Unpack(buf, unpParticipant)
		if err != nil {
			log.Error("Unpack Fail: F1Participant ", err.Error())
			return false, err
		}
		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpParticpantData)
			if err != nil {
				log.Error("Unpack Fail: F1ParticipantData ", err.Error())
				return false, err
			}
			log.Debugf("F1Participant unpacked: %v\n%+v\n", i, unpParticpantData)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + strings.Join(unpParticpantData.valueStrings(), ",")

		}
		context.SetOutput(ovOutputData, outputHeader+arraystring)

	case 5: //Car Setups
		fmt.Println("five")
		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpCarSetupData)
			if err != nil {
				log.Error("Unpack Fail: F1CarSetupData ", err.Error())
				return false, err
			}
			log.Debugf("F1CarSetupData unpacked: %v\n%+v\n", i, unpCarSetupData)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + strings.Join(unpCarSetupData.valueStrings(), ",")
		}
		context.SetOutput(ovOutputData, outputHeader+arraystring)

	case 6: //Car Telemetery
		fmt.Println("six")
		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpCarTelemetry)
			if err != nil {
				log.Error("Unpack Fail: F1CarTelemetry ", err.Error())
				return false, err
			}
			log.Debugf("Car Array unpacked: %v\n%+v\n", i, unpCarTelemetry)
			arrayfields := unpCarTelemetry.valueStrings()
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + strings.Join(arrayfields, ",")

		}

		err = struc.Unpack(buf, unpCarTelemetryExtra)
		if err != nil {
			log.Error("Unpack Fail: F1CarTelemetryExtra ", err.Error())
			return false, err
		}
		// Send all fields
		fieldsstring := "|" + strings.Join(unpCarTelemetryExtra.valueStrings(), ",")
		context.SetOutput(ovOutputData, outputHeader+fieldsstring+arraystring)

	case 7: //Car Status
		fmt.Println("seven")

		arraystring := ""

		for i := 0; i <= 19; i++ {
			err = struc.Unpack(buf, unpCarStatus)
			if err != nil {
				log.Error("Unpack Fail: F1CarStatus ", err.Error())
				return false, err
			}
			log.Debugf("CarStatus unpacked: %v\n%+v\n", i, unpCarStatus)
			arraystring = arraystring + fmt.Sprintf("|%v,", i) + strings.Join(unpCarStatus.valueStrings(), ",")

		}
		context.SetOutput(ovOutputData, outputHeader+arraystring)

	default:
		fmt.Println("Error")
		return false, fmt.Errorf("F1 Data: Undefined packet ID %v", unpHeader.PacketID)
	}

	return true, nil
}

// func (f F1Data) valueStrings() []string {
// 	v := reflect.ValueOf(f)
// 	ss := make([]string, v.NumField())
// 	for i := range ss {
// 		typeField := v.Type().Field(i)
// 		if strings.HasPrefix(typeField.Name, "Filler") {
// 			ss[i] = fmt.Sprintf("%v", "-1")
// 		} else {
// 			switch v.Field(i).Kind() {
// 			case reflect.Float32, reflect.Float64:
// 				x := v.Field(i).Float()
// 				ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
// 			default:
// 				ss[i] = fmt.Sprintf("%v", v.Field(i))
// 			}

// 		}
// 	}
// 	return ss
// }
func (f F1CarMotion) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1CarMotionExtra) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1LapData) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1Session) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1Event) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1ParticipantData) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = strings.Trim(fmt.Sprintf("%v", v.Field(i)), "\x00")
		}

	}
	return ss
}
func (f F1SetupData) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1CarTelemetryData) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1CarTelemetryDataExtra) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
func (f F1CarStatus) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		//typeField := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.Float32, reflect.Float64:
			x := v.Field(i).Float()
			ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
		default:
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
