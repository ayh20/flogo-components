# Read a File

Read a file and get the content.

This version is an update on github.com/retgits/flogo-components/activity/readfile which was using the depreated package ioutil

## Installation

```bash
flogo install github.com/ayh20/flogo-components/activity/readfile
```
Link for flogo web:
```
https://github.com/ayh20/flogo-components/activity/readfile
```

## Schema
Inputs and Outputs:

```json
{
"inputs": [
        {
            "name": "filename",
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
| Input    | Description                                                                 |
|:---------|:----------------------------------------------------------------------------|
| filename | The name of the file you want to read (like `data.txt` or `./tmp/data.txt`) |

## Ouputs
| Output      | Description             |
|:------------|:------------------------|
| result      | The content of the file |