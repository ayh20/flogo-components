package readdir

import (
	"github.com/project-flogo/core/data/coerce"
)

// Input data structure
type Input struct {
	DirName string `md:"dirname"`
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"dirname": i.DirName,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.DirName, err = coerce.ToString(values["dirname"])

	if err != nil {
		return err
	}
	return nil
}

//Output data structure
type Output struct {
	Result interface{} `md:"result"`
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
	o.Result, err = coerce.ToObject(values["result"])
	if err != nil {
		return err
	}

	return nil
}
