package types

import (
	"encoding/json"
	"fmt"
)

const (
	nonStandard = "non-standard"
)

// This file is a star-wars intro

type Character struct {
	Age         *int         `json:"age"`
	Aliases     *[]string    `json:"aliases"`
	Birthday    *[2]int      `json:"birthday"`
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
	Dims      *[2]int `json:"dims"`
	Id        *string `json:"id"`
	Sexual    *uint8  `json:"sexual"` // 0 - 2
	Url       *string `json:"url"`
	Violence  *uint8  `json:"violence"` // 0 - 2
	VoteCount *int    `json:"votecount"`
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

type Release struct {
	Id         *string            `json:"id"`
	Title      *string            `json:"title"`
	AltTitle   *string            `json:"alltitle"`
	Languages  *[]Language        `json:"languages"`
	Platforms  *[]string          `json:"platforms"`
	Media      *[]Media           `json:"media"`
	Vns        *[]string          `json:"vns"`
	Producers  *[]ReleaseProducer `json:"producers"`
	Images     *[]ReleaseImage    `json:"images"`
	Released   *ReleaseDate       `json:"released"`
	Minage     *int               `json:"minage"`
	Patch      bool               `json:"patch"`
	Freeware   bool               `json:"freeware"`
	Uncensored *bool              `json:"uncensored"`
	Resolution *Resolution        `json:"resolution"`
	Engine     *string            `json:"engine"`
	Voiced     *int               `json:"voiced"` // 1 = not voiced, 2 = only ero scenes voiced, 3 = partially voiced, 4 = fully voiced.
	Notes      *string            `json:"notes"`
	Gtin       *string            `json:"gtin"`
	Catalog    *string            `json:"catalog"`
	Extlinks   *[]ExtLink         `json:"extlinks"`
}

type Vn struct {
	Id            *string       `json:"id"`
	Title         *string       `json:"title"`
	AltTitle      *string       `json:"alttitle"`
	Titles        *[]Title      `json:"titles"`
	Aliases       *[]string     `json:"aliases"`
	Olang         *string       `json:"olang"`
	DevStatus     *int          `json:"dev_status"` // 0 - 2
	Released      *ReleaseDate  `json:"released"`
	Languages     *[]string     `json:"languages"`
	Platforms     *[]string     `json:"platforms"`
	Image         *Image        `json:"image"`
	Length        *int          `json:"length"`
	LengthMinutes *int          `json:"length_minutes"`
	LengthVotes   *int          `json:"length_votes"`
	Description   *string       `json:"description"`
	Average       *int          `json:"average"`
	Rating        *int          `json:"rating"`
	VoteCount     *int          `json:"vote_count"`
	Screenshots   *[]Screenshot `json:"screenshots"`
	Relations     *[]Relation   `json:"relations"`
	Tags          *[]TagVn      `json:"tags"`
	Developers    *[]Producer   `json:"developers"`
	Editions      *[]Edition    `json:"editions"`
	Staff         *[]StaffVn    `json:"staff"`
	Va            *[]Va         `json:"va"`
	Extlinks      *[]ExtLink    `json:"extlinks"`
}

type Title struct {
	Lang     *string `json:"lang"`
	Title    *string `json:"title"`
	Latin    *string `json:"latin"`
	Official *bool   `json:"official"`
	Main     *bool   `json:"main"`
}

type Image struct {
	Id            *string `json:"id"`
	Url           *string `json:"url"`
	Dims          *[2]int `json:"dims"`
	Sexual        *int    `json:"sexual"`
	Violence      *int    `json:"violence"`
	VoteCount     *int    `json:"votecount"`
	Thumbnail     *string `json:"thumbnail"`
	ThumbnailDims *[2]int `json:"thumbnail_dims"`
}

type ReleaseImage struct {
	*Image
	Type      *string     `json:"type"`
	Vn        *string     `json:"vn"`
	Languages *[]Language `json:"languages"`
	Photo     *bool       `json:"photo"`
}

type Screenshot struct {
	*Image
	Release *Release `json:"release"`
}

type Relation struct {
	*Vn
	Relation         *string `json:"relation"`
	RelationOfficial *string `json:"relation_official"`
}

type TagVn struct {
	*Tag
	Rating  *int  `json:"rating"`
	Spoiler *int  `json:"spoiler"`
	Lie     *bool `json:"lie"`
}

type Edition struct {
	Eid      *int    `json:"eid"`
	Lang     *string `json:"lang"`
	Name     *string `json:"name"`
	Official *bool   `json:"official"`
}

type Tag struct {
	*Common
	Category   *string `json:"category"`
	Searchable *bool   `json:"searchable"`
	Applicable *bool   `json:"applicable"`
	VnCount    *int    `json:"vn_count"`
}

type Language struct {
	Lang  *string `json:"lang"`
	Title *string `json:"title"`
	Latin *string `json:"latin"`
	MTL   *bool   `json:"mtl"`
	Main  *bool   `json:"main"`
}

type Media struct {
	Medium *string `json:"medium"`
	Qty    *int    `json:"qty"`
}

type ReleaseProducer struct {
	*Producer
	Developer *string `json:"developer"`
	Publisher *string `json:"publisher"`
}

type Staff struct {
	Id          *string       `json:"id"`
	Name        *string       `json:"name"`
	Description *string       `json:"description"`
	Aid         *int          `json:"aid"`
	IsMain      *bool         `json:"ismain"`
	Original    *string       `json:"original"`
	Lang        *string       `json:"lang"`
	Gender      *string       `json:"gender"`
	Extlinks    *[]ExtLink    `json:"extlinks"`
	Aliases     *[]StaffAlias `json:"aliases"`
}

type StaffVn struct {
	*Staff
	Eid  *int    `json:"eid"`
	Role *string `json:"role"`
	Note *string `json:"note"`
}

type StaffAlias struct {
	Aid    *int    `json:"aid"`
	Name   *string `json:"name"`
	Latin  *string `json:"latin"`
	IsMain *bool   `json:"ismain"`
}

type Va struct {
	Character *Character `json:"character"`
	Staff     *Staff     `json:"staff"`
	Note      *string    `json:"note"`
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

type Quote struct {
	Id        *string    `json:"id"`
	Quote     *string    `json:"quote"`
	Score     *int       `json:"score"`
	Vn        *Vn        `json:"vn"`
	Character *Character `json:"character"`
}

type Resolution struct {
	Type *string
	Res  *[2]int
}

func (rs *Resolution) UnmarshalJSON(data []byte) error {
	var i interface{}
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}

	// if its string, its non-standard
	str, ok := i.(string)
	if ok && str == nonStandard {
		nonStand := nonStandard

		rs.Type = &nonStand
		rs.Res = nil
		return nil
	}

	// if its array, it contains 2 ints (width, height)
	arr, ok := i.([]interface{})
	if ok {
		if len(arr) == 2 {
			width, okw := arr[0].(float64)
			height, okh := arr[1].(float64)

			if !okw || !okh {
				return fmt.Errorf("invalid resolution: %s", i)
			}

			rs.Type = nil
			resArray := [2]int{int(width), int(height)}
			rs.Res = &resArray
			return nil
		}
	}

	return fmt.Errorf("invalid resolution: %s", i)
}
