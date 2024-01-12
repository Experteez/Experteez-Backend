package entity

type Partner struct {
	User
	Description string    `json:"description"`
	Photo       string    `json:"photo"`
	Projects    []Project `gorm:"foreignKey:PartnerID"`
}
