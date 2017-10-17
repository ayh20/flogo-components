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
### Mashery
Set the whitelist to the following value to whitelist all the Mashery SaaS POPs
"64.94.14.0/27,64.94.228.128/28,216.52.39.0/24,216.52.244.96/27,216.133.249.0/24,23.23.79.128/25,107.22.159.192/28,54.82.131.0/25 ,75.101.137.168,75.101.142.168,75.101.146.168,75.101.141.43,75.101.129.141,174.129.251.74,174.129.251.80,50.18.151.192/28,50.112.119.192/28,54.193.255.0/25 ,204.236.130.149 ,204.236.130.201,204.236.130.207,176.34.239.192/28,54.247.111.192/26 ,54.93.255.128/27 ,54.252.79.192/27,54.251.88.0/27,69.71.111.140,69.71.111.141,207.126.59.91,207.126.59.94,165.254.103.205,165.254.103.203,70.34.228.92,70.34.228.93,4.53.108.203,4.53.108.205,208.72.116.130,208.72.116.131,200.85.152.87,200.85.152.89,200.155.158.42,200.155.158.43,187.45.223.91,187.45.223.93,165.254.103.205,165.254.103.203,213.130.49.203,213.130.49.205,213.198.94.38,213.198.94.39,212.72.53.203,212.72.53.205,87.236.193.132,87.236.193.137,93.94.105.60,93.94.105.75,103.19.90.28,103.19.90.29,103.15.105.253,103.15.105.254,103.248.191.19,123.100.230.144,123.100.230.146,123.100.230.148,123.100.230.150,110.50.254.174,110.50.254.177"

To get the Caller IP you will need to add the X-Forwarded-For header in the RecieveHTTPMessage activity. To do this go to the Output Settings and scroll down to the Headers section. Add a new string header with the name X-Forwarded-For.

Then in the BlackWhiteList Activity assign the ipaddress field from $TriggerData.headers.X-Forwarded-For
