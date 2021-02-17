package kafka

import (
	"encoding/hex"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestPlain(t *testing.T) {

	settings := &Settings{BrokerUrls: "localhost:9092", Topic: "syslog"}

	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	arr, _ := hex.DecodeString("e30701040103c655e025b6e12cb300000000000000001353535441a009d74101")
	tc.SetInput("message", arr)

	done, err := act.Eval(tc)
	assert.Nil(t, err)
	assert.True(t, done)
}
