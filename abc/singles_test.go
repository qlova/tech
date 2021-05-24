package abc_test

import (
	"testing"

	"qlova.org/should"
	"qlova.tech/abc"
)

func TestSingles(t *testing.T) {
	for _, plural := range plurals {
		should.Be(plural[0])(abc.ToSingle(plural[1])).Test(t)
	}
}
