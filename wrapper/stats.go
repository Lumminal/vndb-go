package wrapper

import (
	"context"
	"fmt"
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

func (c *VNDBClient) GetStats(ctx context.Context) (*Stats, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/stats", c.BaseUrl), nil)
	if err != nil {
		return nil, err
	}

	var stats Stats
	err = c.SendRequest(req, &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
