package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

const x = 1

func expressionsWithConstNamedPlus() somecolor.Color { return somecolor.Green + x }

func expressionsWithConstNamedMinus() somecolor.Color { return somecolor.Green - x }

func expressionsWithConstNamedMul() somecolor.Color { return somecolor.Green * x }

func expressionsWithConstNamedDiv() somecolor.Color { return somecolor.Green / x }

func expressionsWithConstNamedMod() somecolor.Color { return somecolor.Green % x }

func expressionsWithConstNamedShiftLeft() somecolor.Color { return somecolor.Green << x }

func expressionsWithConstNamedShiftRight() somecolor.Color { return somecolor.Green >> x }

func expressionsWithConstNamedBinAnd() somecolor.Color { return somecolor.Green & x }

func expressionsWithConstNamedBinAndTild() somecolor.Color { return somecolor.Green &^ x }

func expressionsWithConstNamedBinOr() somecolor.Color { return somecolor.Green | x }

func expressionsWithConstNamedBinXOr() somecolor.Color { return somecolor.Green ^ x }

func expressionsWithConstNamedLess() bool { return somecolor.Green < x }

func expressionsWithConstNamedLessEq() bool { return somecolor.Green <= x }

func expressionsWithConstNamedLessMore() bool { return somecolor.Green > x }

func expressionsWithConstNamedLessMoreEq() bool { return somecolor.Green >= x }
