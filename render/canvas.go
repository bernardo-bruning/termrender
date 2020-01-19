package render

import "image/color"

type Canvas interface {
	SetPoint(position Vector, color color.Color)
}
