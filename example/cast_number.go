package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func castNumber() somecolor.Color {
	x := somecolor.Color(1000)
	return x
}
