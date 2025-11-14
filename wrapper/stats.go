package wrapper

import (
	"context"
	"encoding/json"
	"net/http"
)

type Stats struct {
	Chars     int `json:"chars"`
	Producers int `json:"producers"`
	Releases  int `json:"releases"`
	Staff     int `json:"staff"`
	Tags      int `json:"tags"`
	Traits    int `json:"traits"`
	Vn        int `json:"vn"`
}

func GetStats(client *http.Client, ctx context.Context) (*Stats, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", BaseUrl+"/stats", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var stats Stats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}

	return &stats, nil
}
