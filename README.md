# go-enum-untyped-block

Basic enum
```go
type Color uint

const (
  UndefinedColor Color = iota
  Red
  Green
  Blue
)
```

## Issues

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
const BadColor Color = 1000
```

### Expressions with operators and untyped constants

```go
const x = 5

func main() {
  var e Color = Green + 5
}
```

### Functions returning untyped constants

```go
func badfunc() Color {
  return 1000
}
```

## References

* https://go.dev/ref/spec
