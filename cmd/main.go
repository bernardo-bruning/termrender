package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	canvas := pixel.NewCanvas()
	for !canvas.Win.Closed() {
		for i := 0; i < 5; i++ {
			triangle := render.NewRandTriangle(0, 800)
			r := uint8(rand.Intn(255))
			g := uint8(rand.Intn(255))
			b := uint8(rand.Intn(255))
			triangle.Draw(canvas, color.RGBA{R: r, G: g, B: b})
		}
		pixel.Render(canvas)
		time.Sleep(time.Second * 10)
		//time.Sleep(time.Second * 5)
	}
}

func main() {
	pixelgl.Run(run)
}
