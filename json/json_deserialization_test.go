package json

import (
	"encoding/json"
	"strings"
	"testing"
)

//
// How to get the key value from a json string in Go
// https://stackoverflow.com/a/17453121/4381330
//
// Convert a JSON to map in Go (Golang)
// https://golangbyexample.com/json-to-map-golang/
//
// How to use JSON with Go [best practices]
// https://yourbasic.org/golang/json-example/
//

/* ----- json.Unmarshal ----- */

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
	// You have to use PascalCase to unmarshall the JSON, using camelCase to
	// unmarshall it won't work, e.g. a string value will be filled as "".
	Name  string
	Count int
}

func TestDeserializeAsStruct(t *testing.T) {
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
	apple := Stock{Name: "apple", Count: 1}
	banana := Stock{Name: "banana", Count: 2}
	if fruits[0] != apple || fruits[1] != banana {
		t.Error()
	}
}

func TestDeserializeAsRawMessage(t *testing.T) {
	// language=JSON
	content := `{
  "my_index": {
    "mappings": {
      "properties": {
        "msg": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        }
      }
    }
  }
}
`
	var indexMappingsMap map[string]json.RawMessage
	err := json.Unmarshal([]byte(content), &indexMappingsMap)
	if err != nil {
		t.Error()
	}
	if len(indexMappingsMap) != 1 {
		t.Error()
	}

	mappings := indexMappingsMap["my_index"]
	if mappings == nil {
		t.Error()
	}
}

/* ----- json.Decoder ----- */

func TestDecodeAsMap(t *testing.T) {
	// Given
	// language=JSON
	content := `{"apple": 1, "banana": 2}`
	reader := strings.NewReader(content)
	decoder := json.NewDecoder(reader)

	// When
	var fruits map[string]int
	err := decoder.Decode(&fruits)

	// Then
	if err != nil {
		t.Error(err)
	}
	if fruits["apple"] != 1 || fruits["banana"] != 2 {
		t.Error()
	}
}
