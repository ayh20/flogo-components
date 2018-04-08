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

For an up to date list of valid IP's go to https://support.mashery.com/docs/read/proxy_information/Security_Options or https://support.mashery.com/files/TIBCOMasheryIPs.txt

The following IP list is valid of 14th April 2018

Set the whitelist to the following value to whitelist ALL the Mashery SaaS POPs

"64.94.14.0/27,64.94.228.128/28,216.52.39.0/24,216.52.244.96/27,216.133.249.0/24,23.23.79.128/25,107.22.159.192/28,54.82.131.0/25,75.101.137.168,75.101.142.168,75.101.146.168,75.101.141.43,75.101.129.141,174.129.251.74,174.129.251.80,50.18.151.192/28,50.112.119.192/28,54.193.255.0/25,204.236.130.149,204.236.130.201,204.236.130.207,176.34.239.192/28,54.247.111.192/26,54.93.255.128/27,54.252.79.192/27,54.251.88.0/27,18.231.105.96/28,69.71.111.140,69.71.111.141,207.126.59.91,207.126.59.94,165.254.103.205,165.254.103.203,70.34.228.92,70.34.228.93,4.53.108.203,4.53.108.205,208.72.116.130,208.72.116.131,38.104.3.42,38.122.138.22,75.149.229.162,173.205.4.198,74.202.23.214,75.149.229.118,201.131.127.138,63.237.255.222,209.249.94.34,205.204.93.122,75.149.229.82,152.179.93.94,129.250.199.94,129.250.199.90,75.149.228.230,200.85.152.87,200.85.152.89,200.155.158.42,200.155.158.43,187.45.223.91,187.45.223.93,165.254.103.205,165.254.103.203,186.250.242.26,200.55.243.124,200.55.243.125,201.216.249.66,187.45.179.28,190.117.62.122,185.31.158.244,177.52.180.125,190.112.220.162,200.6.122.211,191.235.90.217,213.130.49.203,213.130.49.205,213.198.94.38,213.198.94.39,212.72.53.203,212.72.53.205,87.236.193.132,87.236.193.137,93.94.105.60,93.94.105.75,185.10.229.160/28,62.103.152.167,195.93.242.2,195.93.242.3,93.17.191.204,93.115.86.209,185.212.169.67,185.206.224.89,37.29.2.10,52.57.11.200,52.57.155.186,52.56.71.32,93.189.33.16,93.189.33.18,159.122.189.182/31,94.180.111.234,159.8.89.144/28,94.72.18.106,212.69.167.117,212.69.167.122,103.19.90.28,103.19.90.29,103.15.105.253,103.15.105.254,103.248.191.19,123.100.230.144,123.100.230.146,123.100.230.148,123.100.230.150,110.50.254.174,110.50.254.177,210.48.32.17,210.48.32.18,122.155.166.196,122.155.166.197,123.103.13.84,14.140.39.202,86.96.201.100,86.96.201.103,23.99.122.185,213.74.78.220,213.74.76.132,202.76.226.4,202.76.226.5,83.222.220.129,52.66.74.180,211.175.216.2,140.206.200.66,140.206.200.68,117.18.236.184,54.252.112.123,175.41.253.70"

To get the Caller IP you will need to add the X-Forwarded-For header in the RecieveHTTPMessage activity. To do this go to the Output Settings and scroll down to the Headers section. Add a new string header with the name X-Forwarded-For.

Then in the BlackWhiteList Activity assign the ipaddress field from $TriggerData.headers.X-Forwarded-For
