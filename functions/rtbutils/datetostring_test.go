package rtbutils

import (
	"fmt"
	"testing"
	"time"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFndatetofndatetostring_Eval(t *testing.T) {
	f := &fndatetostring{}

	v, err := function.Eval(f, time.Now(), "02-Jan-2006")

	assert.Nil(t, err)
	//assert.Equal(t, Time.Format("02-Jan-2006"), v)
	fmt.Println(v)
	v, err = function.Eval(f, time.Now(), "2006-01-02T15:04:05Z07:00")

	assert.Nil(t, err)
	fmt.Println(v)
}
