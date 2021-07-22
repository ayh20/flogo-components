package rtbutils

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnzipit_Eval(t *testing.T) {
	f := &fnzipdata{}
	f2 := &fnunzipdata{}
	v, err := function.Eval(f, `{ "data": "ABC", "data2": 123 }`)

	assert.Nil(t, err)

	v, err = function.Eval(f2, v)

	assert.Nil(t, err)
	assert.Equal(t, `{ "data": "ABC", "data2": 123 }`, v)
}
