//Package gpu is a high-level GPU rendering API.
package gpu

import (
	"sync"

	"qlova.tech/dsl"
	"qlova.tech/rgb"
	"qlova.tech/rgb/texture"
	"qlova.tech/rgb/led"
	"qlova.tech/xyz/vertex"
)

// Driver is a GPU driver that enables GPU rendering.
type Driver struct {
	NewFrame   func(color rgb.Color)
	NewMesh    func(vertices vertex.Array, hints ...vertex.Hint) (vertex.Reader, Pointer, error)
	NewTexture func(image texture.Data, hints ...texture.Hint) (texture.Reader, Pointer, error)
	NewProgram func(vert, frag dsl.Shader, hints ...dsl.Hint) (dsl.Reader, Pointer, error)

	SetLighting func(lights ...led.Light)

	Draw func(program, mesh Pointer)
	Sync func()
}

//driver is the current driver.
var driver Driver
var drivers = make(map[string]func() (Driver, error))
var mutex sync.Mutex

// Register a new gpu Opener.
func Register(name string, opener func() (Driver, error)) {
	mutex.Lock()
	defer mutex.Unlock()
	drivers[name] = opener
}

// Pointer is a pointer to a GPU resource.
// Drivers are free to use this as they wish.
type Pointer [3]uint64

// Error is an error returned by the GPU package.
type Error struct {
	string
}

// Error implements error.
func (e Error) Error() string {
	return e.string
}

// ErrNoDriver is returned when no driver is available.
var ErrNoDriver = Error{"no gpu.Driver available, did you import one?"}

// Open attempts to open a GPU driver for rendering operations.
// returns ErrNoDriver if no driver is available. Pass hints
// to influence which driver is used for rendering when there
// are multiple candidates.
func Open(hints ...string) error {
	mutex.Lock()
	defer mutex.Unlock()

	try := func(name string) error {
		if opener, ok := drivers[name]; ok {
			var err error
			driver, err = opener()
			if err != nil {
				return err
			}
			return nil
		}
		return nil
	}

	for _, name := range hints {
		if err := try(name); err == nil {
			return nil
		}
	}

	for name := range drivers {
		if err := try(name); err == nil {
			return nil
		}
	}

	return ErrNoDriver
}

// Sync flushes any pending operations to the GPU
// and returns true.
func Sync() bool {
	driver.Sync()
	return true
}

// NewFrame clears the output display with the given color.
// Until the next call to Draw, all drawing operations will
// be shaded by the value of Shader at the time of the call.
func NewFrame(color rgb.Color) {
	driver.NewFrame(color)
}

//Mesh is a reference to a mesh uploaded to the GPU.
type Mesh struct {
	reader  vertex.Reader
	pointer Pointer
}

// NewMesh returns a new Mesh from the given vertex array and
// vertex hints.
func NewMesh(vertices vertex.Array, hints ...vertex.Hint) (Mesh, error) {
	reader, pointer, err := driver.NewMesh(vertices, hints...)
	return Mesh{reader, pointer}, err
}

// Program is a reference to a program uploaded to the GPU.
// Programs are used to draw Meshes and specify how they
// will be renderered.
type Program struct {
	reader  dsl.Reader
	pointer Pointer
}

// NewProgram returns a new Program from the given vertex
// and fragment functions. The vertex function is called
// once for each vertex in the mesh. The fragment function
// is called once for each fragment in the mesh. Hints can
// be provided to configure rendering behaviour.
func NewProgram(vert, frag dsl.Shader, hints ...dsl.Hint) (Program, error) {
	binary, pointer, err := driver.NewProgram(vert, frag, hints...)
	return Program{binary, pointer}, err
}

// Draw draws the given mesh using the given program.
// Until the next call to gpu.Draw, the rendering order
// of subsequent calls to Draw are undefined.
func (p Program) Draw(m Mesh) {
	driver.Draw(p.pointer, m.pointer)
}

// Texture is a reference to a texture uploaded to the GPU.
type Texture struct {
	dsl.Texture

	reader  texture.Reader
	pointer Pointer
}

// NewTexture returns a new Texture from the given texture data and hints.
func NewTexture(image texture.Data, hints ...texture.Hint) (Texture, error) {
	reader, pointer, err := driver.NewTexture(image, hints...)
	return Texture{nil, reader, pointer}, err
}

// SetLighting sets the lighting for the next frame.
func SetLighting(lights ...led.Light) {
	driver.SetLighting(lights...)
}
