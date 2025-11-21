package vndb_go

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/Lumminal/vndb-go/helper"
	"github.com/Lumminal/vndb-go/types"
)

// This file includes all query related functions and structs

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
	Query       *types.Query
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

func (cq *BaseQuery) Filters(filter types.Filter) {
	cq.Query.Filters = filter
}

func (cq *BaseQuery) Sort(sort ...string) {
	for _, str := range sort {
		if !helper.Contains(cq.AllowedSort, str) {
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedVnSort},
	}
}

func (vnq *VNQuery) Get(ctx context.Context) ([]types.Vn, error) {
	resp, err := vnq.Client.Post(ctx, "vn", vnq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Vn{}, nil
	}

	var vns []types.Vn
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedCharacterSort},
	}
}

func (cq *CharQuery) Get(ctx context.Context) ([]types.Character, error) {
	resp, err := cq.Client.Post(ctx, "character", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Character{}, nil
	}

	var chars []types.Character
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedProducerSort},
	}
}

func (cq *ProducerQuery) Get(ctx context.Context) ([]types.Producer, error) {
	resp, err := cq.Client.Post(ctx, "producer", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Producer{}, nil
	}

	var prods []types.Producer
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedReleaseSort},
	}
}

func (cq *ReleaseQuery) Get(ctx context.Context) ([]types.Release, error) {
	resp, err := cq.Client.Post(ctx, "release", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Release{}, nil
	}

	var releases []types.Release
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedStaffSort},
	}
}

func (cq *StaffQuery) Get(ctx context.Context) ([]types.Staff, error) {
	resp, err := cq.Client.Post(ctx, "staff", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Staff{}, nil
	}

	var staff []types.Staff
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedTagSort},
	}
}

func (cq *TagQuery) Get(ctx context.Context) ([]types.Tag, error) {
	resp, err := cq.Client.Post(ctx, "tag", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Tag{}, nil
	}

	var tags []types.Tag
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedTraitSort},
	}
}

func (cq *TraitQuery) Get(ctx context.Context) ([]types.Trait, error) {
	resp, err := cq.Client.Post(ctx, "trait", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Trait{}, nil
	}

	var traits []types.Trait
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
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedQuoteSort},
	}
}

func (cq *QuoteQuery) Get(ctx context.Context) ([]types.Quote, error) {
	resp, err := cq.Client.Post(ctx, "quote", cq.Query)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.Quote{}, nil
	}

	var quotes []types.Quote
	err = json.Unmarshal(resp.Results, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

type UListQueryRequest struct {
	BaseQuery
}

type UListQuery struct {
	BaseQuery
	User string `json:"user,omitempty"`
}

func (ul *UListQuery) SetUser(id string) {
	ul.User = id
}

func (ul *UListQuery) Get(ctx context.Context) ([]types.UList, error) {

	ulistquery := &types.UlistQueryRequest{
		Query: *ul.Query,
		User:  ul.User,
	}

	resp, err := ul.Client.PostUlist(ctx, "ulist", ulistquery)
	if err != nil {
		return nil, err
	}

	if resp.Results == nil || len(resp.Results) == 0 {
		return []types.UList{}, nil
	}

	var ulists []types.UList
	err = json.Unmarshal(resp.Results, &ulists)
	if err != nil {
		return nil, err
	}

	return ulists, nil
}

func NewUlistQuery(client *VNDBClient) *UListQuery {
	return &UListQuery{
		BaseQuery: BaseQuery{Query: &types.Query{}, Client: client, AllowedSort: allowedUlistSort},
	}
}
