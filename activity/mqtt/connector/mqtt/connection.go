package mqttconnection

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"

	b64 "encoding/base64"
	"io/ioutil"
	"strings"

	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/connection"
	"github.com/project-flogo/core/support/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var logmqtt = log.ChildLogger(log.RootLogger(), "mqttayh20-connection")
var factory = &mqttFactory{}

// Settings struct
type Settings struct {
	Name        string `md:"name,required"`
	Description string `md:"description"`
	Broker      string `md:"broker"`
	ID          string `md:"id"`
	User        string `md:"user"`
	Password    string `md:"password"`
	EnableTLS   bool   `md:"enabletls"`
	CertStore   string `md:"certstore"`
	Thing       string `md:"thing"`
}

func init() {
	err := connection.RegisterManagerFactory(factory)
	if err != nil {
		panic(err)
	}
}

type mqttFactory struct {
}

func (*mqttFactory) Type() string {
	return "mqtt"
}

func (*mqttFactory) NewManager(settings map[string]interface{}) (connection.Manager, error) {
	sharedConn := &mqttSharedConfigManager{}
	var err error
	logmqtt.Debug("Get shared config")
	sharedConn.config, err = getmqttClientConfig(settings)

	if err != nil {
		return nil, err
	}
	if sharedConn.mqttclient != nil {
		logmqtt.Debug("Reuse Shared")
		return sharedConn, nil
	}

	logmqtt.Debug("Create new config")
	opts := mqtt.NewClientOptions()

	opts.AddBroker(sharedConn.config.Broker)
	opts.SetClientID(sharedConn.config.ID)
	opts.SetUsername(sharedConn.config.User)
	opts.SetPassword(sharedConn.config.Password)

	if sharedConn.config.EnableTLS {
		//set tls config
		if sharedConn.config.Thing == "device" {
			tlsConfig := NewTLSConfig(sharedConn.config.CertStore)
			opts.SetTLSConfig(tlsConfig)
		} else {
			tlsConfig := AWSTLSConfig(sharedConn.config.Thing, sharedConn.config.CertStore)
			opts.SetTLSConfig(tlsConfig)
		}
	}

	client := mqtt.NewClient(opts)

	logmqtt.Debugf("MQTT Publisher connecting to broker %v using client ID %v", sharedConn.config.Broker, sharedConn.config.ID)

	// validate and return connection
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logmqtt.Debugf("%v\n", token.Error())
		return nil, token.Error()
	} else {
		logmqtt.Debugf("%v\n", token)
	}

	logmqtt.Debug("MQTT client connection returned")

	sharedConn.mqttclient = &client

	return sharedConn, nil
}

// mqttSharedConfigManager Structure
type mqttSharedConfigManager struct {
	config     *Settings
	name       string
	mqttclient *mqtt.Client
}

// Type of SharedConfigManager
func (k *mqttSharedConfigManager) Type() string {
	return "mqtt"
}

// GetConnection ss
func (k *mqttSharedConfigManager) GetConnection() interface{} {
	return k.mqttclient
}

// GetConfig ss
func (k *mqttSharedConfigManager) GetConfig() interface{} {
	return k.config
}

// ReleaseConnection ss
func (k *mqttSharedConfigManager) ReleaseConnection(connection interface{}) {

}

// Start connection manager
func (k *mqttSharedConfigManager) Start() error {
	return nil
}

// Stop connection manager
func (k *mqttSharedConfigManager) Stop() error {
	logmqtt.Debug("Cleaning up client connection cache")
	return nil
}

// GetSharedConfiguration function to return MongoDB connection manager
func GetSharedConfiguration(conn interface{}) (connection.Manager, error) {
	var cManager connection.Manager
	var err error
	cManager, err = coerce.ToConnection(conn)
	if err != nil {
		return nil, err
	}
	return cManager, nil
}

func getmqttClientConfig(settings map[string]interface{}) (*Settings, error) {
	connectionConfig := &Settings{}
	err := metadata.MapToStruct(settings, connectionConfig, false)
	if err != nil {
		return nil, err
	}
	return connectionConfig, nil
}

// parse cert

func parseCert(cert string) string {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(cert), &m)
	if err != nil {
		logmqtt.Errorf("=======Error Parsing Certificate for SSL handshake=====", err)
	}
	content := m["content"].(string)
	lastBin := strings.LastIndex(content, "base64,")
	endofString := len(content)
	sEnc := content[lastBin+7 : endofString]
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	return (string(sDec))
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
