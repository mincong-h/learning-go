package core

import "testing"

func TestStructs(t *testing.T) {
	a := Point{x: 4.0, y: 0.0}
	b := Point{x: 0.0, y: 3.0}
	if a.distanceTo(&b) != 5.0 {
		t.Error()
	}

	paris := City{Country: "France", Name: "Paris"}
	user := User{Name: "Tom", City: paris}

	if user.City.Country != "France" {
		t.Error()
	}
	if user.City.Name != "Paris" {
		t.Error()
	}
	if user.Name != "Tom" {
		t.Error()
	}
	if paris.repr() != "Paris, France" {
		t.Error()
	}
	if user.repr() != "Tom from Paris, France" {
		t.Error()
	}
}
