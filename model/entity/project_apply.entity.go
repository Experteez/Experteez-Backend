package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	StatusOnProgress Status = "Pending"
	StatusAccepted   Status = "Accepted"
	StatusRejected   Status = "Rejected"
)

type ProjectApply struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	ProjectID uuid.UUID      `json:"project_id"`
	TalentID  uuid.UUID      `json:"talent_id"`
	Status    Status         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
