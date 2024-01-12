package entity

type Talent struct {
	User
	Bio    string `json:"bio"`
	Points int    `json:"points"`
	Photo  string `json:"photo"`
}
