package f1telemetry

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/lunixbochs/struc"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

type F1SourceData struct {
	A                  float32 `struc:"float32,little"`
	B                  float32 `struc:"float32,little"`
	C                  float32 `struc:"float32,little"`
	D                  float32 `struc:"float32,little"`
	E                  float32 `struc:"float32,little"`
	F                  float32 `struc:"float32,little"`
	G                  float32 `struc:"float32,little"`
	H                  float32 `struc:"float32,little"`
	I                  float32 `struc:"float32,little"`
	J                  float32 `struc:"float32,little"`
	K                  float32 `struc:"float32,little"`
	L                  float32 `struc:"float32,little"`
	M                  float32 `struc:"float32,little"`
	N                  float32 `struc:"float32,little"`
	O                  float32 `struc:"float32,little"`
	P                  float32 `struc:"float32,little"`
	Q                  float32 `struc:"float32,little"`
	R                  float32 `struc:"float32,little"`
	S                  float32 `struc:"float32,little"`
	T                  float32 `struc:"float32,little"`
	U                  float32 `struc:"float32,little"`
	V                  float32 `struc:"float32,little"`
	W                  float32 `struc:"float32,little"`
	X                  float32 `struc:"float32,little"`
	Y                  float32 `struc:"float32,little"`
	Z                  float32 `struc:"float32,little"`
	AaC                float32 `struc:"float32,little"`
	AbD                float32 `struc:"float32,little"`
	AcD                float32 `struc:"float32,little"`
	AdD                float32 `struc:"float32,little"`
	AeD                float32 `struc:"float32,little"`
	AfD                float32 `struc:"float32,little"`
	AgA                float32 `struc:"float32,little"`
	AhB                float32 `struc:"float32,little"`
	AiC                float32 `struc:"float32,little"`
	AjD                float32 `struc:"float32,little"`
	AkD                float32 `struc:"float32,little"`
	BA                 float32 `struc:"float32,little"`
	BB                 float32 `struc:"float32,little"`
	BC                 float32 `struc:"float32,little"`
	BD                 float32 `struc:"float32,little"`
	BE                 float32 `struc:"float32,little"`
	BF                 float32 `struc:"float32,little"`
	BG                 float32 `struc:"float32,little"`
	BH                 float32 `struc:"float32,little"`
	BI                 float32 `struc:"float32,little"`
	BJ                 float32 `struc:"float32,little"`
	BK                 float32 `struc:"float32,little"`
	BL                 float32 `struc:"float32,little"`
	BM                 float32 `struc:"float32,little"`
	BN                 float32 `struc:"float32,little"`
	BO                 float32 `struc:"float32,little"`
	BP                 float32 `struc:"float32,little"`
	BQ                 float32 `struc:"float32,little"`
	BR                 float32 `struc:"float32,little"`
	BS                 float32 `struc:"float32,little"`
	BT                 float32 `struc:"float32,little"`
	BU                 float32 `struc:"float32,little"`
	BV                 float32 `struc:"float32,little"`
	BW                 float32 `struc:"float32,little"`
	BX                 float32 `struc:"float32,little"`
	BY                 float32 `struc:"float32,little"`
	BZ                 float32 `struc:"float32,little"`
	CA                 float32 `struc:"float32,little"`
	CB                 float32 `struc:"float32,little"`
	CC                 float32 `struc:"float32,little"`
	CD                 float32 `struc:"float32,little"`
	CE                 float32 `struc:"float32,little"`
	CF                 float32 `struc:"float32,little"`
	CG                 float32 `struc:"float32,little"`
	CH                 float32 `struc:"float32,little"`
	CI                 float32 `struc:"float32,little"`
	CJ                 float32 `struc:"float32,little"`
	CK                 float32 `struc:"float32,little"`
	CL                 float32 `struc:"float32,little"`
	CM                 float32 `struc:"float32,little"`
	AlD                uint8   `struc:"uint8,little"`
	AmD                uint8   `struc:"uint8,little"`
	AnD                uint8   `struc:"uint8,little"`
	AoA                uint8   `struc:"uint8,little"`
	ApB                uint8   `struc:"uint8,little"`
	AqC                uint8   `struc:"uint8,little"`
	ArD                uint8   `struc:"uint8,little"`
	AsD                uint8   `struc:"uint8,little"`
	AtD                uint8   `struc:"uint8,little"`
	AuD                uint8   `struc:"uint8,little"`
	AvD                uint8   `struc:"uint8,little"`
	AwA                uint8   `struc:"uint8,little"`
	AxB                uint8   `struc:"uint8,little"`
	AyC                uint8   `struc:"uint8,little"`
	AzD                uint8   `struc:"uint8,little"`
	BaD                uint8   `struc:"uint8,little"`
	BbD                uint8   `struc:"uint8,little"`
	BcD                uint8   `struc:"uint8,little"`
	BdD                uint8   `struc:"uint8,little"`
	BeA                uint8   `struc:"uint8,little"`
	BfB                uint8   `struc:"uint8,little"`
	BgC                uint8   `struc:"uint8,little"`
	BhD                uint8   `struc:"uint8,little"`
	BiD                uint8   `struc:"uint8,little"`
	BjD                float32 `struc:"float32,little"` // Doh !
	BkD                uint8   `struc:"uint8,little"`
	BlD                uint8   `struc:"uint8,little"`
	BmD                uint8   `struc:"uint8,little"`
	BmE                uint8   `struc:"uint8,little"`
	BmF                uint8   `struc:"uint8,little"`
	Filler1            []byte  `struc:"[900]byte"`      // cars data array
	Yaw                float32 `struc:"float32,little"` // NEW (v1.8)
	Pitch              float32 `struc:"float32,little"` // NEW (v1.8)
	Roll               float32 `struc:"float32,little"` // NEW (v1.8)
	XLocalVelocity     float32 `struc:"float32,little"` // NEW (v1.8) Velocity in local space
	YLocalVelocity     float32 `struc:"float32,little"` // NEW (v1.8) Velocity in local space
	ZLocalVelocity     float32 `struc:"float32,little"` // NEW (v1.8) Velocity in local space
	SuspAccelerationRL float32 `struc:"float32,little"` // NEW (v1.8) RL, RR, FL, FR
	SuspAccelerationRR float32 `struc:"float32,little"`
	SuspAccelerationFL float32 `struc:"float32,little"`
	SuspAccelerationFR float32 `struc:"float32,little"`
	AngAccX            float32 `struc:"float32,little"` // NEW (v1.8) angular acceleration x-component
	AngAccY            float32 `struc:"float32,little"` // NEW (v1.8) angular acceleration y-component
	AngAccZ            float32 `struc:"float32,little"` // NEW (v1.8) angular acceleration z-component
}

