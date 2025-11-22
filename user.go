package vndb_go

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
