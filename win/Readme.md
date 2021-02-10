# Win

A Go module for creating surfaces/windows for GPU rendering.

## Wishlist:
Things that we would love to see in the future with this module:  
(please contribute)

* HTML5 canvas driver.
* Android/IOS mobile driver.

## HelloCube Example

```go
package main

import (
	"log"

	"qlova.tech/gpu"
	"qlova.tech/win"

	"qlova.tech/gpu/models"

	_ "qlova.tech/gpu/driver/opengl"
	_ "qlova.tech/win/driver/glfw"
)

func main() {
	if err := win.Open(); err != nil {
		log.Fatalln("could not open a window: ", err)
	}

	if err := gpu.Open(); err != nil {
		log.Fatalln("could not open the gpu: ", err)
	}

	cube := gpu.NewMesh(models.Cube()).Model()

	if err := gpu.Upload(); err != nil {
		log.Fatalln("gpu upload failed: ", err)
	}

	gpu.Viewport = gpu.Position(3, 3, 3).LookAt(gpu.Position(0, 0, 0))

	var t gpu.Transform
	for win.Update() {
		gpu.DrawModel(&cube, &t)

		if err := gpu.Sync(); err != nil {
			log.Fatalln("there was an error syncing the gpu: ", err)
		}
	}
}
```

**License**  
This work is subject to the terms of the Qlova Public
License, Version 2.0. If a copy of the QPL was not distributed with this
work, You can obtain one at https://license.qlova.org/v2

The QPL is compatible with the AGPL which is why both licenses are provided within this repository.