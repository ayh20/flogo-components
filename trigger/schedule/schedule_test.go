package timer

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

const testConfig string = `{
	"id": "flogo-schedule",
	"ref": "github.com/ayh20/flogo-components/trigger/schedule",
	"handlers": [
	  {
		"settings":{
			"startDay": "Tuesday",
			"startTime": "13:50"
		},
		"action":{
			"id":"dummy"
		}
	  }
	]
  }
  `

// const testConfig string = `{
// 	"id": "flogo-timer",
// 	"ref": "github.com/project-flogo/contrib/trigger/timer",
// 	"handlers": [
// 	  {
// 		"settings":{
// 			"repeatInterval" : "1s"
// 		},
// 		"action":{
// 			"id":"dummy"
// 		}
// 	  }
// 	]
//   }
//   `

func TestInitOk(t *testing.T) {
	f := &Factory{}
	tgr, err := f.New(nil)
	assert.Nil(t, err)
	assert.NotNil(t, tgr)
}

func TestTimerTrigger_Initialize(t *testing.T) {
	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
		//do nothing
	})}

	trg, err := test.InitTrigger(f, config, actions)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	time.Sleep(time.Second * 120)
	assert.Nil(t, err)
	err = trg.Stop()
	assert.Nil(t, err)

}
