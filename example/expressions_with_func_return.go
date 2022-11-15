package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func expFunc() somecolor.Color { return somecolor.Blue }

func expressionsWithFuncPlus() somecolor.Color { return somecolor.Green + expFunc() }

func expressionsWithFuncMinus() somecolor.Color { return somecolor.Green - expFunc() }

func expressionsWithFuncMul() somecolor.Color { return somecolor.Green * expFunc() }

func expressionsWithFuncDiv() somecolor.Color { return somecolor.Green / expFunc() }

func expressionsWithFuncMod() somecolor.Color { return somecolor.Green % expFunc() }

func expressionsWithFuncShiftLeft() somecolor.Color { return somecolor.Green << expFunc() }

func expressionsWithFuncShiftRight() somecolor.Color { return somecolor.Green >> expFunc() }

func expressionsWithFuncBinAnd() somecolor.Color { return somecolor.Green & expFunc() }

func expressionsWithFuncBinAndTild() somecolor.Color { return somecolor.Green &^ expFunc() }

func expressionsWithFuncBinOr() somecolor.Color { return somecolor.Green | expFunc() }

func expressionsWithFuncBinXOr() somecolor.Color { return somecolor.Green ^ expFunc() }

func expressionsWithFuncLess() bool { return somecolor.Green < expFunc() }

func expressionsWithFuncLessEq() bool { return somecolor.Green <= expFunc() }

func expressionsWithFuncLessMore() bool { return somecolor.Green > expFunc() }

func expressionsWithFuncLessMoreEq() bool { return somecolor.Green >= expFunc() }
