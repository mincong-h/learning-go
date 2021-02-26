package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO: how to handle custom JSON property mapping, e.g. cluster_name => ClusterName?
type ClusterSummary struct {
	Name         string
	Cluster_Name string
	Cluster_Uuid string
	Version      ClusterVersion
	Tagline      string
}

type ClusterVersion struct {
	Number                              string
	Build_Flavor                        string
	Build_Type                          string
	Build_Hash                          string
	Build_Date                          string
	Build_Snapshot                      bool
	Lucene_Version                      string
	Minimum_Wire_Compatibility_Version  string
	Minimum_Index_Compatibility_Version string
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
