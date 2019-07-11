package udp

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

var jsonMetadata = getJSONMetadata()

func getJSONMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("descriptor.json")
	if err != nil {
		panic("No Json Metadata found for descriptor.json path")
	}
	return string(jsonMetadataBytes)
}

// Run Once, Start Immediately
const testConfig1 string = `{
  "name": "udp",
  "settings": {
		"port": "22600",
		"multicastGroup": "224.192.32.19"
  },
  "handlers": [
    {
      "action": {
		  "id" : "dummy"
	  },
      "settings": {
        "handler_setting": "xxx"
      }
    }
  ]
}`

// Listen for F1-2018 data
const testConfig2 string = `{
  "name": "udp",
  "settings": {
		"port": "20777",
		"multicastGroup": ""
  },
  "handlers": [
    {
		"action": {
			"id" : "dummy"
		},
		"settings": {
		  "handler_setting": "xxx"
		}
	  }
  ]
}`

func TestTrigger_Initialize(t *testing.T) {
	f := &Factory{}

	t.Log("Running test udp trigger")

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig1), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
		//do nothing

	})}

	trg, err := test.InitTrigger(f, config, actions)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	assert.Nil(t, err)
	err = trg.Stop()
	assert.Nil(t, err)

}

//192.168.1.19
/*
type TestRunner struct {
}

var Test action.Runner

// Run implements action.Runner.Run
func (tr *TestRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {
	log.Infof("Ran Action (Run): %v", uri)
	return 0, nil, nil
}

func (tr *TestRunner) RunAction(ctx context.Context, act action.Action, options map[string]interface{}) (results map[string]*data.Attribute, err error) {
	log.Infof("Ran Action (RunAction): %v", act)
	return nil, nil
}

func (tr *TestRunner) Execute(ctx context.Context, act action.Action, inputs map[string]*data.Attribute) (results map[string]*data.Attribute, err error) {
	log.Infof("Ran Action (Execute): %v", act)
	return nil, nil
}
*/

/*
//TODO: Fix Test
func TestUDPTrigger(t *testing.T) {
	log.Info("Testing UDP")
	config := trigger.Config{}

	//  Owl PV monitor test
	//json.Unmarshal([]byte(testConfig1), &config)

	// F1-2017 Telemtery
	json.Unmarshal([]byte(testConfig2), &config)

	f := &udpTriggerFactory{}
	f.metadata = trigger.NewMetadata(jsonMetadata)
	tgr := f.New(&config)
	runner := &TestRunner{}
	//tgr.Init(runner)
	tgr.Initialize()
	tgr.Start()
	defer tgr.Stop()
	log.Infof("Press CTRL-C to quit")
	for {
	}
}
*/
