package rtbutils

import (
	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnjsontostring{})
}

type fnjsontostring struct {
}

func (fnjsontostring) Name() string {
	return "jsontostring"
}

func (fnjsontostring) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

func (fnjsontostring) Eval(params ...interface{}) (interface{}, error) {

	// Unmarshal JSON data
	type People struct {
		Name string
	}

	var people []People
	err := json.Unmarshal([]byte(params[0].(string)), &people)
	sep := params[2].(string)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	total := ""
	for _, elem := range people {
		if total == "" {
			total = elem.Name
		} else {
			total = total + sep + elem.Name
		}
	}

	return total, nil
}
