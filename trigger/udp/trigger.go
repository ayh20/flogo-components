package udp

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"

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
	lock           sync.Mutex
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

	if t.multicastGroup == "" {
		t.connection, err = net.ListenUDP("udp", t.address)
	} else {
		t.connection, err = net.ListenMulticastUDP("udp", nil, t.address)
	}

	if err != nil {
		return err
	}

	t.connection.SetReadBuffer(maxDatagramSize * 20)

	t.logger.Infof("Binding to %v : %v", t.address.Network(), t.address.Port)

	return err
}

// Start starts the trigger
func (t *Trigger) Start() error {
	t.logger.Info("Starting to listen for incoming udp messages")

	go t.ReadLoop()

	t.logger.Infof("Started listener on Port - %s", t.settings.Port)
	return nil
}

// ReadLoop process inbound messages
func (t *Trigger) ReadLoop() {
	for {
		t.ReadFromUDP()
	}
}

func (t *Trigger) ReadFromUDP() {
	t.lock.Lock()
	defer t.lock.Unlock()
	buf := make([]byte, maxDatagramSize)
	n, addr, err := t.connection.ReadFromUDP(buf)

	// Read ok ?
	if err != nil {
		//t.logger.Errorf("ReadFromUDP failed: %v", err)
		errString := err.Error()
		if !strings.Contains(errString, "use of closed network connection") {
			t.logger.Error("Error ReadFromUDP failed: ", err.Error())
		}
		return
	}

	//t.logger.Debugf("Received %v from %v", payload, addr)

	trgData := make(map[string]interface{})
	trgData["payload"] = string(buf[0:n])
	trgData["buffer"] = buf[0:n]
	trgData["address"] = addr.IP.String()

	t.logger.Debug("Processing handlers")
	for _, handler := range t.handlers {
		_, err := handler.Handle(context.Background(), trgData)
		if err != nil {
			t.logger.Error("Error starting action: ", err.Error())
		}
	}
}

// Stop implements trigger.Trigger.Stop
func (t *Trigger) Stop() error {
	// stop the trigger
	if t.connection != nil {
		t.connection.Close()
	}

	t.logger.Info("Stopped UDP trigger")

	return nil
}
