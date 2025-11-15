package wrapper

import (
	"context"
	"net/http"
)

type Schema struct {
	ApiFields SchemaApiFields `json:"api_fields"`
	ExtLinks  ExtLinkSchema   `json:"extlinks"`
}

type SchemaApiFields struct {
	Character Character `json:"/character"`
	Producer  Producer  `json:"/producer"`
	Trait     Trait     `json:"/trait"`
}

type Character struct {
	Age         *int         `json:"age"`
	Aliases     []*string    `json:"aliases"`
	Birthday    [][2]*int    `json:"birthday"`
	BloodType   *string      `json:"blood_type"`
	Bust        *int         `json:"bust"`
	Cup         *string      `json:"cup"`
	Description *string      `json:"description"`
	Gender      *string      `json:"gender"`
	Height      *int         `json:"height"`
	Hips        *int         `json:"hips"`
	Id          *string      `json:"id"`
	Image       *CharImage   `json:"image"`
	Name        *string      `json:"name"`
	Original    *string      `json:"original"`
	Sex         *string      `json:"sex"`
	Traits      TraitsSchema `json:"traits"`
}

type CharImage struct {
	Dims      [][2]*int `json:"dims"`
	Id        *string   `json:"id"`
	Sexual    *uint8    `json:"sexual"` // 0 - 2
	Url       *string   `json:"url"`
	Violence  *uint8    `json:"violence"` // 0 - 2
	VoteCount *int      `json:"votecount"`
}

type TraitsSchema struct {
	Inherit string `json:"_inherit"`
	Lie     *bool  `json:"lie"`
	Spoiler *int   `json:"spoiler"` // 0 - 2
}

type Producer struct {
	*Common
	ExtLinks ExtLink `json:"extlinks"`
	Lang     *string `json:"lang"`
	Original *string `json:"original"`
	Type     *string `json:"type"`
}

// ExtLink
//
// Links to external websites
type ExtLink struct {
	Id    *string `json:"id"`
	Label *string `json:"label"`
	Name  *string `json:"name"`
	Url   *string `json:"url"`
}

type Trait struct {
	*Common
	Applicable *bool   `json:"applicable"`
	CharCount  *int    `json:"char_count"`
	GroupId    *string `json:"group_id"`
	GroupName  *string `json:"group_name"`
	Searchable *bool   `json:"searchable"`
	Sexual     *bool   `json:"sexual"`
}

// Common
//
// Includes common variables
type Common struct {
	Aliases     []*string `json:"aliases"`
	Description *string   `json:"description"`
	Id          *string   `json:"id"`
	Name        *string   `json:"name"`
}

type ExtLinkSchema struct {
	Producer []*ExtLink `json:"/producer"`
}

// GetSchema
//
// Gets the Schema
func (c *VNDBClient) GetSchema(ctx context.Context) (*Schema, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", BaseUrl+"/schema", nil)
	if err != nil {
		return nil, err
	}

	var schema Schema
	err = c.sendRequest(req, &schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
