package f1telemetry2019

import (
	"github.com/project-flogo/core/data/coerce"
)

// Input data structure
type Input struct {
	Buffer []byte `md:"buffer"` // the UDP data packet
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"buffer": i.Buffer,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Buffer, err = coerce.ToBytes(values["buffer"])
	if err != nil {
		return err
	}
	return nil
}

//Output data structure
type Output struct {
	MsgType int    `md:"msgtype"` // The data format type of this UDP packet
	Data    string `md:"data"`    // The formatted CSV like data
}

//ToMap Output mapper
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"msgtype": o.MsgType,
		"data":    o.Data,
	}
}

//FromMap Output  from map
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
