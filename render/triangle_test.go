package render_test

import (
	"github.com/bernardo-bruning/termrender/render"
	"github.com/bernardo-bruning/termrender/render/termui"
	"math/rand"
	"testing"
)

func random() float64 {
	MIN := 0.0
	MAX := 2000.0
	return MIN + rand.Float64()*(MAX-MIN)
}

func newRandTriangle() render.Triangle {
	return render.NewTriangle(
		render.Vector{random(), random(), random()},
		render.Vector{random(), random(), random()},
		render.Vector{random(), random(), random()},
	)
}

func BenchmarkRasterize(b *testing.B) {
	canvas := termui.NewCanvas()
	triangle := newRandTriangle()

	b.Run("BenchmarkRasterizeByIntersection", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			triangle.RasterizeByIntersection(canvas, 1)
		}
	})

	b.Run("RasterizeByLine", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			triangle.RasterizeByLine(canvas, 1)
		}
	})
}

func BenchmarkRasterizeByIntersection(b *testing.B) {
	for n := 0; n < b.N; n++ {
		canvas := termui.NewCanvas()
		triangle := newRandTriangle()
		triangle.RasterizeByIntersection(canvas, 1)
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
