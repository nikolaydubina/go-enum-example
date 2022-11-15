package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func expressionsWithItselfPlus() somecolor.Color { return somecolor.Green + somecolor.Green }

func expressionsWithItselfMinus() somecolor.Color { return somecolor.Green - somecolor.Green }

func expressionsWithItselfMul() somecolor.Color { return somecolor.Green * somecolor.Green }

func expressionsWithItselfDiv() somecolor.Color { return somecolor.Green / somecolor.Green }

func expressionsWithItselfMod() somecolor.Color { return somecolor.Green % somecolor.Green }

func expressionsWithItselfShiftLeft() somecolor.Color { return somecolor.Green << somecolor.Green }

func expressionsWithItselfShiftRight() somecolor.Color { return somecolor.Green >> somecolor.Green }

func expressionsWithItselfBinAnd() somecolor.Color { return somecolor.Green & somecolor.Green }

func expressionsWithItselfBinAndTild() somecolor.Color { return somecolor.Green &^ somecolor.Green }

func expressionsWithItselfBinOr() somecolor.Color { return somecolor.Green | somecolor.Green }

func expressionsWithItselfBinXOr() somecolor.Color { return somecolor.Green ^ somecolor.Green }

func expressionsWithItselfLess() bool { return somecolor.Green < somecolor.Green }

func expressionsWithItselfLessEq() bool { return somecolor.Green <= somecolor.Green }

func expressionsWithItselfLessMore() bool { return somecolor.Green > somecolor.Green }

func expressionsWithItselfLessMoreEq() bool { return somecolor.Green >= somecolor.Green }
