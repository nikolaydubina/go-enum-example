### How to make strict Enum in Go?

If you want to harden your Enum with
* compile time block of accidental arithmetics
* compile time block of implicit cast of untyped constants
* compile time block of all operators except `==` and `!=`

Simply
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
- inside of package can override values
- outside of package can swap enum values

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
