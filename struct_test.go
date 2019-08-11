package main

import "math"
import "testing"

func TestStructs(t *testing.T) {
	dist := func(p, q *Point) float64 {
		xdiff := p.x - q.x
		ydiff := p.y - q.y
		return math.Sqrt(xdiff*xdiff + ydiff*ydiff)
	}
	a := Point{x: 4.0, y: 0.0}
	b := Point{x: 0.0, y: 3.0}
	if dist(&a, &b) != 5.0 {
		t.Error()
	}
}

type Point struct {
	x float64
	y float64
}
