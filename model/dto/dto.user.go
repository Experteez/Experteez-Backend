package dto

import (
	"github.com/google/uuid"
	"time"
)

type UserRegisterRequestDTO struct {
	FullName  string `json:"full_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
}

type UserTalentRegisterResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	Role      string    `json:"role"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birthDate"`
	Bio       *string   `json:"bio,omitempty"`
	Points    *int      `json:"points,omitempty"`
	Photo     *string   `json:"photo,omitempty"`
}
