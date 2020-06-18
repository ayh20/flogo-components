package rtbutils

import (
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnURLencode{})
}

type fnURLencode struct {
}

func (fnURLencode) Name() string {
	return "URLencode"
}

func (fnURLencode) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnURLencode) Eval(params ...interface{}) (interface{}, error) {
	return url.QueryEscape(params[0].(string)), nil
}
