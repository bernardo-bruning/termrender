package render_test

import "testing"

import "github.com/bernardo-bruning/termrender/render"

import "math"

func TestAngle(t *testing.T) {
	scenarios := []struct {
		a render.Vector
		b render.Vector
		expected float64
	} {
		{
			a: render.NewVector(1, 1),
			b: render.NewVector(10, 10),
			expected: 1,
		},
	}

	for _, scenario := range scenarios {
		a := scenario.a
		b := scenario.b
		result := a.Angle(b)
		expected := 1.

		if !equalFloat(result, expected) {
			t.Fatalf("Expected %f but got %f", expected, result)
		}
	}
}

func equalFloat(a, b float64) bool {
	epsilon := 0.00000001
	return math.Abs(a - b) < epsilon && math.Abs(b - a) < epsilon
}
