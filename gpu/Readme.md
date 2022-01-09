# GPU

A high-level Go module for high-performance GPU operations.
This package currently is being developed against OpenGL 4.6 
\+ `ARB_bindless_texture`, however this package 
has been built so that other GPU drivers can be developed.

This package aims to be as performant as possible by
minimizing the number of C drawcalls being made (multidraw 
indirect) and aims to use cache-friendly data structures and 
patterns. At the same-time, this package is meant to be 
easier to work with than low-level GPU drivers such as 
`OpenGL`/`Vulkan`. Common 3D workflows are
standardised and opionated.

## Shaders
Shaders are written in Go using a DSL that will be translated
into the appropriate shading language for any supported gpu drivers, currently bare-bones GLSL is supported.

(Check out the source for the `gpu.Textured` type, which implements
a basic texturing shader)


## Wishlist / Roadmap
Things that we would love to see in the future with this module:  
(please contribute)

* Vulkan driver
* WebGL/WebGPU driver
* OpenGLES driver
* GPU Skeletal Animations
* CPU/GPU Frustum Culling
* Proper shadow rendering
* Transparency Handling