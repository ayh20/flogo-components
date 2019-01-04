package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/eclipse/paho.mqtt.golang"
)

// log is the default package logger
var log = logger.GetLogger("activity-ayh20-mqtt-tls")

const (
	broker    = "broker"
	topic     = "topic"
	qos       = "qos"
	payload   = "message"
	id        = "id"
	user      = "user"
	password  = "password"
	enabletls = "enabletls"
	certstore = "certstore"
	thing     = "thing"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval

	brokerInput := context.GetInput(broker)

	ivbroker, ok := brokerInput.(string)
	if !ok {
		context.SetOutput("result", "BROKER_NOT_SET")
		return true, fmt.Errorf("broker not set")
	}

	topicInput := context.GetInput(topic)

	ivtopic, ok := topicInput.(string)
	if !ok {
		context.SetOutput("result", "TOPIC_NOT_SET")
		return true, fmt.Errorf("topic not set")
	}

	payloadInput := context.GetInput(payload)
	if payloadInput == nil {
		context.SetOutput("result", "PAYLOAD_NOT_SET")
		return true, fmt.Errorf("payload not set")
	}

	ivpayload := makeMsg(payloadInput)
	log.Debugf("Created Message: %v", ivpayload)

	ivqos, ok := context.GetInput(qos).(int)
	if !ok {
		context.SetOutput("result", "QOS_NOT_SET")
		return true, fmt.Errorf("qos not set")
	}

	idInput := context.GetInput(id)

	ivID, ok := idInput.(string)
	if !ok {
		context.SetOutput("result", "CLIENTID_NOT_SET")
		return true, fmt.Errorf("client id not set")
	}

	userInput := context.GetInput(user)

	ivUser, ok := userInput.(string)
	if !ok {
		//User not set, use default
		ivUser = ""
	}

	passwordInput := context.GetInput(password)

	ivPassword, ok := passwordInput.(string)
	if !ok {
		//Password not set, use default
		ivPassword = ""
	}

	// Get settings for TLS (store location, thing name)
	tlsInput := context.GetInput(enabletls)

	tlsEnabled, ok := tlsInput.(bool)
	if !ok {
		context.SetOutput("result", "ENABLETLS_NOT_SET")
		return true, fmt.Errorf("ENABLE TLS not set")
	}

	certstoreInput := context.GetInput(certstore)

	ivCertStore, ok := certstoreInput.(string)
	if !ok {
		//Store not set, use default
		ivCertStore = ""
	}

	thingInput := context.GetInput(thing)

	ivThing, ok := thingInput.(string)
	if !ok || ivThing == "" {
		//thing not set, use default
		ivThing = "device"
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(ivbroker)
	opts.SetClientID(ivID)
	opts.SetUsername(ivUser)
	opts.SetPassword(ivPassword)

	if tlsEnabled {
		//set tls config
		tlsConfig := NewTLSConfig(ivThing, ivCertStore)
		opts.SetTLSConfig(tlsConfig)
	}

	client := mqtt.NewClient(opts)

	log.Debugf("MQTT Publisher connecting to broker %v using client ID %v", ivbroker, ivID)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("%v\n", token.Error())
		panic(token.Error())
	}

	log.Debugf("MQTT Publisher connected, sending message on topic %v", ivtopic)
	token := client.Publish(ivtopic, byte(ivqos), false, ivpayload)
	token.Wait()

	client.Disconnect(250)
	log.Debugf("MQTT Publisher disconnected")
	context.SetOutput("result", "OK")

	return true, nil
}

func makeMsg(msgData interface{}) string {

	returnData := ""
	b, _ := json.Marshal(msgData)
	returnData = string(b)

	log.Debugf("MakeMsg returning data: %v", returnData)

	return returnData
}

// NewTLSConfig creates a TLS configuration for the specified 'thing'
func NewTLSConfig(thingName string, thingDir string) *tls.Config {

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
