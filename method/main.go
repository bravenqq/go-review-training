package main

import (
	"image/color"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type ColoredPoint struct {
	Point
	Color color.Color
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{X: p.X - q.X, Y: p.Y - q.Y}
}

type Path []Point

// same thing, but as a method of the Point type
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (path Path) TranslateBy(offset Point, add bool) {
	//Point类型方法的方法值
	var op func(p Point, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for _, point := range path {
		op(point, offset)
	}
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Values map[string][]string

// Get returns the first value associated with the given key,
// or "" if there are none.
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	// red := color.RGBA{255, 0, 0, 255}
	// blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{Point: Point{1, 1}}
	q := ColoredPoint{Point: Point{5, 4}}
	p.Distance(q.Point)
}
