package termui

import (
	"github.com/bernardo-bruning/termrender/render"
	ui "github.com/gizak/termui/v3"
)

type CanvasTermUI struct {
	inner *ui.Canvas
}

func NewCanvas() render.Canvas {
	c := ui.NewCanvas()
	c.SetRect(0, 0, 132, 100)
	return &CanvasTermUI{inner: c}
}

func (canvas *CanvasTermUI) size() render.Vector {
	size := canvas.inner.GetRect().Size()
	return render.NewVector(float64(size.X), float64(size.Y))
}

func (canvas *CanvasTermUI) SetPoint(position render.Vector, color int) {
	pointer := position.ToPointer()
	canvas.inner.SetPoint(pointer, ui.Color(color))
}
