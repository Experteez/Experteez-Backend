package dto

import (
	"Experteez-Backend/model/entity"
	"time"

	"github.com/google/uuid"
)

type ProjectRegisterRequestDTO struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Deadline    string    `json:"deadline" validate:"required"`
	PartnerID   uuid.UUID `json:"partner_id" validate:"required"`
}

type ProjectRegisterResponseDTO struct {
	ID           uuid.UUID             `json:"id"`
	Message      string                `json:"message"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Deadline     time.Time             `json:"deadline"`
	PartnerID    uuid.UUID             `json:"partner_id"`
	Mentors      []entity.Mentor       `json:"mentors"`
	Talents      []entity.Talent       `json:"talents"`
	Applications []entity.ProjectApply `json:"applications"`
}

type ProjectGetResponseDTO struct {
	ID           uuid.UUID             `json:"id"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Deadline     time.Time             `json:"deadline"`
	PartnerID    uuid.UUID             `json:"partner_id"`
	Mentors      []entity.Mentor       `json:"mentors"`
	Talents      []entity.Talent       `json:"talents"`
	Applications []entity.ProjectApply `json:"applications"`
}

type ProjectApplyGetResponseDTO struct {
	ID        uuid.UUID     `json:"id"`
	ProjectID uuid.UUID     `json:"project_id"`
	Talent    entity.Talent `json:"talent"`
	Status    string        `json:"status"`
}
