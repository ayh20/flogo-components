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
	AlD int `struc:"byte"`
	AmD int `struc:"byte"`
	AnD int `struc:"byte"`
	AoA int `struc:"byte"`
	ApB int `struc:"byte"`
	AqC int `struc:"byte"`
	ArD int `struc:"byte"`
	AsD int `struc:"byte"`
	AtD int `struc:"byte"`
	AuD int `struc:"byte"`
	AvD int `struc:"byte"`
	AwA int `struc:"byte"`
	AxB int `struc:"byte"`
	AyC int `struc:"byte"`
	AzD int `struc:"byte"`
	BaD int `struc:"byte"`
	BbD int `struc:"byte"`
	BcD int `struc:"byte"`
	BdD int `struc:"byte"`
	BeA int `struc:"byte"`
	BfB int `struc:"byte"`
	BgC int `struc:"byte"`
	BhD int `struc:"byte"`
	BiD int `struc:"byte"`
	BjD int `struc:"byte"`
	BkD int `struc:"byte"`
	BlD int `struc:"byte"`
	BmD int `struc:"byte"`
}

func TestEvalQuality(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing ")
	//test1

	var buf bytes.Buffer
	src := &F1SourceData{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1}
	err := struc.Pack(&buf, src)
	fmt.Printf("error code: %v \n", err)
	fmt.Printf("struct : \n %+v \n", src)

	// convert to byte array and print
	x := buf.Bytes()
	// for _, v := range x {
	// 	fmt.Printf("%v - %v \n", v, x[v])
	// }
	fmt.Printf("buffer : \n %s \n", x)

	tc.SetInput("buffer", x)

	o := &F1SourceData{}
	err = struc.Unpack(&buf, o)
	fmt.Printf("struct : \n %+v \n", o)

	fmt.Println("#######   call routine ")
	act.Eval(tc)

	rdata := tc.GetOutput("data")

	fmt.Printf("csv data: %v \n", rdata)

	if tc.GetOutput("data") == nil {
		t.Fail()
	}

}
