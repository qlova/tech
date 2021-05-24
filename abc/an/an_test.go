package an_test

import (
	"testing"

	"qlova.org/should"
	"qlova.tech/abc/an"
)

func TestPackage(t *testing.T) {

	shouldBeAn := []string{
		"apple",
		"Umbrella",
		"Apple",
		"Umbrella",
		"IOU",
		"FFA prodigy",
		"honor",
		"honor-bound",
		"hours",
		"heiresses",
		"honored",
		"heir's",
		"Hour",

		"a",
		"f",
		"h",
		"l",
		"m",
		"n",
		"r",
		"s",
		"x",
	}

	shouldBeA := []string{
		"banana",
		"Banana",
		"UFO",
		"CEO",
		"euro",
		"ukelele",
		"ouija board",
		"b",
	}

	for _, testcase := range shouldBeAn {
		should.Be("an " + testcase)(an.A(testcase)).Test(t)
	}
	for _, testcase := range shouldBeA {
		should.Be("a " + testcase)(an.A(testcase)).Test(t)
	}
}
