package rtbutils

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnstringtodate_Eval(t *testing.T) {
	f := &fnstringtodate{}

	v, err := function.Eval(f, "10-Feb-2021", "02-Jan-2006")

	assert.Nil(t, err)
	//assert.Equal(t, Time.Format("02-Jan-2006"), v)
	fmt.Println(v)
	v, err = function.Eval(f, "July 29, 2020", "January 02, 2006")

	assert.Nil(t, err)
	fmt.Println(v)
}