// F1CarArray2 - Struct for the unpacking of the UDP data format (Car data array)
type F1CarArray2 struct {
	X                 float32 `struc:"float32,little"` // world co-ordinates of vehicle
	Y                 float32 `struc:"float32,little"` // world co-ordinates of vehicle
	Z                 float32 `struc:"float32,little"` // world co-ordinates of vehicle
	LastLapTime       float32 `struc:"float32,little"`
	CurrentLapTime    float32 `struc:"float32,little"`
	BestLapTime       float32 `struc:"float32,little"`
	Sector1Time       float32 `struc:"float32,little"`
	Sector2Time       float32 `struc:"float32,little"`
	LapDistance       float32 `struc:"float32,little"`
	DriverID          uint8   `struc:"uint8,little"`
	TeamID            uint8   `struc:"uint8,little"`
	CarPosition       uint8   `struc:"uint8,little"` // UPDATED: track positions of vehicle
	CurrentLapNum     uint8   `struc:"uint8,little"`
	TyreCompound      uint8   `struc:"uint8,little"` // compound of tyre – 0 = ultra soft, 1 = super soft, 2 = soft, 3 = medium, 4 = hard, 5 = inter, 6 = wet
	InPits            uint8   `struc:"uint8,little"` // 0 = none, 1 = pitting, 2 = in pit area
	Sector            uint8   `struc:"uint8,little"` // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid uint8   `struc:"uint8,little"` // current lap invalid - 0 = valid, 1 = invalid
	Penalties         uint8   `struc:"uint8,little"` // NEW: accumulated time penalties in seconds to be added
}

