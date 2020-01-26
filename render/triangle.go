package render

import (
	"image/color"
	"image/draw"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Triangle struct {
	a  Vector
	b  Vector
	c  Vector
	ab Vector
	ac Vector
}

var rander *rand.Rand = rand.New(rand.NewSource(time.Now().Unix()))

func random(min, max float64) float64 {
	return min + rander.Float64()*(max-min)
}

func NewRandTriangle(min, max float64) Triangle {
	return NewTriangle(
		Vector{random(min, max), random(min, max), random(min, max)},
		Vector{random(min, max), random(min, max), random(min, max)},
		Vector{random(min, max), random(min, max), random(min, max)},
	)
}

func NewTriangle(a, b, c Vector) Triangle {
	return Triangle{a: a, b: b, c: c, ab: b.Sub(a), ac: c.Sub(a)}
}

func swap(v1 Vector, v2 Vector) (Vector, Vector) {
	return v2, v1
}

func sortVectorsByY(triangle Triangle) (Vector, Vector, Vector) {
	a := triangle.a
	b := triangle.b
	c := triangle.c

	if a.Y > b.Y {
		a, b = swap(a, b)
	}

	if a.Y > c.Y {
		a, c = swap(a, c)
	}

	if b.Y > c.Y {
		b, c = swap(b, c)
	}

	return a, b, c
}

func (line Line) lineSweeping(canvas draw.Image, alpha, beta Line, color color.Color) Line {
	len := beta.LenVertical()

	if len == 0 {
		line.target = beta.target
		line.Draw(canvas, color)
		return line
	}

	for y := 0.; y < len; y++ {
		line.source = alpha.NextVertical(line.source)
		line.target = beta.NextVertical(line.target)
		line = calculate(line)
		line.Draw(canvas, color)
	}

	return line
}

func (triangle *Triangle) Barycentric(point Vector) Vector {
	v0 := triangle.ab
	v1 := triangle.ac
	v2 := point.Sub(triangle.a)

	v0.Z = 0
	v1.Z = 0
	v2.Z = 0

	d00 := v0.Dot(v0)
	d01 := v0.Dot(v1)
	d11 := v1.Dot(v1)
	d20 := v2.Dot(v0)
	d21 := v2.Dot(v1)

	denom := 1 / (d00*d11 - d01*d01)
	v := (d11*d20 - d01*d21) * denom
	w := (d00*d21 - d01*d20) * denom
	u := 1 - v - w
	return Vector{X: v, Y: w, Z: u}
}

func (triangle *Triangle) Intersection(point Vector) bool {
	b := triangle.Barycentric(point)
	return b.X >= 0 && b.Y >= 0 && b.Z >= 0
}

func (triangle Triangle) RasterizeByLine(canvas draw.Image, color color.Color) {
	a, b, c := sortVectorsByY(triangle)

	alpha := NewLine(a, c)
	beta := NewLine(a, b)
	teta := NewLine(b, c)

	line := NewLine(a, a)
	line = line.lineSweeping(canvas, alpha, beta, color)
	line.lineSweeping(canvas, alpha, teta, color)

	alpha.Draw(canvas, color)
	beta.Draw(canvas, color)
	teta.Draw(canvas, color)
}

func (triangle Triangle) RasterizeByIntersection(canvas draw.Image, color color.Color) {
	start := triangle.a.Min(triangle.b).Min(triangle.c)
	end := triangle.a.Max(triangle.b).Max(triangle.c)

	for x := start.X; x <= end.X; x++ {
		for y := start.Y; y <= end.Y; y++ {
			point := Vector{X: x, Y: y, Z: 0}
			if triangle.Intersection(point) {
				canvas.Set(point.ToPointer().X, point.ToPointer().Y, color)
			}
		}
	}
}

func (triangle Triangle) RasterizeByIntersectionParallel(canvas draw.Image, color color.Color) {
	start := triangle.a.Min(triangle.b).Min(triangle.c)
	end := triangle.a.Max(triangle.b).Max(triangle.c)
	var wg sync.WaitGroup
	numCpu := runtime.NumCPU()
	batch := int(end.X) - int(start.X)/numCpu
	for i := 0; i < numCpu; i++ {
		wg.Add(1)
		go func(i int) {
			for x := start.X + float64(batch*i); x <= end.X+float64(batch*(i+1)); x++ {
				for y := start.Y; y <= end.Y; y++ {
					point := Vector{X: x, Y: y, Z: 0}
					if triangle.Intersection(point) {
						canvas.Set(point.ToPointer().X, point.ToPointer().Y, color)
					}
				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func (triangle Triangle) Draw(canvas draw.Image, color color.Color) {
	triangle.RasterizeByIntersection(canvas, color)
}
