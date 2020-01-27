package render

import (
	"image/color"
	"image/draw"
	"math/rand"
)

type Mesh struct {
	Triangles []Triangle
}

func NewMesh(triangles []Triangle) Mesh {
	return Mesh{Triangles: triangles}
}

func (m Mesh) Add(v Vector) Mesh {
	for i := range m.Triangles {
		m.Triangles[i].Add(v)
	}
	return m
}

func (m Mesh) Draw(dst draw.Image) {
	for _, triangle := range m.Triangles {
		r := uint8(rand.Intn(255))
		g := uint8(rand.Intn(255))
		b := uint8(rand.Intn(255))
		color := color.RGBA{r, g, b, 0xff}
		triangle.Draw(dst, color)
	}
}
