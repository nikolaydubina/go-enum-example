package color_test

import (
	"testing"

	"github.com/nikolaydubina/go-enum-example/color-string"
)

var colors = []color.Color{
	color.Red,
	color.Blue,
	color.Green,
	color.Purple,
}

type Apple struct {
	Name  string
	Size  int
	Color color.Color
}

//go:noinline
func callOne(i int, c color.Color) color.Color {
	if colors[(i+17)%len(colors)] == c {
		return color.Purple
	}
	return c
}

//go:noinline
func callOneApple(i int, a Apple, c color.Color) Apple {
	if a.Color == c {
		return Apple{Color: color.Purple}
	}
	return a
}

func BenchmarkEnumPassFunction(b *testing.B) {
	var a color.Color

	b.ResetTimer()
	b.Run("call_one", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			x := colors[n%len(colors)]
			a = callOne(n, x)
		}
	})

	b.ResetTimer()
	b.Run("call_one_apple", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			x := colors[n%len(colors)]
			y := colors[(n%31)%len(colors)]
			apple := callOneApple(n, Apple{Color: y}, x)
			a = apple.Color
		}
	})

	// blocking optimizations
	has := false
	for _, q := range colors {
		if a == q {
			has = true
			break
		}
	}
	if !has {
		b.Error()
	}
}
