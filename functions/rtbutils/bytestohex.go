package rtbutils

import (
	"encoding/hex"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnbytestohex{})
}

type fnbytestohex struct {
}

func (fnbytestohex) Name() string {
	return "bytestohex"
}

func (fnbytestohex) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeBytes}, false
}

func (fnbytestohex) Eval(params ...interface{}) (interface{}, error) {
	inp, err := coerce.ToBytes(params[0])

	if err != nil {
		return nil, fmt.Errorf("first parameter [%+v] must be []byte", params[0])
	}

	return hex.EncodeToString(inp), nil
}
