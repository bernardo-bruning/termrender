package main

import (
	"fmt"
	"image"
	"os"
	"time"

	"github.com/bernardo-bruning/termrender/loader/obj"
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
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	mesh, err := obj.Load(file)
	mesh = mesh.Mul(render.NewVectorFromScalar(-4)).Add(render.NewVector(400, 400, 0))
	fmt.Println("Numbers of triangles", len(mesh.Triangles))
	if err != nil {
		panic(err)
	}

	fps := time.Now()
	fpsIterator := 0
	for !win.Closed() {
		mesh.Draw(img)
		pixel.Render(win, img)
		img = image.NewRGBA(image.Rect(0, 0, int(win.Bounds().W()), int(win.Bounds().H())))
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
