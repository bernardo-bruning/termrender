package render

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
)

type Canvas struct {
	inner    *ui.Canvas
	debugger bool
}

func NewCanvas() Canvas {
	c := ui.NewCanvas()
	c.SetRect(0, 0, 132, 100)
	return Canvas{inner: c}
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

func (canvas *Canvas) Triangle(triangle Triangle, color int) {
	a, b, c := sortVectorsByY(triangle)

	canvas.Line(a, b, 1)
	canvas.Line(b, c, 2)
	canvas.Line(c, a, 3)
}

func (canvas *Canvas) Size() Vector {
	size := canvas.inner.GetRect().Size()
	return NewVector(float64(size.X), float64(size.Y))
}

func (canvas *Canvas) SetPoint(position Vector, color int) {
	if canvas.debugger {
		fmt.Println(position)
		return
	}
	pointer := position.ToPointer()
	canvas.inner.SetPoint(pointer, ui.Color(color))
}

func (canvas *Canvas) Line(source, target Vector, color int) {
	line := NewLine(source, target)
	len := line.Len()
	position := source
	for i := 0.; i < len; i++ {
		position = line.Next(position)
		canvas.SetPoint(position, color)
	}
}
