package main

import (
	"image"
	"math"
)

type Vector struct {
	X, Y float64
}

func NewVector(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (source Vector) Sub(target Vector) Vector {
	return NewVector(source.X-target.X, source.Y-target.Y)
}

func (target Vector) Len() float64 {
	return math.Sqrt(math.Abs(target.X*target.X + target.Y*target.Y))
}

func (target Vector) Normalize() Vector {
	len := target.Len()
	return NewVector(target.X/len, target.Y/len)
}

func (target Vector) Invert() Vector {
	return NewVector(target.Y, target.X)
}

func (target Vector) Min(source Vector) Vector {
	return NewVector(math.Min(source.X, target.X), math.Min(source.Y, target.Y))
}

func (target Vector) Max(source Vector) Vector {
	return NewVector(math.Max(source.X, target.X), math.Max(source.Y, target.Y))
}

func (target *Vector) Diff(source Vector) Vector {
	return target.Max(source).Sub(target.Min(source))
}

func (v *Vector) ToPointer() image.Point {
	return image.Pt(int(v.X), int(v.Y))
}
