package main

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

/*
	Prerequisite:

		Create an index with a new document if the target index does not exist:

		curl -X POST http://localhost:9200/my_index/_doc/?pretty \
			 -H 'Content-Type: application/json' \
			 -d '{"msg": "hello world!"}'

	Execution:

		go run elasticsearch/get_indices.go
*/
func main() {
	es, _ := elasticsearch.NewDefaultClient()
	response, err := es.Indices.Get([]string{"_all"})
	if err != nil {
		panic(err)
	}
	indices := make(map[string]json.RawMessage)
	decoder := json.NewDecoder(response.Body)

	decodeErr := decoder.Decode(&indices)
	if decodeErr != nil {
		panic(decodeErr)
	}
	indexNames := make([]string, len(indices))
	i := 0
	for indexName := range indices {
		indexNames[i] = indexName
		i++
	}
	log.Printf("Found %d indices: %s", len(indexNames), indexNames)
}
