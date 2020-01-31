package render

import (
	"image"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func NewVectorFromScalar(scalar float64) Vector {
	return Vector{X: scalar, Y: scalar, Z: scalar}
}

func NewVector(x float64, y float64, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}

func (source Vector) Sub(target Vector) Vector {
	source.X -= target.X
	source.Y -= target.Y
	source.Z -= target.Z

	return source
}

func (source Vector) Add(target Vector) Vector {
	return NewVector(source.X+target.X, source.Y+target.Y, source.Z+target.Z)
}

func (source Vector) MulScalar(Scalar float64) Vector {
	return source.Mul(NewVectorFromScalar(Scalar))
}

func (source Vector) DivScalar(Scalar float64) Vector {
	return source.Div(NewVectorFromScalar(Scalar))
}

func (source Vector) Dot(target Vector) float64 {
	vector := source.Mul(target)
	return vector.X + vector.Y + vector.Z
}

func (source Vector) Mul(target Vector) Vector {
	return NewVector(source.X*target.X, source.Y*target.Y, source.Z*target.Z)
}

func (source Vector) Div(target Vector) Vector {
	if target.X != 0 {
		source.X /= target.X
	} else {
		source.X = 0
	}
	if target.Y != 0 {
		source.Y /= target.Y
	} else {
		source.Y = 0
	}
	if target.Z != 0 {
		source.Z /= target.Z
	} else {
		source.Z = 0
	}
	return source
}

func (source Vector) Angle(target Vector) float64 {
	return source.Dot(target) / (source.Len() * target.Len())
}

func (target Vector) Len() float64 {
	return math.Sqrt(math.Abs(target.Dot(target)))
}

func (target Vector) Unit() Vector {
	len := target.Len()
	if len == 0 {
		return NewVectorFromScalar(0)
	}

	return target.DivScalar(len)
}

func (target Vector) Normalize() Vector {
	len := target.Len()
	return target.Div(NewVectorFromScalar(len))
}

func (target Vector) NormalizeY() Vector {
	return target.Div(NewVectorFromScalar(target.Y))
}

func (target Vector) NormalizeX() Vector {
	return target.Div(NewVectorFromScalar(target.X))
}

func (target Vector) Invert() Vector {
	return NewVector(target.Y, target.X, target.Z)
}

func (target Vector) Min(source Vector) Vector {
	return NewVector(math.Min(source.X, target.X), math.Min(source.Y, target.Y), math.Min(source.Z, target.Z))
}

func (target Vector) Max(source Vector) Vector {
	return NewVector(math.Max(source.X, target.X), math.Max(source.Y, target.Y), math.Max(source.Z, target.Z))
}

func (target *Vector) Diff(source Vector) Vector {
	return target.Max(source).Sub(target.Min(source))
}

func (v *Vector) ToPointer() image.Point {
	return image.Pt(int(v.X), int(v.Y))
}

func (v Vector) RotateY(rotation float64) Vector {
	v.X = v.X*math.Cos(rotation) + v.Z*math.Sin(rotation)
	v.Z = -v.X*math.Sin(rotation) + v.Z*math.Cos(rotation)
	return v
}

func (v Vector) RotateX(rotation float64) Vector {
	v.Y = v.Y*math.Cos(rotation) - v.Z*math.Sin(rotation)
	v.Z = v.Y*math.Sin(rotation) + v.Z*math.Cos(rotation)
	return v
}

func (v Vector) RotateZ(rotation float64) Vector {
	v.X = v.X*math.Cos(rotation) - v.Y*math.Sin(rotation)
	v.Y = v.X*math.Sin(rotation) + v.Y*math.Cos(rotation)
	return v
}

func (source Vector) Cross(target Vector) Vector {
	source.X = source.Y*target.Z - target.Z - source.Y
	source.Y = source.Z*target.X - target.Z - source.X
	source.Z = source.X*target.Y - target.X - source.Y
	return source
}
