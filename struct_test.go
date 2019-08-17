package main

import "testing"

func TestStructs(t *testing.T) {
	a := Point{x: 4.0, y: 0.0}
	b := Point{x: 0.0, y: 3.0}
	if a.distanceTo(&b) != 5.0 {
		t.Error()
	}
}
