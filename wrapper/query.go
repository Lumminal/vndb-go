package wrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Query struct {
	Filters           interface{} `json:"filters,omitempty"`
	Fields            string      `json:"fields,omitempty"`
	Sort              string      `json:"sort,omitempty"`
	Reverse           bool        `json:"reverse,omitempty"`
	Results           int         `json:"results,omitempty"`
	Page              int         `json:"page,omitempty"`
	User              *User       `json:"user,omitempty"`
	Count             bool        `json:"count,omitempty"`
	CompactFilters    bool        `json:"compact_filters,omitempty"`
	NormalizedFilters bool        `json:"normalized_filters,omitempty"`
}

type VNResponse struct {
	Results           json.RawMessage `json:"results"`
	More              bool            `json:"more"`
	Count             int             `json:"count"`
	CompactFilters    string          `json:"compact_filters"`
	NormalizedFilters []interface{}   `json:"normalized_filters"`
}

// TODO: FILTER FLAGS
func (c *VNDBClient) Query(ctx context.Context, endpoint string, q *Query) (*VNResponse, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, endpoint)

	body, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var result VNResponse

	err = c.sendRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
