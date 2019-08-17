package main

import "math"

// -- Point

type Point struct {
	x float64
	y float64
}

func (p *Point) distanceTo(q *Point) float64 {
	x2 := (p.x - q.x) * (p.x - q.x)
	y2 := (p.y - q.y) * (p.y - q.y)
	return math.Sqrt(x2 + y2)
}

// -- User and Location (embeded types)

type User struct {
	Name     string
	Location Location
}

type Location struct {
	Country string
	City    string
}
