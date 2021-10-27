//Package gui provides immediate-mode GUI rendering on-top of the gpu and win packages.
package gui

import (
	"fmt"

	"github.com/splizard/imgui"
	"qlova.tech/gpu"
	"qlova.tech/win"
)

var opened bool

var font struct {
	data          []byte
	width, height int32
}

type u32Colors []uint32

func (c u32Colors) Name() string {
	return "colors"
}

var mesh gpu.Mesh

//Open intialises and/or updates the GUI context, it returns false if
//the user wants to exit.
func Open() bool {
	if !opened {
		opened = true
		imgui.CreateContext(nil)
		imgui.GetIO().DisplaySize = *imgui.NewImVec2(float32(win.Width), float32(win.Height))

		imgui.GetIO().Fonts.GetTexDataAsAlpha8(&font.data, &font.width, &font.height, nil)

	} else {
		imgui.Render()

		drawdata := imgui.GetDrawData()

		if mesh.Nil() {
			//We copy everything into the gpu
			//representation, this is not very
			//efficient and we want to use the
			//gpu.Mesh type in the imgui package.

			var mesh gpu.Mesh

			var vertices gpu.Vertices
			var uvs gpu.UVs
			var colors gpu.Colors

			for _, list := range drawdata.CmdLists {
				for _, cmd := range list.VtxBuffer {
					vertices = append(vertices, list.VtxBuffer)
				}

			}

		}

	}

	imgui.NewFrame()

	return true
}

func Text(s string) {
	imgui.Text(s)
}

func Textf(s string, args ...interface{}) {
	imgui.Text(fmt.Sprintf(s, args...))
}
