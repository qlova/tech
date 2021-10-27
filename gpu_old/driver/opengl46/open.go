package opengl46

import (
	"errors"
	"fmt"

	"qlova.tech/gpu"
	"qlova.tech/gpu/driver/opengl"

	"github.com/go-gl/gl/v4.6-core/gl"
)

func init() {
	gpu.Register("OpenGL 4.6", gpu.Driver(func() (gpu.Context, error) {
		if err := open(); err != nil {
			return gpu.Context{}, err
		}

		return gpu.Context{
			Load: load,
			Draw: draw,
			Sync: sync,
		}, nil
	}))
}

var opened bool

//GPU limits need to be known so that we can work around them.
var (
	maxTextureSize,
	maxUniformBlockSize,
	maxVertexTextureUnits,
	maxArrayTextureLayers int32
)

//shadowMap framebuffer.
var shadowMapFBO uint32

//shadowMap texture.
var shadowMap gpu.Texture
var shadowMapTex uint32

const shadowSize = 2048

func checkErrors() {
	var count int
	var errors []uint32
	for {
		err := gl.GetError()
		if err == 0 {
			break
		}
		if err != 0 {
			count++
		}
		errors = append(errors, err)
	}

	if count > 0 {
		panic(fmt.Errorf("%v gl error(s): %v", count, errors))
	}
}

func open() error {
	//If open has already been called then this is a noop.
	if opened {
		return nil
	}

	if err := gl.Init(); err != nil {
		//If there is an error here, we should fallback to an earlier OpenGL version.

		return fmt.Errorf("gl.Init failed: %w", err)
	}

	//Query GPU limits.
	gl.GetIntegerv(gl.MAX_UNIFORM_BLOCK_SIZE, &opengl.MaxUniformBlockSize)
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &opengl.MaxTextureSize)
	gl.GetIntegerv(gl.MAX_VERTEX_UNIFORM_COMPONENTS, &opengl.MaxVertexUniformComponents)
	gl.GetIntegerv(gl.MAX_VERTEX_UNIFORM_VECTORS, &opengl.MaxVertexUniformVectors)

	opengl.Active = true

	gl.GetIntegerv(gl.MAX_ARRAY_TEXTURE_LAYERS, &maxArrayTextureLayers)
	gl.GetIntegerv(gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS, &maxVertexTextureUnits)

	gl.Enable(gl.DEPTH_TEST)

	//Create the shadowmap.
	gl.GenFramebuffers(1, &shadowMapFBO)

	var tex uint32
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA16F,
		shadowSize, shadowSize, 0, gl.RGBA, gl.FLOAT, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	//pointer := gl.GetTextureHandleARB(tex)
	//gl.MakeTextureHandleResidentARB(pointer)
	//shadowMap = Texture(pointer)
	shadowMapTex = tex

	gl.BindFramebuffer(gl.FRAMEBUFFER, shadowMapFBO)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, tex, 0)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return errors.New("FRAMEBUFFER")
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	//Set("shadowmap", shadowMap)

	return nil
}
