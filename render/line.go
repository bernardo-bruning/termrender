package render

type Line struct {
	source, target, vector, normalization Vector
}

func NewLine(source, target Vector) Line {
	vector := target.Sub(source)
	return Line{
		vector:        vector,
		source:        source,
		target:        target,
		normalization: vector.Normalize(),
	}
}

func (line Line) Normalize() Vector {
	return line.normalization
}

func (line Line) Len() float64 {
	return line.vector.Len()
}

func (line Line) Next(current Vector) Vector {
	return current.Add(line.normalization)
}
