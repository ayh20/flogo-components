package tcpudp

import (
	"bytes"
	"errors"
	"io"
	"net"
	"strings"
	"time"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{}, &Reply{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

// Factory is a kafka trigger factory
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

// Trigger is a kafka trigger
type Trigger struct {
	settings    *Settings
	handlers    []trigger.Handler
	listener    net.Listener
	logger      log.Logger
	connections []net.Conn
}

// Initialize initializes the trigger
func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	host := t.settings.Host
	port := t.settings.Port
	t.handlers = ctx.GetHandlers()
	t.logger = ctx.Logger()

	if port == "" {
		return errors.New("Valid port must be set")
	}

	listener, err := net.Listen(t.settings.Network, host+":"+port)
	if err != nil {
		return err
	}

	t.listener = listener

	return err
}

// Start starts the kafka trigger
func (t *Trigger) Start() error {

	go t.waitForConnection()
	t.logger.Infof("Started listener on Port - %s, Network - %s", t.settings.Port, t.settings.Network)
	return nil
}

func (t *Trigger) waitForConnection() {
	for {
		// Listen for an incoming connection.
		conn, err := t.listener.Accept()
		if err != nil {
			errString := err.Error()
			if !strings.Contains(errString, "use of closed network connection") {
				t.logger.Error("Error accepting connection: ", err.Error())
			}
			return
		} else {
			t.logger.Debugf("Handling new connection from client - %s", conn.RemoteAddr().String())
			// Handle connections in a new goroutine.
			go t.handleNewConnection(conn)
		}
	}
}

func (t *Trigger) handleNewConnection(conn net.Conn) {

	//Gather connection list for later cleanup
	t.connections = append(t.connections, conn)

	for {

		if t.settings.TimeOut > 0 {
			t.logger.Info("Setting timeout: ", t.settings.TimeOut)
			conn.SetDeadline(time.Now().Add(time.Duration(t.settings.TimeOut) * time.Millisecond))
		}

		output := &Output{}

		var buf bytes.Buffer
		_, err := io.Copy(&buf, conn)
		if err != nil {
			errString := err.Error()
			if !strings.Contains(errString, "use of closed network connection") {
				t.logger.Error("Error reading data from connection: ", err.Error())
			} else {
				t.logger.Info("Connection is closed.")
			}
			if nerr, ok := err.(net.Error); !ok || !nerr.Timeout() {
				// Return if not timeout error
				return
			}
		} else {
			output.Data = buf.Bytes()
		}

	}
}

// Stop implements ext.Trigger.Stop
func (t *Trigger) Stop() error {

	for i := 0; i < len(t.connections); i++ {
		t.connections[i].Close()
	}

	t.connections = nil

	if t.listener != nil {
		t.listener.Close()
	}

	t.logger.Info("Stopped listener")

	return nil
}
