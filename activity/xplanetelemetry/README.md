# xplanetelemetry
This activity Decodes telemetry data from X Plane 11, and places it in to a "CSV" "record". 
This is designed as a demonstration for integrating X-Plane .. Flogo edge .... TIBCO Streaming, TIBCO Data Streams and TIBCO Spotfire

The Flogo app needs to read the data in from UDP and feed the byte stream into this component for decoding ... it's then passed to Streambase via some form of messaging (MQTT in my Demo)

The UDP data is a formatted with a header and sentences, each sentence is a 4 byte index (Byte 1 is the value + 3 bytes padding) followed by 8, 4 byte floating point numbers in little Endian format (Float32)

Information about the format of the raw data can be found here: http://www.nuclearprojects.com/xplane/xplaneref.html

## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/xplanetelemetry
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
| Output      | Description                                 |
|:------------|:--------------------------------------------|
| data        | CSV formatted car data for current driver   |
| msgtype     | Message type - Currently unused set to 1    |

## Configuration Examples
### Simple
To be added....
