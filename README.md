# go-enum-check

Basic enum
```go
type Color uint

const (
  Undefined Color = iota
  Red
  Green
  Blue
)
```

`go vet` analyzer. Checks AST for possible issues.

File that contains directive for enum marking will be skipped.
```go
//enum:Color
type Color uint

const (
  Undefined Color = iota
  Red
  Green
  Blue
)

### Usage Notes

This tool needs to be run over all packages.
As of 2022-11-12, I did not figure out how to get full AST that will go to your binary.
Thus, you need to run this tool over all packages you include in your binary.
If you vendor or have monorepo, this is easier to do.

### TODO!!!!!

what to do when some package uses symbol from other package enum?

what to do with renamed packages?

what to do with collision of package names?

what to do when some package redeclares enum type with new name?

## Issues

Classes of issues:
- create new value
- convert value to enum
- non-comparison operators

### Assign variable on init

```go
func myfunc() {
  ...
  var b Color = 100
  ...
}
```

### New constants

```go
package badpacakge

const BadColor colors.Color = 1000
```

### Cast number

```go
var x = Color(1000)
```

### Expressions with operators and untyped constants

```go
const x = 5

func main() {
  var e Color = Green + 5
  var b Color = Green - x
}
```

### Functions returning untyped constants

```go
func badfunc() Color {
  return 1000
}
```

### Non-comparison operators

```go
var e Color = Green + Red
var e Color = Green * Red
```

## References

* https://go.dev/ref/spec
