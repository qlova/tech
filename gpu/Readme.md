# GPU
A high-level Go module for cross-platform GPU operations.

## Shaders
Shaders are written in Go using DSL that will be translated
into the appropriate shading language for any supported gpu drivers, currently bare-bones GLSL (100/110) is supported.

(Check out the source for the `gpu.Textured` type, which implements
a basic texturing shader)


## Wishlist / Roadmap
Things that we would love to see in the future with this module:  
(please contribute)

* Vulkan driver
* Order-independent transparency
* WebGPU driver
* GPU Skeletal Animations
* CPU/GPU Frustum Culling
* Proper shadow rendering
* Transparency Handling
* Reflections
* VR/AR support
* Raytracing