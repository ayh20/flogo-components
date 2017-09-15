# compare
This activity gives you the ability to input a datetime string and add or subtract a number of days,Hours,Minutes or seconds to that date


## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/datemath
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
      {
        "name": "date",
        "required": true,
        "type": "string"
      },
      {
        "name": "amount",
        "required": true,
        "type": "string"
      },
      {
        "name": "unit",
        "required": true,
        "type": "string",
        "allowed" : ["Day", "Hour", "Min", "Sec"]
      },
      {
        "name": "function",
        "required": true,
        "type": "string",
        "allowed" : ["Add", "Subtract"]
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
| Setting     | Description                                         |
|:------------|:----------------------------------------------------|
| date        | the date we are going perform the operation on      |
| amount      | the amount of hours/mins etc to add or subtract     |
| unit        | the unit of operation ("Day", "Hour", "Min", "Sec") |
| function    | the type of operation ("Add", "Subtract")           |

## Outputs
| Output      | Description                             |
|:------------|:----------------------------------------|
| result      | the result as a date                    |

## Configuration Examples
### Simple
Add and subtract from a date:
