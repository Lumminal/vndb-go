package wrapper

// This file is a star-wars intro

type SchemaApiFields struct {
	Character Character `json:"/character"`
	Producer  Producer  `json:"/producer"`
	Trait     Trait     `json:"/trait"`
	Vn        Vn        `json:"/vn"`
	// Release   Release   `json:"/release"`
}

type Character struct {
	Age         *int         `json:"age"`
	Aliases     []*string    `json:"aliases"`
	Birthday    [][2]*int    `json:"birthday"` // todo: is there something better?
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
	Id        *string            `json:"id"`
	Title     *string            `json:"title"`
	AltTitle  *string            `json:"alltitle"`
	Languages *[]Language        `json:"languages"`
	Platforms *[]string          `json:"platforms"`
	Media     *[]Media           `json:"media"`
	Vns       *[]string          `json:"vns"`
	Producers *[]ReleaseProducer `json:"producers"`
	// todo: finish
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
	Inherit   *string `json:"_inherit"`
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
