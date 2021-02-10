# Utils Functions for RTB project

This package adds a set of functions that can be used in Flogo versions >= 0.9.0.

Installing them in the Web UI means that they will show up as functions that can be used in the mapper
These will get added to the core project at some point

## Installation

```CLI
flogo install github.com/ayh20/flogo-components/functions/rtbutils
```

Link for flogo web:

```
https://github.com/ayh20/flogo-components/functions/rtbutils
```

## Functions

| Name         | Decription                           | Sample                                                           |
| :----------- | :----------------------------------- | :--------------------------------------------------------------- |
| urlencode    | Make the passed string url safe      | rtbutils.urlencode(\"Hello World\")                              |
| pathencode   | Make the path string url safe        | rtbutils.pathencode(\"path with?reserved+characters\")           |
| stringtodate | convert a date string to a date type | rtbutils.stringtodate(\"10 February 2021\", \"02 January 2006\") |
