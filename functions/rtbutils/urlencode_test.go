package rtbutils

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnConcat_Eval(t *testing.T) {
	f := &fnurlencode{}
	v, err := function.Eval(f, "Hello World")

	assert.Nil(t, err)
	assert.Equal(t, "Hello+World", v)

	v, err = function.Eval(f, "this \\ \" data")

	assert.Nil(t, err)
	assert.Equal(t, "this+%5C+%22+data", v)
}
