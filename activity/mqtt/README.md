# Publish MQTT Message with TLS support (AWS IoT/Eclipse Hono/Bosch IoT Suite)
This activity provides your flogo application the ability to publish a message on an MQTT topic.

This activity is based on the MQTT activity produced by Jan van der Lugt https://github.com/jvanderl/flogo-components/tree/master/incubator/activity/mqtt and the TLS MQTT trigger from Anshul Sharmas https://github.com/anshulsharmas/flogo-contrib/tree/master/trigger/mqtt 

Take care when using AWS IoT to get the correct certs/keys and Policy configuration. Topic, Client Name and keys must all match the policy definition.

Note This is the first version of the connector for the purposes of testing the new connection object structure... doc to be updated soon to refect changes

## Installation

```bash
flogo install github.com/ayh20/flogo-components/activity/mqtt/connection

flogo install github.com/ayh20/flogo-components/activity/mqtt/publish
```
Link for flogo web:
```
https://github.com/ayh20/flogo-components/activity/mqtt/connection
https://github.com/ayh20/flogo-components/activity/mqtt/publish
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
   {
      "name": "broker",
      "type": "string",
      "required": true
    },
    {
      "name": "id",
      "type": "string"
    },
    {
      "name": "user",
      "type": "string"
    },
    {
      "name": "password",
      "type": "string"
    },
    {
      "name": "enabletls",
      "type": "boolean"
    },
    {
      "name": "certstore",
      "type": "string"
    },
    {
      "name": "thing",
      "type": "string"
    },
    {
      "name": "topic",
      "type": "string",
      "required": true
    },
    {
      "name": "qos",
      "type": "integer",
      "required": true,
      "allowed" : ["0", "1", "2"]
    },
    {
      "name": "message",
      "type": "any",
      "required": true
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
## Settings

See activity-test.go for sample SSL/TLS parameters ...

| Setting   | Description    |
|:----------|:---------------|
| broker    | the MQTT Broker/AWS IoT URI (tcp://[hostname]:[port]) or ssl://[hostname]:8883 |
| id        | The MQTT Client ID (Must be a valid name in the AWS Policy files) |         
| user      | The UserID used when connecting to the MQTT IoT broker (Not AWS)|
| password  | The Password used when connecting to the MQTT broker (Not AWS) |
| certstore | For AWS TLS keys location directory, otherwise this is the server's TLS cert file |
| thing     | Blank unless connecting to AWS, then the thing name used for locating the correct TLS certs/keys in the certstore dir |
| topic     | Topic on which the message is published (Valid AWS policy entry required) |
| qos       | MQTT Quality of Service |
| message   | The message payload |


## Configuration Examples
### Simple
Configure a task in flow to publish a "hello world" message on MQTT topic called "flogo":

```json
{
  "id": 2,
  "name": "Publish MQTT Message",
  "type": 1,
  "activityType": "mqtt",
  "attributes": [
    {
      "name": "broker",
      "value": "tcp://localhost:1883",
      "type": "string"
    },
    {
      "name": "id",
      "value": "testmqtt",
      "type": "string"
    },
    {
      "name": "user",
      "value": "",
      "type": "string"
    },
    {
      "name": "password",
      "value": "",
      "type": "string"
    },
    {
      "name": "enabletls",
      "type": "boolean",
      "value": false
    },
    {
      "name": "certstore",
      "value": "",
      "type": "string"
    },
    {
      "name": "thing",
      "value": "",
      "type": "string"
    },
    {
      "name": "topic",
      "value": "flogo",
      "type": "string"
    },
    {
      "name": "qos",
      "value": "0",
      "type": "integer"
    },
    {
      "name": "message",
      "value": "Hello World",
      "type": "string"
    }
  ]
}
```