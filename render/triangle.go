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

func (canvas *Canvas) lineSweeping(line Line, alpha, beta Line, color int) Line {
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

func (triangle Triangle) Draw(canvas *Canvas, color int) {

	a, b, c := sortVectorsByY(triangle)

	alpha := NewLine(a, c)
	beta := NewLine(a, b)
	teta := NewLine(b, c)

	line := NewLine(a, a)
	line = canvas.lineSweeping(line, alpha, beta, color)
	canvas.lineSweeping(line, alpha, teta, color)

	alpha.Draw(canvas, color)
	beta.Draw(canvas, color)
	teta.Draw(canvas, color)
}
