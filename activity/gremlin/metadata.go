package gremlin

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings - add comments
type Settings struct {
	GremlinUrls string `md:"gremlinUrls,required"` // The gremlin server to connect to
	User        string `md:"user"`                 // If connecting to a SASL enabled port, the user id to use for authentication
	Password    string `md:"password"`             // If connecting to a SASL enabled port, the password to use for authentication
}

// Input - add comments
type Input struct {
	Query string `md:"query,required"` // The query to run
}

// ToMap -
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"query": i.Query,
	}
}

// FromMap - add comments
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Query, err = coerce.ToString(values["query"])
	return err
}

// Output - add comments
type Output struct {
	Result interface{} `md:"result"` // Json object returned from the server
}

// ToMap - add comments
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}

// FromMap - add comments
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Result, err = coerce.ToObject(values["result"])
	if err != nil {
		return err
	}

	return nil
}
