# Read a directiory

Read a directory and get the content.

returns alist of files and directories

## Installation

```bash
flogo install github.com/ayh20/flogo-components/activity/readdir
```

Link for flogo web:

```
https://github.com/ayh20/flogo-components/activity/readdir
```

## Schema

Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "dirname",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "object"
    }
  ]
}
```

## Inputs

| Input    | Description                                                           |
| :------- | :-------------------------------------------------------------------- |
| filename | The name of the directory you want to read (like `C:\tmp` or `./tmp`) |

## Ouputs

| Output | Description                                 |
| :----- | :------------------------------------------ |
| result | A json object containing the directory data |
