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

	ovOutput     = "data"
	ovOutput2    = "array"
	ovOutputtype = "msgtype"
)

// F1Data - Struct for the unpacking of the UDP data format
type F1Data struct {
	Time                 float32 `struc:"float32,little"` //F
	LapTime              float32 `struc:"float32,little"` //F
	LapDistance          float32 `struc:"float32,little"` //F
	TotalDistance        float32 `struc:"float32,little"` //F
	X                    float32 `struc:"float32,little"` // World space position F
	Y                    float32 `struc:"float32,little"` // World space position F
	Z                    float32 `struc:"float32,little"` // World space position F
	Speed                float32 `struc:"float32,little"` // Speed of car in MPH F
	Xv                   float32 `struc:"float32,little"` // Velocity in world space F
	Yv                   float32 `struc:"float32,little"` // Velocity in world space F
	Zv                   float32 `struc:"float32,little"` // Velocity in world space F
	Xr                   float32 `struc:"float32,little"` // World space right direction F
	Yr                   float32 `struc:"float32,little"` // World space right direction F
	Zr                   float32 `struc:"float32,little"` // World space right direction F
	Xd                   float32 `struc:"float32,little"` // World space forward direction F
	Yd                   float32 `struc:"float32,little"` // World space forward direction F
	Zd                   float32 `struc:"float32,little"` // World space forward direction F
	SuspPosRL            float32 `struc:"float32,little"` // Suspension position RL, RR, FL, FR  F array
	SuspPosRR            float32 `struc:"float32,little"`
	SuspPosFL            float32 `struc:"float32,little"`
	SuspPosFR            float32 `struc:"float32,little"`
	SuspVelRL            float32 `struc:"float32,little"` // Suspension velocity RL, RR, FL, FR  F array
	SuspVelRR            float32 `struc:"float32,little"`
	SuspVelFL            float32 `struc:"float32,little"`
	SuspVelFR            float32 `struc:"float32,little"`
	WheelspeedRL         float32 `struc:"float32,little"` // Wheel Speed RL, RR, FL, FR  F array
	WheelspeedRR         float32 `struc:"float32,little"`
	WheelspeedFL         float32 `struc:"float32,little"`
	WheelspeedFR         float32 `struc:"float32,little"`
	Throttle             float32 `struc:"float32,little"` // F
	Steer                float32 `struc:"float32,little"` // F
	Brake                float32 `struc:"float32,little"` // F
	Clutch               float32 `struc:"float32,little"` // F
	Gear                 float32 `struc:"float32,little"` // F
	Gforcelat            float32 `struc:"float32,little"` // F
	Gforcelon            float32 `struc:"float32,little"` // F
	Lap                  float32 `struc:"float32,little"` // F
	EngineRate           float32 `struc:"float32,little"` // F
	Slipronativesupport  float32 `struc:"float32,little"` // F	// SLI Pro support
	Carposition          float32 `struc:"float32,little"` // F	// car race position
	Kerslevel            float32 `struc:"float32,little"` // F	// kers energy left
	Kersmaxlevel         float32 `struc:"float32,little"` // F	// kers maximum energy
	Drs                  float32 `struc:"float32,little"` // F	// 0 = off, 1 = on
	Tractioncontrol      float32 `struc:"float32,little"` // F	// 0 (off) - 2 (high)
	Antilockbrakes       float32 `struc:"float32,little"` // F	// 0 (off) - 1 (on)
	Fuelintank           float32 `struc:"float32,little"` // F	// current fuel mass
	Fuelcapacity         float32 `struc:"float32,little"` // F	// fuel capacity
	Inpits               float32 `struc:"float32,little"` // F	// 0 = none, 1 = pitting, 2 = in pit area
	Sector               float32 `struc:"float32,little"` // F	// 0 = sector1, 1 = sector2, 2 = sector3
	Sector1time          float32 `struc:"float32,little"` // F	// time of sector1 (or 0)
	Sector2time          float32 `struc:"float32,little"` // F	// time of sector2 (or 0)
	BrakestempRL         float32 `struc:"float32,little"` // brakes temperature (centigrade)
	BrakestempRR         float32 `struc:"float32,little"`
	BrakestempFL         float32 `struc:"float32,little"`
	BrakestempFR         float32 `struc:"float32,little"`
	TyrespressureRL      float32 `struc:"float32,little"` // tyres pressure PSI
	TyrespressureRR      float32 `struc:"float32,little"`
	TyrespressureFL      float32 `struc:"float32,little"`
	TyrespressureFR      float32 `struc:"float32,little"`
	Teaminfo             float32 `struc:"float32,little"` // F	// team ID
	Totallaps            float32 `struc:"float32,little"` // F	// total number of laps in this race
	Tracksize            float32 `struc:"float32,little"` // F	// track size meters
	Lastlaptime          float32 `struc:"float32,little"` // F	// last lap time
	Maxrpm               float32 `struc:"float32,little"` // cars max RPM, at which point the rev limiter will kick in
	Idlerpm              float32 `struc:"float32,little"` // cars idle RPM
	Maxgears             float32 `struc:"float32,little"` // maximum number of gears
	SessionType          float32 `struc:"float32,little"` // 0 = unknown, 1 = practice, 2 = qualifying, 3 = race
	DrsAllowed           float32 `struc:"float32,little"` // 0 = not allowed, 1 = allowed, -1 = invalid / unknown
	Tracknumber          float32 `struc:"float32,little"` // -1 for unknown, 0-21 for tracks
	VehicleFIAFlags      float32 `struc:"float32,little"` // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	Era                  float32 `struc:"float32,little"` // era, 2017 (modern) or 1980 (classic)
	Enginetemperature    float32 `struc:"float32,little"` // engine temperature (centigrade)
	Gforcevert           float32 `struc:"float32,little"` // vertical g-force component
	Angvelx              float32 `struc:"float32,little"` // angular velocity x-component
	Angvely              float32 `struc:"float32,little"` // angular velocity y-component
	Angvelz              float32 `struc:"float32,little"` // angular velocity z-component
	TyrestemperatureRL   byte    `struc:"byte,little"`    // tyres temperature (centigrade)
	TyrestemperatureRR   byte    `struc:"byte,little"`
	TyrestemperatureFL   byte    `struc:"byte,little"`
	TyrestemperatureFR   byte    `struc:"byte,little"`
	TyreswearRL          byte    `struc:"byte,little"` // tyre wear percentage
	TyreswearRR          byte    `struc:"byte,little"`
	TyreswearFL          byte    `struc:"byte,little"`
	TyreswearFR          byte    `struc:"byte,little"`
	Tyrecompound         byte    `struc:"byte,little"` // compound of tyre – 0 = ultra soft, 1 = super soft, 2 = soft, 3 = medium, 4 = hard, 5 = inter, 6 = wet
	Frontbrakebias       byte    `struc:"byte,little"` // front brake bias (percentage)
	Fuelmix              byte    `struc:"byte,little"` // fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	CurrentLapInvalid    byte    `struc:"byte,little"` // current lap invalid - 0 = valid, 1 = invalid
	TyresdamageRL        byte    `struc:"byte,little"` // tyre damage (percentage)
	TyresdamageRR        byte    `struc:"byte,little"`
	TyresdamageFL        byte    `struc:"byte,little"`
	TyresdamageFR        byte    `struc:"byte,little"`
	Frontleftwingdamage  byte    `struc:"byte,little"`    // front left wing damage (percentage)
	Frontrightwingdamage byte    `struc:"byte,little"`    // front right wing damage (percentage)
	Rearwingdamage       byte    `struc:"byte,little"`    // rear wing damage (percentage)
	Enginedamage         byte    `struc:"byte,little"`    // engine damage (percentage)
	Gearboxdamage        byte    `struc:"byte,little"`    // gear box damage (percentage)
	Exhaustdamage        byte    `struc:"byte,little"`    // exhaust damage (percentage)
	Pitlimiterstatus     byte    `struc:"byte,little"`    // pit limiter status – 0 = off, 1 = on
	Pitspeedlimit        byte    `struc:"byte,little"`    // pit speed limit in mph
	Sessiontimeleft      float32 `struc:"float32,little"` // NEW: time left in session in seconds
	Revlightspercent     byte    `struc:"byte,little"`    // NEW: rev lights indicator (percentage)
	Isspectating         byte    `struc:"byte,little"`    // NEW: whether the player is spectating
	Spectatorcarindex    byte    `struc:"byte,little"`    // NEW: index of the car being spectated
	NumCars              byte    `struc:"byte,little"`    // number of cars in data
	PlayerCarIndex       byte    `struc:"byte,little"`
	Filler1              []byte  `struc:"[900]byte"`      // cars data array
	Yaw                  float32 `struc:"float32,little"` // NEW (v1.8)
	Pitch                float32 `struc:"float32,little"` // NEW (v1.8)
	Roll                 float32 `struc:"float32,little"` // NEW (v1.8)
	XLocalVelocity       float32 `struc:"float32,little"` // NEW (v1.8) Velocity in local space
	YLocalVelocity       float32 `struc:"float32,little"` // NEW (v1.8) Velocity in local space
	ZLocalVelocity       float32 `struc:"float32,little"` // NEW (v1.8) Velocity in local space
	SuspAccelerationRL   float32 `struc:"float32,little"` // NEW (v1.8) RL, RR, FL, FR
	SuspAccelerationRR   float32 `struc:"float32,little"`
	SuspAccelerationFL   float32 `struc:"float32,little"`
	SuspAccelerationFR   float32 `struc:"float32,little"`
	AngAccX              float32 `struc:"float32,little"` // NEW (v1.8) angular acceleration x-component
	AngAccY              float32 `struc:"float32,little"` // NEW (v1.8) angular acceleration y-component
	AngAccZ              float32 `struc:"float32,little"` // NEW (v1.8) angular acceleration z-component
}

