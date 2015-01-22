package main

import (
	"azul3d.org/gfx.v1"
	"azul3d.org/gfx/window.v2"
	"fmt"
	"image"
)

var black = gfx.Color{0, 0, 0, 1}
var white = gfx.Color{1, 1, 1, 1}

var whiteness = true

func gfxloop(w window.Window, r gfx.Renderer) {

	// Start with white background
	r.Clear(image.Rect(0, 0, 0, 0), white)

	// Spawn event handling coroutine
	go handleEvents(w, r)

	for /*ever*/ {

		// Render the window
		r.Render()

	}

}

func handleEvents(w window.Window, r gfx.Renderer) {

	// Create our events channel
	events := make(chan window.Event, 256)

	// Notify channel when event occurs
	w.Notify(events, window.AllEvents)

	// Wait for events, react accordingly
	for event := range events {

		// React to window move and resize
		switch event.(type) {
		case window.Resized:
			if whiteness {
				fmt.Println("painting black")
				whiteness = false
				r.Clear(image.Rect(0, 0, 0, 0), black)
			} else {
				fmt.Println("painting white")
				whiteness = true
				r.Clear(image.Rect(0, 0, 0, 0), white)
			}
		case window.Moved:
			if whiteness {
				whiteness = false
				r.Clear(image.Rect(0, 0, 0, 0), black)
			} else {
				whiteness = true
				r.Clear(image.Rect(0, 0, 0, 0), white)
			}
		}

	}
}

func main() {
	window.Run(gfxloop, nil)
}
