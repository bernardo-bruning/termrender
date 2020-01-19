package termui

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	ui "github.com/gizak/termui/v3"
)

type CanvasTermUI struct {
	inner *ui.Canvas
}

func NewCanvas() draw.Image {
	c := ui.NewCanvas()
	c.SetRect(0, 0, 132, 100)
	return &CanvasTermUI{inner: c}
}

func (canvas *CanvasTermUI) ColorModel() color.Model {
	return color.NRGBAModel
}

func (canvas *CanvasTermUI) Bounds() image.Rectangle {
	size := canvas.inner.GetRect().Size()
	return image.Rect(0, 0, size.X, size.Y)
}

func (canvas *CanvasTermUI) At(x, y int) color.Color {
	return canvas.inner.At(x, y)
}

func (canvas *CanvasTermUI) Set(x, y int, color color.Color) {
	pointer := image.Point{x, y}
	if pointer.X >= 0 && pointer.Y >= 0 {
		canvas.inner.SetPoint(pointer, ui.Color(rand.Intn(10)))
	}
}
