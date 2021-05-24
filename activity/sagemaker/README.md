# Amazon Invoke Sagemaker

Simple Sagemaker function invoke

## Installation

```bash
flogo install github.com/ayh20/flogo-components/activity/sagemaker
```

Link for flogo web:

```
https://github.com/ayh20/flogo-components/activity/sagemaker
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
      "name": "endpointname",
      "type": "string",
      "required": true
    },
    {
      "name": "body",
      "type": "string",
      "required": true
    },
    {
      "name": "contenttype",
      "type": "string",
      "required": true
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

## Inputs

| Input              | Description                                                       |
| :----------------- | :---------------------------------------------------------------- |
| awsAccessKeyID     | Your AWS Access Key                                               |
| awsSecretAccessKey | Your AWS Secret Key                                               |
| assumeRole         | Flag to say if you need to switch role to run the task            |
| roleARN            | The Role ARN that is to be switched too                           |
| roleSessionName    | A name for the active session                                     |
| awsRegion          | The AWS region your S3 bucket is in                               |
| endpointname       | The name of the Sagemaker enpoint                                 |
| body               | The data to be passed to the Sagemaker deployed model             |
| contenttype        | The format of the body data ie "text/csv" or ""application/json"" |

## Ouputs

| Output | Description                                                          |
| :----- | :------------------------------------------------------------------- |
| result | The result will contain the value returned from the Sagemaker invoke |
