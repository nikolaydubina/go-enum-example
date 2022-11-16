### How to make strict Enum in Go?

To harden Enum with,
* compile time block of accidental arithmetics
* compile time block of implicit cast of untyped constants
* compile time block of all operators except `==` and `!=`
* compile time block of creaing new values

With zero overhead.

Simply,
* wrap into struct
* do not export field
* make separate package

```go
package color

type Color struct{ c uint }

var (
	Undefined = Color{}
	Red       = Color{1}
	Green     = Color{2}
	Blue      = Color{3}
)
```

There are still uncovered cases, albeit they are very unlikely and easy to spot
- outside of package can swap enum values
- (if not separate pacakge) inside of package can override values

### Why not use `string`?

```go
type Color string

const(
	Red Color = "red"
	Blue Color = "blue"
)
```

Passing strings in Go is done by referneces.
String content is not copied in function calls nor assignments.

Problems
 * string comparison can take up to `O(N)`, which is very common operation on Enums
 * strings in struct fields result to more indirection as compared to `int` or `struct`
 * if it is not wrapped into struct, it will also leak comparison and concatenation operators and casts
 * practically, it is `4x` ~ `5x` slower


### Why not `uint` and `iota`?

Following two alternatives

```go
type Color uint

const(
	Red Color = iota
	Blue
)
```

This is very common way.
It is also efficient.

Problems,
 * leaks all arithmetic operators
 * leaks cast
 * easily mixed with untyped constants in very many places: function returns, channles, switch, casts, assignemnts, etc.

### Benchmarks

All versions have zero mallocs and memory transfer.

```bash
go test -bench=. -count=3 ./color > struct                
go test -bench=. -count=3 ./color-int > int   
go test -bench=. -count=3 ./color-string > string
benchstat -split="XYZ" struct int string 
```

```
name \ time/op                          struct       int          string
EnumPassFunction/call_one-10            2.12ns ± 0%  2.14ns ± 0%   2.52ns ± 0%
EnumPassFunction/call_one_apple-10      2.37ns ± 0%  2.44ns ± 0%   8.62ns ± 0%
EnumPassFunction_Big/call_one-10        2.19ns ± 0%  2.14ns ± 0%   6.59ns ± 0%
EnumPassFunction_Big/call_one_apple-10  2.39ns ± 0%  2.38ns ± 0%  11.19ns ± 0%
```

### What is Enum?

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

### Related Tools

* https://github.com/nishanths/exhaustive - performs exhaustive checks in switch; does not check for type conversions; widely used; however does not support struct enums
* https://github.com/loov/enumcheck - focuses on performing exhastive checks; also does constant expressions validations; as of 2022-11-15, work in progress after 3 years

### References

* https://en.wikipedia.org/wiki/Enumerated_type
* https://go.dev/ref/spec
* https://www.w3schools.com/java/java_enums.asp
* https://en.cppreference.com/w/cpp/language/enum
* https://en.cppreference.com/w/c/language/enum
* https://docs.python.org/3/library/enum.html
