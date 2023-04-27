<!--
title: Gremlin
weight: 4616
-->

# Gremlin Activity

This activity invokes gremlin commands to the connected GraphDB.

### Flogo CLI

```bash
flogo install github.com/ayh20/flogo-components/activity/gremlin
```

## Configuration

### Settings:

| Name       | Type   | Description                                                                                                                                                                                           |
| :--------- | :----- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| gremlinUrls | string | The brokers of the Kafka cluster to connect to - **_REQUIRED_**                                                                                                                                                                                                               |
| user       | string | If connecting to a SASL enabled port, the user id to use for authentication                                                                                                                           |
| password   | string | If connecting to a SASL enabled port, the password to use for authentication                                                                                                                          |

### Input:

| Name    | Type   | Description               |
| :------ | :----- | :------------------------ |
| query  | string    | The query to run |

### Output:

| Name      | Type  | Description                                            |
| :-------- | :---- | :----------------------------------------------------- |
| result | object | the result object from gremlin |

