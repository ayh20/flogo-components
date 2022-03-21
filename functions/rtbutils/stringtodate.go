package rtbutils

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnstringtodate{})
}

type fnstringtodate struct {
}

func (fnstringtodate) Name() string {
	return "stringtodate"
}

func (fnstringtodate) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnstringtodate) Eval(params ...interface{}) (interface{}, error) {

	date, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("format date first argument must be string")
	}
	format, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("format date second argument must be string")
	}

	t, err := time.Parse(format, date)

	if err != nil {
		return nil, fmt.Errorf("error occured: %w", err)
	}

	return t, nil
}
