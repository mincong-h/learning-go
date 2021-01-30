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

```
curl -X POST http://localhost:9200/my_index/_doc/?pretty \
  -H 'Content-Type: application/json' \
  -d '{"msg": "hello world!"}'
```

## Run

From project root, you can run the following samples:

```sh
> go run elasticsearch/*.go info
2021/01/29 21:54:17 7.10.0
2021/01/29 21:54:17 [200 OK] {
  "name" : "7550ef2d3518",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "N1U2tJhoQ_Kvclx7R3bKmA",
  "version" : {
    "number" : "7.8.0",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "757314695644ea9a1dc2fecd26d1a43856725e65",
    "build_date" : "2020-06-14T19:35:50.234439Z",
    "build_snapshot" : false,
    "lucene_version" : "8.5.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
 <nil>
```

```sh
> go run elasticsearch/*.go get_indices
2021/01/30 18:54:54 Found 1 indices: [my_index]
```