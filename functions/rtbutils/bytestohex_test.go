package rtbutils

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnbytestohex_Eval(t *testing.T) {
	f := &fnbytestohex{}

	arr, _ := hex.DecodeString("e507010401030000000000000000000000000000000013ff4255544e5000000003000000")
	v, err := function.Eval(f, arr)

	assert.Nil(t, err)
	fmt.Print("BytesToHex - ")
	fmt.Println(v)

}
