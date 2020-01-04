package render

type Triangle struct {
	a Vector
	b Vector
	c Vector
}

func NewTriangle(a, b, c Vector) Triangle {
	return Triangle{a, b, c}
}
