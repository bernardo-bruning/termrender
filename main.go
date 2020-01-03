package main

import (
	"fmt"
	"image"
	"log"
	"math"

	ui "github.com/gizak/termui/v3"
)

type Vector struct {
	X, Y float64
}

func (source Vector) Sub(target Vector) Vector {
	return NewVector(source.X-target.X, source.Y-target.Y)
}

func NewVector(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (target Vector) Len() float64 {
	return math.Sqrt(math.Abs(target.X*target.X + target.Y*target.Y))
}

func (target Vector) Normalize() Vector {
	len := target.Len()
	return NewVector(target.X/len, target.Y/len)
}

func (target Vector) Invert() Vector {
	return NewVector(target.Y, target.X)
}

func (target Vector) Min(source Vector) Vector {
	return NewVector(math.Min(source.X, target.X), math.Min(source.Y, target.Y))
}

func (target Vector) Max(source Vector) Vector {
	return NewVector(math.Max(source.X, target.X), math.Max(source.Y, target.Y))
}

func (target *Vector) Diff(source Vector) Vector {
	return target.Max(source).Sub(target.Min(source))
}

func (v *Vector) ToPointer() image.Point {
	return image.Pt(int(v.X), int(v.Y))
}

type Triangle struct {
	a Vector
	b Vector
	c Vector
}

type Canvas struct {
	inner    *ui.Canvas
	debugger bool
}

func NewCanvas() Canvas {
	c := ui.NewCanvas()
	c.SetRect(0, 0, 132, 100)
	return Canvas{inner: c}
}

func Swap(v1 Vector, v2 Vector) (Vector, Vector) {
	return v2, v1
}

func (canvas *Canvas) Triangle(triangle Triangle, color int) {
	a := triangle.a
	b := triangle.b
	c := triangle.c

	if a.Y > b.Y {
		a, b = Swap(a, b)
	}

	if a.Y > c.Y {
		a, c = Swap(a, c)
	}

	if b.Y > c.Y {
		b, c = Swap(b, c)
	}

	canvas.Line(a, b, 1)
	canvas.Line(b, c, 2)
	canvas.Line(c, a, 3)
}

func (canvas *Canvas) Size() Vector {
	size := canvas.inner.GetRect().Size()
	return NewVector(float64(size.X), float64(size.Y))
}

func (canvas *Canvas) SetPoint(position Vector, color int) {
	if canvas.debugger {
		fmt.Println(position)
		return
	}
	// if position.X >= 0 && position.Y >= 0 {
	// 	return
	// }
	pointer := position.ToPointer()
	canvas.inner.SetPoint(pointer, ui.Color(color))
}

func (canvas *Canvas) Line(source, target Vector, color int) {
	vector := target.Sub(source)
	len := vector.Len()
	norm := vector.Normalize()
	position := source
	for i := 0.; i < len; i++ {
		position.X = source.X + i*norm.X
		position.Y = source.Y + i*norm.Y
		canvas.SetPoint(position, color)
	}
	// diff := source.Diff(target)
	// min := source.Min(target)
	// max := math.Max(diff.X, diff.Y)
	// for i := 0.; i < max; i++ {
	// 	x := min.X + (i/max)*diff.X
	// 	y := min.Y + (i/max)*diff.Y
	// 	if x >= 0 && y >= 0 {
	// 		canvas.SetPoint(NewVector(x, y), color)
	// 	}
	// }
}

func main() {
	if err := ui.Init(); err != nil {
		log.Printf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	canvas := NewCanvas()
	//canvas.debugger = true

	//canvas.Line(Vector{30, 50}, Vector{33, 13}, 4)
	//canvas.Line(Vector{33, 13}, Vector{20, 40}, 4)
	//canvas.Line(Vector{10, 10}, Vector{30, 40}, 4)
	canvas.Triangle(Triangle{Vector{10, 10}, Vector{20, 10}, Vector{20, 30}}, 1)
	canvas.Triangle(Triangle{Vector{30, 50}, Vector{30, 13}, Vector{20, 40}}, 13)

	ui.Render(canvas.inner)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
