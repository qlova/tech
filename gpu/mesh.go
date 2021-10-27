package gpu

//Mesh is essentially an opaque slice of vertex/index data on the GPU.
//meshes are designed to be rendered by a Shader.
type Mesh struct {
	data    Pointer
	offset  uint32
	count   uint32
	voffset uint32
	_       uint32 //reserved
}
