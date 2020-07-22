# Schedule
This trigger is based on the time trigger and provides the same functionality at that. In addition it allows a task to be scheduled at a specific time and day once, or repeating weekly, or repeating each day.

## Installation

```bash
flogo install github.com/ayh20/flogo-components/trigger/schedule
```
Link for flogo web:
```
https://github.com/ayh20/flogo-components/trigger/schedule
```


## Configuration

### Handler Settings:
| Name           | Type   | Description
|:---            | :---   | :---     
| startDelay     | string | The start delay (ex. 1m, 1h, etc.), immediate if not specified
| repeatInterval | string | The repeat interval (ex. 1m, 1h, etc.), doesn't repeat if not specified
| startDay       | string | The day(s) the process should run (ex. Monday, Tuesday, Everyday.)
| startTime      | string | The time the process will be run


## Example Configurations

Triggers are configured via the triggers.json of your application. The following are some example configuration of the Timer Trigger.

### One off scheduled at a given time
Configure the Trigger to run at Sunday 1am.

```json
{
  "triggers": [
    {
      "id": "flogo-schedule",
      "ref": "github.com/project-flogo/contrib/trigger/schedule",
      "handlers": [
        {
          "settings":{
            "startDay": "Sunday",
            "startTime": "01:00"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Repeating task scheduled at a given time
Configure the Trigger to run a flow every Monday at 2pm.

```json
{
  "triggers": [
    {
      "id": "flogo-schedule",
      "ref": "github.com/project-flogo/contrib/trigger/schedule",
      "handlers": [
        {
          "settings":{
            "startDay": "Monday",
            "startTime": "14:00",
            "repeatInterval" : "1w"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Only once and immediate
Configure the Trigger to run a flow immediately

```json
{
  "triggers": [
    {
      "id": "flogo-schedule",
      "ref": "github.com/project-flogo/contrib/trigger/schedule",
      "handlers": [
        {
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Only once with a delay
Configure the Trigger to run a flow once with a delay of one minute.  "startDelay" settings format = "[hours]h[minutes]m[seconds]s"

```json
{
  "triggers": [
    {
      "id": "flogo-schedule",
      "ref": "github.com/project-flogo/contrib/trigger/schedule",
      "handlers": [
        {
          "settings": {
            "startDelay": "1m"
          },
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Repeating
Configure the Trigger to run a flow repeating every 10 minutes. "repeatInterval" settings format = "[hours]h[minutes]m[seconds]s"

```json
{
  "triggers": [
    {
      "id": "flogo-schedule",
      "ref": "github.com/project-flogo/contrib/trigger/schedule",
      "handlers": [
        {
          "settings": {
            "repeatInterval": "10m"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```

### Repeating with start delay
Configure the Trigger to run a flow every minute, with a delayed start of 10 minutes and 30 seconds.

```json
{
  "triggers": [
    {
      "id": "flogo-schedule",
      "ref": "github.com/project-flogo/contrib/trigger/schedule",
      "handlers": [
        {
          "settings": {
            "repeatInterval": "1m",
            "startDelay": "10m30s"
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:myflow"
            }
          }
        }
      ]
    }
  ]
}
```
