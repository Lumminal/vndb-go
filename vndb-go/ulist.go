package vndb_go

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const (
	ulistUrl = "ulist_labels"
)

type UList struct {
	Id       string          `json:"id"`
	Added    int64           `json:"added"` // int64 so its compatible with time.Unix()
	Voted    *int64          `json:"voted"`
	LastMod  int64           `json:"lastmod"`
	Vote     *int            `json:"vote"`
	Started  *time.Time      `json:"started"`
	Finished *time.Time      `json:"finished"`
	Notes    *string         `json:"notes"`
	Labels   *[]Label        `json:"labels"`
	Vn       *Vn             `json:"vn"`
	Releases *[]ReleaseUList `json:"releases"`
}

type GetUList struct {
	Labels []UListLabel `json:"labels"`
}

type UListLabel struct {
	Id      int    `json:"id"`
	Private bool   `json:"private"`
	Label   string `json:"label"`
	Count   int    `json:"count"`
}

type Label struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type ReleaseUList struct {
	*Release
	ListStatus int `json:"list_status"`
}

type UListQueryRequest struct {
	BaseQuery
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

	url := fmt.Sprintf("%s?user=%s", ulistUrl, userToSearch)

	var ulist GetUList
	err := c.Get(ctx, url, &ulist)
	if err != nil {
		return nil, err
	}

	return &ulist, nil
}

type UListQuery struct {
	BaseQuery
	User string `json:"user,omitempty"`
}

func UlistQuery(client *VNDBClient) *UListQuery {
	return &UListQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedUlistSort},
	}
}

func (ul *UListQuery) SetUser(id string) {
	ul.User = id
}

func (cq *UListQuery) Get(ctx context.Context) ([]UList, error) {

	ulistquery := &UlistQueryRequest{
		Query: *cq.Query,
		User:  cq.User,
	}

	resp, err := cq.Client.PostUlist(ctx, "ulist", ulistquery)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []UList{}, nil
	}

	var ulists []UList
	err = json.Unmarshal(resp.Results, &ulists)
	if err != nil {
		return nil, err
	}

	return ulists, nil
}
