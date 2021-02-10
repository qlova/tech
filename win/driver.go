package win

type Driver interface {
	Open() error
	//Fullscreen() error

	Update() bool

	DeltaTime() float32

	//SetTitle(title string) error

	//SetPosition(x, y int) error
	//Position() (x, y int)

	//Size() (width, height int)
	//Resize(width, height int) error

	//Mouse() (x, y int)
	Button(name string) bool

	//Event handling.
	//OnChar(func(char string))
	//OnButton(name string, pressed bool) bool
	//OnMouse(func(button string, pressed bool))
	//OnScroll(func(amount int))
	//OnResize(func(width, height int))
	//OnDrop(func(path string))

	//Copy(string) error
	//Paste() string

	Close()
}
