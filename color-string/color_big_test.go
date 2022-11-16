package color_test

import (
	"testing"

	"github.com/nikolaydubina/go-enum-example/color-string"
)

var colorsBig = []color.ColorBig{
	color.RedBig,
	color.BlueBig,
	color.GreenBig,
	color.PurpleBig,
}

type AppleBig struct {
	Name  string
	Size  int
	Color color.ColorBig
}

//go:noinline
func callOneBig(i int, c color.ColorBig) color.ColorBig {
	if colorsBig[(i+17)%len(colors)] == c {
		return color.PurpleBig
	}
	return c
}

//go:noinline
func callOneAppleBig(i int, a AppleBig, c color.ColorBig) AppleBig {
	if a.Color == c {
		return AppleBig{Color: color.PurpleBig}
	}
	return a
}

func BenchmarkEnumPassFunction_Big(b *testing.B) {
	var a color.ColorBig

	b.ResetTimer()
	b.Run("call_one", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			x := colorsBig[n%len(colors)]
			a = callOneBig(n, x)
		}
	})

	b.ResetTimer()
	b.Run("call_one_apple", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			x := colorsBig[n%len(colors)]
			y := colorsBig[(n%31)%len(colors)]
			apple := callOneAppleBig(n, AppleBig{Color: y}, x)
			a = apple.Color
		}
	})

	// blocking optimizations
	has := false
	for _, q := range colorsBig {
		if a == q {
			has = true
			break
		}
	}
	if !has {
		b.Error()
	}
}
