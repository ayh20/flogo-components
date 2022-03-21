package rtbutils

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnjsontocsv_Eval(t *testing.T) {
	f := &fnjsontocsv{}

	var testdata = `[{
        "vegetable": "carrot",
        "fruit": "banana",
        "rank": 1
    },
    {
        "vegetable": "potato",
        "fruit": "strawberry",
        "rank": 2
    }]`

	v, err := function.Eval(f, testdata)
	assert.Nil(t, err)
	fmt.Println(v)

}
