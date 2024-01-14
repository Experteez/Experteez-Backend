package entity

type Mentor struct {
	User
	Company   string   `json:"company"`
	Specialty string   `json:"specialty"`
	Bio       string   `json:"bio"`
	Photo     string   `json:"photo"`
	Talents   []Talent `gorm:"many2many:mentor_talents;"` // many to many relationship
}
