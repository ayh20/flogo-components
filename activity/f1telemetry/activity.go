package f1telemetry

import (
	"encoding/binary"
	"fmt"
	"reflect"
	//"strconv"
	//"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	//"github.com/lunixbochs/struc"
	"gopkg.in/restruct.v1"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-f1telemetry")

const (
	ivInput = "buffer"

	ovOutput = "data"
)

type F1Data struct {
	Time                 float32 `struc:"float32"` //F
	LapTime              float32 `struc:"float32"` //F
	LapDistance          float32 `struc:"float32"` //F
	TotalDistance        float32 `struc:"float32"` //F
	X                    float32 `struc:"float32"` // World space position F
	Y                    float32 `struc:"float32"` // World space position F
	Z                    float32 `struc:"float32"` // World space position F
	Speed                float32 `struc:"float32"` // Speed of car in MPH F
	Xv                   float32 `struc:"float32"` // Velocity in world space F
	Yv                   float32 `struc:"float32"` // Velocity in world space F
	Zv                   float32 `struc:"float32"` // Velocity in world space F
	Xr                   float32 `struc:"float32"` // World space right direction F
	Yr                   float32 `struc:"float32"` // World space right direction F
	Zr                   float32 `struc:"float32"` // World space right direction F
	Xd                   float32 `struc:"float32"` // World space forward direction F
	Yd                   float32 `struc:"float32"` // World space forward direction F
	Zd                   float32 `struc:"float32"` // World space forward direction F
	Suspposrl            float32 `struc:"float32"` // RL, RR, FL, FR  F array
	Suspposrr            float32 `struc:"float32"`
	Suspposfl            float32 `struc:"float32"`
	Suspposfr            float32 `struc:"float32"`
	Suspvelrl            float32 `struc:"float32"` // RL, RR, FL, FR  F array
	Suspvelrr            float32 `struc:"float32"`
	Suspvelfl            float32 `struc:"float32"`
	Suspvelfr            float32 `struc:"float32"`
	Wheelspeedrl         float32 `struc:"float32"`
	Wheelspeedrr         float32 `struc:"float32"`
	Wheelspeedfl         float32 `struc:"float32"`
	Wheelspeedfr         float32 `struc:"float32"`
	Throttle             float32 `struc:"float32"` // F
	Steer                float32 `struc:"float32"` // F
	Brake                float32 `struc:"float32"` // F
	Clutch               float32 `struc:"float32"` // F
	Gear                 float32 `struc:"float32"` // F
	Gforcelat            float32 `struc:"float32"` // F
	Gforcelon            float32 `struc:"float32"` // F
	Lap                  float32 `struc:"float32"` // F
	EngineRate           float32 `struc:"float32"` // F
	Slipronativesupport  float32 `struc:"float32"` // F	// SLI Pro support
	Carposition          float32 `struc:"float32"` // F	// car race position
	Kerslevel            float32 `struc:"float32"` // F	// kers energy left
	Kersmaxlevel         float32 `struc:"float32"` // F	// kers maximum energy
	Drs                  float32 `struc:"float32"` // F	// 0 = off, 1 = on
	Tractioncontrol      float32 `struc:"float32"` // F	// 0 (off) - 2 (high)
	Antilockbrakes       float32 `struc:"float32"` // F	// 0 (off) - 1 (on)
	Fuelintank           float32 `struc:"float32"` // F	// current fuel mass
	Fuelcapacity         float32 `struc:"float32"` // F	// fuel capacity
	Inpits               float32 `struc:"float32"` // F	// 0 = none, 1 = pitting, 2 = in pit area
	Sector               float32 `struc:"float32"` // F	// 0 = sector1, 1 = sector2, 2 = sector3
	Sector1time          float32 `struc:"float32"` // F	// time of sector1 (or 0)
	Sector2time          float32 `struc:"float32"` // F	// time of sector2 (or 0)
	Brakestemprl         float32 `struc:"float32"` // brakes temperature (centigrade)
	Brakestemprr         float32 `struc:"float32"`
	Brakestempfl         float32 `struc:"float32"`
	Brakestempfr         float32 `struc:"float32"`
	Tyrespressurerl      float32 `struc:"float32"` // tyres pressure PSI
	Tyrespressurerr      float32 `struc:"float32"`
	Tyrespressurefl      float32 `struc:"float32"`
	Tyrespressurefr      float32 `struc:"float32"`
	Teaminfo             float32 `struc:"float32"` // F	// team ID
	Totallaps            float32 `struc:"float32"` // F	// total number of laps in this race
	Tracksize            float32 `struc:"float32"` // F	// track size meters
	Lastlaptime          float32 `struc:"float32"` // F	// last lap time
	Maxrpm               float32 `struc:"float32"` // cars max RPM, at which point the rev limiter will kick in
	Idlerpm              float32 `struc:"float32"` // cars idle RPM
	Maxgears             float32 `struc:"float32"` // maximum number of gears
	SessionType          float32 `struc:"float32"` // 0 = unknown, 1 = practice, 2 = qualifying, 3 = race
	DrsAllowed           float32 `struc:"float32"` // 0 = not allowed, 1 = allowed, -1 = invalid / unknown
	Tracknumber          float32 `struc:"float32"` // -1 for unknown, 0-21 for tracks
	VehicleFIAFlags      float32 `struc:"float32"` // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	Era                  float32 `struc:"float32"` // era, 2017 (modern) or 1980 (classic)
	Enginetemperature    float32 `struc:"float32"` // engine temperature (centigrade)
	Gforcevert           float32 `struc:"float32"` // vertical g-force component
	Angvelx              float32 `struc:"float32"` // angular velocity x-component
	Angvely              float32 `struc:"float32"` // angular velocity y-component
	Angvelz              float32 `struc:"float32"` // angular velocity z-component
	TyrestemperatureRL   int     `struc:"int8"`    // tyres temperature (centigrade)
	TyrestemperatureRR   int     `struc:"int8"`
	TyrestemperatureFL   int     `struc:"int8"`
	TyrestemperatureFR   int     `struc:"int8"`
	TyreswearRL          int     `struc:"int8"` // tyre wear percentage
	TyreswearRR          int     `struc:"int8"`
	TyreswearFL          int     `struc:"int8"`
	TyreswearFR          int     `struc:"int8"`
	Tyrecompound         int     `struc:"int8"` // compound of tyre – 0 = ultra soft, 1 = super soft, 2 = soft, 3 = medium, 4 = hard, 5 = inter, 6 = wet
	Frontbrakebias       int     `struc:"int8"` // front brake bias (percentage)
	Fuelmix              int     `struc:"int8"` // fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	CurrentLapInvalid    int     `struc:"int8"` // current lap invalid - 0 = valid, 1 = invalid
	TyresdamageRL        int     `struc:"int8"` // tyre damage (percentage)
	TyresdamageRR        int     `struc:"int8"`
	TyresdamageFL        int     `struc:"int8"`
	TyresdamageFR        int     `struc:"int8"`
	Frontleftwingdamage  int     `struc:"int8"` // front left wing damage (percentage)
	Frontrightwingdamage int     `struc:"int8"` // front right wing damage (percentage)
	Rearwingdamage       int     `struc:"int8"` // rear wing damage (percentage)
	Enginedamage         int     `struc:"int8"` // engine damage (percentage)
	Gearboxdamage        int     `struc:"int8"` // gear box damage (percentage)
	Exhaustdamage        int     `struc:"int8"` // exhaust damage (percentage)
	Pitlimiterstatus     int     `struc:"int8"` // pit limiter status – 0 = off, 1 = on
	Pitspeedlimit        int     `struc:"int8"` // pit speed limit in mph
	Sessiontimeleft      int     `struc:"int8"` // NEW: time left in session in seconds
	Revlightspercent     int     `struc:"int8"` // NEW: rev lights indicator (percentage)
	Isspectating         int     `struc:"int8"` // NEW: whether the player is spectating
	Spectatorcarindex    int     `struc:"int8"` // NEW: index of the car being spectated
}

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// f1telemetry is an Activity that takes in data from a byte stream and interprets it as data from F1 2017
//
// inputs : {input1, input2, datatype, comparemode}
// outputs: result (bool)
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
	fmt.Println("Starting")

	//var buf bytes.Buffer
	input, _ := context.GetInput(ivInput).([]byte)
	//buf := bytes.NewBuffer(input)

	fmt.Printf("input : \n %s \n", input)

	//unpackedData := &F1Data{}

	var unpackedData F1Data

	fmt.Println("Unpack")
	restruct.Unpack(input, binary.LittleEndian, &unpackedData)

	fmt.Println("print")
	fmt.Printf("struct : \n %+v \n", unpackedData)

	//err = struc.Unpack(buf, unpackedData)

	fmt.Println("results")

	// Write the CSV rows.
	fields := unpackedData.ValueStrings()
	fmt.Printf("data : %v \n", fields)
	context.SetOutput(ovOutput, fields)

	return true, nil
}
func (f F1Data) ValueStrings() []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		ss[i] = fmt.Sprintf("%v", v.Field(i))
	}
	return ss
}
