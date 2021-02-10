# Parse CSV

This activity can be used to parse comma separated text strings into arrays

## Installation

### Flogo CLI

```bash
flogo install github.com/mellistibco/flogo-activities/activities/parsecsv
```

## Schema

Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "fieldNames",
      "type": "array",
      "required": true
    },
    {
      "name": "csv",
      "type": "string",
      "required": false
    },
    {
      "name": "file",
      "type": "string",
      "required": false
    }
  ],
  "output": [
    {
      "name": "output",
      "type": "array"
    }
  ]
}
```

## Settings

| Setting    | Required | Description                                                                                         |
| :--------- | :------- | :-------------------------------------------------------------------------------------------------- |
| fieldNames | True     | The expected fields from the csv (the headers, will be used for the field name in the JSON object.) |
| csv        | False    | The csv text (field1,field2,field3)                                                                 |
| file       | False    | The optional location to a CSV file on disk                                                         |

## Example

The below example will parse the supplied text.

```json
{
  "id": "parsecsv_1",
  "name": "parsecsv",
  "description": "Parse CSV into a Flogo Array",
  "activity": {
    "ref": "github.com/mellistibco/flogo-activities/activities/parsecsv",
    "input": {
      "fieldNames": ["field1", "field2", "field3"],
      "csv": "data1,data2,data3\ndata11,data22,data33"
    }
  }
}
```

The output of the above sample will be an array of objects:

```json
[
  {
    "field1": "data1",
    "field2": "data3",
    "field3": "data3"
  },
  {
    "field1": "data11",
    "field2": "data22",
    "field3": "data33"
  }
]
```
