package json

import (
	"encoding/json"
	"testing"
)

func TestDeserializeAsMap(t *testing.T) {
	content := `{"apple": 1, "banana": 2}`
	var fruits map[string]int
	err := json.Unmarshal([]byte(content), &fruits)
	if err != nil {
		t.Error()
	}
	if fruits["apple"] != 1 && fruits["banana"] != 2 {
		t.Error()
	}
}
