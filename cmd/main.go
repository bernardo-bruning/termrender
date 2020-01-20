package main

import (
	"image"
	"image/color"
	"math/rand"

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
	for !win.Closed() {
		for i := 0; i < 200; i++ {
			triangle := render.NewRandTriangle(0, 800)
			r := uint8(rand.Intn(255))
			g := uint8(rand.Intn(255))
			b := uint8(rand.Intn(255))
			triangle.Draw(img, color.RGBA{R: r, G: g, B: b})
		}
		pixel.Render(win, img)
		img = image.NewRGBA(image.Rect(0, 0, int(win.Bounds().W()), int(win.Bounds().H())))
	}
}

func main() {
	pixelgl.Run(run)
}
