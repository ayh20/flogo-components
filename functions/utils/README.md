# Utils Functions
This package adds a set of functions that can be used in Flogo versions >= 0.9.0.

Installing them in the Web UI means that they will show up as functions that can be used in the mapper


## Installation

```CLI
flogo install github.com/ayh20/flogo-components/functions/utils
```
Link for flogo web:
```
https://github.com/ayh20/flogo-components/functions/utils
```

## Functions

| Name         | Decription             | Sample                                                |
|:-------------|:-----------------------|:------------------------------------------------------|
| matchregex   | Match input against regular expression | utils.matchregex(\"p([a-z]+)ch\", \"peach\")" |
| replaceregex | Replace data in a string based on a regular expression match  | utils.replaceregex(\"p([a-z]+)ches\", \" I hate peaches !\", \"apples\")" |
| contains     | Contains reports whether substr is within s. |  utils.contains(\"seafood\", \"foo\")" |
| containsany  | ContainsAny reports whether any Unicode code points in chars are within s. |  utils.containsany(\"failure\", \"u & i\")" |
| count        | Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s. | utils.count(\"cheese\", \"e\")" |
| index        | Index returns the index of the first instance of substr in s, or -1 if substr is not present in s. | utils.index(\"cheese\", \"e\")" |
| indexany     | IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s. | utils.indexany(\"chicken\", \"aeiouy\")" |
| lastindex    | LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s. | utils.lastindex(\"go gopher\", \"go\")" |
| repeat       | Repeat returns a new string consisting of count copies of the string s. | utils.repeat(\"na\", 3)" |
| replace      | Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.  If n < 0, there is no limit on the number of replacements | utils.replace(\"oink oink oink\", \"k\", \"ky\", 2)" |
| replaceall   | ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new. | utils.replaceall(\"oink oink oink\", \"oink\", \"moo\")" |
| tolower      | ToLower returns a copy of the string s with all Unicode letters mapped to their lower case. | utils.tolower(\"Hello World\")" |
| toupper      | ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case. | utils.toupper(\"Hello World\")" |
| trim         | Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed. | utils.trim(\"¡¡¡Hello, Gophers!!!\", \"!¡\")" |
| trimleft     | TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed. | utils.trimleft(\"¡¡¡Hello, Gophers!!!\", \"!¡\")" |
| trimright    | TrimRight returns a slice of the string s with all leading Unicode code points contained in cutset removed. | utils.trimright(\"¡¡¡Hello, Gophers!!!\", \"!¡\")" |
| trimprefix   | TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged. | utils.trimprefix(\"¡¡¡Hello, Gophers!!!\", \"¡¡¡Hello\")" |
| trimsuffix   | TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged. | utils.trimsuffix(\"¡¡¡Hello, Gophers!!!\", \"Gophers!!!\")" |
| split        | Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators. | utils.split(\"a,b,c\", \",\")" |