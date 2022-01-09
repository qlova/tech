package gpu

import (
	"errors"

	"qlova.tech/mat/mat4"
	"qlova.tech/vec/vec4"
)

//Driver is able to open return a context.
type Driver func() (Context, error)

//Pointer is an opaque reference to a GPU memory location.
type Pointer struct {
	uint64
}

func (p Pointer) Value() uint64 {
	return p.uint64
}

//Update updates the given buffer's data, resizing the buffer if needed.
type Update struct {
	Pointer Pointer
	Data    interface{}
}

//Variable is a variable on the GPU.
type Variable struct {
	Name  string
	Value interface{}
}

//Set sets a uniform variable on the GPU.
//The value is read after a Sync, so only the last value is used.
func Set(name string, value interface{}) error {
	_, err := context.Load(Variable{name, value})
	return err
}

//Context provides an interface to a context on the GPU.
type Context struct {
	//Buffer uploads the given data into a suitable buffer on the GPU.
	//If the type is unsupported, it will return an error. An opaque reference to the buffer is returned.
	//A buffer can be updated by passing an Update type to this function with the Buffer that needs updating.
	//If 'data' is nil, the entire context is cleared and reset to an empty state.
	Load func(data interface{}) (uint64, error)

	//Draw schedules the drawing of the given mesh with the given transform and drawing options.
	//If the mesh has an invalid/outdated buffer then this may return an error.
	Draw func(Mesh, Transform, DrawOptions) error

	//Sync waits for any pending buffer operations on the GPU to complete and then waits for all pending drawing operations to complete.
	Sync func() error

	version     uint64
	meshBuffers map[mode]*meshBuffer
}

//Upload uploads any scheduled meshes to the GPU.
func (context *Context) Upload() error {
	for _, buf := range context.meshBuffers {
		if buf.changed {
			_, err := context.Load(Update{buf.id, buf.attributes})
			if err != nil {
				return err
			}

			buf.changed = false
		}
	}
	return nil
}

//Upload uploads any scheduled meshes to the GPU.
func Upload() error { return context.Upload() }

var driver string
var drivers = make(map[string]Driver)

//Register registers a new GPU driver.
func Register(name string, d Driver) {
	drivers[name] = d
}

var context Context

//Open opens the GPU ready for uploading data and rendering.
//You can provide a hint to select the name of the driver to use.
func Open(hints ...string) error {
	if len(hints) > 0 {
		open, ok := drivers[hints[0]]
		if ok {
			ctx, err := open()
			if err == nil {
				driver = hints[0]
				context = ctx
				return nil
			}
			return err
		}
		return errors.New("gpu driver " + hints[0] + " not found")
	}

	if len(drivers) == 0 {
		return errors.New("no drivers available, please import one")
	}

	var ErrorString string = "failed to open a gpu.Context\n"

	for name, open := range drivers {
		ctx, err := open()
		if err == nil {
			driver = name
			context = ctx
			return nil
		}
		ErrorString += "\t" + name + ":" + err.Error()
	}

	return errors.New(ErrorString)
}

//DrawOptions that affect how a mesh is drawn.
type DrawOptions uint64

//DrawOptions.
const (
	Wireframe = 1 << iota
	FrontFaceCulling
	NoShadows
	Clear
)

type mode struct {
	//drawing mode, ie triangles, lines.
	draw    uint16
	indexed uint16

	options DrawOptions

	//hash of the attributes.
	hash uint64
}

//Frame is a definition for a frame.
type Frame struct {
	ClearColor vec4.Type
}

//Frames clears the screen to black and then returns true.
func Frames() bool {
	context.Load(Frame{})
	return true
}

//Sync applies all sceduled drawing operations.
func Sync() error { return context.Sync() }

var Camera mat4.Type
