package example

import (
	somecolor "github.com/nikolaydubina/enumcheck/example/color"
)

func expressionsWithConstPlus() somecolor.Color { return somecolor.Green + 5 }

func expressionsWithConstMinus() somecolor.Color { return somecolor.Green - 1 }

func expressionsWithConstMul() somecolor.Color { return somecolor.Green * 2 }

func expressionsWithConstDiv() somecolor.Color { return somecolor.Green / 2 }

func expressionsWithConstMod() somecolor.Color { return somecolor.Green % 2 }

func expressionsWithConstShiftLeft() somecolor.Color { return somecolor.Green << 2 }

func expressionsWithConstShiftRight() somecolor.Color { return somecolor.Green >> 2 }

func expressionsWithConstBinAnd() somecolor.Color { return somecolor.Green & 0 }

func expressionsWithConstBinAndTild() somecolor.Color { return somecolor.Green &^ 0 }

func expressionsWithConstBinOr() somecolor.Color { return somecolor.Green | 0 }

func expressionsWithConstBinXOr() somecolor.Color { return somecolor.Green ^ 0 }

func expressionsWithConstLess() bool { return somecolor.Green < 0 }

func expressionsWithConstLessEq() bool { return somecolor.Green <= 0 }

func expressionsWithConstLessMore() bool { return somecolor.Green > 0 }

func expressionsWithConstLessMoreEq() bool { return somecolor.Green >= 0 }
