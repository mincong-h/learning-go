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

	if u != 1 { t.Error() }
	if u16 != 1 { t.Error() }
	if u32 != 1 { t.Error() }
	if u64 != 1 { t.Error() }

	if i != 1 { t.Error() }
	if i8 != 1 { t.Error() }
	if i16 != 1 { t.Error() }
	if i32 != 1 { t.Error() }
	if i64 != 1 { t.Error() }

	var f32 float32 = 1.0
	var f64 = 1.0

	if f32 != 1.0 { t.Error() }
	if f64 != 1.0 { t.Error() }
}
