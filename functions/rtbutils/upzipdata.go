package rtbutils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnunzipdata{})
}

type fnunzipdata struct {
}

func (fnunzipdata) Name() string {
	return "unzipdata"
}

func (fnunzipdata) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeBytes}, false
}

func (fnunzipdata) Eval(params ...interface{}) (interface{}, error) {
	inp, err := coerce.ToBytes(params[0])

	if err != nil {
		return nil, fmt.Errorf("unzip function first parameter [%+v] must be []byte", params[0])
	}

	reader := bytes.NewReader(inp)
	gzreader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, fmt.Errorf("unzip unable to process data")
	}
	output, err := ioutil.ReadAll(gzreader)
	if err != nil {
		return nil, fmt.Errorf("unzip unable to read result data")
	}
	return string(output), nil
}
