package udp

import (
	"context"
	"errors"
	"net"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{settings: s}, nil
}

// Trigger is a stub for your Trigger implementation
type Trigger struct {
	settings       *Settings
	handlers       []trigger.Handler
	logger         log.Logger
	connection     *net.UDPConn
	address        *net.UDPAddr
	port           string
	multicastGroup string
}

const (
	maxDatagramSize = 8192
)

// Initialize initializes the trigger
func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	port := t.settings.Port
	multicastGroup := t.settings.MulticastGroup
	t.handlers = ctx.GetHandlers()
	t.logger = ctx.Logger()

	if port == "" {
		return errors.New("Valid port must be set")
	}

	t.logger.Infof("Set UDP port to %v", port)

	var err error
	t.address, err = net.ResolveUDPAddr("udp", multicastGroup+":"+port)

	if err != nil {
		// Handle error
		t.logger.Errorf("Error resolving address %v", err)
		return nil
	}
	t.port = port
	t.multicastGroup = multicastGroup

	t.logger.Infof("Binding to %v : %v", t.address.Network(), t.address.Port)

	return err
}

// Start starts the trigger
func (t *Trigger) Start() error {
	t.logger.Info("Starting to listen for incoming udp messages")

	var err error
	if t.multicastGroup == "" {
		t.connection, err = net.ListenUDP("udp", t.address)
	} else {
		t.connection, err = net.ListenMulticastUDP("udp", nil, t.address)
	}

	if err != nil {
		return err
	}

	t.connection.SetReadBuffer(maxDatagramSize)

	for {
		buf := make([]byte, maxDatagramSize)

		output := &Output{}

		n, addr, err := t.connection.ReadFromUDP(buf)

		// Read ok ?
		if err != nil {
			t.logger.Errorf("ReadFromUDP failed: %v", err)
		}

		t.logger.Debug("after ReadFromUDP")
		payload := string(buf[0:n])
		payloadB := buf[0:n]

		t.logger.Debugf("Received %v from %v", payload, addr)

		trgData := make(map[string]interface{})
		trgData["payload"] = payload
		trgData["buffer"] = payloadB

		output.Payload = payload
		output.Buffer = payloadB

		t.logger.Debug("Processing handlers")
		for _, handler := range t.handlers {
			results, err := handler.Handle(context.Background(), trgData)
			if err != nil {
				t.logger.Error("Error starting action: ", err.Error())
			}
			t.logger.Debugf("Ran Handler: [%s]", handler)
			t.logger.Debugf("Results: [%v]", results)
		}
	}
}

// Stop implements trigger.Trigger.Start
func (t *Trigger) Stop() error {
	// stop the trigger
	if t.connection != nil {
		t.connection.Close()
	}

	t.logger.Info("Stopped UDP trigger")

	return nil
}
