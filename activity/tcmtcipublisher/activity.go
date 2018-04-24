package tcmtcipublisher

import (

	//"strconv"
	//"strings"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jvanderl/tib-eftl"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-TCMPublisher")

const (
	ivMessage = "message"
	ivKey     = "key"
	ivURL     = "url"
	ivChannel = "channel"
	ovResult  = "result"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// TCMPublisher is an Activity that  ..........................
// inputs : {input1, input2, datatype, comparemode}
// outputs: result (bool)
type TCMPublisher struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &TCMPublisher{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *TCMPublisher) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *TCMPublisher) Eval(context activity.Context) (done bool, err error) {

	// Get the runtime values
	inmessage, _ := context.GetInput(ivMessage).(string)
	inkey, _ := context.GetInput(ivKey).(string)
	inurl, _ := context.GetInput(ivURL).(string)
	inchannel, _ := context.GetInput(ivChannel).(string)
	uuid, err := newUUID()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// connection options
	activityLog.Info("Set Options")
	opts := &eftl.Options{
		Password: inkey, ClientID: uuid,
	}

	// channel on which to receive connection errors
	errChan := make(chan error, 1)

	// connect to TIBCO Cloud Messaging
	activityLog.Info("Connect")
	conn, err := eftl.Connect(inurl, opts, errChan)
	if err != nil {
		activityLog.Errorf("connect failed: %v", err)
	}

	// disconnect from TIBCO Cloud Messaging when done
	defer conn.Disconnect()

	// Listen for connection errors.
	go func() {
		for err := range errChan {
			activityLog.Errorf("connection error: %v", err)
		}
	}()

	// publish a message to TIBCO Cloud Messaging
	activityLog.Info("Publish Now")

	b := []byte(inmessage)

	var f interface{}
	err3 := json.Unmarshal(b, &f)
	if err3 != nil {
		activityLog.Errorf("Unmarshal fail: %v", err3)
	}

	m := f.(map[string]interface{})

	myMsg := eftl.Message{"_dest": inchannel}

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			myMsg[k] = vv
		case float64, int:
			myMsg[k] = vv
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			myMsg[k] = vv
		}
	}

	err2 := conn.Publish(myMsg)

	if err2 != nil {
		activityLog.Errorf("publish failed: %v", err2)
	}

	activityLog.Info("Published")

	context.SetOutput(ovResult, "Send OK")

	return true, nil
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
