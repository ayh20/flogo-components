# UDP
This trigger provides your flogo application a stream of UDP data from the specificed port

## Installation

```bash
flogo install github.com/ayh20/flogo-components/trigger/udp
```
Link for flogo web:
```
https://github.com/ayh20/flogo-components/trigger/udp
```

## Schema
Outputs and Endpoint:

```json
{
"settings":[
    {
      "name": "port",
      "type": "integer"
    },
    {
      "name": "multicast_group",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "payload",
      "type": "string"
    }
  ],
  "handler": {
    "settings": [
      {
        "name": "handler_setting",
        "type": "string"
      }
    ]
}
```
## Settings
| Setting   | Description    |
|:----------|:---------------|
| port      | port to listen on |
| multicast_group    | listen group for Mukticast messages |

## Ouputs
| Output   | Description    |
|:---------|:---------------|
| payload    | The raw data from the message |

## Handlers
| Setting   | Description    |
|:----------|:---------------|
| N/A       | awaiting better understanding  |


## Example Configuration

Triggers are configured via the triggers.json of your application. The following is and example configuration of the GPIO Trigger.

### Only once and immediate
Configure the Trigger to run a flow when pin 7 becomes high, check every 0.5 seconds
```json
{
  "name": "udp",
  "settings": {
		"port": 20777,
		"multicast_group": ""
  },
  "handlers": [
    {
      "actionId": "local://testFlow2",
      "settings": {
        "handler_setting": "xxx"
      }
    }
  ]
}}
```
