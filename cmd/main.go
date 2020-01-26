package main

import (
	"image"

	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/pixel"
	p "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Term render",
		Bounds: p.R(0, 0, 800, 800),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, int(win.Bounds().W()), int(win.Bounds().H())))

	triangles := make([]render.Triangle, 100)
	for i := 0; i < 1000; i++ {
		triangle := render.NewRandTriangle(0, 800)
		triangles = append(triangles, triangle)
	}

	mesh := render.NewMesh(triangles)

	for !win.Closed() {
		mesh.Draw(img)
		pixel.Render(win, img)
		img = image.NewRGBA(image.Rect(0, 0, int(win.Bounds().W()), int(win.Bounds().H())))
	}
}

func main() {
	pixelgl.Run(run)
}
