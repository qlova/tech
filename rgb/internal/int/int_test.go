package int

import "testing"

func Test_Int(t *testing.T) {
	if Encode3(Decode3(100)) != 100 {
		t.Fail()
	}
	if Encode4(Decode4(23432424)) != 23432424 {
		t.Fail()
	}
}
