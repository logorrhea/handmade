package main

import (
	"azul3d.org/gfx.v1"
	"azul3d.org/gfx/window.v2"
	"image"
)

func gfxloop(w window.Window, r gfx.Renderer) {

	for {

		// Clear the entire area (empty rectangle means "the whole area")
		r.Clear(image.Rect(0, 0, 0, 0), gfx.Color{0, 0, 0, 1})

		// Render hte whole frame
		r.Render()
	}

}

func main() {
	window.Run(gfxloop, nil)
}
