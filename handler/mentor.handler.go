package handler

import (
	"Experteez-Backend/database"
	"Experteez-Backend/model/dto"
	"Experteez-Backend/model/entity"
	"Experteez-Backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func MentorAssign(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error verifying token",
			"error":   err.Error(),
		})
	}

	var existingTalent entity.Talent
	if err := database.DB.Where("email = ?", claims["email"]).First(&existingTalent).Error; err != nil {
		return handleError(c, "Error getting talent", err)
	}

	mentor, project, talent, err := getMentorProjectTalent(c)
	if err != nil {
		return handleError(c, "Error getting mentor, project, or talent", err)
	}

	if isMentorAssigned(mentor, project) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Mentor already assigned to project",
		})
	}

	if isTalentAssigned(talent, project) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Talent already assigned to project",
		})
	}

	assignMentorTalentToProject(mentor, talent, project)

	return c.Status(200).JSON(fiber.Map{
		"message": "Mentor assigned successfully",
	})
}

func getMentorProjectTalent(c *fiber.Ctx) (*entity.Mentor, *entity.Project, *entity.Talent, error) {
	id := c.Params("id")
	idProject := c.Params("idProject")

	var mentor entity.Mentor
	if err := database.DB.Where("id = ?", id).First(&mentor).Error; err != nil {
		return nil, nil, nil, err
	}

	var project entity.Project
	if err := database.DB.Where("id = ?", idProject).First(&project).Error; err != nil {
		return nil, nil, nil, err
	}

	var talent entity.Talent
	return &mentor, &project, &talent, nil
}

func isMentorAssigned(mentor *entity.Mentor, project *entity.Project) bool {
	var mentors []entity.Mentor
	database.DB.Model(&project).Association("Mentors").Find(&mentors)

	for _, m := range mentors {
		if m.ID == mentor.ID {
			return true
		}
	}

	return false
}

func isTalentAssigned(talent *entity.Talent, project *entity.Project) bool {
	var talents []entity.Talent
	database.DB.Model(&project).Association("Talents").Find(&talents)

	for _, t := range talents {
		if t.ID == talent.ID {
			return true
		}
	}

	return false
}

func assignMentorTalentToProject(mentor *entity.Mentor, talent *entity.Talent, project *entity.Project) {
	database.DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&project).Association("Mentors").Append(mentor)
		tx.Model(&project).Association("Talents").Append(talent)
		tx.Model(mentor).Association("Projects").Append(project)
		tx.Model(mentor).Association("Talents").Append(talent)
		return nil
	})
}

func handleError(c *fiber.Ctx, message string, err error) error {
	return c.Status(400).JSON(fiber.Map{
		"message": message,
		"error":   err.Error(),
	})
}
