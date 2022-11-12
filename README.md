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
