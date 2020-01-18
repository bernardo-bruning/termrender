package render

type Triangle struct {
	a Vector
	b Vector
	c Vector
}

func NewTriangle(a, b, c Vector) Triangle {
	return Triangle{a, b, c}
}

func swap(v1 Vector, v2 Vector) (Vector, Vector) {
	return v2, v1
}

func sortVectorsByY(triangle Triangle) (Vector, Vector, Vector) {
	a := triangle.a
	b := triangle.b
	c := triangle.c

	if a.Y > b.Y {
		a, b = swap(a, b)
	}

	if a.Y > c.Y {
		a, c = swap(a, c)
	}

	if b.Y > c.Y {
		b, c = swap(b, c)
	}

	return a, b, c
}

func (line Line) lineSweeping(canvas Canvas, alpha, beta Line, color int) Line {
	len := beta.LenVertical()

	if len == 0 {
		line.target = beta.target
		line.Draw(canvas, color)
		return line
	}

	for y := 0.; y < len; y++ {
		line.source = alpha.NextVertical(line.source)
		line.target = beta.NextVertical(line.target)
		line = calculate(line)
		line.Draw(canvas, color)
	}

	return line
}

func (triangle Triangle) Barycentric(point Vector) Vector {
	v0 := triangle.b.Sub(triangle.a)
	v1 := triangle.c.Sub(triangle.a)
	v2 := point.Sub(triangle.a)

	d00 := v0.Dot(v0)
	d01 := v0.Dot(v1)
	d11 := v1.Dot(v1)
	d20 := v2.Dot(v0)
	d21 := v2.Dot(v1)

	denom := 1 / (d00*d11 - d01*d01)
	v := (d11*d20 - d01*d21) * denom
	w := (d00*d21 - d01*d20) * denom
	u := 1 - v - w
	return Vector{X: v, Y: w, Z: u}
}

func (triangle Triangle) Intersection(point Vector) bool {
	b := triangle.Barycentric(point)
	return b.X >= 0 && b.Y >= 0 && b.Z >= 0
}

func (triangle Triangle) Draw(canvas Canvas, color int) {

	a, b, c := sortVectorsByY(triangle)

	alpha := NewLine(a, c)
	beta := NewLine(a, b)
	teta := NewLine(b, c)

	line := NewLine(a, a)
	line = line.lineSweeping(canvas, alpha, beta, color)
	line.lineSweeping(canvas, alpha, teta, color)

	alpha.Draw(canvas, color)
	beta.Draw(canvas, color)
	teta.Draw(canvas, color)
}
