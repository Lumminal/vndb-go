package vndb_go

type FilterField string
type Filter []interface{}

// Note: Some filter fields don't accept some operations
// Refer to the API to know which ones are allowed

const (
	ID               FilterField = "id"
	Released         FilterField = "released"
	Search           FilterField = "search"
	Lang             FilterField = "lang"
	Olang            FilterField = "olang"
	Platform         FilterField = "platform"
	Length           FilterField = "length"
	Rating           FilterField = "rating"
	VoteCount        FilterField = "votecount"
	HasDescription   FilterField = "has_description"
	HasAnime         FilterField = "has_anime"
	HasScreenshot    FilterField = "has_screenshot"
	HasReview        FilterField = "has_review"
	DevStatus        FilterField = "devstatus"
	TagFilter        FilterField = "tag"
	Dtag             FilterField = "dtag"
	AnimeId          FilterField = "anime_id"
	LabelFilter      FilterField = "label"
	ReleaseFilter    FilterField = "release"
	CharacterFilter  FilterField = "character"
	StaffFilter      FilterField = "staff"
	DeveloperFilter  FilterField = "developer"
	ResolutionFilter FilterField = "resolution"
	ResolutionAspect FilterField = "resolution_aspect"
	Minage           FilterField = "minage"
	Medium           FilterField = "medium"
	Voiced           FilterField = "voiced"
	Engine           FilterField = "engine"
	Rtype            FilterField = "rtype"
	Extlink          FilterField = "extlink"
	Patch            FilterField = "patch"
	Freeware         FilterField = "freeware"
	Uncensored       FilterField = "uncensored"
	Official         FilterField = "official"
	HasEro           FilterField = "has_ero"
	VnFilter         FilterField = "vn"
	ProducerFilter   FilterField = "producer"
	BloodType        FilterField = "blood_type"
	Sex              FilterField = "sex"
	SexSpoil         FilterField = "sex_spoil"
	Gender           FilterField = "gender"
	GenderSpoil      FilterField = "gender_spoil"
	Height           FilterField = "height"
	Weight           FilterField = "weight"
	Bust             FilterField = "bust"
	Waist            FilterField = "waist"
	Hips             FilterField = "hips"
	Cup              FilterField = "cup"
	Age              FilterField = "age"
	TraitFilter      FilterField = "trait"
	DTrait           FilterField = "dtrait"
	Birthday         FilterField = "birthday"
	Seiyuu           FilterField = "seiyuu"
	Aid              FilterField = "aid"
	Role             FilterField = "role"
	IsMain           FilterField = "is_main"
	Category         FilterField = "category"
	Random           FilterField = "random"
)

func (f FilterField) Equal(compareWith string) Filter {
	return Filter{f, "=", compareWith}
}

func (f FilterField) GreaterThan(compareWith string) Filter {
	return Filter{f, ">", compareWith}
}

func (f FilterField) LessThan(compareWith string) Filter {
	return Filter{f, "<", compareWith}
}

func (f FilterField) GreaterOrEqualThan(compareWith string) Filter {
	return Filter{f, ">=", compareWith}
}

func (f FilterField) LessOrEqualThan(compareWith string) Filter {
	return Filter{f, "<=", compareWith}
}

func (f FilterField) NotEqual(compareWith string) Filter {
	return Filter{f, "!=", compareWith}
}

func And(filter ...Filter) Filter {
	result := Filter{"and"}
	for _, fltr := range filter {
		result = append(result, fltr)
	}
	return result
}

func Or(filter ...Filter) Filter {
	result := Filter{"or"}
	for _, fltr := range filter {
		result = append(result, fltr)
	}
	return result
}
