package render_test

import (
	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/termui"
	"testing"
)

func BenchmarkTriangle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		canvas := termui.NewCanvas()
		triangle := render.NewTriangle(render.Vector{10, 10}, render.Vector{15, 10}, render.Vector{10, 30})
		triangle.Draw(canvas, 1)
	}
}
