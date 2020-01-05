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

func (canvas *Canvas) lineSweeping(source, target Vector, alpha, beta Line, color int) (Vector, Vector) {
	len := beta.LenVertical()

	if len == 0 {
		target = beta.target
		canvas.Line(source, target, color)
		return source, target
	}

	for y := 0.; y < len; y++ {
		source = alpha.NextVertical(source)
		target = beta.NextVertical(target)
		canvas.Line(source, target, color)
	}

	return source, target
}

func (canvas *Canvas) Triangle(triangle Triangle, color int) {

	a, b, c := sortVectorsByY(triangle)

	alpha := NewLine(a, c)
	beta := NewLine(a, b)
	teta := NewLine(b, c)

	source := a
	target := a
	source, target = canvas.lineSweeping(source, target, alpha, beta, color)
	canvas.lineSweeping(source, target, alpha, teta, color)

	canvas.Line(a, b, color)
	canvas.Line(b, c, color)
	canvas.Line(c, a, color)
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
