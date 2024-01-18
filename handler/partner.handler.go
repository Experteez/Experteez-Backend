package handler

import (
	"Experteez-Backend/database"
	"Experteez-Backend/model/dto"
	"Experteez-Backend/model/entity"

	"github.com/gofiber/fiber/v2"
)

func PartnerHandlerGetAll(c *fiber.Ctx) error {
	var partners []entity.Partner
	results := database.DB.Find(&partners)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := make([]dto.PartnerGetResponseDTO, len(partners))

	for i, partner := range partners {
		var projects []entity.Project
		database.DB.Model(&partner).Association("Projects").Find(&projects)

		responseDTO[i] = dto.PartnerGetResponseDTO{
			ID:          partner.ID,
			FullName:    partner.FullName,
			Email:       partner.Email,
			Description: partner.Description,
			Photo:       partner.Photo,
			Projects:    projects,
		}
	}

	return c.Status(200).JSON(responseDTO)
}
