package render

import "image/color"

type Canvas interface {
	SetPoint(x, y int, color color.Color)
}
