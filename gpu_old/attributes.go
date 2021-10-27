package gpu

import (
	"path"
	"reflect"
	"strings"
	mutex "sync"
)

/*
	Meshes are generally composed of vertices which contain a set of attributes.
	For example positions, UVs, colors, normals, etc. Some of the most common
	attributes are included here but you can also define your own attributes.

	Each field must be a fixed size type, the driver will attempt to convert
	your data to the closest representation. See the driver packages for more
	information.

	type MeshVertex struct {
		Position  Vec3
		Normal    Vec3
		UV        Vec2
		Color     Vec4

		CustomAttribute1 Vec3
	}

	In your vertex shader, you will then refer to these attributes.
*/

var ihash uint64
var hashes = make(map[string]uint64)
var syncHash mutex.Mutex

//Attributes is one or more Attribute elements.
type Attributes []Attribute

//hash returns a unique string for the given attributes.
func (a Attributes) hash() uint64 {
	var key string
	for _, attr := range a {
		key += reflect.TypeOf(attr).String()
		key += ";"
	}

	syncHash.Lock()
	defer syncHash.Unlock()

	if hash, ok := hashes[key]; ok {
		return hash
	}

	ihash++
	hashes[key] = ihash

	return ihash
}

//Attribute of a vertex or mesh.
type Attribute interface{}

//AttributeName returns the name of an Attribute.
func AttributeName(attr Attribute) string {
	type named interface{ Name() string }

	if n, ok := attr.(named); ok {
		return n.Name()
	}

	return strings.ToLower(path.Ext(reflect.TypeOf(attr).String())[1:])
}

//Vertices geometry of a mesh.
type Vertices [][3]float32

//Name of the attribute is "position"
func (Vertices) Name() string {
	return "position"
}

//Normals of a mesh.
type Normals [][3]float32

//Name of the attribute is "normal"
func (Normals) Name() string {
	return "normal"
}

//UVs of a mesh.
type UVs [][2]float32

//Name of the attribute is "uv"
func (UVs) Name() string {
	return "uv"
}

//Colors of a mesh.
type Colors [][3]uint8

//Name of the attribute is "color"
func (Colors) Name() string {
	return "color"
}

//Indicies of a mesh.
type Indicies []uint32

//Weights of a mesh.
type Weights [][4]float32

//Name of the attribute is "weight"
func (Weights) Name() string {
	return "weight"
}

//Joints of a mesh.
type Joints [][4]uint8

//Name of the attribute is "joint"
func (Joints) Name() string {
	return "joint"
}
