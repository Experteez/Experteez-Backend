package handler

import (
	"Experteez-Backend/database"
	"Experteez-Backend/model/dto"
	"Experteez-Backend/model/entity"
	"Experteez-Backend/utils"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ProjectHandlerGetAll (c *fiber.Ctx) error {
	var projects []entity.Project
	results := database.DB.Find(&projects)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message":   results.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"results": len(projects),
		"projects": projects,
	})
}

func ProjectRegister (c *fiber.Ctx) error {
	project := new(dto.ProjectRegisterRequestDTO)

	if err := c.BodyParser(project); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new Project",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(project)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new Project",
			"error":   errValidate.Error(),
		})
	}

	deadline, err := time.Parse("2006-01-02", project.Deadline)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing Deadline",
			"error":   err.Error(),
		})
	}

	if deadline.Before(time.Now()) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Deadline has already passed",
		})
	}

	authHeader := c.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error verifying token",
			"error":   err.Error(),
		})
	}

	var user entity.User
	err = database.DB.Where("email = ?", claims["email"]).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting user",
			"error":   err.Error(),
		})
	}

	newProject := entity.Project{
		Name:        project.Name,
		Description: project.Description,
		Deadline:    deadline,
		PartnerID:   user.ID,
	}

	newProjectRes := database.DB.Create(&newProject)

	// Program will be returning error because violates FK constraint (fk_partners_project)
	if newProjectRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new Project",
			"error":   newProjectRes.Error.Error(),
		})
	}

	responseDTO := dto.ProjectRegisterResponseDTO{
		ID: 			newProject.ID,
		Message:   		"New Project created successfully",
		Name:	   		newProject.Name,
		Description:	newProject.Description,
		Deadline:		newProject.Deadline,
		PartnerID: 		newProject.PartnerID,
	}

	return c.Status(201).JSON(responseDTO)
}