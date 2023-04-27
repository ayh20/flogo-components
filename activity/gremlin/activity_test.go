package gremlin

import (
	"fmt"
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

	settings := &Settings{GremlinUrls: "wss://qbe-cosmosdb-gml.gremlin.cosmos.azure.com:443", User: "/dbs/vehicle_events_db/colls/vehicle_events_graph", Password: "AjEPsihz85GUAEwMTJYc4CqXOKn5c1PH6oe7Ggk55mzqSmSWBYYAgSqwWkaDJezQAN3DSXrst2zZACDbvJ9ZHQ=="}

	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("query", `g.V()`)

	done, err := act.Eval(tc)
	res := tc.GetOutput("result")
	fmt.Print(res)
	assert.Nil(t, err)
	assert.True(t, done)
}
