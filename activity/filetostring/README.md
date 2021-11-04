# Read a file to a string

This activity reads the contents of a file and places it in a single string field.
It removes all CR/LF chars from the string

## Installation

### Flogo CLI

```bash
flogo install github.com/ayh20/flogo-components/activity/filetostring
```

## Schema

Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "filename",
      "type": "string",
      "required": false
    }
  ],
  "output": [
    {
      "name": "output",
      "type": "string"
    }
  ]
}
```

## Settings

| Setting | Required | Description |
| :--------- | :------- | :-------------------------------------------------------------------------------------------------- | |
| filename | True | The fully qualified filename |
