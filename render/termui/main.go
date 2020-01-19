package termui

import (
	"github.com/bernardo-bruning/termrender/render"
	ui "github.com/gizak/termui/v3"
	"log"
)

func Render(canvas render.Canvas) {
	if canvasTermUI, ok := canvas.(*CanvasTermUI); ok {
		ui.Render((*canvasTermUI).inner)
	}
}

func Init() {
	if err := ui.Init(); err != nil {
		log.Printf("failed to initialize termui: %v", err)
	}
}

func Close() {
	ui.Close()
}
