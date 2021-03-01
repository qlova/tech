package cmd

/*
Copyright (c) 2016, Iñigo López de Heredia
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the {organization} nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
*/

import (
	"image"
	"image/color"
	"time"

	_ "image/jpeg" //support jpg
	_ "image/png"  //support png

	"os"
	"syscall"
	"unsafe"

	"github.com/nsf/termbox-go"
)

func rgb(c color.Color) (uint16, uint16, uint16) {
	r, g, b, _ := c.RGBA()
	// Reduce color values to the range [0, 15]
	return uint16(r >> 8), uint16(g >> 8), uint16(b >> 8)
}

// termColor converts a 24-bit RGB color into a term256 compatible approximation.
func termColor(r, g, b uint16) uint16 {
	rterm := (((r * 5) + 127) / 255) * 36
	gterm := (((g * 5) + 127) / 255) * 6
	bterm := (((b * 5) + 127) / 255)

	return rterm + gterm + bterm + 16 + 1 // termbox default color offset
}

const defaultRatio float64 = 7.0 / 3.0 // The terminal's default cursor width/height ratio

// load an image stored in the given path
func load(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}

// canvasSize returns the terminal columns, rows, and cursor aspect ratio
func canvasSize() (int, int, float64) {
	var size [4]uint16
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdout.Fd()), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&size)), 0, 0, 0); err != 0 {
		panic(err)
	}
	rows, cols, width, height := size[0], size[1], size[2], size[3]

	var whratio = defaultRatio
	if width > 0 && height > 0 {
		whratio = float64(height/rows) / float64(width/cols)
	}

	return int(cols), int(rows), whratio
}

// scales calculates the image scale to fit within the terminal width/height
func scale(imgW, imgH, termW, termH int, whratio float64) float64 {
	hr := float64(imgH) / (float64(termH) * whratio)
	wr := float64(imgW) / float64(termW)
	return max(hr, wr, 1)
}

// imgArea calcuates the approximate rectangle a terminal cell takes up
func imgArea(termX, termY int, imgScale, whratio float64) (int, int, int, int) {
	startX, startY := float64(termX)*imgScale, float64(termY)*imgScale*whratio
	endX, endY := startX+imgScale, startY+imgScale*whratio

	return int(startX), int(startY), int(endX), int(endY)
}

// avgRGB calculates the average RGB color within the given
// rectangle, and returns the [0,1] range of each component.
func avgRGB(img image.Image, startX, startY, endX, endY int) (uint16, uint16, uint16) {
	var total = [3]uint16{}
	var count uint16
	for x := startX; x < endX; x++ {
		for y := startY; y < endY; y++ {
			if (!image.Point{x, y}.In(img.Bounds())) {
				continue
			}
			r, g, b := rgb(img.At(x, y))
			total[0] += r
			total[1] += g
			total[2] += b
			count++
		}
	}

	r := total[0] / count
	g := total[1] / count
	b := total[2] / count
	return r, g, b
}

// max returns the maximum value
func max(values ...float64) float64 {
	var m float64
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}

func draw(img image.Image) {
	// Get terminal size and cursor width/height ratio
	width, height, whratio := canvasSize()

	bounds := img.Bounds()
	imgW, imgH := bounds.Dx(), bounds.Dy()

	imgScale := scale(imgW, imgH, width, height, whratio)

	// Resize canvas to fit scaled image
	width, height = int(float64(imgW)/imgScale), int(float64(imgH)/(imgScale*whratio))

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Calculate average color for the corresponding image rectangle
			// fitting in this cell. We use a half-block trick, wherein the
			// lower half of the cell displays the character ▄, effectively
			// doubling the resolution of the canvas.
			startX, startY, endX, endY := imgArea(x, y, imgScale, whratio)

			r, g, b := avgRGB(img, startX, startY, endX, (startY+endY)/2)
			colorUp := termbox.Attribute(termColor(r, g, b))

			r, g, b = avgRGB(img, startX, (startY+endY)/2, endX, endY)
			colorDown := termbox.Attribute(termColor(r, g, b))

			termbox.SetCell(x, y, '▄', colorDown, colorUp)
		}
	}
	termbox.Flush()
}

func display(img image.Image) {
	err := termbox.Init()
	if err == nil {
		defer termbox.Close()
		termbox.SetOutputMode(termbox.Output256)
	}

	draw(img)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Ch == 'q' {
				return
			}
		case termbox.EventResize:
			draw(img)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
