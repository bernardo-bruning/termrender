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
	zbuffer := make([][]bool, dst.Bounds().Dx())
	for i := range zbuffer {
		zbuffer[i] = make([]bool, dst.Bounds().Dy())
	}

	for _, triangle := range m.triangles {
		r := uint8(rand.Intn(255))
		g := uint8(rand.Intn(255))
		b := uint8(rand.Intn(255))
		color := color.RGBA{r, g, b, 0xff}
		bound := triangle.Bounds()
		for x := bound.Min.X; x < bound.Max.X; x++ {
			for y := bound.Min.Y; y < bound.Max.Y; y++ {
				if zbuffer[x][y] {
					continue
				}

				point := Vector{X: float64(x), Y: float64(y)}
				if triangle.Intersection(point) {
					dst.Set(x, y, color)
					zbuffer[x][y] = true
				}
			}
		}
		//triangle.Draw(dst, color.RGBA{R: r, G: g, B: b})
	}
}
