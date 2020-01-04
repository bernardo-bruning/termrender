package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

type Triangle struct {
	a Vector
	b Vector
	c Vector
}

func main() {
	if err := ui.Init(); err != nil {
		log.Printf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	canvas := NewCanvas()
	//canvas.debugger = true

	//canvas.Line(Vector{30, 50}, Vector{33, 13}, 4)
	//canvas.Line(Vector{33, 13}, Vector{20, 40}, 4)
	//canvas.Line(Vector{10, 10}, Vector{30, 40}, 4)
	canvas.Triangle(Triangle{Vector{10, 10}, Vector{20, 10}, Vector{20, 30}}, 1)
	canvas.Triangle(Triangle{Vector{30, 50}, Vector{30, 13}, Vector{20, 40}}, 13)

	ui.Render(canvas.inner)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
