package rtbutils

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fntimestamp{})
}

type fntimestamp struct {
}

func (fntimestamp) Name() string {
	return "timestamp"
}

func (fntimestamp) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fntimestamp) Eval(params ...interface{}) (interface{}, error) {

	format, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Date Format argument must be string")
	}

	t := time.Now().Format(format)

	return t, nil
}
