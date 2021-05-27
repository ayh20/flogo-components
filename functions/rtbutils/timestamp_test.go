package rtbutils

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFntimestamp_Eval(t *testing.T) {
	f := &fntimestamp{}

	v, err := function.Eval(f, "02-Jan-2006")

	assert.Nil(t, err)
	//assert.Equal(t, Time.Format("02-Jan-2006"), v)
	fmt.Println(v)
	v, err = function.Eval(f, "2006-01-02 15:04:05.999")

	assert.Nil(t, err)
	fmt.Println(v)
}
