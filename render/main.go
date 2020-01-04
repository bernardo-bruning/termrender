package render

import (
	ui "github.com/gizak/termui/v3"
	"log"
)

func Render(canvas Canvas) {
	ui.Render(canvas.inner)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
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
