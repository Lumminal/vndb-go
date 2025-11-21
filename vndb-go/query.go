package vndb_go

import (
	"encoding/json"
)

type Query struct {
	Filters           Filter `json:"filters,omitempty"`
	Fields            string `json:"fields,omitempty"`
	Sort              string `json:"sort,omitempty"`
	Reverse           bool   `json:"reverse,omitempty"`
	Results           int    `json:"results,omitempty"`
	Page              int    `json:"page,omitempty"`
	User              *User  `json:"user,omitempty"`
	Count             bool   `json:"count,omitempty"`
	CompactFilters    bool   `json:"compact_filters,omitempty"`
	NormalizedFilters bool   `json:"normalized_filters,omitempty"`
}

type UlistQueryRequest struct {
	Query
	User string `json:"user,omitempty"`
}

type VNResponse struct {
	Results           json.RawMessage `json:"results"`
	More              bool            `json:"more"`
	Count             int             `json:"count"`
	CompactFilters    string          `json:"compact_filters"`
	NormalizedFilters []interface{}   `json:"normalized_filters"`
}
