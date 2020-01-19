package main

import (
	"math/rand"
	"time"

	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/termui"
	ui "github.com/gizak/termui/v3"
)

func main() {
	defer termui.Close()
	closed := false
	termui.Init()

	go func() {
		for e := range ui.PollEvents() {
			if e.Type == ui.KeyboardEvent {
				closed = true
				break
			}
		}
	}()

	for !closed {
		canvas := termui.NewCanvas()
		for i := 0; i < 2000; i++ {
			triangle := render.NewRandTriangle(0, 100)
			triangle.Draw(canvas, rand.Intn(10))
		}

		termui.Render(canvas)
		select {
		case <-time.After(time.Millisecond):
			continue
			//default: continue
		}
	}
}
