# enumcheck

This is `go vet` compatible tool to confirm correctness of strict enum type usage.
It analyses AST to detect issues leading to loss of strictness.
This tool aims to make enums in Go closer to what they are in Java, Rust, Swift and generally match intuition many people have about compiler guarantees about enums.

```go
//enumcheck:def Color
type Color uint

const (
  Undefined Color = iota
  Red
  Green
  Blue
)
```

```bash
enumcheck ./...
```

## Detected Issues

For a full list refer to `/example`.
If you found some case not covered, please raise PR or open an Issue.

#### Assign variable on init

```go
func myfunc() {
  ...
  var b Color = 100
  ...
}
```

#### New constants

```go
package badpacakge

const BadColor colors.Color = 1000
```

#### Cast number

```go
var x = Color(1000)
```

#### Expressions with operators and untyped constants

```go
const x = 5

func main() {
  var e Color = Green + 5
  var b Color = Green - x
}
```

#### Functions returning untyped constants

```go
func badfunc() Color {
  return 1000
}
```

#### Non-Comparison operators

```go
var e Color = Green + Red
var e Color = Green * Red
```

## Usage Notes

This tool needs to be run over all packages.
As of 2022-11-12, I did not figure out how to get full AST that will go to your binary.
Thus, you need to run this tool over all packages you include in your binary.
If you vendor or have monorepo or do not export your enum type or have `internal`, this is easier to do.

### Specifying Definition

File that contains this directive for enum marking will be skipped and type will be selected for checks.

```go
//enumcheck:def Color
type Color uint

const (
  Undefined Color = iota
  Red
  Green
  Blue
)
```

### Skipping Files

File that contains this directive will be skipped.

```go 
//enumcheck:skip
```

## Related Work

* https://github.com/nishanths/exhaustive - performs exhaustive checks in switch; does not check for type conversions; widely used
* https://github.com/loov/enumcheck - focuses on performing exhastive checks; also does constant expressions validations; as of 2022-11-15, work in progress after 3 years

----

## Appendix A: What is Enum?

Enum is a data type consisting of a set of named values.
They are typically constants.
They have comaprison and assignemnt operators.
Underlying represtentation is free up to compiler, but typically is an integer.
Enums are typically prevented from illogical operations such as arithmetic operations.

* `C` - interger; exposed; arithmetics permitted;
* `C#` - integer; not exposed; some arithmetics permitted;
* `C++` - integer; not exposed; not converted; some arithmetics not permitted;
* `Go` - integer; exposed; arithmetics permitted;
* `Python` - integer; arithmetics permitted;
* `Java` - not integer; not converted; arithmetics not permitted; internally integer;
* `Rust` - not integer; not converted; arithmetics not permitted; can be extended to integer
* `Swift` - not integer; not converted; arithmetics not permitted; can be extended to interger

## References

* https://en.wikipedia.org/wiki/Enumerated_type
* https://go.dev/ref/spec
* https://www.w3schools.com/java/java_enums.asp
* https://en.cppreference.com/w/cpp/language/enum
* https://en.cppreference.com/w/c/language/enum
* https://docs.python.org/3/library/enum.html
