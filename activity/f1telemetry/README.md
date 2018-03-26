# f1telemetry
This activity Decodes telemetry data from Codemasters F1-2017, and places it in to a "CSV" "record". 
This is designed as a demonstration for integrating F1 2017 .. Flogo .... TIBCO Streambase and TIBCO Spotfire/Liveview

The Flogo app needs to read the data in from UDP and feed the byte stream into this component for decoding ... it's then passed to Streambase via
some form of messaging

The UDP data is a packed little endian C struct, which is transformed to a Go struct by https://restruct.io/. This uses struct tags to decode the raw data
correctly to the Go struct.

## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/f1telemetry
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
| data        | CSV formatted  race data   |

## Configuration Examples
### Simple
To be added....
