package win_test

import (
	"testing"
	"time"

	"qlova.tech/win"
	_ "qlova.tech/win/driver/glfw"
)

func TestMain(t *testing.T) {
	win.Name = "Hello World"

	for win.Open() {
		time.Sleep(time.Millisecond)
	}

	if win.Error != nil {
		t.Error(win.Error)
	}
}
