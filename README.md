### How to make strict Enum in Go?

To harden Enum with,
* compile time block of accidental arithmetics
* compile time block of implicit cast of untyped constants
* compile time block of all operators except `==` and `!=`
* compile time block of creaing new values
* **zero overhead**

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

There are still uncovered cases, albeit they are very unlikely and easy to spot,
- outside of package can swap enum values
- (if not separate pacakge) inside of package can override values

### Why not `string`?

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
 * can be up to `4x` ~ `5x` slower


### Why not `uint` and `iota`?

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
go test -bench=. -count=3 ./color > doc/struct
go test -bench=. -count=3 ./color-int > doc/int
go test -bench=. -count=3 ./color-string > doc/string
benchstat -split="XYZ" doc/struct doc/int doc/string
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
* https://godbolt.org

### Assembly

Overall, `int` and `struct { int }` versions are same and very efficient.
Version with `string`, as expected, has much more code and jumps, presumably for string comparison logic with shortcuts.

`int`

```s
main_callOne_pc101:
        CALL    runtime.panicdivide(SB)
        XCHGL   AX, AX
        TEXT    main.callOneApple(SB), NOSPLIT|ABIInternal, $0-48
        MOVQ    BX, main.a+16(FP)
        FUNCDATA        $0, gclocals·IuErl7MOXaHVn7EZYWzfFA==(SB)
        FUNCDATA        $1, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
        FUNCDATA        $5, main.callOneApple.arginfo1(SB)
        FUNCDATA        $6, main.callOneApple.argliveinfo(SB)
        PCDATA  $3, $1
        CMPB    SIB, R8B
        JNE     main_callOneApple_pc23
        XORL    AX, AX
        XORL    BX, BX
        MOVQ    BX, CX
        MOVL    $3, DI
        RET
```

`struct`

```s
main_callOne_pc101:
        CALL    runtime.panicdivide(SB)
        XCHGL   AX, AX
        TEXT    main.callOneApple(SB), NOSPLIT|ABIInternal, $0-48
        MOVQ    BX, main.a+16(FP)
        FUNCDATA        $0, gclocals·IuErl7MOXaHVn7EZYWzfFA==(SB)
        FUNCDATA        $1, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
        FUNCDATA        $5, main.callOneApple.arginfo1(SB)
        FUNCDATA        $6, main.callOneApple.argliveinfo(SB)
        PCDATA  $3, $1
        CMPB    SIB, R8B
        JNE     main_callOneApple_pc25
        MOVBLZX main.Purple(SB), DI
        XORL    AX, AX
        XORL    BX, BX
        MOVQ    BX, CX
        RET
```

`string`
```s
main_callOneApple_pc0:
        TEXT    main.callOneApple(SB), ABIInternal, $72-64
        CMPQ    SP, 16(R14)
        PCDATA  $0, $-2
        JLS     main_callOneApple_pc234
        PCDATA  $0, $-1
        SUBQ    $72, SP
        MOVQ    BP, 64(SP)
        LEAQ    64(SP), BP
        MOVQ    R9, main.c+128(FP)
        FUNCDATA        $0, gclocals·J+YAdREO0hCD8EYeU6UDCw==(SB)
        FUNCDATA        $1, gclocals·yROwgZmxcEjQO7qZUR29ZQ==(SB)
        FUNCDATA        $5, main.callOneApple.arginfo1(SB)
        FUNCDATA        $6, main.callOneApple.argliveinfo(SB)
        PCDATA  $3, $1
        MOVQ    BX, main.a+88(SP)
        MOVQ    CX, main.a+96(SP)
        MOVQ    DI, main.a+104(SP)
        MOVQ    SI, main.a+112(SP)
        MOVQ    R8, main.a+120(SP)
        MOVUPS  X15, main.~r0+24(SP)
        MOVUPS  X15, main.~r0+32(SP)
        MOVUPS  X15, main.~r0+48(SP)
        MOVQ    main.a+120(SP), CX
        MOVQ    main.a+112(SP), AX
        CMPQ    R10, CX
        JNE     main_callOneApple_pc105
        MOVQ    R9, BX
        PCDATA  $1, $1
        NOP
        CALL    runtime.memequal(SB)
        TESTB   AL, AL
        JNE     main_callOneApple_pc170
main_callOneApple_pc105:
        MOVQ    main.a+88(SP), DX
        MOVQ    DX, main.~r0+24(SP)
        MOVUPS  main.a+96(SP), X0
        MOVUPS  X0, main.~r0+32(SP)
        MOVUPS  main.a+112(SP), X0
        MOVUPS  X0, main.~r0+48(SP)
        MOVQ    main.~r0+24(SP), AX
        MOVQ    main.~r0+32(SP), BX
        MOVQ    main.~r0+40(SP), CX
        MOVQ    main.~r0+48(SP), DI
        MOVQ    main.~r0+56(SP), SI
        MOVQ    64(SP), BP
        ADDQ    $72, SP
        RET
main_callOneApple_pc170:
        MOVUPS  X15, main.~r0+24(SP)
        MOVUPS  X15, main.~r0+32(SP)
        MOVUPS  X15, main.~r0+48(SP)
        LEAQ    go.string."purple"(SB), DI
        MOVQ    DI, main.~r0+48(SP)
        MOVQ    $6, main.~r0+56(SP)
        MOVQ    main.~r0+24(SP), AX
        XORL    BX, BX
        MOVQ    BX, CX
        MOVL    $6, SI
        MOVQ    64(SP), BP
        ADDQ    $72, SP
        RET
main_callOneApple_pc234:
        NOP
        PCDATA  $1, $-1
        PCDATA  $0, $-2
        MOVQ    AX, 8(SP)
        MOVQ    BX, 16(SP)
        MOVQ    CX, 24(SP)
        MOVQ    DI, 32(SP)
        MOVQ    SI, 40(SP)
        MOVQ    R8, 48(SP)
        MOVQ    R9, 56(SP)
        MOVQ    R10, 64(SP)
        CALL    runtime.morestack_noctxt(SB)
        MOVQ    8(SP), AX
        MOVQ    16(SP), BX
        MOVQ    24(SP), CX
        MOVQ    32(SP), DI
        MOVQ    40(SP), SI
        MOVQ    48(SP), R8
        MOVQ    56(SP), R9
        MOVQ    64(SP), R10
        PCDATA  $0, $-1
        NOP
        JMP     main_callOneApple_pc0
main_main_pc0:
        TEXT    main.main(SB), ABIInternal, $72-0
        CMPQ    SP, 16(R14)
        PCDATA  $0, $-2
        JLS     main_main_pc88
        PCDATA  $0, $-1
        SUBQ    $72, SP
        MOVQ    BP, 64(SP)
        LEAQ    64(SP), BP
        FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        MOVQ    main..stmp_1(SB), BX
        MOVQ    main..stmp_1+8(SB), CX
        MOVQ    main..stmp_1+16(SB), DI
        MOVQ    main..stmp_1+24(SB), SI
        MOVQ    main..stmp_1+32(SB), R8
        MOVL    $10, AX
        LEAQ    go.string."blue"(SB), R9
        MOVL    $4, R10
        PCDATA  $1, $0
        CALL    main.callOneApple(SB)
        MOVQ    64(SP), BP
        ADDQ    $72, SP
        RET
```
