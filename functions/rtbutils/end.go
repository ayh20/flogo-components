package rtbutils

import (
	"os"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnend{})
}

type fnend struct {
}

func (fnend) Name() string {
	return "end"
}

func (fnend) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnend) Eval(params ...interface{}) (interface{}, error) {
	os.Exit(0)
	return "", nil
}
