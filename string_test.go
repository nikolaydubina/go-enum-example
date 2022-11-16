package example

import (
	"fmt"
	"reflect"
	"unsafe"
)

type A string

const (
	red A = "red"
)

func x(s string) {
	hdB := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("x: 0x%x\n", hdB.Data)
}

func ExampleString() {
	a := "Hello, world!"
	b := a

	hdA := (*reflect.StringHeader)(unsafe.Pointer(&a))
	fmt.Printf("String A: 0x%x\n", hdA.Data)

	x(a)

	hdB := (*reflect.StringHeader)(unsafe.Pointer(&b))
	fmt.Printf("String B: 0x%x\n", hdB.Data)

	x(b)

	b += " and some more"

	hdB = (*reflect.StringHeader)(unsafe.Pointer(&b))
	fmt.Printf("String B after append: 0x%x\n", hdB.Data)

	x(b)
}
