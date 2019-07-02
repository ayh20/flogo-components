package f1telemetry2019

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Buffer []byte `md:"buffer"` // the UDP data packet
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"buffer": i.Buffer,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Buffer, err = coerce.ToBytes(values["buffer"])
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	MsgType int    `md:"msgtype"` // The data format type of this UDP packet
	Data    string `md:"data"`    // The formatted CSV like data
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"msgtype": o.MsgType,
		"data":    o.Data,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.MsgType, err = coerce.ToInt(values["msgtype"])
	if err != nil {
		return err
	}

	o.Data, err = coerce.ToString(values["data"])
	if err != nil {
		return err
	}

	return nil
}
