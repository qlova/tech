//Package gpu is a high-level GPU rendering API.
package gpu

import (
	"image"
	"reflect"
	"sort"
	"sync"
	"unsafe"

	"qlova.tech/dsl"
	"qlova.tech/rgb"
	"qlova.tech/xyz"
)

var (
	// Camera is a standard camera position and
	// rotation, recognised by drivers and shaders.
	Camera xyz.Transform

	// Projection is the desired projection matrix. May
	// be overridden by driver if rendering to a
	// steroscopic display.
	Projection xyz.Transform

	// Transform applied to the model before rendering.
	Transform xyz.Transform
)

type MeshHint uint64

type TextureHint uint64

type ProgramHint uint64

type AnimationHint uint64

type Raycast struct {
	Position  xyz.Vector
	Direction xyz.Vector

	hit uint32 //draw call ID that was hit.
}

func (ray *Raycast) Hit(drawcall uint32) {
	ray.hit = drawcall
}

//TODO
type Frames interface {

	//Length returns the number of frames.
	Length() int

	//Frame returns the bone transformations for the given frame.
	Frame(int) []xyz.Transform
}

// Driver is a GPU driver that enables GPU rendering.
type Driver struct {

	// NewFrame begins a new frame, clearing it with the
	// given color.
	NewFrame func(color rgb.Color)

	// NewMesh returns a new Mesh from the given vertices.
	NewMesh func(vertices xyz.Vertices, hints ...MeshHint) (unsafe.Pointer, error)

	// NewTexture returns a new Texture from the given image.
	NewTexture func(image image.Image, hints ...TextureHint) (unsafe.Pointer, error)

	// NewProgram returns a new Program from the given DSL shaders.
	NewProgram func(vert, frag dsl.Shader, hints ...ProgramHint) (unsafe.Pointer, error)

	// NewAnimation returns a new Animation from the given animation frames.
	NewAnimation func(frames Frames, hints ...AnimationHint) (unsafe.Pointer, error)

	// SetRaycasts sets the rays that need to be tested for intersection
	// with drawn meshes in the next frame. Before the end of the next Sync
	// you must call Hit on each ray with the drawcall of the mesh that was hit.
	SetRaycasts func(raycasts []Raycast)

	// SetLighting sets the lighting that should be applied to drawcalls next
	// frame.
	SetLighting func(lights []Light)

	// Draw makes a numbered 'draw call' that uses the given program to draw
	// the given mesh.
	Draw func(program, mesh unsafe.Pointer)

	// Sync flushes any pending operations to the GPU.
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
	ptr unsafe.Pointer
}

// NewMesh returns a new Mesh from the given vertex array and
// vertex hints.
func NewMesh(vertices xyz.Vertices, hints ...MeshHint) (Mesh, error) {
	ptr, err := driver.NewMesh(vertices, hints...)
	return Mesh{ptr: ptr}, err
}

// Program is a reference to a program uploaded to the GPU.
// Programs are used to draw Meshes and specify how they
// will be renderered.
type Program struct {
	ptr unsafe.Pointer
}

// NewProgram returns a new Program from the given vertex
// and fragment functions. The vertex function is called
// once for each vertex in the mesh. The fragment function
// is called once for each fragment in the mesh. Hints can
// be provided to configure rendering behaviour.
func NewProgram(vert, frag dsl.Shader, hints ...ProgramHint) (Program, error) {
	ptr, err := driver.NewProgram(vert, frag, hints...)
	return Program{ptr}, err
}

// Draw draws the given mesh using the given program.
// Until the next call to gpu.Draw, the rendering order
// of subsequent calls to Draw are undefined.
func (p Program) Draw(m Mesh) {
	driver.Draw(p.ptr, m.ptr)
}

// Texture is a reference to a texture uploaded to the GPU.
type Texture struct {
	dsl.Texture

	ptr unsafe.Pointer
}

