package color

import (
	"fmt"
	"reflect"
	"unsafe"
)

type A string

func x(s string) {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%s: 0x%x\n", s, p.Data)
}

func ExampleStringPointers() {
	a := "Hello, world!"
	b := a

	pa := (*reflect.StringHeader)(unsafe.Pointer(&a))
	fmt.Printf("a: 0x%x\n", pa.Data)

	x(a)

	pb := (*reflect.StringHeader)(unsafe.Pointer(&b))
	fmt.Printf("b: 0x%x\n", pb.Data)

	x(b)

	b += " and some more"

	pbMutated := (*reflect.StringHeader)(unsafe.Pointer(&b))
	fmt.Printf("String B after append: 0x%x\n", pbMutated.Data)

	x(b)
}
