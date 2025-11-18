package wrapper

import "time"

type UList struct {
	Id       string          `json:"id"`
	Added    int             `json:"added"`
	Voted    *int            `json:"voted"`
	LastMod  time.Time       `json:"lastmod"`
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
