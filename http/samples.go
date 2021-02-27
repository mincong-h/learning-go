package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ClusterSummary struct {
	Name        string         `json:"name"`
	ClusterName string         `json:"cluster_name"`
	ClusterUuid string         `json:"cluster_uuid"`
	Version     ClusterVersion `json:"version"`
	Tagline     string         `json:"tagline"`
}

type ClusterVersion struct {
	Number                           string `json:"number"`
	BuildFlavor                      string `json:"build_flavor"`
	BuildType                        string `json:"build_type"`
	BuildHash                        string `json:"build_hash"`
	BuildDate                        string `json:"build_date"`
	BuildSnapshot                    bool   `json:"build_snapshot"`
	LuceneVersion                    string `json:"lucene_version"`
	MinimumWireCompatibilityVersion  string `json:"minimum_wire_compatibility_version"`
	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`
}

func GetClusterSummarySync() (ClusterSummary, error) {
	return getClusterSummarySync(http.DefaultClient, "http://localhost:9200")
}

func getClusterSummarySync(httpClient *http.Client, url string) (ClusterSummary, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return ClusterSummary{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ClusterSummary{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ClusterSummary{}, fmt.Errorf("failed to read response: %v", err)
	}

	var summary ClusterSummary
	err = json.Unmarshal(body, &summary)
	if err != nil {
		return ClusterSummary{}, fmt.Errorf("failed to unmarshal cluster summary: %s", body)
	}

	return summary, nil
}

func getClusterSummaryAsync(ctx context.Context, httpClient *http.Client, url string) (ClusterSummary, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return ClusterSummary{}, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return ClusterSummary{}, fmt.Errorf("failed to perform request: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ClusterSummary{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ClusterSummary{}, fmt.Errorf("failed to read response: %v", err)
	}

	var summary ClusterSummary
	err = json.Unmarshal(body, &summary)
	if err != nil {
		return ClusterSummary{}, fmt.Errorf("failed to unmarshal cluster summary: %s", body)
	}

	return summary, nil
}
