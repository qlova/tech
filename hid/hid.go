//Package hid provides a platform-agnostic interface to human interface devices.
package hid

import "qlova.tech/key"

var Mouse MouseState
var Keyboard KeyboardState
var Joystick JoystickState
var Clipboard ClipboardState

type MouseState struct {
	X, Y float32

	Wheel complex64

	Buttons uint16
}

func (mouse *MouseState) IsPressed(button uint16) bool {
	return mouse.Buttons&(1<<button) != 0
}

func (mouse *MouseState) Press(button uint16) {
	mouse.Buttons |= 1 << button
}

func (mouse *MouseState) Release(button uint16) {
	mouse.Buttons &^= 1 << button
}

type KeyboardState struct {
	Keys key.State

	Text string
}

type JoystickState struct {
	Axis    []float32
	Buttons uint64
}

type ClipboardState struct {
	Data []byte
}
