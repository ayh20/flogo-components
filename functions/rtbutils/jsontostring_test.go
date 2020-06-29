package rtbutils

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnjsontostring_Eval(t *testing.T) {
	f := &fnjsontostring{}
	//v, err := function.Eval(f, '[{\"name\":\"Laureano Leyva (Director, Sales) \"},{\"name\":\"Victor Strauss (VP, Global Renewals) \"},{\"name\":\"Roger Seamons (Acoount Executive) \"},{\"name\":\"Griselda Sanchez (Sales Ops Analyst) \"},{\"name\":\"Jaime Varela Calderon (Staff Solutions Consultant) \"},{\"name\":\"Armando Gimbernat (Enterprise Account Executive) \"},{\"name\":\"Bruna Amato (Partner Manager - LATAM) \"}]", "name", ", ")

	//text := `{"text":[{"query":"Laureano Leyva (Director, Sales) "},{"query":"Victor Strauss (VP, Global Renewals) "},{"query":"Roger Seamons (Acoount Executive) "},{"query":"Griselda Sanchez (Sales Ops Analyst) "},{"query":"Jaime Varela Calderon (Staff Solutions Consultant) "},{"query":"Armando Gimbernat (Enterprise Account Executive) "},{"query":"Bruna Amato (Partner Manager - LATAM) "}]}`
	//birdJSON := `{"species": "pigeon","description": "likes to perch on rocks"}`
	birdJSON := `[{"name":"Laureano Leyva"},{"name":"Victor Strauss"},{"name":"Roger Seamons"},{"name":"Griselda Sanchez"},{"name":"Jaime Varela Calderon"},{"name":"Armando Gimbernat"},{"name":"Bruna Amato"}]`
	v, err := function.Eval(f, birdJSON, "query", ", ")
	assert.Nil(t, err)
	assert.Equal(t, "Laureano Leyva, Victor Strauss, Roger Seamons, Griselda Sanchez, Jaime Varela Calderon, Armando Gimbernat, Bruna Amato", v)

}
