//Package opengl reports driver information when a opengl driver has been opened with the gpu package.
package opengl

var (
	//Active is true if the driver is active, if this value is false, then variables in this package should be treated as junk.
	Active bool

	//MaxTextureSize gives a rough estimate of the largest texture that OpenGL can handle.
	MaxTextureSize int32 = -1

	//MaxUniformBlockSize in basic machine units of a uniform block.
	MaxUniformBlockSize int32 = -1

	//MaxVertexUniformComponents is the maximum number of individual floating-point, integer, or boolean values that can be held in uniform variable storage for a vertex shader.
	MaxVertexUniformComponents int32 = -1

	//MaxVertexUniformVectors is the maximum number of 4-vectors that may be held in uniform variable storage for the vertex shader.
	MaxVertexUniformVectors int32 = -1
)