func TestEvalQuality(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing ")
	//test1

	var buf bytes.Buffer
	var buf2 bytes.Buffer

	// Not working
	// arr := make([]F1CarArray2, 0)
	// f1d := F1CarArray2{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// arr = append(arr, f1d)
	// f1d = F1CarArray2{11.1, 22.1, 33.1, 44.1, 55.1, 66.1, 77.1, 88.1, 99.1, 110, 111, 112, 113, 114, 115, 116, 117, 118}
	// arr = append(arr, f1d)
	// f1d = F1CarArray2{12.1, 23.1, 34.1, 45.1, 56.1, 67.1, 78.1, 89.1, 99.9, 210, 211, 212, 213, 214, 215, 216, 217, 218}
	// arr = append(arr, f1d)

	arr, _ := hex.DecodeString("1cd33ac361126240685cdec2000000009191893f0000000000000000000000003697ac45220211010000020000247608c3e9755440fcd833c1000000009191893f0000000000000000000000001458af42100003010000000000fafd05c31dab4940948db0c1000000009191893f000000000000000000000000625d9e42000104010000000000cd200dc30c2c4940c36710c2000000009191893f00000000000000000000000022b57d42060106010000000000180323c3f95c4a40c73d9fc2000000009191893f00000000000000000000000091406b41070b0c010000000000d07a42c320c971408170fcc2000000009191893f0000000000000000000000003b10ac45120513010000020000514614c3b3754a405e0949c2000000009191893f000000000000000000000000e5383e42050608010000000000875e38c374715d408e9ef3c2000000009191893f0000000000000000000000004054ac45020212010000020000fab925c374b6554009b18ac2000000009191893f0000000000000000000000002f55b54123070b010000000000efec16c3e0ad55409f1520c2000000009191893f000000000000000000000000c0eb5d42030707010000000000ffe930c381a25140fa0dd7c2000000009191893f00000000000000000000000011d5ac450e0b1001000002000006ce29c352aa4b409857bac2000000009191893f0000000000000000000000002755ad4501080e0100000200005cc71bc3d43a4a40cefd82c2000000009191893f0000000000000000000000004eb7f4410a030a01000000000090ad33c3da705940c109c2c2000000009191893f0000000000000000000000003916ad4514030f010000020000c6a81ec3ce21564081ef5ac2000000009191893f00000000000000000000000002831b42210609010000000000bdb50fc3da7354404802cdc1000000009191893f00000000000000000000000028258f42160005010000000000ddc901c397385540a2a90940000000009191893f0000000000000000000000005b47cd420904010100000000009f7b2cc3dc0e564088dfa4c2000000009191893f0000000000000000000000004cb0fd4017080d0100000000005746fec23ad1494086cc05c1000000009191893f000000000000000000000000fb0bbd420f04020100000000005f403fc30d9c7040e4a907c3000000009191893f00000000000000000000000082d8ab451f0514010000020000")

	err := struc.Pack(&buf2, arr)
	fmt.Printf("error code: %v \n", err)

	fmt.Printf("Car data Array: \n%+x", buf2)
	//fmt.Printf("arr length = %v \n", buf2.Len())

	// working !!!
	src := &F1SourceData{1.11111, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
		buf2.Bytes(), 1.08, 10.8, 101.0, 0.111, 0.0120, 0.00113, 114, 115, 116, 117, 118, 119, 120}

	err = struc.Pack(&buf, src)
	fmt.Printf("error code: %v \n", err)
	fmt.Printf("FULL Car data: \n%+x", buf)
	//fmt.Printf("arr length = %v \n", buf.Len())
	//fmt.Printf("struct : \n %+v \n", src)

	// convert to byte array and print
	x := buf.Bytes()
	// for _, v := range x {
	// 	fmt.Printf("%v - %v \n", v, x[v])
	// }
	//fmt.Printf("buffer : \n %s \n", x)

	tc.SetInput("buffer", x)

	o := &F1SourceData{}
	err = struc.Unpack(&buf, o)
	//fmt.Printf("struct : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rdata := tc.GetOutput("cardata")
	rdata2 := tc.GetOutput("cararray")

	fmt.Printf("csv data: %v \n", rdata)
	fmt.Printf("csv array: %v \n", rdata2)

	if tc.GetOutput("cardata") == nil {
		t.Fail()
	}

}
