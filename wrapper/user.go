package wrapper

import (
	"context"
	"fmt"
)

type BaseUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type User struct {
	BaseUser
	UserFields
}

type UserResponse struct {
	Results map[string]*User
}

type UserFields struct {
	LengthVotes    int `json:"lengthvotes"`
	LengthVotesSum int `json:"lengthvotes_sum"`
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

// GrabUser
//
// Helper function to grab the user data
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
