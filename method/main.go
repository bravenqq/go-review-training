package main

type Point struct {
	X float32
	Y float32
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{X: p.X - q.X, Y: p.Y - q.Y}
}

type Path []Point

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

func main() {

}
