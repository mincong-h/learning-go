package main

import "testing"
import "strings"

func TestStrings(t *testing.T) {
	if !strings.Contains("test", "es") {
		t.Error()
	}
	if strings.Count("test", "t") != 2 {
		t.Error()
	}
	if !strings.HasPrefix("test", "te") {
		t.Error()
	}
	if !strings.HasSuffix("test", "st") {
		t.Error()
	}
	if strings.Index("test", "e") != 1 {
		t.Error()
	}
	if strings.Index("test", "a") != -1 {
		t.Error()
	}
	if strings.Join([]string{"a", "b"}, ",") != "a,b" {
		t.Error()
	}
	if strings.Repeat("6", 3) != "666" {
		t.Error()
	}
}
