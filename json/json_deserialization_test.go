package json

import (
	"encoding/json"
	"reflect"
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

type Stock struct {
	// You have to use PascalCase to unmarshall the JSON, using camelCase to
	// unmarshall it won't work, e.g. a string value will be filled as "".
	Name  string
	Count int
}

/* ----- json.Unmarshal ----- */

func TestUnmarshalAsMap(t *testing.T) {
	// Given
	// language=JSON
	content := `{"apple": 1, "banana": 2}`

	// When
	var fruits map[string]int
	err := json.Unmarshal([]byte(content), &fruits)

	// Then
	if err != nil {
		t.Error(err)
	}
	eq := reflect.DeepEqual(fruits, map[string]int{
		"apple":  1,
		"banana": 2,
	})
	if !eq {
		t.Error()
	}
}

func TestUnmarshalAsStruct(t *testing.T) {
	// Given
	// language=JSON
	content := `[
	{"name": "apple", "count": 1},
	{"name": "banana", "count": 2}
]`

	// When
	var fruits []Stock
	err := json.Unmarshal([]byte(content), &fruits)

	// Then
	if err != nil {
		t.Error(err)
	}
	apple := Stock{Name: "apple", Count: 1}
	banana := Stock{Name: "banana", Count: 2}
	if fruits[0] != apple || fruits[1] != banana {
		t.Error()
	}
}

func TestUnmarshalAsRawMessage(t *testing.T) {
	// Given
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

	// When
	var indexMappingsMap map[string]json.RawMessage
	err := json.Unmarshal([]byte(content), &indexMappingsMap)

	// Then
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
	eq := reflect.DeepEqual(fruits, map[string]int{
		"apple":  1,
		"banana": 2,
	})
	if !eq {
		t.Error()
	}
}

func TestDecodeAsStruct(t *testing.T) {
	// Given
	// language=JSON
	content := `[
	{"name": "apple", "count": 1},
	{"name": "banana", "count": 2}
]`
	reader := strings.NewReader(content)
	decoder := json.NewDecoder(reader)

	// When
	var fruits []Stock
	err := decoder.Decode(&fruits)

	// Then
	if err != nil {
		t.Error(err)
	}
	apple := Stock{Name: "apple", Count: 1}
	banana := Stock{Name: "banana", Count: 2}
	if fruits[0] != apple || fruits[1] != banana {
		t.Error()
	}
}

func TestDecodeAsRawMessage(t *testing.T) {
	// Given
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
	reader := strings.NewReader(content)
	decoder := json.NewDecoder(reader)

	// When
	var indexMappingsMap map[string]json.RawMessage
	err := decoder.Decode(&indexMappingsMap)

	// Then
	if err != nil {
		t.Error(err)
	}
	if len(indexMappingsMap) != 1 {
		t.Error()
	}

	mappings := indexMappingsMap["my_index"]
	if mappings == nil {
		t.Error()
	}
}
