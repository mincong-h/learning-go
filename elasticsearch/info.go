package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func Info() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
