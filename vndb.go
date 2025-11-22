package vndb_go

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	BaseUrl   = "https://api.vndb.org/kana"
	VNDBToken = "VNDB_TOKEN"
)

// VNDBClient
//
// Includes a client that makes use of VNDB's api via several functions
// For more info, check: https://api.vndb.org/kana
type VNDBClient struct {
	BaseUrl     string
	token       string
	HttpClient  *http.Client
	RateLimiter *rate.Limiter
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
	limiter := rate.NewLimiter(rate.Limit(200.0/300.0), 200) // 200 requests per 5 minutes

	return &VNDBClient{
		BaseUrl: BaseUrl,
		token:   token,
		HttpClient: &http.Client{
			Timeout: 3 * time.Second,
		},
		RateLimiter: limiter,
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

	err := c.RateLimiter.Wait(req.Context())
	if err != nil {
		return err
	}

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

// Use this instead of Post when working with ulists
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

// GetUser
//
// Gets the user without any fields specified
func (c *VNDBClient) GetUser(username string, ctx context.Context) (*User, error) {
	var url = fmt.Sprintf("user?q=%s", username)

	usr, err := GrabUser(url, username, ctx, c)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

// GetUserWithFields
//
// Gets the user with specified fields
//   - `lv` : If true gets lengthvotes
//   - `lvsum` : If true gets lengthvotes_sum
func (c *VNDBClient) GetUserWithFields(username string, ctx context.Context, lv, lvsum bool) (*User, error) {
	var url = fmt.Sprintf("user?q=%s&fields=", username)

	switch lv {
	case true:
		if lvsum {
			url += "lengthvotes,lengthvotes_sum"
			break
		}
		url += "lengthvotes"
	case false:
		if lvsum {
			url += "lengthvotes_sum"
			break
		}
	}

	usr, err := GrabUser(url, username, ctx, c)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (c *VNDBClient) GetUListLabels(ctx context.Context, userId *string) (*GetUList, error) {
	var userToSearch string

	// if nil, we pass the current user using the client
	if userId == nil {
		auth, err := c.GetAuthInfo(context.Background(), c.token)
		if err != nil {
			return nil, err
		}

		userToSearch = auth.Id
	} else {
		userToSearch = *userId
	}

	url := fmt.Sprintf("%s?user=%s", UlistUrl, userToSearch)

	var ulist GetUList
	err := c.Get(ctx, url, &ulist)
	if err != nil {
		return nil, err
	}

	return &ulist, nil
}

func (c *VNDBClient) GetStats(ctx context.Context) (*Stats, error) {
	var stats Stats
	err := c.Get(ctx, "stats", &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (c *VNDBClient) GetAuthInfo(ctx context.Context, token string) (*AuthInfo, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, AuthUrl)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.token))

	var authInfo AuthInfo
	err = c.SendRequestWithToken(req, &authInfo, token)
	if err != nil {
		return nil, err
	}

	return &authInfo, nil
}

// GrabUser
//
// Function to grab the user data
func GrabUser(url, username string, ctx context.Context, c *VNDBClient) (*User, error) {
	var usr UserResponse
	usr.Results = make(map[string]*User)

	err := c.Get(ctx, url, &usr.Results)
	if err != nil {
		return nil, err
	}

	user, ok := usr.Results[username]
	if !ok {
		return nil, fmt.Errorf("user %s not found", username)
	}

	return user, nil
}
