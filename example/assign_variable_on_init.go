package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func assignVariableOnInit() somecolor.Color {
	var x somecolor.Color = 1000
	return x
}
