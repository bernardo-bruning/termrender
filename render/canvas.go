package render

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
)

type Canvas struct {
	inner    *ui.Canvas
	debugger bool
}

func NewCanvas() *Canvas {
	c := ui.NewCanvas()
	c.SetRect(0, 0, 132, 100)
	return &Canvas{inner: c}
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
