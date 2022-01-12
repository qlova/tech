//Package gpu is a high-level GPU rendering API.
package gpu

import (
	"image/color"
	"sync"

	"qlova.tech/gpu/dsl"
	"qlova.tech/gpu/internal/core"
	"qlova.tech/gpu/texture"
	"qlova.tech/gpu/vertex"
)

// Driver is a GPU driver that enables GPU rendering.
type Driver struct {
	NewFrame   func(c color.Color)
	NewMesh    func(vertices vertex.Array, hints ...vertex.Hint) (vertex.Reader, Pointer, error)
	NewTexture func(data texture.Data, hints ...texture.Hint) (texture.Reader, Pointer, error)
	NewProgram func(vert, frag func(Core), hints ...Hint) (Binary, Pointer, error)

	SetShader func(shader func(Core))

	Draw func(program, mesh Pointer)
	Sync func()
}

//driver is the current driver.
var driver Driver
var drivers map[string]func() (Driver, error)
var mutex sync.Mutex

// Register a new gpu Opener.
func Register(name string, opener func() (Driver, error)) {
	mutex.Lock()
	defer mutex.Unlock()
	drivers[name] = opener
}

// Core is a processing core on the GPU and can
// be instructed on how to render a Mesh.
// Whenever a mesh is drawn, the gpu core will be
// passed vertex attributes, which can then
// adjust the attributes and produce an output.
//
// The core can be instructed using a GLSL-style
// DSL. Checkout the dsl package for more info.
type Core = dsl.Core

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

// Draw flushes any pending operations to the GPU
// and returns true.
func Draw() bool {
	driver.Sync()
	return true
}

// NewFrame clears the output display with the given color.
// Until the next call to Draw, all drawing operations will
// be shaded by the value of Shader at the time of the call.
func NewFrame(c color.Color) {
	driver.NewFrame(c)
}

// SetShader sets the shader used to determine the shading
// for a fragment. By default, no shading calculation
// is performed and the fragment will be set to the
// output of the fragment's program. This call will have
// take effect after the next call to gpu.Draw.
func SetShader(shader func(Core)) {
	driver.SetShader(shader)
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

// Hint is a hint that can be used to configure the
// behaviour of the GPU when rendering.
type Hint uint64

// Hints.
const (
	Cull Hint = 1 << iota
	Front
	Back
	Blend
	Wireframe
	Shaded
)

//Binary is a compiled Program.
type Binary interface {
	Data() []byte
}

// Program is a reference to a program uploaded to the GPU.
// Programs are used to draw Meshes and specify how they
// will be renderered.
type Program struct {
	reader  Binary
	pointer Pointer
}

// NewProgram returns a new Program from the given vertex
// and fragment functions. The vertex function is called
// once for each vertex in the mesh. The fragment function
// is called once for each fragment in the mesh. Hints can
// be provided to configure rendering behaviour.
func NewProgram(vert, frag func(Core), hints ...Hint) (Program, error) {
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
type Texture = core.Texture

// NewTexture returns a new Texture from the given texture data and hints.
func NewTexture(data texture.Data, hints ...texture.Hint) (Texture, error) {
	reader, pointer, err := driver.NewTexture(data, hints...)
	return core.NewTexture(reader, core.Pointer(pointer)), err
}
