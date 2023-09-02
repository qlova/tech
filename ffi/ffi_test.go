package ffi_test

import (
	"fmt"
	"math"
	"testing"

	"qlova.tech/abi"
	"qlova.tech/lib/std"
)

func init() {
	if err := std.Link(); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	std.Program.Exit(abi.Int(m.Run()))
}

func TestLibc(t *testing.T) {
	fmt.Println(std.Char.IsAlphaNumeric('a'))

	fmt.Println(std.Float.Sqrt(2))

	fmt.Println(std.Float.Frexp(2.2))

	fmt.Println(std.Time.Now(nil))

	fmt.Println(std.Locale.Get())

	std.Program.OnExit(func() {
		fmt.Println("exiting...")
	})
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
