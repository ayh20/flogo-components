# f1telemetry

This activity Decodes telemetry data from EA F1®22 , and places it in to a "CSV" "record".
This is designed as a demonstration for integrating F1®22 .. Flogo .... TIBCO Streambase and TIBCO Spotfire/Liveview

The Flogo app needs to read the data in from UDP and feed the byte stream into this component for decoding ... it's then passed to Streambase via some form of messaging (MQTT in my Demo)

The UDP data is a packed little endian C struct, which is transformed to a Go struct by https://github.com/lunixbochs/struc. This uses struct tags to decode the raw data correctly to the Go struct.

Information about the format of the raw data can be found here: https://answers.ea.com/t5/General-Discussion/F1-22-UDP-Specification/td-p/11551274

## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/f1telemetry2022
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
    },
    {
      "name": "source",
      "required": true,
      "type": "string"
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

| Setting | Description                                               |
| :------ | :-------------------------------------------------------- |
| buffer  | The raw UDP data                                          |
| source  | The name of the data source (ip address if not specified) |

## Outputs

| Output  | Description                                     |
| :------ | :---------------------------------------------- |
| data    | CSV formatted car data for current driver       |
| msgtype | Message type from the game for optional routing |

## Configuration Examples

### Simple

To be added....
