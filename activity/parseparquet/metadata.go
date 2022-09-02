package parseparquet

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
	ParquetFile string `md:"parquetFile,required"`
	MaxRows     int    `md:"maxRows"`
	InitRow     int    `md:"initRow"`
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"parquetFile": r.ParquetFile,
		"maxRows":     r.MaxRows,
		"initRow":     r.InitRow,
	}
}
func (r *Input) FromMap(values map[string]interface{}) error {

	var err error
	r.ParquetFile, err = coerce.ToString(values["parquetFile"])
	if err != nil {
		return err
	}
	r.MaxRows, err = coerce.ToInt(values["maxRows"])
	if err != nil {
		return err
	}
	r.InitRow, err = coerce.ToInt(values["initRow"])
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	Output string `md:"output"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutput": o.Output,
	}
}
func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["output"])
	o.Output = strVal
	return nil
}
