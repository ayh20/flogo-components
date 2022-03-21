package rtbutils

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fndatetostring{})
}

type fndatetostring struct {
}

func (fndatetostring) Name() string {
	return "datetostring"
}

func (fndatetostring) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fndatetostring) Eval(params ...interface{}) (interface{}, error) {

	date, err := coerce.ToDateTime(params[0])
	if err != nil {
		return nil, fmt.Errorf("format date first argument must be a datetime")
	}
	format, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("format date second argument must be string")
	}

	t := date.Format(format)

	//if err != nil {
	//	return nil, fmt.Errorf("Error Occured: %w", err)
	//}

	return t, nil
}
