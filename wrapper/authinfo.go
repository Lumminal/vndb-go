package wrapper

import (
	"context"
	"fmt"
	"net/http"
)

const (
	authUrl = "authinfo"
)

type AuthInfo struct {
	BaseUser
	Permissions []string `json:"permissions"`
}

func (c *VNDBClient) GetAuthInfo(ctx context.Context, token string) (*AuthInfo, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, authUrl)

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
