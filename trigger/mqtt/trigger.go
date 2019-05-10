package mqtt

import (
	"context"

	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"

	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/eclipse/paho.mqtt.golang"
)

// log is the default package logger
var log = logger.GetLogger("trigger-ayh20-mqtt-tls")

// default handlername
var defaultHandler string

// MqttTrigger is simple MQTT trigger
type MqttTrigger struct {
	metadata       *trigger.Metadata
	client         mqtt.Client
	config         *trigger.Config
	handlers       []*trigger.Handler
	topicToHandler map[string]*trigger.Handler
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MQTTFactory{metadata: md}
}

// MQTTFactory MQTT Trigger factory
type MQTTFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *MQTTFactory) New(config *trigger.Config) trigger.Trigger {
	return &MqttTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *MqttTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize implements trigger.Initializable.Initialize
func (t *MqttTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	return nil
}

// Start implements trigger.Trigger.Start
func (t *MqttTrigger) Start() error {

	log.Debugf("MQTT Trigger: %v", "Starting")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(t.config.GetSetting("broker"))
	opts.SetClientID(t.config.GetSetting("id"))
	opts.SetUsername(t.config.GetSetting("user"))
	opts.SetPassword(t.config.GetSetting("password"))

	if t.config.Settings["keepalive"] != "" {
		k, err := data.CoerceToInteger(t.config.Settings["keepalive"])
		if err != nil {
			log.Error("Error converting \"keepalive\" to an integer ", err.Error())
			return err
		}

		opts.SetKeepAlive(time.Duration(k) * time.Second)
	}

	ac, err := data.CoerceToBoolean(t.config.Settings["autoreconnect"])
	if err != nil {
		log.Error("Error converting \"autoreconnect\" to a boolean ", err.Error())
		return err
	}
	opts.SetAutoReconnect(ac)

	b, err := data.CoerceToBoolean(t.config.Settings["cleansess"])
	if err != nil {
		log.Error("Error converting \"cleansess\" to a boolean ", err.Error())
		return err
	}
	opts.SetCleanSession(b)

	if storeType := t.config.Settings["store"]; storeType != ":memory:" && storeType != "" {
		opts.SetStore(mqtt.NewFileStore(t.config.GetSetting("store")))
	}

	// Get settings for TLS (store location, thing name)
	tlsEnabled, err := data.CoerceToBoolean(t.config.Settings["enabletls"])
	if err != nil {
		log.Error("Error converting \"enabletls\" to a boolean ", err.Error())
		return err
	}

	log.Debugf("MQTT Trigger: %v", "Values set, do TLS")

	ivCertStore := t.config.GetSetting("certstore")

	ivThing := t.config.GetSetting("thing")

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		//topic := msg.Topic()

		log.Debugf("MQTT Trigger: %v", "In pub handler")
		//TODO we should handle other types, since mqtt message format are data-agnostic
		payload := string(msg.Payload())
		log.Debug("Received msg:", payload)

		// add a hack here .... if the topic doesn't exist use the first/default one ---- this should be a patten matcher
		handler, found := t.topicToHandler[defaultHandler]
		if found {
			t.RunHandler(handler, payload)
		} else {
			log.Errorf("handler for topic '%s' not found", defaultHandler)
		}
	})

	//set tls config
	if tlsEnabled {
		// if thing is not set, this indicates that this is not AWS IoT
		if ivThing == "" {
			tlsConfig := NewTLSConfig(ivCertStore)
			opts.SetTLSConfig(tlsConfig)
		} else {
			tlsConfig := AWSTLSConfig(ivThing, ivCertStore)
			opts.SetTLSConfig(tlsConfig)
		}
	}

	log.Debugf("Opts for mqtt client: %+v", opts)

	log.Debugf("MQTT Trigger: %v", "Create handler")
	client := mqtt.NewClient(opts)
	t.client = client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	i, err := data.CoerceToDouble(t.config.Settings["qos"])
	if err != nil {
		log.Error("Error converting \"qos\" to an integer ", err.Error())
		return err
	}

	t.topicToHandler = make(map[string]*trigger.Handler)

	d := 1
	for _, handler := range t.handlers {
		log.Debugf("MQTT Trigger: %v", "Start handlers")
		topic := handler.GetStringSetting("topic")

		if token := t.client.Subscribe(topic, byte(i), nil); token.Wait() && token.Error() != nil {
			log.Errorf("Error subscribing to topic %s: %s", topic, token.Error())
			return token.Error()
		}
		log.Debugf("Subscribed to topic: %s, will trigger handler: %s", topic, handler)
		t.topicToHandler[topic] = handler
		// set the first handler as the default
		if d == 1 {
			defaultHandler = topic
			d = 0
		}

	}

	return nil
}

// Stop implements ext.Trigger.Stop
func (t *MqttTrigger) Stop() error {
	//unsubscribe from topic
	for _, handlerCfg := range t.config.Handlers {
		log.Debug("Unsubscribing from topic: ", handlerCfg.GetSetting("topic"))
		if token := t.client.Unsubscribe(handlerCfg.GetSetting("topic")); token.Wait() && token.Error() != nil {
			log.Errorf("Error unsubscribing from topic %s: %s", handlerCfg.Settings["topic"], token.Error())
		}
	}

	t.client.Disconnect(250)

	return nil
}

// RunHandler runs the handler and associated action
func (t *MqttTrigger) RunHandler(handler *trigger.Handler, payload string) {

	trgData := make(map[string]interface{})
	trgData["message"] = payload

	results, err := handler.Handle(context.Background(), trgData)

	if err != nil {
		log.Error("Error starting action: ", err.Error())
	}

	log.Debugf("Ran Handler: [%s]", handler)

	var replyData interface{}

	if len(results) != 0 {
		dataAttr, ok := results["data"]
		if ok {
			replyData = dataAttr.Value()
		}
	}

	if replyData != nil {
		dataJSON, err := json.Marshal(replyData)
		if err != nil {
			log.Error(err)
		} else {
			replyTo := handler.GetStringSetting("topic")
			if replyTo != "" {
				t.publishMessage(replyTo, string(dataJSON))
			}
		}
	}
}

func (t *MqttTrigger) publishMessage(topic string, message string) {

	log.Debug("ReplyTo topic: ", topic)
	log.Debug("Publishing message: ", message)

	qos, err := strconv.Atoi(t.config.GetSetting("qos"))
	if err != nil {
		log.Error("Error converting \"qos\" to an integer ", err.Error())
		return
	}
	if len(topic) == 0 {
		log.Warn("Invalid empty topic to publish to")
		return
	}
	token := t.client.Publish(topic, byte(qos), false, message)

	sent := token.WaitTimeout(5000 * time.Millisecond)
	if !sent {
		// Timeout occurred
		log.Errorf("Timeout occurred while trying to publish to topic '%s'", topic)
		return
	}
}

// NewTLSConfig creates a TLS configuration using passed cert
func NewTLSConfig(certstore string) *tls.Config {

	// Import root CA
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(certstore)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
	}
}

// AWSTLSConfig creates a TLS configuration for the specified 'thing'
func AWSTLSConfig(thingName string, thingDir string) *tls.Config {

	thingLoc := thingDir

	if thingLoc != "" {
		thingLoc += "/"
	}

	// Import root CA
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(thingLoc + "root-CA.crt")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair for the specified 'thing'
	cert, err := tls.LoadX509KeyPair(thingLoc+thingName+".cert.pem", thingLoc+thingName+".private.key")
	if err != nil {
		panic(err)
	}

	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}

	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
}
