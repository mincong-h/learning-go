package main

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
)

/*
	Prerequisite:

		Create an index with a new document if the target index does not exist:

		curl -X POST http://localhost:9200/my_index/_doc/?pretty \
			 -H 'Content-Type: application/json' \
			 -d '{"msg": "hello world!"}'

	Execution:

		go run elasticsearch/*.go get_indices
		go run elasticsearch/*.go get_indices http://localhost:9200
*/
func GetIndices() {
	var es *elasticsearch.Client
	var cfgErr error

	url := GetUrl()
	if url == "" {
		es, cfgErr = elasticsearch.NewDefaultClient()
	} else {
		es, cfgErr = elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{url},
		})
	}
	if cfgErr != nil {
		panic(cfgErr)
	}

	response, httpErr := es.Indices.Get([]string{"_all"})
	if httpErr != nil {
		panic(httpErr)
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

func GetUrl() string {
	var lastArg = os.Args[len(os.Args)-1]
	if lastArg != "get_indices" {
		return lastArg
	}
	return ""
}
