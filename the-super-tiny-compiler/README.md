# the-super-tiny-compiler

## Usage

Input: `(add 2 (subtract 10 5))`

Output: `add(2, subtract(10, 5));`

## Tests

```Bash
go test
go test -v
go test -cover
go test -bench .
```
