package box_test

import (
	"testing"

	"qlova.tech/box"
)

type Point struct {
	X, Y float64
}

func Test_Main(t *testing.T) {

	var a = Point{2, 7}

	bytes, err := box.Marshal(&a)
	if err != nil {
		t.Fatal(err)
	}

	var b Point

	if err := box.Unmarshal(bytes, &b); err != nil {
		t.Fatal(err)
	}

	if b != a {
		t.Fatal("failed to box")
	}
}
