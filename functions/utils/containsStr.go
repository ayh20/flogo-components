package utils

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnContainsStr{})
}

type fnContainsStr struct {
}

func (fnContainsStr) Name() string {
	return "contains"
}

func (fnContainsStr) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (fnContainsStr) Eval(params ...interface{}) (interface{}, error) {
	return strings.Contains(params[0].(string), params[1].(string)), nil
}