// NewTexture returns a new Texture from the given texture data and hints.
func NewTexture(img image.Image, hints ...TextureHint) (Texture, error) {
	ptr, err := driver.NewTexture(img, hints...)
	return Texture{nil, ptr}, err
}

// SetLighting sets the lighting for the next frame.
func SetLighting(lights ...Light) {
	driver.SetLighting(lights)
}

// Attributes can be used to specify vertex data by
// specifying the vertex data for each attribute.
// Each Attribute's Slice must be the same length.
type Attributes map[dsl.Attribute]slice

// Length returns the number of vertices.
func (a Attributes) Length() int {
	var keys = make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	var length = -1
	for _, key := range keys {
		attr := a[dsl.Attribute(key)]
		l := len(attr.buffer) / attr.size
		if length == -1 {
			length = l
		} else {
			if l < length {
				length = l
			}
		}
	}

	if length == -1 {
		return 0
	}

	return length
}

func (a Attributes) Indexed() (int, bool) {
	return 0, false
}

// Layout returns the vertex attribute layout for
// the Attributes, each attribute is tightly packed
// in a seperate buffer.
func (a Attributes) Layout() []xyz.Pointer {
	var pointers = make([]xyz.Pointer, 0, len(a))

	var keys = make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for i, key := range keys {
		attribute := dsl.Attribute(key)
		data := a[attribute]

		pointers = append(pointers, xyz.Pointer{
			Attribute: string(attribute),
			Kind:      data.kind,
			Buffer:    uint(i),
			Count:     uint(data.count),
		})
	}

	return pointers
}

// Buffers returns the buffers used to store the
// vertex data of the Attributes.
func (a Attributes) Buffers() []xyz.Buffer {
	var buffers = make([]xyz.Buffer, 0, len(a))

	var keys = make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for _, key := range keys {
		attribute := dsl.Attribute(key)
		data := a[attribute]

		buffers = append(buffers, data.buffer)
	}

	return buffers
}

//Slice is a typed slice of vertex data.
//Sliced on a single Attribute.
type slice struct {
	buffer xyz.Buffer
	size   int //element size.
	kind   reflect.Kind
	count  int
}

type data interface {
	//basic types.
	//We can accept single values or up to 4 values.
	~bool | ~[1]bool | ~[2]bool | ~[3]bool | ~[4]bool |
		~int8 | ~[1]int8 | ~[2]int8 | ~[3]int8 | ~[4]int8 |
		~int16 | ~[1]int16 | ~[2]int16 | ~[3]int16 | ~[4]int16 |
		~int32 | ~[1]int32 | ~[2]int32 | ~[3]int32 | ~[4]int32 |
		~uint8 | ~[1]uint8 | ~[2]uint8 | ~[3]uint8 | ~[4]uint8 |
		~uint16 | ~[1]uint16 | ~[2]uint16 | ~[3]uint16 | ~[4]uint16 |
		~uint32 | ~[1]uint32 | ~[2]uint32 | ~[3]uint32 | ~[4]uint32 |
		~float32 | ~[1]float32 | ~[2]float32 | ~[3]float32 | ~[4]float32 |
		~float64 | ~[1]float64 | ~[2]float64 | ~[3]float64 | ~[4]float64 |

		~complex64 | ~[1]complex64 | ~[2]complex64
}

//Data returns the given data slice as a Slice.
func Data[T data](data []T) slice {
	var elem = &data[0]
	var rtype = reflect.TypeOf(elem)
	var kind = rtype.Elem().Kind()
	var count = 1
	if kind == reflect.Array {
		count = rtype.Elem().Len()
		kind = rtype.Elem().Elem().Kind()
	}

	return slice{
		kind:   kind,
		count:  count,
		size:   int(unsafe.Sizeof(elem)),
		buffer: unsafe.Slice((*byte)(unsafe.Pointer(elem)), unsafe.Sizeof(data[0])*uintptr(len(data))),
	}
}
