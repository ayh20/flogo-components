# f1telemetry
This activity Decodes telemetry data from Codemasters F1-2019, and places it in to a "CSV" "record". 
This is designed as a demonstration for integrating F1 2019 .. Flogo .... TIBCO Streambase and TIBCO Spotfire/Liveview

The Flogo app needs to read the data in from UDP and feed the byte stream into this component for decoding ... it's then passed to Streambase via some form of messaging (MQTT in my Demo)

The UDP data is a packed little endian C struct, which is transformed to a Go struct by https://github.com/lunixbochs/struc. This uses struct tags to decode the raw data correctly to the Go struct.

Information about the format of the raw data can be found here: https://forums.codemasters.com/topic/38920-f1-2019-udp-specification/

## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/f1telemetry2019
```

## Schema
Inputs and Outputs: 

```json
{
 },
  "inputs":[
    {
      "name": "buffer",
      "required": true,
      "type": "any"
    }
  ],
  "outputs": [
    {
      "name": "msgtype",
      "type": "integer"
    },
    {
      "name": "data",
      "type": "string"
    }
  ]
}
```

## Settings
| Setting     | Description       |
|:------------|:------------------|
| buffer      | The raw UDP data   |

## Outputs
| Output      | Description                             |
|:------------|:----------------------------------------|
| data        | CSV formatted car data for current driver   |
| msgtype     | Message type from the game for optional routing   |

## Configuration Examples
### Simple
To be added....
