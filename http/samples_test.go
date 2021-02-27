package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetClusterSummarySync(t *testing.T) {
	// Given
	// language=JSON
	content := `{
		"name": "4455175953c5",
		"cluster_name": "docker-cluster",
		"cluster_uuid": "yFfZWsvCT6ODAosBBMB8AA",
		"version": {
			"number": "7.8.0",
			"build_flavor": "default",
			"build_type": "docker",
			"build_hash": "757314695644ea9a1dc2fecd26d1a43856725e65",
			"build_date": "2020-06-14T19:35:50.234439Z",
			"build_snapshot": false,
			"lucene_version": "8.5.1",
			"minimum_wire_compatibility_version": "6.8.0",
			"minimum_index_compatibility_version": "6.0.0-beta1"
		},
		"tagline": "You Know, for Search"
	}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	}))
	defer server.Close()

	// When
	actualSummary, err := getClusterSummarySync(server.Client(), server.URL)

	// Then
	expectedSummary := ClusterSummary{
		Name:        "4455175953c5",
		ClusterName: "docker-cluster",
		ClusterUuid: "yFfZWsvCT6ODAosBBMB8AA",
		Version: ClusterVersion{
			Number:                           "7.8.0",
			BuildFlavor:                      "default",
			BuildType:                        "docker",
			BuildHash:                        "757314695644ea9a1dc2fecd26d1a43856725e65",
			BuildDate:                        "2020-06-14T19:35:50.234439Z",
			BuildSnapshot:                    false,
			LuceneVersion:                    "8.5.1",
			MinimumWireCompatibilityVersion:  "6.8.0",
			MinimumIndexCompatibilityVersion: "6.0.0-beta1",
		},
		Tagline: "You Know, for Search",
	}

	if err != nil || actualSummary != expectedSummary {
		t.Error()
	}
}
