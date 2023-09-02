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
		&std.Float,
		&std.Char,
		&std.Time,
		&std.Locale,
	); err != nil {
		panic(err)
	}
}

func TestLibc(t *testing.T) {
	fmt.Println(std.Char.IsAlphaNumeric('a'))

	fmt.Println(std.Float.Sqrt(2))

	fmt.Println(std.Float.Frexp(2.2))

	fmt.Println(std.Time.Now(nil))

	fmt.Println(std.Locale.Get())

}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sqrt(2)
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		std.Float.Sqrt(2)
	}
}
