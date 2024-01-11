package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID          uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Deadline    time.Time      `json:"deadline"`
	PartnerID   uuid.UUID      `json:"partner_id"`
	Mentors     []Mentor       `gorm:"many2many:mentor_projects;"` // many to many relationship
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
