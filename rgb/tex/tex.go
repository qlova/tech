package tex

// Hint for how the gpu should handle the texture.
type Hint uint64

// Hints
const (
	// Read enables the texture to be read.
	Read Hint = 1 << iota

	// Write enables the texture to be written to.
	// This may be implemented as a framebuffer.
	Write

	// Mirror the repetition of the texture.
	// Combine with X, Y, Z to set axis.
	Mirror

	// Clamp the texture, so that it
	// is not repeated.
	// Combine with X, Y, Z to set axis.
	Clamp

	// Crop the texture so that samples
	// outside of the texture are blank.
	// Combine with X, Y, Z to set axis.
	Crop

	// Nearest filtering, scaling will
	// scale the pixels.
	Nearest

	// Linear filtering, scaling will
	// blur the texture.
	Linear

	// Directions used to specify axis
	// of certain Hints.
	X
	Y
	Z

	// Zoom can be combined with
	// Nearest or Linear to specify
	// the filtering when the texture
	// is magnified.
	Zoom

	// Generate mipmaps for this texture.
	// Combine with Nearest or Linear to
	// specify the filtering for the
	// mipmap.
	Mipmap

	// Texture is 3D
	Cube

	// Texture is 1D.
	Flat

	// The GPU should compress the texture using
	// an internal compression format.
	Compress
)

//Reader TODO
type Reader interface{}

// Format is a texture format.
// The imaginary part of the format is either a
// specifier for the format's version/number
// or the number of bits per channel.
// If the imaginary part is negative then
// a linear SRGB format is assumed, otherwise
// SRGB is the default colour space.
//
// If the real part of the format is negative
// then an alpha channel is included.
type Format complex128

// Formats
const (
	Alpha = -1

	// Raw formats.
	RGB Format = iota + 1
	BGR

	// Common formats.
	PNG
	JPG
	QOI
	GIF
	DDS

	DTX //S3TC (GL_COMPRESSED_RGB_S3TC_DXT1_EXT/GL_COMPRESSED_RGBA_S3TC_DXT5_EXT)

	//Red Green Texture Compression
	RTC  //GL_COMPRESSED_RED_RGTC1_EXT
	RGTC //GL_COMPRESSED_RED_GREEN_RGTC2_EXT

	//Ericsson Texture Compression
	ETC //(GL_ETC1_RGB8_OES/GL_COMPRESSED_RGBA8_ETC2_EAC)
	EAC //(GL_COMPRESSED_R11_EAC/GL_COMPRESSED_RG11_EAC)

	BPTC //BPTC (GL_COMPRESSED_RGBA_BPTC_UNORM)

	PVRTC14 //PowerVR Texture compression (GL_COMPRESSED_RGB_PVRTC_4BPPV1_IMG/GL_COMPRESSED_RGBA_PVRTC_4BPPV1_IMG)
	PVRTC24 //PowerVR Texture compression (GL_COMPRESSED_RGBA_PVRTC_4BPPV2_IMG)

	ASTC //ASTC Texture Compression (GL_COMPRESSED_RGBA_ASTC_4x4_KHR)

	AMD //AMD Texture Compression (GL_ATC_RGB_AMD/GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD)

	FXT1 //(GL_COMPRESSED_RGB_FXT1_3DFX)

)

// Data holds image data.
type Data interface {

	// TextureData returns a byte slice of texture data
	// in the returned format. The returned format
	// must appear in the list of formats provided to
	// this function. Otherwise Texture should return
	// an error.
	TextureData(...Format) (Format, []byte, error)

	// TextureSize returns the size of the texture.
	TextureSize() (int, int)
}
