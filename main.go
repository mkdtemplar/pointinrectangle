package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

type rectangle struct {
	a point
	b point
	c point
	d point
}

func newRectangle(a point, b point, c point, d point) *rectangle {
	return &rectangle{a: a, b: b, c: c, d: d}
}

func distance(a, b point) float64 {
	return math.Abs(math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2)))
}

func areaTriangle(a, b, c point) float64 {
	// Heron formula
	distanceAB := distance(a, b)
	distanceBC := distance(b, c)
	distanceCA := distance(c, a)

	s := (distanceAB + distanceBC + distanceCA) / 2.0

	return math.Abs(math.Sqrt(s * (s - distanceAB) * (s - distanceBC) * (s - distanceCA)))
}

func (r *rectangle) areaRectangle() float64 {
	distanceAB := math.Abs(distance(r.a, r.b))
	distanceBC := math.Abs(distance(r.b, r.c))

	return distanceAB * distanceBC
}

func (r *rectangle) checkPoint(p point) bool {
	PAB := areaTriangle(p, r.a, r.b)
	PBC := areaTriangle(p, r.b, r.c)
	PCD := areaTriangle(p, r.c, r.d)
	PAD := areaTriangle(p, r.a, r.d)

	trianglesArea := PAB + PBC + PCD + PAD
	fmt.Println("Triangles area: ", trianglesArea)

	rectangleArea := r.areaRectangle()
	fmt.Println("Rectangle area : ", rectangleArea)

	return trianglesArea <= rectangleArea

}
func main() {

	rectangle := newRectangle(point{1, 2}, point{5, 2}, point{5, 6}, point{1, 6})

	fmt.Println(rectangle.areaRectangle())
	fmt.Println(rectangle.checkPoint(point{3, 4}))

	if rectangle.checkPoint(point{3, 4}) {
		fmt.Printf("Point %v is inside of the rectangle %v\n", point{3, 4}, rectangle)
	} else {
		fmt.Printf("Point %v is not inside of the rectangle %v\n", point{3, 4}, rectangle)
	}

}
