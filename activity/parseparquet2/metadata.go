package parseparquet2

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
	ParquetFile string `md:"filename,required"`
	MaxRows     int    `md:"maxrows"`
	InitRow     int    `md:"initrow"`
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"filename": r.ParquetFile,
		"maxrows":  r.MaxRows,
		"initrow":  r.InitRow,
	}
}
func (r *Input) FromMap(values map[string]interface{}) error {

	var err error
	r.ParquetFile, err = coerce.ToString(values["filename"])
	if err != nil {
		return err
	}
	r.MaxRows, err = coerce.ToInt(values["maxrows"])
	if err != nil {
		return err
	}
	r.InitRow, err = coerce.ToInt(values["initrow"])
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	Result string `md:"result"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}
func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["result"])
	o.Result = strVal
	return nil
}
