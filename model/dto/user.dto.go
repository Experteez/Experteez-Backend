package dto

import (
	"Experteez-Backend/model/entity"

	"github.com/google/uuid"
)

type UserRegisterRequestDTO struct {
	FullName  string `json:"full_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
}

type TalentRegisterRequestDTO struct {
	UserRegisterRequestDTO
	Bio    string `json:"bio"`
	Points int    `json:"points"`
	Photo  string `json:"photo"`
}

type TalentRegisterResponseDTO struct {
	Message   string    `json:"message"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	BirthDate string `json:"birthDate"`
	Bio       string    `json:"bio"`
	Points    int       `json:"points"`
	Photo     string    `json:"photo"`
}

type PartnerRegisterRequestDTO struct {
	UserRegisterRequestDTO
	Description string           `json:"description" validate:"required"`
	Photo       string           `json:"photo"`
	Projects    []entity.Project `json:"projects"`
}

type PartnerRegisterResponseDTO struct {
	Message     string           `json:"message"`
	ID          uuid.UUID        `json:"id"`
	Role        string           `json:"role"`
	FullName    string           `json:"fullName"`
	Email       string           `json:"email"`
	BirthDate   string       `json:"birthDate"`
	Description string           `json:"description"`
	Photo       string           `json:"photo"`
	Projects    []entity.Project `json:"projects"`
}

type MentorRegisterRequestDTO struct {
	UserRegisterRequestDTO
	Company   string          `json:"company" validate:"required"`
	Specialty string          `json:"specialty" validate:"required"`
	Bio       string          `json:"bio"`
	Photo     string          `json:"photo"`
	Talents   []entity.Talent `json:"talents"`
}

type MentorRegisterResponseDTO struct {
	Message   string          `json:"message"`
	ID        uuid.UUID       `json:"id"`
	Role      string          `json:"role"`
	FullName  string          `json:"fullName"`
	Email     string          `json:"email"`
	BirthDate string      `json:"birthDate"`
	Company   string          `json:"company"`
	Specialty string          `json:"specialty"`
	Bio       string          `json:"bio"`
	Photo     string          `json:"photo"`
	Talents   []entity.Talent `json:"talents"`
}
