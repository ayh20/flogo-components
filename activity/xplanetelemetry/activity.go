package xplanetelemetry

import (
	"github.com/project-flogo/core/activity"

	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
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

	output := &Output{}

	// Get the runtime values
	ctx.Logger().Debug("Starting")

	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	ctx.Logger().Debugf("input : \n %x \n", in.Buffer)

	// Get byte stream from input
	b := bytes.NewBuffer(in.Buffer)

	// testing
	//var b bytes.Buffer // A Buffer needs no initialization.
	//b.Write([]byte{0x35, 0xce, 0x86, 0x41, 0x33, 0x33, 0x9f, 0x41, 0x00, 0xc0, 0x79, 0xc4, 0x7e, 0x13, 0x73, 0x3d, 0x89, 0xd1, 0x43, 0x3d, 0xd6, 0x4b, 0x75, 0x3d, 0x00, 0x00, 0x80, 0x3f})

	// Create byte array to hold data
	var headerdata []byte
	var internaluse []byte
	var indexdata []byte
	var floatdata []byte

	floatdata = make([]byte, 4)
	headerdata = make([]byte, 4)
	internaluse = make([]byte, 1)
	indexdata = make([]byte, 4)
	var resultdata string

	//peel off header
	b.Read(headerdata)
	b.Read(internaluse)
	b.Read(indexdata)

	ctx.Logger().Debugf("header : %+v flag : %+v index : %+v\n", headerdata, internaluse, indexdata)

	for {
		_, err := b.Read(floatdata)

		if err == io.EOF {
			resultdata = resultdata
			break
		}

		float := math.Float32frombits(binary.LittleEndian.Uint32(floatdata))
		resultdata = resultdata + fmt.Sprintf("%+v,", float)
		//fmt.Printf("\n float= %+v", float)
	}

	output.Data = resultdata
	output.MsgType = 1
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
