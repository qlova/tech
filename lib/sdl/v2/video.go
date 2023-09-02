/*

  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.

*/

package sdl

import (
	"qlova.tech/abi"
)

type BlendMode abi.Enum

const (
	BlendOff = 0x00000000 /**< no blending dstRGBA = srcRGBA */
	Blend    = 0x00000001 /**< alpha blending dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA)) dstA = srcA + (dstA * (1-srcA)) */
	BlendAdd = 0x00000002 /**< additive blending dstRGB = (srcRGB * srcA) + dstRGB dstA = dstA */
	BlendMod = 0x00000004 /**< color modulate dstRGB = srcRGB * dstRGB dstA = dstA */
	BlemdMul = 0x00000008 /**< color multiply dstRGB = (srcRGB * dstRGB) + (dstRGB * (1-srcA)) dstA = dstA */
)

type DisplayMode struct {
	Format      abi.Uint32
	W           abi.Int
	H           abi.Int
	RefreshRate abi.Int
	DriverData  abi.UnsafePointer
}

type Renderer abi.Opaque[Renderer]

var Video struct {
	Lib

	GetRenderDrawBlendMode func(Renderer, *BlendMode) abi.Error `ffi:"SDL_GetRenderDrawBlendMode"`
}

var Surfaces struct {
	Lib

	GetBlendMode func(Surface, *BlendMode) abi.Error `ffi:"SDL_GetSurfaceBlendMode"`
}

type GraphicsLibraryAttribute abi.Enum

const (
	RedSize GraphicsLibraryAttribute = iota
	GreenSize
	BlueSize
	AlphaSize
	BufferSize
	DoubleBuffer
	DepthSize
	StencilSize
	AccumRedSize
	AccumGreenSize
	AccumBlueSize
	Stereo
	MultiSampleBuffers
	MultiSampleSamples
	AcceleratedVisual
	RetainedBacking
	ContextMajorVersion
	ContextMinorVersion
	ContextFlags
	ContextProfileMask
	ShareWithCurrentContext
	FramebufferCapableSRGB
	ContextReleaseBehavior
	ContextEGL
)
