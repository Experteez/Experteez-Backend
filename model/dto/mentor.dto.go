package dto

import (
	"Experteez-Backend/model/entity"

	"github.com/google/uuid"
)

type MentorGetResponseDTO struct {
	ID        uuid.UUID       `json:"id"`
	FullName  string          `json:"name"`
	Email     string          `json:"email"`
	Company   string          `json:"company"`
	Specialty string          `json:"specialty"`
	Bio       string          `json:"bio"`
	Photo     string          `json:"photo"`
	Talents   []entity.Talent `json:"talents"`
}
