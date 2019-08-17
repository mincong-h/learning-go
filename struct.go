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

// -- Embedded Types: User and City

type User struct {
	Name string
	City City
}

type City struct {
	Country string
	Name    string
}

// -- Interface: Callable
// To implement an interface in Go, we just need
// implement all the methods in the interface.
type Representable interface {
	repr() string
}

func (u *User) repr() string {
	return u.Name + " from " + u.City.repr()
}

func (c *City) repr() string {
	return c.Name + ", " + c.Country
}
