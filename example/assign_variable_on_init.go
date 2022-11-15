package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func assignVariableOnInit() somecolor.Color {
	var x somecolor.Color = 1000
	return x
}

func assignConstant() somecolor.Color {
	var x somecolor.Color
	x = 1000
	return x
}
