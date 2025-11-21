package types

const (
	UlistUrl = "ulist_labels"
)

type UList struct {
	Id       string          `json:"id"`
	Added    int64           `json:"added"` // int64 so its compatible with time.Unix()
	Voted    *int64          `json:"voted"`
	LastMod  int64           `json:"lastmod"`
	Vote     *int            `json:"vote"`
	Started  *string         `json:"started"`
	Finished *string         `json:"finished"`
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
