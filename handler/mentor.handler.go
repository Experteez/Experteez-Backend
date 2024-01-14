package handler

import (
	"Experteez-Backend/database"
	"Experteez-Backend/model/dto"
	"Experteez-Backend/model/entity"

	"github.com/gofiber/fiber/v2"
)

func MentorHandlerGetAll(c *fiber.Ctx) error {
	var mentors []entity.Mentor
	results := database.DB.Find(&mentors)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := make([]dto.MentorGetResponseDTO, len(mentors))

	for i, mentor := range mentors {
		responseDTO[i] = dto.MentorGetResponseDTO{
			ID:        mentor.ID,
			FullName:  mentor.FullName,
			Email:     mentor.Email,
			Company:   mentor.Company,
			Specialty: mentor.Specialty,
			Bio:       mentor.Bio,
			Photo:     mentor.Photo,
			Talents:   mentor.Talents,
		}
	}

	return c.Status(200).JSON(responseDTO)
}
