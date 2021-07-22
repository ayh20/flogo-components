package rtbutils

import (
	"bytes"
	"compress/gzip"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnzipdata{})
}

type fnzipdata struct {
}

func (fnzipdata) Name() string {
	return "zipdata"
}

func (fnzipdata) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnzipdata) Eval(params ...interface{}) (interface{}, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	gz.Write([]byte(params[0].(string)))
	gz.Close()
	return buf.Bytes(), nil
}
