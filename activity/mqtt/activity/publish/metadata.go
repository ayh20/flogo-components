package mqtt

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings structure
type Settings struct {
	Connection  string `md:"connection,required"`
	Topic       string `md:"topic"`
	QOS         int    `md:"qos"`
	JSONPayload bool   `md:"jsonpayload"`
}

// Input data structure
type Input struct {
	Message string `md:"message"`
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": i.Message,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}
	return nil
}

//Output data structure
type Output struct {
	Result string `md:"result"` // The formatted CSV like data
}

//ToMap Output mapper
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}

//FromMap Output  from map
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}

	return nil
}
