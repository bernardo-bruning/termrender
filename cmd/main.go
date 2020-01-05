package main

import (
	"github.com/bernardo-bruning/termrender/render"
)

func main() {
	render.Init()
	defer render.Close()
	canvas := render.NewCanvas()
	triangleA := render.NewTriangle(render.Vector{10, 10}, render.Vector{15, 10}, render.Vector{10, 30})
	triangleB := render.NewTriangle(render.Vector{20, 20}, render.Vector{42, 24}, render.Vector{25, 39})
	triangleA.Draw(canvas, 1)
	triangleB.Draw(canvas, 1)

	render.Render(canvas)
}
