package main

import "testing"

func TestIntegers(t *testing.T) {
	// unsiged
	var u uint = 1
	var u16 uint16 = 1
	var u32 uint32 = 1
	var u64 uint64 = 1

	// signed
	var i = 1
	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1

	if u != 1 {
		t.Error()
	}
	if u16 != 1 {
		t.Error()
	}
	if u32 != 1 {
		t.Error()
	}
	if u64 != 1 {
		t.Error()
	}

	if i != 1 {
		t.Error()
	}
	if i8 != 1 {
		t.Error()
	}
	if i16 != 1 {
		t.Error()
	}
	if i32 != 1 {
		t.Error()
	}
	if i64 != 1 {
		t.Error()
	}

	var f32 float32 = 1.0
	var f64 = 1.0

	if f32 != 1.0 {
		t.Error()
	}
	if f64 != 1.0 {
		t.Error()
	}
}

func TestStrings(t *testing.T) {
	if len("Go") != 2 {
		t.Error()
	}
	if "Go"[0] != 'G' {
		t.Error()
	}
	if "Go"+"lang" != "Golang" {
		t.Error()
	}
}

func TestArrays(t *testing.T) {
	arr := [6]int{0, 1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			t.Error()
		}
	}
	for _, value := range arr {
		if value < 0 {
			t.Error()
		}
	}
}

func TestSlices(t *testing.T) {
	// make
	s1 := make([]int, 3, 5)
	for _, value := range s1 {
		if value != 0 {
			t.Error()
		}
	}
	if len(s1) != 3 {
		t.Error(len(s1))
	}
	if cap(s1) != 5 {
		t.Error(cap(s1))
	}

	// append
	s2 := []int{0}
	s3 := append(s2, 0, 0)
	for i, value := range s1 {
		if value != s3[i] {
			t.Errorf("%d != %d", value, s3[i])
		}
	}

	// copy
	s4 := make([]int, 2)
	copy(s1, s4)
	if len(s4) != 2 {
		t.Error(len(s4))
	}
	if cap(s4) != 2 {
		t.Error(len(s4))
	}
	for _, value := range s4 {
		if value != 0 {
			t.Error(value)
		}
	}
}

func TestMaps(t *testing.T) {
	// make
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	if map1["one"] != 1 || map1["two"] != 2 {
		t.Error()
	}

	// delete
	delete(map1, "one")
	if _, exist := map1["one"]; exist {
		t.Error()
	}
	if _, exist := map1["two"]; !exist {
		t.Error()
	}

	// anther declaration
	map2 := map[string]int{
		"one": 1,
		"two": 2,
	}
	if map2["one"] != 1 || map2["two"] != 2 {
		t.Error()
	}
}
