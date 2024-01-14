package dto

import (
	"Experteez-Backend/model/entity"

	"github.com/google/uuid"
)

type PartnerGetResponseDTO struct {
	ID          uuid.UUID        `json:"id"`
	FullName    string           `json:"name"`
	Email       string           `json:"email"`
	Description string           `json:"description"`
	Photo       string           `json:"photo"`
	Projects    []entity.Project `json:"projects"`
}
