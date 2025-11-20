package vndb_go

import (
	"context"
	"encoding/json"
	"log"
	"strings"
)

type AllowedSort []string

const (
	IDSort         string = "id"
	TitleSort      string = "title"
	ReleasedSort   string = "released"
	NameSort       string = "name"
	RatingSort     string = "rating"
	VoteCountSort  string = "votecount"
	SearchRankSort string = "searchrank"
	VnCountSort    string = "vn_count"
	CharCountSort  string = "char_count"
	ScoreSort      string = "score"
	VotedSort      string = "voted"
	VoteSort       string = "vote"
	AddedSort      string = "added"
	LastModSort    string = "lastmod"
	StartedSort    string = "started"
	FinishedSort   string = "finished"
)

var (
	allowedVnSort        = []string{IDSort, TitleSort, ReleasedSort, RatingSort, VoteCountSort, SearchRankSort}
	allowedReleaseSort   = []string{IDSort, TitleSort, ReleasedSort, SearchRankSort}
	allowedProducerSort  = []string{IDSort, NameSort, SearchRankSort}
	allowedCharacterSort = []string{IDSort, NameSort, SearchRankSort}
	allowedStaffSort     = []string{IDSort, NameSort, SearchRankSort}
	allowedTagSort       = []string{IDSort, NameSort, VnCountSort, SearchRankSort}
	allowedTraitSort     = []string{IDSort, NameSort, CharCountSort, SearchRankSort}
	allowedQuoteSort     = []string{IDSort, ScoreSort}
	allowedUlistSort     = []string{IDSort, TitleSort, ReleasedSort, RatingSort, VoteCountSort, VotedSort, VoteSort, AddedSort, LastModSort, StartedSort, FinishedSort, SearchRankSort}
)

// BaseQuery - Start
//
// Holds common setter functions
type BaseQuery struct {
	Query       *Query
	Client      *VNDBClient
	AllowedSort AllowedSort
}

func (cq *BaseQuery) Fields(fields ...string) {
	cq.Query.Fields = strings.Join(fields, ",")
}

func (cq *BaseQuery) Results(number int) {
	cq.Query.Results = number
}

func (cq *BaseQuery) Page(number int) {
	cq.Query.Page = number
}

func (cq *BaseQuery) Reverse(reverse bool) {
	cq.Query.Reverse = reverse
}

func (cq *BaseQuery) CompactFilters(cf bool) {
	cq.Query.CompactFilters = cf
}

func (cq *BaseQuery) NormalizedFilters(nf bool) {
	cq.Query.NormalizedFilters = nf
}

func (cq *BaseQuery) Filters(filter Filter) {
	cq.Query.Filters = filter
}

func (cq *BaseQuery) Sort(sort ...string) {
	for _, str := range sort {
		if !Contains(cq.AllowedSort, str) {
			log.Printf("Warning: Sorting not allowed: %s", str)
		}
	}

	cq.Query.Sort = strings.Join(sort, ",")
}

// BaseQuery - End

// VNQuery - Start
type VNQuery struct {
	BaseQuery
}

func NewVnQuery(client *VNDBClient) *VNQuery {
	return &VNQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedVnSort},
	}
}

func (vnq *VNQuery) Get(ctx context.Context) ([]Vn, error) {
	resp, err := vnq.Client.Post(ctx, "vn", vnq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Vn{}, nil
	}

	var vns []Vn
	err = json.Unmarshal(resp.Results, &vns)
	if err != nil {
		return nil, err
	}

	return vns, nil
}

// VNQuery - End

// CharQuery - Start
type CharQuery struct {
	BaseQuery
}

func NewCharacterQuery(client *VNDBClient) *CharQuery {
	return &CharQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedCharacterSort},
	}
}

func (cq *CharQuery) Get(ctx context.Context) ([]Character, error) {
	resp, err := cq.Client.Post(ctx, "character", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Character{}, nil
	}

	var chars []Character
	err = json.Unmarshal(resp.Results, &chars)
	if err != nil {
		return nil, err
	}

	return chars, nil
}

// CharQuery - End

// ProducerQuery - Start
type ProducerQuery struct {
	BaseQuery
}

func NewProducerQuery(client *VNDBClient) *ProducerQuery {
	return &ProducerQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedProducerSort},
	}
}

func (cq *ProducerQuery) Get(ctx context.Context) ([]Producer, error) {
	resp, err := cq.Client.Post(ctx, "producer", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Producer{}, nil
	}

	var prods []Producer
	err = json.Unmarshal(resp.Results, &prods)
	if err != nil {
		return nil, err
	}

	return prods, nil
}

// ProducerQuery - End

// ReleaseQuery - Start
type ReleaseQuery struct {
	BaseQuery
}

func NewReleaseQuery(client *VNDBClient) *ReleaseQuery {
	return &ReleaseQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedReleaseSort},
	}
}

func (cq *ReleaseQuery) Get(ctx context.Context) ([]Release, error) {
	resp, err := cq.Client.Post(ctx, "release", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Release{}, nil
	}

	var releases []Release
	err = json.Unmarshal(resp.Results, &releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

// ReleaseQuery - End

// StaffQuery - Start
type StaffQuery struct {
	BaseQuery
}

func NewStaffQuery(client *VNDBClient) *StaffQuery {
	return &StaffQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedStaffSort},
	}
}

func (cq *StaffQuery) Get(ctx context.Context) ([]Staff, error) {
	resp, err := cq.Client.Post(ctx, "staff", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Staff{}, nil
	}

	var staff []Staff
	err = json.Unmarshal(resp.Results, &staff)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

// StaffQuery - End

// TagQuery - Start
type TagQuery struct {
	BaseQuery
}

func NewTagQuery(client *VNDBClient) *TagQuery {
	return &TagQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedTagSort},
	}
}

func (cq *TagQuery) Get(ctx context.Context) ([]Tag, error) {
	resp, err := cq.Client.Post(ctx, "tag", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Tag{}, nil
	}

	var tags []Tag
	err = json.Unmarshal(resp.Results, &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// TagQuery - End

type TraitQuery struct {
	BaseQuery
}

func NewTraitQuery(client *VNDBClient) *TraitQuery {
	return &TraitQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedTraitSort},
	}
}

func (cq *TraitQuery) Get(ctx context.Context) ([]Trait, error) {
	resp, err := cq.Client.Post(ctx, "trait", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Trait{}, nil
	}

	var traits []Trait
	err = json.Unmarshal(resp.Results, &traits)
	if err != nil {
		return nil, err
	}

	return traits, nil
}

type QuoteQuery struct {
	BaseQuery
}

func NewQuoteQuery(client *VNDBClient) *QuoteQuery {
	return &QuoteQuery{
		BaseQuery: BaseQuery{Query: &Query{}, Client: client, AllowedSort: allowedQuoteSort},
	}
}

func (cq *QuoteQuery) Get(ctx context.Context) ([]Quote, error) {
	resp, err := cq.Client.Post(ctx, "quote", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []Quote{}, nil
	}

	var quotes []Quote
	err = json.Unmarshal(resp.Results, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

// Contains
//
// Helper to check if a string is contained in a slice/array
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
