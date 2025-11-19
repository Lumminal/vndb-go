package wrapper

import (
	"context"
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
	var stats Stats
	err := c.Get(ctx, "stats", &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
