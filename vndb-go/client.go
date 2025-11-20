package vndb_go

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// VNDBClient
//
// Includes a client that makes use of VNDB's api via several functions
// For more info, check: https://api.vndb.org/kana
type VNDBClient struct {
	BaseUrl    string
	token      string
	HttpClient *http.Client
}

// errorResponse
//
// Code is the status code, while Message is the error message.
// Used when an error happens during sendRequest, with a bad status code.
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewVndbClient
//
// Returns a VNDBClient for use in operations
// You can pass an empty token if you don't have one
func NewVndbClient(token string) *VNDBClient {
	return &VNDBClient{
		BaseUrl: BaseUrl,
		token:   token,
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SendRequest
//
// Sends a request towards VNDB api and allows us to fetch the stats.
//   - `req` : The request to handle (e.g. from http.NewRequest() or http.NewRequestWithContext() )
//   - `v` : The interface to decode the info into (e.g. if you pass a reference of a Stats struct, Stats will now hold info taken from the API)
func (c *VNDBClient) SendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.token))
	req.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

// SendRequestWithToken
//
// Same as sendRequest, except you specify the token you want to use
func (c *VNDBClient) SendRequestWithToken(req *http.Request, v interface{}, token string) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func (c *VNDBClient) Get(ctx context.Context, endpoint string, out interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", c.BaseUrl, endpoint), nil)
	if err != nil {
		return err
	}

	err = c.SendRequest(req, &out)
	if err != nil {
		return err
	}

	return nil
}

func (c *VNDBClient) Post(ctx context.Context, endpoint string, q *Query) (*VNResponse, error) {
	body, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s", c.BaseUrl, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var result VNResponse
	err = c.SendRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *VNDBClient) PostUlist(ctx context.Context, endpoint string, q *UlistQueryRequest) (*VNResponse, error) {
	body, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s", c.BaseUrl, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var result VNResponse
	err = c.SendRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
