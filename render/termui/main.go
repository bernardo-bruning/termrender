package termui

import (
	"image/draw"
	"log"

	ui "github.com/gizak/termui/v3"
)

func Render(canvas draw.Image) {
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
