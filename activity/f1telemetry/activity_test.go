package f1telemetry

import (
	"bytes"
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
	A   float32 `struc:"float32,little"`
	B   float32 `struc:"float32,little"`
	C   float32 `struc:"float32,little"`
	D   float32 `struc:"float32,little"`
	E   float32 `struc:"float32,little"`
	F   float32 `struc:"float32,little"`
	G   float32 `struc:"float32,little"`
	H   float32 `struc:"float32,little"`
	I   float32 `struc:"float32,little"`
	J   float32 `struc:"float32,little"`
	K   float32 `struc:"float32,little"`
	L   float32 `struc:"float32,little"`
	M   float32 `struc:"float32,little"`
	N   float32 `struc:"float32,little"`
	O   float32 `struc:"float32,little"`
	P   float32 `struc:"float32,little"`
	Q   float32 `struc:"float32,little"`
	R   float32 `struc:"float32,little"`
	S   float32 `struc:"float32,little"`
	T   float32 `struc:"float32,little"`
	U   float32 `struc:"float32,little"`
	V   float32 `struc:"float32,little"`
	W   float32 `struc:"float32,little"`
	X   float32 `struc:"float32,little"`
	Y   float32 `struc:"float32,little"`
	Z   float32 `struc:"float32,little"`
	AaC float32 `struc:"float32,little"`
	AbD float32 `struc:"float32,little"`
	AcD float32 `struc:"float32,little"`
	AdD float32 `struc:"float32,little"`
	AeD float32 `struc:"float32,little"`
	AfD float32 `struc:"float32,little"`
	AgA float32 `struc:"float32,little"`
	AhB float32 `struc:"float32,little"`
	AiC float32 `struc:"float32,little"`
	AjD float32 `struc:"float32,little"`
	AkD float32 `struc:"float32,little"`
	BA  float32 `struc:"float32,little"`
	BB  float32 `struc:"float32,little"`
	BC  float32 `struc:"float32,little"`
	BD  float32 `struc:"float32,little"`
	BE  float32 `struc:"float32,little"`
	BF  float32 `struc:"float32,little"`
	BG  float32 `struc:"float32,little"`
	BH  float32 `struc:"float32,little"`
	BI  float32 `struc:"float32,little"`
	BJ  float32 `struc:"float32,little"`
	BK  float32 `struc:"float32,little"`
	BL  float32 `struc:"float32,little"`
	BM  float32 `struc:"float32,little"`
	BN  float32 `struc:"float32,little"`
	BO  float32 `struc:"float32,little"`
	BP  float32 `struc:"float32,little"`
	BQ  float32 `struc:"float32,little"`
	BR  float32 `struc:"float32,little"`
	BS  float32 `struc:"float32,little"`
	BT  float32 `struc:"float32,little"`
	BU  float32 `struc:"float32,little"`
	BV  float32 `struc:"float32,little"`
	BW  float32 `struc:"float32,little"`
	BX  float32 `struc:"float32,little"`
	BY  float32 `struc:"float32,little"`
	BZ  float32 `struc:"float32,little"`
	CA  float32 `struc:"float32,little"`
	CB  float32 `struc:"float32,little"`
	CC  float32 `struc:"float32,little"`
	CD  float32 `struc:"float32,little"`
	CE  float32 `struc:"float32,little"`
	CF  float32 `struc:"float32,little"`
	CG  float32 `struc:"float32,little"`
	CH  float32 `struc:"float32,little"`
	CI  float32 `struc:"float32,little"`
	CJ  float32 `struc:"float32,little"`
	CK  float32 `struc:"float32,little"`
	CL  float32 `struc:"float32,little"`
	CM  float32 `struc:"float32,little"`
	AlD byte     `struc:"byte"`
	AmD byte     `struc:"byte"`
	AnD byte     `struc:"byte"`
	AoA byte     `struc:"byte"`
	ApB byte     `struc:"byte"`
	AqC byte     `struc:"byte"`
	ArD byte     `struc:"byte"`
	AsD byte     `struc:"byte"`
	AtD byte     `struc:"byte"`
	AuD byte     `struc:"byte"`
	AvD byte     `struc:"byte"`
	AwA byte     `struc:"byte"`
	AxB byte     `struc:"byte"`
	AyC byte     `struc:"byte"`
	AzD byte     `struc:"byte"`
	BaD byte     `struc:"byte"`
	BbD byte     `struc:"byte"`
	BcD byte     `struc:"byte"`
	BdD byte     `struc:"byte"`
	BeA byte     `struc:"byte"`
	BfB byte     `struc:"byte"`
	BgC byte     `struc:"byte"`
	BhD byte     `struc:"byte"`
	BiD byte     `struc:"byte"`
	BjD float32 `struc:"float32,little"` // Doh !
	BkD byte     `struc:"byte"`
	BlD byte     `struc:"byte"`
	BmD byte     `struc:"byte"`
	BmE byte     `struc:"byte"`
	BmF byte     `struc:"byte"`
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

func TestEvalQuality(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing ")
	//test1

	var buf bytes.Buffer
	arr := make([]byte, 900)

	src := &F1SourceData{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
		arr, 1.06, 10.7, 108.0, 0.109, 0.0110, 0.00111, 112, 113, 114, 115, 116, 117, 118	}

	err := struc.Pack(&buf, src)
	fmt.Printf("error code: %v \n", err)
	fmt.Printf("struct : \n %+v \n", src)

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

	rdata := tc.GetOutput("data")

	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("data") == nil {
		t.Fail()
	}

}
