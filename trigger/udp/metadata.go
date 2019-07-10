package udp

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Port           string `md:"port"`           // The UDP port to listen on
	MulticastGroup string `md:"multicastGroup"` // The multicast group for multicast messages
}

type HandlerSettings struct {
}

type Output struct {
	Payload string `md:"payload"` // The data received from the connection
	Buffer  []byte `md:"buffer"`  // The raw data received from the connection
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"payload": o.Payload,
		"buffer":  o.Buffer,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Payload, err = coerce.ToString(values["payload"])
	if err != nil {
		return err
	}

	o.Buffer, err = coerce.ToBytes(values["buffer"])
	if err != nil {
		return err
	}

	return nil
}
