package cgo_test

import (
	"fmt"
	"math"
	"testing"

	"qlova.tech/cgo"
	"qlova.tech/cgo/std"
)

func init() {
	if err := cgo.Set(&std.Math, "libm.so.6"); err != nil {
		panic(err)
	}
	if err := cgo.Set(&std.Char, "libc.so.6"); err != nil {
		panic(err)
	}
	if err := cgo.Set(&std.Clock, "libc.so.6"); err != nil {
		panic(err)
	}
}

func TestLibc(t *testing.T) {
	fmt.Println(std.Char.IsAlphaNumeric('a'))

	fmt.Println(std.Math.Sqrt(2))

	fmt.Println(std.Math.Frexp(2.2))

	fmt.Println(std.Clock.Time(nil))
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sqrt(2)
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		std.Math.Sqrt(2)
	}
}
