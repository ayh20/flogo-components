package mqtt

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

var logquery = log.ChildLogger(log.RootLogger(), "mqttayh20-activity")

func init() {
	err := activity.Register(&Activity{}, New)
	if err != nil {
		logquery.Errorf("MQTT Init error : %s ", err.Error())
	}
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {

	logquery.Debug("In MQTT activity - New")

	settings := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	if settings.Connection != "" {
		logquery.Debug("Got a connection")
		mcon, toConnerr := coerce.ToConnection(settings.Connection)
		if toConnerr != nil {
			return nil, toConnerr
		}
		client := mcon.GetConnection().(*mqtt.Client)
		act := &Activity{client: client, jsonpayload: settings.JSONPayload, qos: settings.QOS, topic: settings.Topic}

		return act, nil
	}

	//logquery.Debug("Exiting MQTT activity - New")

	return nil, nil
}

// Activity is a MQTT with TLS activity
type Activity struct {
	client      *mqtt.Client
	jsonpayload bool
	qos         int
	topic       string
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

//Cleanup method
func (a *Activity) Cleanup() error {

	logquery.Debugf("Cleaning up MQTT activity")

	//ctx, cancel := ctx.WithTimeout(ctx.Background(), 30*time.Second)
	//defer cancel()

	// terminate the  mqtt session - not sure if tghis makes any sense ?
	//return a.client.Disconnect(ctx)
	return nil
}

// Eval implements activity.Activity.Eval
//func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	// Get the runtime values
	logquery.Debug("Starting Eval")

	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	output := &Output{}

	var ivpayload = ""
	if a.jsonpayload {
		ivpayload = makeMsg(ctx, in.Message)
		logquery.Debugf("Created Message: %v", ivpayload)
	} else {
		ivpayload = in.Message
	}

	client := *a.client

	logquery.Debugf("MQTT Publisher sending message %v on topic:%v qos:%v ", in.Message, a.topic, a.qos)
	token := client.Publish(a.topic, byte(a.qos), false, ivpayload)
	token.Wait()

	output.Result = "OK"
	ctx.SetOutputObject(output)

	logquery.Debug("End Eval")

	return true, nil
}

func makeMsg(ctx activity.Context, msgData interface{}) string {

	returnData := ""
	b, _ := json.Marshal(msgData)
	returnData = string(b)

	logquery.Debugf("MakeMsg returning data: %v", returnData)

	return returnData
}
