package example_test

import (
	"fmt"

	somecolor "github.com/nikolaydubina/go-enum-example/color"
)

func ExampleSwapExportedVar() {
	somecolor.Purple = somecolor.Red
	fmt.Println(somecolor.Purple)
	// Output: red
}

func ExampleImplicitAssignmentToNonExportedStruct() {
	// does not compile
	//somecolor.Blue = somecolor.Color{5}
}

func ExampleUntypedConstantsAndOperatorsNotAllowed() {
	// does not compile
	// x := somecolor.Blue + 5
}

func ExampleMap() {
	m := map[somecolor.Color]string{}
	m[somecolor.Blue] = "blue"
	m[somecolor.Green] = "more green"
	m[somecolor.Red] = "more more red"
	fmt.Println(m[somecolor.Red])
	// Output: more more red
}

func ExampleComparable() {
	a := somecolor.Blue
	b := somecolor.Green
	fmt.Println(a == b)
	// Output: false
}
