package render

import (
	"image/draw"
	"math"

	"golang.org/x/image/colornames"
)

type Mesh struct {
	Triangles      []Triangle
	TextureMapping []Triangle
}

func NewMesh(triangles []Triangle) Mesh {
	return Mesh{Triangles: triangles}
}

func NewMeshWithTexture(triangles, textureMapping []Triangle) Mesh {
	return Mesh{Triangles: triangles, TextureMapping: textureMapping}
}

func (m Mesh) Add(v Vector) Mesh {
	triangles := make([]Triangle, len(m.Triangles))
	for i := range m.Triangles {
		triangles[i] = m.Triangles[i].Add(v)
	}
	m.Triangles = triangles
	return m
}

func (m Mesh) Mul(v Vector) Mesh {
	triangles := make([]Triangle, len(m.Triangles))
	for i := range m.Triangles {
		triangles[i] = m.Triangles[i].Mul(v)
	}

	m.Triangles = triangles
	return m
}

func (m Mesh) apply(f func(Triangle) Triangle) Mesh {
	triangles := make([]Triangle, len(m.Triangles))
	for i := range m.Triangles {
		triangles[i] = f(m.Triangles[i])
	}
	m.Triangles = triangles
	return m
}

func (m Mesh) RotateX(rotation float64) Mesh {
	return m.apply(func(t Triangle) Triangle {
		return t.RotateX(rotation)
	})
}

func (m Mesh) RotateY(rotation float64) Mesh {
	return m.apply(func(t Triangle) Triangle {
		return t.RotateY(rotation)
	})
}

func (m Mesh) RotateZ(rotation float64) Mesh {
	return m.apply(func(t Triangle) Triangle {
		return t.RotateZ(rotation)
	})
}

func (m Mesh) Draw(dst draw.Image) {
	lightZ := 400.0
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
						color.B += uint8(z/2 + lightZ)
						dst.Set(point.ToPointer().X, point.ToPointer().Y, color)
						zbuffer[y+x*dst.Bounds().Dy()] = z
					}
				}
			}
		}
	}
}
