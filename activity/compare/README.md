# Compare
This activity allows you to compare two values and return a true/false result. Data is passed as a string numeric and operation ( i.e. "=", "<=", ">",.... ).


## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/compare
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "input1",
      "required": true,
      "type": "string"
    },
    {
      "name": "input2",
      "required": true,
      "type": "string"
    },
    {
      "name": "comparemode",
      "required": true,
      "type": "string",
      "allowed" : [">", "<", "=", "==", ">=", "<=", "!=" ]
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "boolean"
    }
  ]
}
```

## Settings
| Setting     | Description       |
|:------------|:------------------|
| input1      | The first value   |
| input2      | The second value  |
| comparemode | Compare operation |

## Outputs
| Output      | Description                             |
|:------------|:----------------------------------------|
| result      | Bool result based on input comparison   |

## Configuration Examples
### Simple
Configure a task to compare two values:
