package main

import (
	"github.com/bernardo-bruning/termrender/render"
)

func main() {
	render.Init()
	defer render.Close()
	canvas := render.NewCanvas()
	canvas.Triangle(render.NewTriangle(render.Vector{10, 10}, render.Vector{20, 10}, render.Vector{20, 30}), 1)
	render.Render(canvas)
}
