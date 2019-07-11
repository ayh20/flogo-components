package mqtt

import (
	"github.com/project-flogo/core/data/coerce"
)

// Input data structure
type Input struct {
	Broker string `md:"broker"`
	ID string `md:"id"`
	User string `md:"user"`
	Password string `md:"password"`
	EnableTLS bool `md:"enabletls"`
	CertStore string `md:"certstore"`
	Thing string `md:"thing"`
	Topic string `md:"topic"`
	QOS int `md:"qos"`
	Message string `md:"message"`
	JSONPayload bool `md:"jsonpayload"`
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"broker": i.Broker,
		"id": i.ID,
		"user": i.User,
		"password": i.Password,
		"enabletls": i.EnableTLS,
		"certstore": i.CertStore,
		"thing": i.Thing,
		"topic": i.Topic,
		"qos": i.QOS,
		"message": i.Message,
		"jsonpayload": i.JSONPayload,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Broker, err = coerce.ToString(values["broker"])
	if err != nil {
		return err
	}
	i.ID, err = coerce.ToString(values["id"])
	if err != nil {
		return err
	}
	i.User, err = coerce.ToString(values["user"])
	if err != nil {
		return err
	}
	i.Password, err = coerce.ToString(values["password"])
	if err != nil {
		return err
	}
	i.EnableTLS, err = coerce.ToBool(values["enabletls"])
	if err != nil {
		return err
	}
	i.CertStore, err = coerce.ToString(values["certstore"])
	if err != nil {
		return err
	}
	i.Thing, err = coerce.ToString(values["thing"])
	if err != nil {
		return err
	}
	i.Topic, err = coerce.ToString(values["topic"])
	if err != nil {
		return err
	}
	i.QOS, err = coerce.ToInt(values["qos"])
	if err != nil {
		return err
	}
	i.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}
	i.JSONPayload, err = coerce.ToBool(values["jsonpayload"])
	if err != nil {
		return err
	}
	return nil
}

//Output data structure
type Output struct {
	Result    string `md:"result"`    // The formatted CSV like data
}

//ToMap Output mapper
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}

//FromMap Output  from map
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}

	return nil
}
