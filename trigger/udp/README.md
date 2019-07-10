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
      "type": "string"
    },
    {
      "name": "multicastGroup",
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
| multicastGroup    | listen group for Mukticast messages |

## Ouputs
| Output   | Description    |
|:---------|:---------------|
| payload    | The raw data from the message |

## Handlers
| Setting   | Description    |
|:----------|:---------------|
| N/A       | awaiting better understanding  |


## Example Configuration

Triggers are configured via the triggers.json of your application. The following is and example configuration of the UDP Trigger.

### Read UDP Data 
Configure the Trigger to capture all data on a given port 
```json
{
  "name": "udp",
  "settings": {
		"port": "20777",
		"multicastGroup": ""
  },
  "handlers": [
    {
      "action": {
		  "id" : "local://testFlow2"
	  },
      "settings": {
        "handler_setting": "xxx"
      }
    }
  ]
}}
```
Configure the Trigger to capture data for a multicast group/port
```json
{
  "name": "udp",
  "settings": {
		"port": "22600",
		"multicastGroup": "224.192.32.19"
  },
  "handlers": [
    {
      "action": {
		  "id" : "local://testFlow2"
	  },
      "settings": {
        "handler_setting": "xxx"
      }
    }
  ]
}}
```