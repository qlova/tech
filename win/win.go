//Package win provides a simple way to open a gpu-compatible window in Go.
package win

import "errors"

var ErrDriverNotSet = errors.New("driver not set")

//The following global variables act as an implicit window description
//for the 'main' window. Since most of the time, you will only be opening
//a single window, this is the most convenient way to control/observe the
//state of it. Keep in mind that access to these variables is not thread
//safe, so you should only use them in a single goroutine.
var (
	ID int

	Name string

	Width, Height int

	Fullscreen bool

	Closed bool

	Error error
)

//Window description.
type Window struct {
	ID int

	Name string

	Width, Height int

	Fullscreen bool

	Closed bool

	Error error
}

//Open returns true if the window description is open, if it is not open, it will be opened.
func (d *Window) Open() bool {
	return Driver(d)
}

//Driver is the current window driver, it will be set by driver packages during initialisation.
var Driver func(*Window) bool

var implicit Window

//Open returns true if the implicit window is open, if it is not open, it will be opened.
func Open() bool {
	implicit = Window{
		ID:         ID,
		Name:       Name,
		Width:      Width,
		Height:     Height,
		Fullscreen: Fullscreen,
		Closed:     Closed,
	}

	if Driver == nil {
		Error = ErrDriverNotSet
		return false
	}

	result := Driver(&implicit)

	ID = implicit.ID
	Name = implicit.Name
	Width = implicit.Width
	Height = implicit.Height
	Fullscreen = implicit.Fullscreen
	Closed = implicit.Closed

	return result
}
