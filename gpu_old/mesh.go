package gpu

import (
	"errors"
	"math"
	"reflect"

	"qlova.tech/cad"
	"qlova.tech/ray"
	"qlova.tech/vec/vec3"
)

type meshBuffer struct {
	id Pointer

	vao uint64

	//stored so that the mesh can be read.
	vertices [][3]float32
	indicies []uint32

	attributes Attributes

	changed bool
}

//Mesh is a 3D model on the GPU.
type Mesh struct {
	draw func(Mesh, Transform, DrawOptions) error
	data *meshBuffer

	gpuBuffer Pointer
	offset    uint32
	count     uint32
	voffset   uint32

	indexed bool

	//version is used to prevent use-after free behaviour.
	version uint64

	//Bounding shapes used for frustrum culling.
	Bounding struct {
		Sphere float32
		cad.Box
	}

	Transform
	Material
}

func (m Mesh) Indexed() bool {
	return m.indexed
}

func (m Mesh) Count() uint32 {
	return m.count
}

func (m Mesh) Offset() uint32 {
	return m.offset
}

func (m Mesh) VOffset() uint32 {
	return m.voffset
}

func (m Mesh) Pointer() Pointer {
	return m.gpuBuffer
}

//Nil return true if the mesh is nil.
func (m Mesh) Nil() bool {
	return m.count == 0
}

//NewMesh returns a new mesh from the given vertices.
func NewMesh(a ...Attribute) (Mesh, error) {
	return context.NewMesh(a...)
}

//NewMesh returns a new mesh from the given vertices.
func (context *Context) NewMesh(a ...Attribute) (Mesh, error) {
	if context.Load == nil {
		return Mesh{}, ErrNotOpen
	}
	if context.meshBuffers == nil {
		context.meshBuffers = make(map[mode]*meshBuffer)
	}

	//determine the buffer that this mesh will be placed in.
	//each unique combination of attributes will be placed in a corresponding buffer.
	var m = mode{
		draw: 0,
		hash: Attributes(a).hash(),
	}

	//determine if the mesh is indexed.
	for _, attr := range a {
		if reflect.TypeOf(attr) == reflect.TypeOf(Indicies{}) {
			m.indexed = 1
		}
	}

	buf, ok := context.meshBuffers[m]
	if !ok {
		//Generate an empty attribute buffer.
		id, err := context.Load(Attributes{})
		if err != nil {
			return Mesh{}, err
		}

		buf = &meshBuffer{
			id: Pointer{id},
		}
		context.meshBuffers[m] = buf
	}

	var count, offset, voffset uint32

	//calculate mesh index, offset and count.
	if buf.attributes == nil {
		buf.attributes = a

		for _, attr := range a {
			if reflect.TypeOf(attr) == reflect.TypeOf(Vertices{}) && m.indexed == 0 {
				count = uint32(reflect.ValueOf(attr).Len())
				offset = 0
			}

			if v, ok := attr.(Vertices); ok {
				buf.vertices = v

			}
			if v, ok := attr.(Indicies); ok {
				buf.indicies = v
			}

			if reflect.TypeOf(attr) == reflect.TypeOf(Indicies{}) {
				count = uint32(reflect.ValueOf(attr).Len())
				offset = 0
			}
		}

	} else {
		//Append the data to the exisiting attribute slices.
		for i := range buf.attributes {
			existing := &buf.attributes[i]
			attr := a[i]

			if reflect.TypeOf(attr) == reflect.TypeOf(Vertices{}) {
				if m.indexed == 0 {
					count = uint32(reflect.ValueOf(attr).Len())
					offset = uint32(reflect.ValueOf(*existing).Len()) * 4 * 3
				} else {
					voffset = uint32(reflect.ValueOf(*existing).Len())
				}
			}

			if reflect.TypeOf(attr) == reflect.TypeOf(Indicies{}) {
				count = uint32(reflect.ValueOf(attr).Len())
				offset = uint32(reflect.ValueOf(*existing).Len())
			}

			buf.attributes[i] = reflect.AppendSlice(reflect.ValueOf(*existing), reflect.ValueOf(a[i])).Interface()

			if v, ok := buf.attributes[i].(Vertices); ok {
				buf.vertices = v
			}
			if v, ok := buf.attributes[i].(Indicies); ok {
				buf.indicies = v
			}
		}
	}

	buf.changed = true

	mesh := Mesh{
		draw: context.Draw,

		gpuBuffer: buf.id,

		data: buf,

		indexed: m.indexed != 0,

		//driver references.
		offset:  offset,
		count:   count,
		voffset: voffset,

		version: context.version,

		Transform: NewTransform(),
	}

	//Bounding Box
	for i := range mesh.Triangles() {
		var a, b, c = mesh.Triangle(i)
		var vertices = [3]Vec3{a, b, c}

		for _, vertex := range vertices {
			mesh.Bounding.Box.ExpandByPoint(cad.Point(vertex))
		}
	}

	return mesh, nil
}

//Draw scedules the drawing of the given mesh with the given transform applied.
func (m Mesh) Draw(options DrawOptions, transform *Transform) error {
	if m.draw == nil {
		return errors.New("nil draw")
	}

	if m.Material.shader == nil {
		return ErrNoShader
	}

	return m.draw(m, *transform, options)
}

//Triangles returns an iterator on the triangles of this mesh.
func (m Mesh) Triangles() []struct{} {
	return make([]struct{}, m.count/3)
}

//Triangle returns the i'th Triangle of this mesh.
func (m Mesh) Triangle(i int) (a, b, c Vec3) {
	buf := m.data
	if m.indexed {
		offset := i*3 + int(m.offset)
		return Vec3(buf.vertices[m.voffset+buf.indicies[offset]]),
			Vec3(buf.vertices[m.voffset+buf.indicies[offset+1]]),
			Vec3(buf.vertices[m.voffset+buf.indicies[offset+2]])
	}
	offset := i*3 + int(m.offset)
	return Vec3(buf.vertices[offset]), Vec3(buf.vertices[offset+1]), Vec3(buf.vertices[offset+2])
}

//Raycast checks if the specified ray intersects this ray.
//If ok, the intersection point is returned.
func (m Mesh) Raycast(ray ray.Caster, t Transform) (point Vec3, ok bool) {
	tt := m.Transform.Mul(t)

	inv := tt.Inverse()
	ray.Transform(&inv)

	if _, box := ray.Box(m.Bounding.Box); !box {
		return
	}

	var closest float32 = math.MaxFloat32

	for i := range m.Triangles() {
		if p, hit := ray.Triangle(m.Triangle(i)); hit {
			if dist := vec3.Distance(p, ray.Origin); dist < closest {
				point = p
				ok = true
				closest = dist
			}
		}
	}

	point.Transform(&tt)
	return
}
