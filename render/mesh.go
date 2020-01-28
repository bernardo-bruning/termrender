package render

import (
	"image/draw"
	"math"

	"golang.org/x/image/colornames"
)

type Mesh struct {
	Triangles []Triangle
}

func NewMesh(triangles []Triangle) Mesh {
	return Mesh{Triangles: triangles}
}

func (m Mesh) Add(v Vector) Mesh {
	for i := range m.Triangles {
		m.Triangles[i] = m.Triangles[i].Add(v)
	}
	return m
}

func (m Mesh) Mul(v Vector) Mesh {
	for i := range m.Triangles {
		m.Triangles[i] = m.Triangles[i].Mul(v)
	}
	return m
}

func (m Mesh) Draw(dst draw.Image) {
	zbuffer := make([]float64, dst.Bounds().Dx()*dst.Bounds().Dy())
	for i := range zbuffer {
		zbuffer[i] = math.Inf(-1)
	}

	for _, triangle := range m.Triangles {
		bounds := triangle.Bounds()
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
				point := Vector{X: float64(x), Y: float64(y), Z: 0}
				bc := triangle.Barycentric(point)
				z := bc.X*triangle.a.Z + bc.Y*triangle.b.Z + bc.Z*triangle.c.Z
				if bc.X >= 0 && bc.Y >= 0 && bc.Z >= 0 {
					if z > zbuffer[y+x*dst.Bounds().Dy()] {
						color := colornames.Black
						color.R += uint8(z + 400)
						dst.Set(point.ToPointer().X, point.ToPointer().Y, color)
						zbuffer[y+x*dst.Bounds().Dy()] = z
					}
				}
			}
		}
	}
}
