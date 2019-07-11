package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/project-flogo/core/activity"
)

// Activity is a MQTT with TLS activity
type Activity struct {
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {

	ctx.Logger().Info("In MQTT activity")

	act := &Activity{}
	return act, nil
}

// const (
// 	broker    = "broker"
// 	topic     = "topic"
// 	qos       = "qos"
// 	payload   = "message"
// 	id        = "id"
// 	user      = "user"
// 	password  = "password"
// 	enabletls = "enabletls"
// 	certstore = "certstore"
// 	thing     = "thing"
// )

// Eval implements activity.Activity.Eval
//func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	// Get the runtime values
	ctx.Logger().Debug("Starting")

	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	output := &Output{}

	if in.Broker == "" {
		ctx.Logger().Debug("BROKER_NOT_SET")
		output.Result = "BROKER_NOT_SET"
		ctx.SetOutputObject(output)
		return false, fmt.Errorf("Broker not set")
	}

	if in.Topic == "" {
		ctx.Logger().Debug("TOPIC_NOT_SET")
		output.Result = "TOPIC_NOT_SET"
		ctx.SetOutputObject(output)
		return false, fmt.Errorf("Topic not set")
	}

	var ivpayload = ""
	if in.JSONPayload {
		ivpayload = makeMsg(ctx, in.Message)
		ctx.Logger().Debugf("Created Message: %v", ivpayload)
	} else {
		ivpayload = in.Message
	}

	if in.ID == "" {
		ctx.Logger().Debug("CLIENTID_NOT_SET")
		output.Result = "CLIENTID_NOT_SET"
		ctx.SetOutputObject(output)
		return false, fmt.Errorf("Client ID not set")
	}

	if in.Thing == "" {
		//thing not set, use default and this indicates that this is not AWS
		in.Thing = "device"
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(in.Broker)
	opts.SetClientID(in.ID)
	opts.SetUsername(in.User)
	opts.SetPassword(in.Password)

	if in.EnableTLS {
		//set tls config
		if in.Thing == "device" {
			tlsConfig := NewTLSConfig(in.CertStore)
			opts.SetTLSConfig(tlsConfig)
		} else {
			tlsConfig := AWSTLSConfig(in.Thing, in.CertStore)
			opts.SetTLSConfig(tlsConfig)
		}
	}

	client := mqtt.NewClient(opts)

	ctx.Logger().Debugf("MQTT Publisher connecting to broker %v using client ID %v", in.Broker, in.ID)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("%v\n", token.Error())
		panic(token.Error())
	}

	ctx.Logger().Debugf("MQTT Publisher connected, sending message on topic %v", in.Topic)
	token := client.Publish(in.Topic, byte(in.QOS), false, ivpayload)
	token.Wait()

	client.Disconnect(250)
	ctx.Logger().Debugf("MQTT Publisher disconnected")

	output.Result = "OK"
	ctx.SetOutputObject(output)

	return true, nil
}

func makeMsg(ctx activity.Context, msgData interface{}) string {

	returnData := ""
	b, _ := json.Marshal(msgData)
	returnData = string(b)

	ctx.Logger().Debugf("MakeMsg returning data: %v", returnData)

	return returnData
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
