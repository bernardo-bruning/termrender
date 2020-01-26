package render

import (
	"image/color"
	"image/draw"
	"math/rand"
)

type Mesh struct {
	triangles []Triangle
}

func NewMesh(triangles []Triangle) Mesh {
	return Mesh{triangles: triangles}
}

func (m Mesh) Draw(dst draw.Image) {
	for _, triangle := range m.triangles {
		r := uint8(rand.Intn(255))
		g := uint8(rand.Intn(255))
		b := uint8(rand.Intn(255))
		triangle.Draw(dst, color.RGBA{R: r, G: g, B: b})
	}
}
