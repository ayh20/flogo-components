package rtbutils

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnjsontocsv_Eval(t *testing.T) {
	f := &fnjsontocsv{}

	var testdata = `[{
        "vegetable": "carrot",
        "fruit": "banana",
        "rank": 1
    },
    {
        "vegetable": "potato",
        "fruit": "strawberry",
        "rank": 2
    }]`

	v, err := function.Eval(f, testdata, false)
	assert.Nil(t, err)
	fmt.Println(v)

}
func TestFnjsontocsv2_Eval(t *testing.T) {
	f := &fnjsontocsv{}

	var testdata = `{"aidifficulty":0,"airtemperature":21,"bestlap":false,"brakingassist":1,"drsassist":0,"dynamicracingline":2,"dynamicracinglinetype":1,"ersassist":1,"exclude":false,"formula":0,"gamepaused":0,"gearboxassist":3,"isspectating":0,"networkgame":0,"nummarshalzones":18,"pitassist":1,"pitreleaseassist":1,"pitspeedlimit":80,"seasonlinkidentifier":851129463,"sessionduration":240,"sessionlinkidentifier":851129463,"sessionstarttime":"2022-03-02T00:01:16.33Z","slipronativesupport":0,"spectatorcarindex":255,"steeringassist":0,"trackid":15,"tracktemperature":30,"type":9,"userid":9,"uuid":"baa3aa98-5ae8-4387-8d1e-bc84fdcb4db8","weather":0,"weekendlinkidentifier":851129463}`

	v, err := function.Eval(f, testdata, true)
	assert.Nil(t, err)
	fmt.Println(v)

}
