# Amazon Get Secret

Simple function that retrieves a named secret from AWS Secret manager.
This version ony supports string secrets but can be extended to do binary secrets properly too (returned as an encoded string in this version).

Running this in a AWS Cli configured environment ie Win/Linux where AWS cli is stetup and credentials file is configured means that you DON'T need to supply key/secret as it's picked up. Running under TCI means you need to supply key/secret.

## Installation

```bash
flogo install github.com/ayh20/flogo-components/activity/awsgetsecret
```

Link for flogo web:

```
https://github.com/ayh20/flogo-components/activity/awsgetsecret
```

## Schema

Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "awsAccessKeyID",
      "type": "string",
      "required": true
    },
    {
      "name": "awsSecretAccessKey",
      "type": "string",
      "required": true
    },
    {
      "name": "awsSessionToken",
      "type": "string",
      "required": true
    },
    {
      "name": "assumeRole",
      "type": "boolean",
      "required": true
    },
    {
      "name": "roleARN",
      "type": "string",
      "required": false
    },
    {
      "name": "roleSessionName",
      "type": "string",
      "required": false
    },
    {
      "name": "awsRegion",
      "type": "string",
      "required": true
    },
    {
      "name": "secretARN",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "secret",
      "type": "string"
    }
  ]
}
```

## Inputs

| Input              | Description                                            |
| :----------------- | :----------------------------------------------------- |
| awsAccessKeyID     | Your AWS CLI Access Key                                |
| awsSecretAccessKey | Your AWS CLI Secret Key                                |
| awsSessionToken    | Your AWS CLI Session Token                             |
| assumeRole         | Flag to say if you need to switch role to run the task |
| roleARN            | The Role ARN that is to be switched too                |
| roleSessionName    | A name for the active session                          |
| awsRegion          | The AWS region your S3 bucket is in                    |
| secretARN          | The AWS arn for the secret to be retrieved             |

## Ouputs

| Output | Description                                    |
| :----- | :--------------------------------------------- |
| secret | The result will contain the json secret string |
