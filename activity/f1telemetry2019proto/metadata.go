package f1telemetry2019proto

import (
	"github.com/project-flogo/core/data/coerce"
)

// Input data structure
type Input struct {
	Buffer []byte `md:"buffer"` // the UDP data packet
	Params string `md:"params"` // params
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"buffer": i.Buffer,
		"params": i.Params,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Buffer, err = coerce.ToBytes(values["buffer"])
	if err != nil {
		return err
	}
	i.Params, err = coerce.ToString(values["params"])
	if err != nil {
		return err
	}
	return nil
}

//Output data structure
type Output struct {
	MsgType     int    `md:"msgtype"`     // The data format type of this UDP packet
	Data        []byte `md:"data"`        // The protobuf data
	AuxData     []byte `md:"auxdata"`     // The protobuf data for session record (msgtype=1)
	SessionGUID string `md:"sessionguid"` // The protobuf data
}

//ToMap Output mapper
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"msgtype":     o.MsgType,
		"data":        o.Data,
		"auxdata":     o.AuxData,
		"sessionguid": o.SessionGUID,
	}
}

//FromMap Output  from map
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.MsgType, err = coerce.ToInt(values["msgtype"])
	if err != nil {
		return err
	}

	o.Data, err = coerce.ToBytes(values["data"])
	if err != nil {
		return err
	}

	o.AuxData, err = coerce.ToBytes(values["auxdata"])
	if err != nil {
		return err
	}

	o.SessionGUID, err = coerce.ToString(values["sessionguid"])
	if err != nil {
		return err
	}

	return nil
}

// F1Parms - struct for inbound params
type F1Parms struct {
	FeedNameGUID string `json:"FeedNameGUID"`
	FeedName     string `json:"FeedName"`
	StreamID     string `json:"StreamId"`
	Source       string `json:"Source"`
	Quality      int    `json:"Quality"`
	Frequency    int    `json:"Frequency"`
}
