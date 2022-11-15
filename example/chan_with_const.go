package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func chanWithConst() {
	c := make(chan somecolor.Color, 2)
	c <- 100
}

func chanWithConstType() {
	type C chan somecolor.Color
	c := make(C, 2)
	c <- 100
}

func chanWithConstFunc() {
	f := func() chan somecolor.Color { return make(chan somecolor.Color, 12) }
	f() <- 123
}
