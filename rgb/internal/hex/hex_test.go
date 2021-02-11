package hex

import (
	"testing"
)

func Test_Hex(t *testing.T) {
	if String3("#232323") != "#232323" {
		t.Fail()
	}
	if String3("232323") != "#232323" {
		t.Fail()
	}
	if String3("#gfsm;gf") != "#000000" {
		t.Fail()
	}

	if String4("#23232343") != "#23232343" {
		t.Fail()
	}
	if String4("23232343") != "#23232343" {
		t.Fail()
	}
	if String4("#gfsm;gf") != "#00000000" {
		t.Fail()
	}
}
