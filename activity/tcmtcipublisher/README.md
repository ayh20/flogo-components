# TCM TCI Publisher
This activity sends a message to TIBCO Cloud Messaging in the format that is rquired by the TCI Flogo
TCM Subscriber to be able to recieve

The input message is a valid JSON formatted string that matches the Input format to the TCI Subscribers
definition. 

NOTE - This version only handles single level JSON documents. Next version will handle repearing structures
and child structures

## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/tcmtcipublisher
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "message",
      "required": true,
      "type": "string"
    },
    {
      "name": "key",
      "required": true,
      "type": "string"
    },
    {
      "name": "url",
      "required": true,
      "type": "string"
    },
    {
      "name": "channel",
      "required": true,
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```

## Settings
| Setting     | Description       |
|:------------|:------------------|
| message      | The JSON to be sent in string format   |
| key      | The security key/token to access the TCM channel  |
| url        | The url of the TCM channel |
| Channel        | The channel name specifiec in the TCI Subscriber definition |

## Outputs
| Output      | Description                             |
|:------------|:----------------------------------------|
| result      | Bool result based on input comparison   |

## Configuration Examples
### Simple
Configure a activity to send a message
```
{"message":, `{"_dest":"FlightData","hex":"48433a","flight":"KZR941","lat": 51.522354,"lon":-0.031771,"altitude":4225,"track":236,"speed":148,"messages":76 }`,
	"key", "773b42654dec94de14b659a9c1f01c69",
	"url", "wss://eu.messaging.cloud.tibco.com/tcm/01BKABHMAZKJDJA12PWHDR6WAP/channel",
	"channel", "myChannel"
}
```