package rtbutils

import (
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnpathencode{})
}

type fnpathencode struct {
}

func (fnpathencode) Name() string {
	return "pathencode"
}

func (fnpathencode) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnpathencode) Eval(params ...interface{}) (interface{}, error) {
	return url.PathEscape(params[0].(string)), nil
}
