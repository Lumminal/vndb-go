package wrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type VNDBClient struct {
	BaseUrl    string
	token      string
	HttpClient *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewVndbClient(token string) *VNDBClient {
	return &VNDBClient{
		BaseUrl: BaseUrl,
		token:   token,
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *VNDBClient) sendRequest(req *http.Request, v interface{}) error {
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

func (c *VNDBClient) sendRequestWithToken(req *http.Request, v interface{}, token string) error {
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
