package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func functionRetrunsConst() somecolor.Color {
	return 10
}

func functionRetrunsConstMultiple() (a, b somecolor.Color) {
	return 10, somecolor.Blue
}

func functionRetrunsConstMultipleWithInt() (a somecolor.Color, b uint) {
	return 10, 11
}
