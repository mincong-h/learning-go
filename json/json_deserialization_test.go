package json

import (
	"encoding/json"
	"testing"
)

func TestDeserializeAsMap(t *testing.T) {
	// language=JSON
	content := `{"apple": 1, "banana": 2}`
	var fruits map[string]int
	err := json.Unmarshal([]byte(content), &fruits)
	if err != nil {
		t.Error()
	}
	if fruits["apple"] != 1 || fruits["banana"] != 2 {
		t.Error()
	}
}

type Stock struct {
	name  string
	count int
}

func TestDeserializeAsStruct(t *testing.T)  {
	// language=JSON
	content := `[
	{"name": "apple", "count": 1},
	{"name": "banana", "count": 2}
]`
	var fruits []Stock
	err := json.Unmarshal([]byte(content), &fruits)
	if err != nil {
		t.Error()
	}
	apple := Stock{name: "apple", count: 1}
	banana := Stock{name: "banana", count: 2}
	if fruits[0] != apple || fruits[1] != banana {
		t.Error()
	}
}
