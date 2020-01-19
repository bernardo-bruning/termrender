package main

import (
	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/termui"
)

func main() {
	termui.Init()
	defer termui.Close()
	canvas := termui.NewCanvas()
	triangleA := render.NewTriangle(render.Vector{10, 10, 1}, render.Vector{15, 10, 1}, render.Vector{10, 30, 1})
	triangleB := render.NewTriangle(render.Vector{20, 20, 1}, render.Vector{42, 24, 1}, render.Vector{25, 39, 1})
	triangleA.RasterizeByIntersection(canvas, 1)
	triangleB.RasterizeByIntersection(canvas, 1)

	termui.Render(canvas)
}
