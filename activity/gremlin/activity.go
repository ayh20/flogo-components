package gremlin

import (
	"fmt"

	"github.com/lucasalcantara/gremlin"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is a gremlin activity
type Activity struct {
	client *gremlin.Client
}

// New create a new gremlin activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	auth := gremlin.OptAuthUserPass(settings.User, settings.Password)
	client, err := gremlin.NewClient(settings.GremlinUrls, auth)
	if err != nil {
		ctx.Logger().Errorf("Gremlin new client initialization error: [%s]", err.Error())
		return nil, err
	}

	act := &Activity{client: client}
	return act, nil
}

// Metadata returns the metadata for the gremlin activity
func (*Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements the evaluation of the gremlin activity
func (act *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	if len(input.Query) == 0 {
		return false, fmt.Errorf("no query to run")
	}

	ctx.Logger().Debugf("Run gremlin query")

	data, err := act.client.ExecQuery(input.Query)

	//mydata, err := coerce.ToString(data)
	//fmt.Print(string(data))
	if err != nil {
		return false, fmt.Errorf("query failed for reason [%s]", err.Error())
	}

	//dataout, err := base64.StdEncoding.DecodeString(string(data))
	//if err != nil {
	//	return false, fmt.Errorf("decode failed for reason [%s]", //err.Error())
	//}
	// return data
	output := &Output{}
	output.Result = data

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
