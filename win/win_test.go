package win_test

import (
	"fmt"
	"testing"
	"time"

	"qlova.tech/hid"
	"qlova.tech/win"
	_ "qlova.tech/win/driver/glfw"
)

func TestMain(t *testing.T) {
	win.Name = "Hello World"

	for win.Open() {
		if hid.Mouse.IsPressed(2) {
			fmt.Println(hid.Mouse.X, hid.Mouse.Y)
		}
		time.Sleep(time.Millisecond * 100)
	}

	if win.Error != nil {
		t.Error(win.Error)
	}
}
