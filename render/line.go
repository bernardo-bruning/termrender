package render

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
