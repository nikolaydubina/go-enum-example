package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

type StructWithColor struct {
	C somecolor.Color
}

func structWithColorAssignment() somecolor.Color {
	s := StructWithColor{10}
	return s.C
}

func structWithColorAssignmentNamed() somecolor.Color {
	s := StructWithColor{C: 10}
	return s.C
}

func structWithColorMutate() somecolor.Color {
	var s StructWithColor
	s.C = 100
	return s.C
}
