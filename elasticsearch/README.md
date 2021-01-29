# Elasticsearch

https://www.elastic.co/guide/en/elasticsearch/client/go-api/7.x/installation.html

## Prerequisite

```
docker run \
  --rm \
  -e discovery.type=single-node \
  -p 9200:9200 \
  docker.elastic.co/elasticsearch/elasticsearch:7.8.0
```

## Run

```
go run elasticsearch/main.go
```
