package render

import "image/color"

type Canvas interface {
	Set(x, y int, color color.Color)
}
