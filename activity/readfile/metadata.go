package readfile

import (
	"github.com/project-flogo/core/data/coerce"
)

// Input data structure
type Input struct {
	Filename string `md:"filename"`
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"filename": i.Filename,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error

	i.Filename, err = coerce.ToString(values["filename"])
	if err != nil {
		return err
	}

	return nil
}

//Output data structure
type Output struct {
	Result string `md:"result"`
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
