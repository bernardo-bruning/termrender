package render

import "image/color"

type Line struct {
	source, target, vector, normalization, veticalNormalization, horizontalNormalization Vector
}

func NewLine(source, target Vector) Line {
	vector := target.Sub(source)
	normalization := vector.Normalize()
	veticalNormalization := vector.NormalizeY()
	horizontalNormalization := vector.NormalizeX()
	return Line{
		vector:                  vector,
		source:                  source,
		target:                  target,
		normalization:           normalization,
		veticalNormalization:    veticalNormalization,
		horizontalNormalization: horizontalNormalization,
	}
}

func calculate(line Line) Line {
	line.vector = line.target.Sub(line.source)
	line.normalization = line.vector.Normalize()
	line.veticalNormalization = line.vector.NormalizeY()
	line.horizontalNormalization = line.vector.NormalizeX()
	return line
}

func (line Line) Normalize() Vector {
	return line.normalization
}

func (line Line) Len() float64 {
	return line.vector.Len()
}

func (line Line) LenVertical() float64 {
	return line.vector.Y
}

func (line Line) LenHorizontal() float64 {
	return line.vector.X
}

func (line Line) Next(current Vector) Vector {
	return current.Add(line.normalization)
}

func (line Line) NextVertical(current Vector) Vector {
	return current.Add(line.veticalNormalization)
}

func (line Line) NextHorizontal(current Vector) Vector {
	return current.Add(line.horizontalNormalization)
}

func (line Line) Draw(canvas Canvas, color color.Color) {
	len := line.Len()
	position := line.source
	canvas.SetPoint(position.ToPointer().X, position.ToPointer().Y, color)
	for i := 0.; i < len-1; i++ {
		position = line.Next(position)
		canvas.SetPoint(position.ToPointer().X, position.ToPointer().Y, color)
	}
}
