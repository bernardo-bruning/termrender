package main

import (
	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/termui"
)

func main() {
	termui.Init()
	defer termui.Close()
	canvas := termui.NewCanvas()
	triangleA := render.NewTriangle(render.Vector{10, 10}, render.Vector{15, 10}, render.Vector{10, 30})
	triangleB := render.NewTriangle(render.Vector{20, 20}, render.Vector{42, 24}, render.Vector{25, 39})
	triangleA.Draw(canvas, 1)
	triangleB.Draw(canvas, 1)

	termui.Render(canvas)
}
