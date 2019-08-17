package main

import "testing"

func TestStructs(t *testing.T) {
	a := Point{x: 4.0, y: 0.0}
	b := Point{x: 0.0, y: 3.0}
	if a.distanceTo(&b) != 5.0 {
		t.Error()
	}

	paris := Location{Country: "France", City: "Paris"}
	user := User{Name: "Tom", Location: paris}

	if user.Location.Country != "France" {
		t.Error()
	}
	if user.Location.City != "Paris" {
		t.Error()
	}
	if user.Name != "Tom" {
		t.Error()
	}
}
