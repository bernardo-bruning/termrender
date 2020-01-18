package render

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

func (source Vector) Add(target Vector) Vector {
	return NewVector(source.X+target.X, source.Y+target.Y)
}

func (source Vector) MulEscalar(escalar float64) Vector {
	return source.Mul(NewVector(escalar, escalar))
}

func (source Vector) DivEscalar(escalar float64) Vector {
	return source.Div(NewVector(escalar, escalar))
}

func (source Vector) Dot(target Vector) float64 {
	vector := source.Mul(target)
	return vector.X + vector.Y
}

func (source Vector) Mul(target Vector) Vector {
	return NewVector(source.X*target.X, source.Y*target.Y)
}

func (source Vector) Div(target Vector) Vector {
	return NewVector(source.X/target.Y, source.Y/target.Y)
}

func (source Vector) Angle(target Vector) float64 {
	return source.Dot(target) / (source.Len() * target.Len())
}

func (target Vector) Len() float64 {
	return math.Sqrt(math.Abs(target.X*target.X + target.Y*target.Y))
}

func (target Vector) Unit() Vector {
	len := target.Len()
	if len == 0 {
		return NewVector(0, 0)
	}

	return target.DivEscalar(len)
}

func (target Vector) Normalize() Vector {
	len := target.Len()
	return target.Div(NewVector(len, len))
}

func (target Vector) NormalizeY() Vector {
	return target.Div(NewVector(target.Y, target.Y))
}

func (target Vector) NormalizeX() Vector {
	return target.Div(NewVector(target.X, target.X))
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

func (source Vector) Cross(target Vector) float64 {
	return source.X*target.Y - source.Y*target.Y
}
