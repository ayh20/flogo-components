# compare
This activity allows you to create or validate a JWT token


## Installation

Navigate to your Flogo app directory and enter the following command

```bash
flogo install github.com/ayh20/flogo-components/activity/jwt
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "token",
      "type": "string"
    },
    {
      "name": "header",
      "type": "string"
    },
    {
      "name": "payload",
      "type": "string"
    },
    {
      "name": "secret",
      "type": "string"
    },
    {
      "name": "mode",
      "required": true,
      "type": "string",
      "allowed" : ["Verify", "Sign", "Decrypt"]
    },
    {
      "name": "algorithm",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "token",
      "type": "string"
    },
     {
      "name": "valid",
      "type": "bool"
    }
  ]
}
```
## Settings Mode
| Setting     | Description       |
|:------------|:------------------|
| mode        | Verify or Sign    |


## Settings Verify
| Setting     | Description                                    |
|:------------|:-----------------------------------------------|
| algorithm   | The algorithm name ie HS256, ES512, RS256 etc  |
| secret      | The encryption key (HS*) or public key         |
| token       | The token to be validated                      |


## Settings Sign
| Setting     | Description                                    |
|:------------|:-----------------------------------------------|
| algorithm   | The algorithm name ie HS256, ES512, RS256 etc  |
| secret      | The encryption key (HS*) or public key         |
| header      | The json header (used to validate the request) |
| payload     | The claims string                              |


## Outputs
| Output      | Description                             |
|:------------|:----------------------------------------|
| valid       | Bool result for Verify operation        |
| token       | The token string for Sign operations    |

## Configuration Examples
### Simple
Configure a task to verify or create a JWT:


