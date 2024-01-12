package entity

type Mentor struct {
	User
	Bio       string   `json:"bio"`
	Specialty string   `json:"specialty"`
	Photo     string   `json:"photo"`
	Talents   []Talent `gorm:"many2many:mentor_talents;"` // many to many relationship
}
