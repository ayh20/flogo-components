package rtbutils

import (
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnurlencode{})
}

type fnurlencode struct {
}

func (fnurlencode) Name() string {
	return "urlencode"
}

func (fnurlencode) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnurlencode) Eval(params ...interface{}) (interface{}, error) {
	return url.QueryEscape(params[0].(string)), nil
}
