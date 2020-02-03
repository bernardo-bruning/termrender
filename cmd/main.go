package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/bernardo-bruning/termrender/loader/obj"
	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/pixel"
	p "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Term render",
		Bounds: p.R(0, 0, 800, 500),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, int(win.Bounds().W()), int(win.Bounds().H())))
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	pathTexture := os.Args[2]
	fileTexture, err := os.Open(pathTexture)
	if err != nil {
		panic(err)
	}

	texture, err := png.Decode(fileTexture)
	if err != nil {
		panic(err)
	}

	model, err := obj.Load(file)
	fmt.Println("Numbers of triangles", len(model.Triangles))
	if err != nil {
		panic(err)
	}

	fps := time.Now()
	fpsIterator := 0
	for !win.Closed() {
		model = model.RotateY(0.01)
		mesh := model.Mul(render.NewVector(200, -200, 200)).Add(render.NewVector(400, 400, 0))
		mesh.DrawWithTexture(img, texture)
		pixel.Render(win, img)
		img = image.NewRGBA(image.Rect(0, 0, int(win.Bounds().W()), int(win.Bounds().H())))
		for x := 0; x < img.Bounds().Dx(); x++ {
			for y := 0; y < img.Bounds().Dy(); y++ {
				img.Set(x, y, colornames.White)
			}
		}
		fpsIterator++
		if time.Now().Sub(fps) > time.Second {
			fmt.Printf("fps: %d\n", fpsIterator)
			fpsIterator = 0
			fps = time.Now()
		}
	}
}

func main() {
	pixelgl.Run(run)
}