// F1CarArray - Struct for the unpacking of the UDP data format (Car data array)
type F1CarArray struct {
	X                 float32 `struc:"float32,little"` // world co-ordinates of vehicle
	Y                 float32 `struc:"float32,little"` // world co-ordinates of vehicle
	Z                 float32 `struc:"float32,little"` // world co-ordinates of vehicle
	LastLapTime       float32 `struc:"float32,little"`
	CurrentLapTime    float32 `struc:"float32,little"`
	BestLapTime       float32 `struc:"float32,little"`
	Sector1Time       float32 `struc:"float32,little"`
	Sector2Time       float32 `struc:"float32,little"`
	LapDistance       float32 `struc:"float32,little"`
	DriverID          byte    `struc:"byte"`
	TeamID            byte    `struc:"byte"`
	CarPosition       byte    `struc:"byte"` // UPDATED: track positions of vehicle
	CurrentLapNum     byte    `struc:"byte"`
	TyreCompound      byte    `struc:"byte"` // compound of tyre – 0 = ultra soft, 1 = super soft, 2 = soft, 3 = medium, 4 = hard, 5 = inter, 6 = wet
	InPits            byte    `struc:"byte"` // 0 = none, 1 = pitting, 2 = in pit area
	Sector            byte    `struc:"byte"` // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid byte    `struc:"byte"` // current lap invalid - 0 = valid, 1 = invalid
	Penalties         byte    `struc:"byte"` // NEW: accumulated time penalties in seconds to be added
}

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

	context.SetOutput(ovOutputtype, int(unpHeader.PacketID))

	switch unpHeader.PacketID {
	case 0: //Motion
		fmt.Println("zero")

		// Unpack the 20 item car data array
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
			if arraystring == "" {
				arraystring = fmt.Sprintf("%v", i) + "," + fmt.Sprintf("%v", unpHeader.SessionUID) + "," + fmt.Sprintf("%g", unpHeader.SessionTime) + "," + strings.Join(arrayfields, ",")
			} else {
				arraystring = arraystring + "|" + fmt.Sprintf("%v", i) + "," + fmt.Sprintf("%v", unpHeader.SessionUID) + "," + fmt.Sprintf("%g", unpHeader.SessionTime) + "," + strings.Join(arrayfields, ",")
			}

		}
		context.SetOutput(ovOutput2, arraystring)
		err = struc.Unpack(buf, unpMotionExtra)
		if err != nil {
			log.Error("Unpack Fail: F1CarMotionExtra ", err.Error())
			return false, err
		}
		//TODO: what are we going to do with the data ?
		fields := unpMotionExtra.valueStrings()
		fieldsstring := strings.Join(fields, ",")
		context.SetOutput(ovOutput, fieldsstring)

	case 1: //Session
		fmt.Println("one")
		err = struc.Unpack(buf, unpSession)
		if err != nil {
			log.Error("Unpack Fail: F1Session ", err.Error())
			return false, err
		}
		fields := unpSession.valueStrings()
		fieldsstring := strings.Join(fields, ",")
		context.SetOutput(ovOutput, fmt.Sprintf("%v", unpHeader.SessionUID)+","+fmt.Sprintf("%g", unpHeader.SessionTime)+","+fieldsstring)
		// TODO:  process data ?
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
			arrayfields := unpLapdata.valueStrings()
			if arraystring == "" {
				arraystring = fmt.Sprintf("%v", i) + "," + fmt.Sprintf("%v", unpHeader.SessionUID) + "," + fmt.Sprintf("%g", unpHeader.SessionTime) + "," + strings.Join(arrayfields, ",")
			} else {
				arraystring = arraystring + "|" + fmt.Sprintf("%v", i) + "," + fmt.Sprintf("%v", unpHeader.SessionUID) + "," + fmt.Sprintf("%g", unpHeader.SessionTime) + "," + strings.Join(arrayfields, ",")
			}

		}
		context.SetOutput(ovOutput2, arraystring)
	case 3: //Event
		fmt.Println("three")
		err = struc.Unpack(buf, unpEvent)
		if err != nil {
			log.Error("Unpack Fail: F1Event ", err.Error())
			return false, err
		}
		//fields := F1Event.valueStrings()
		//fieldsstring := strings.Join(fields, ",")
		//context.SetOutput(ovOutput, fieldsstring)

		context.SetOutput(ovOutput, fmt.Sprintf("%v", unpHeader.SessionUID)+","+fmt.Sprintf("%g", unpHeader.SessionTime)+","+unpEvent.EventString)

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
			arrayfields := unpParticpantData.valueStrings()
			if arraystring == "" {
				arraystring = fmt.Sprintf("%v", i) + "," + fmt.Sprintf("%v", unpHeader.SessionUID) + "," + fmt.Sprintf("%g", unpHeader.SessionTime) + "," + strings.Join(arrayfields, ",")
			} else {
				arraystring = arraystring + "|" + fmt.Sprintf("%v", i) + "," + fmt.Sprintf("%v", unpHeader.SessionUID) + "," + fmt.Sprintf("%g", unpHeader.SessionTime) + "," + strings.Join(arrayfields, ",")
			}

		}
		context.SetOutput(ovOutput2, arraystring)

	case 5: //Car Setups
		fmt.Println("five")
	case 6: //Car Telemetery
		fmt.Println("six")
	case 7: //Car Status
		fmt.Println("seven")
	default:
		fmt.Println("Error")
		return false, fmt.Errorf("F1 Data: Undefined packet ID %v", unpHeader.PacketID)
	}

	return true, nil
}
func (f F1Data) valueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		typeField := v.Type().Field(i)
		if strings.HasPrefix(typeField.Name, "Filler") {
			ss[i] = fmt.Sprintf("%v", "-1")
		} else {
			switch v.Field(i).Kind() {
			case reflect.Float32, reflect.Float64:
				x := v.Field(i).Float()
				ss[i] = strconv.FormatFloat(x, 'f', -1, 32)
			default:
				ss[i] = fmt.Sprintf("%v", v.Field(i))
			}

		}
	}
	return ss
}
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
			ss[i] = fmt.Sprintf("%v", v.Field(i))
		}

	}
	return ss
}
