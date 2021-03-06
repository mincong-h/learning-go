package core

import "reflect"
import "testing"

func TestSlice1(t *testing.T) {
	numbers := []string{"a", "b", "c", "d", "e"}
	actual := numbers[:2]
	expected := []string{"a", "b"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error()
	}
}

func TestSlice2(t *testing.T) {
	numbers := []string{"a", "b", "c", "d", "e"}
	actual := numbers[1:2]
	expected := []string{"b"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error()
	}
}

func TestSlice3(t *testing.T) {
	numbers := []string{"a", "b", "c", "d", "e"}
	actual := numbers[1:]
	expected := []string{"b", "c", "d", "e"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error()
	}
}
