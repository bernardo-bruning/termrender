package render_test

import (
	"image"
	"testing"

	"github.com/bernardo-bruning/termrender/render"
	"golang.org/x/image/colornames"
)

func BenchmarkRasterize(b *testing.B) {
	canvas := image.NewRGBA(image.Rect(0, 0, 800, 800))
	triangle := render.NewRandTriangle(0, 800)

	b.Run("RasterizeByIntersectionParallel", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			triangle.RasterizeByIntersectionParallel(canvas, colornames.Red)
		}
	})

	b.Run("RasterizeByIntersection", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			triangle.RasterizeByIntersection(canvas, colornames.Red)
		}
	})

	b.Run("RasterizeByLine", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			triangle.RasterizeByLine(canvas, colornames.Red)
		}
	})
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
