# helm-indent-modifier
Modifies the indent and nindent to specified increment/decrement

## Usage
```bash
$ go install github.com/aweis89/helm-indent-modifier
$ helm-indent-modifier --help 

Usage of helm-indent-modifier:
  -dec int
        decrement indent/nindent by this value
  -end-line int
        ingnore lines after (default 9223372036854775807)
  -file string
        file to modify
  -inc int
        increment indent/nindent by this value
  -start-line int
        ingnore lines before (default 1)
```
