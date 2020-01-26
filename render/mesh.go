package render

import (
	"image/color"
	"image/draw"
)

type Mesh struct {
	triangles []Triangle
}

func (m Mesh) Draw(dst draw.Image) {
	for _, triangle := range m.triangles {
		triangle.Draw(dst, color.White)
	}
}
