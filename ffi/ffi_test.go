package ffi_test

import (
	"fmt"
	"math"
	"testing"

	"qlova.tech/ffi"
	"qlova.tech/lib/std"
)

func init() {
	if err := ffi.Link(
		&std.Math,
		&std.Char,
		&std.Clock,
	); err != nil {
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
