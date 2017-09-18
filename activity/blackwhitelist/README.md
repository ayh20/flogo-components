# Black White List Utility
This activity allows you to validate a IP address against a Whitelist and/or a Blacklist

IP addresses in the list can be in two formats either a absolute IP or an IP range in CIDR format ie 192.168.1.0/24. Each ip needs to be comma separated

ie:

123.221.142.245,123.257.221.0/24,192.168.1.10


## Installation

Navigate to the Flogo app directory [Essential !] and issue the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/blackwhitelist
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
      {
        "name": "whitelist",
        "required": true,
        "type": "string"
      },
      {
        "name": "blacklist",
        "required": true,
        "type": "string"
      },
      {
        "name": "ipaddress",
        "required": true,
        "type": "string"
      }
    ],
    "outputs": [
      {
        "name": "isOK",
        "type": "boolean"
      }
    ]
  }
```

## Settings
| Setting     | Description               |
|:------------|:--------------------------|
| whitelist   | The list of allowed IPs   |
| blacklist   | The list of disallowed    |
| ipaddress   | IP address to be compared |

## Outputs
| Output      | Description                                          |
|:------------|:-----------------------------------------------------|
| isOK        | Bool result signalling that the address is allowed   |

## Configuration Examples
### Simple
Test valid IP...
