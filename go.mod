module github.com/mincong-h/learning-go

require (
	// Installation | Elasticsearch Go Client
	// https://www.elastic.co/guide/en/elasticsearch/client/go-api/7.x/installation.html
	github.com/elastic/go-elasticsearch/v7 v7.10.0

	// Testing
	github.com/stretchr/testify v1.6.1
	k8s.io/api v0.20.4

	// Kubernetes
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
)

go 1.12
