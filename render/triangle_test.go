package render_test

import (
	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/termui"
	"testing"
)

func BenchmarkTriangle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		canvas := termui.NewCanvas()
		triangle := render.NewTriangle(render.Vector{10, 10, 0}, render.Vector{15, 10, 0}, render.Vector{10, 30, 0})
		triangle.Draw(canvas, 1)
	}
}

func TestIntersection(t *testing.T) {
	scenaries := []struct {
		triangle render.Triangle
		point    render.Vector
		expected bool
	}{
		{
			triangle: render.NewTriangle(
				render.NewVector(0, 10, 0),
				render.NewVector(10, 0, 0),
				render.NewVector(0, 0, 0),
			),
			point:    render.NewVector(2, 2, 0),
			expected: true,
		},
		{
			triangle: render.NewTriangle(
				render.NewVector(0, 10, 0),
				render.NewVector(10, 0, 0),
				render.NewVector(0, 0, 0),
			),
			point:    render.NewVector(10, 10, 0),
			expected: false,
		},
		{
			triangle: render.NewTriangle(
				render.NewVector(0, -10, 0),
				render.NewVector(-10, 0, 0),
				render.NewVector(0, 0, 0),
			),
			point:    render.NewVector(-10, 0, 0),
			expected: true,
		},
		{
			triangle: render.NewTriangle(
				render.NewVector(0, -10, 0),
				render.NewVector(-10, 0, 0),
				render.NewVector(0, 0, 0),
			),
			point:    render.NewVector(-10, 10, 0),
			expected: false,
		},
	}

	for _, scenario := range scenaries {
		triangle := scenario.triangle
		point := scenario.point
		if triangle.Intersection(point) != scenario.expected {
			t.Errorf("Error to calculate intersection to point %v", point)
		}
	}
}
