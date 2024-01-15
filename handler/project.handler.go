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

func ProjectHandlerGetAll(c *fiber.Ctx) error {
	var projects []entity.Project
	results := database.DB.Find(&projects)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := make([]dto.ProjectGetResponseDTO, len(projects))

	for i, project := range projects {
		var mentors []entity.Mentor
		var talents []entity.Talent
		var applications []entity.ProjectApply

		database.DB.Model(&project).Association("Mentors").Find(&mentors)
		database.DB.Model(&project).Association("Applications").Find(&applications)
		database.DB.Model(&project).Association("Talents").Find(&talents)

		responseDTO[i] = dto.ProjectGetResponseDTO{
			ID:           project.ID,
			Name:         project.Name,
			Description:  project.Description,
			Deadline:     project.Deadline,
			PartnerID:    project.PartnerID,
			Mentors:      mentors,
			Talents:      talents,
			Applications: applications,
		}
	}

	return c.Status(200).JSON(responseDTO)
}

// Mendapatkan semua project yang belum pernah diambil oleh talent
func ProjectHandlerGetAllAvailable(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error verifying token",
			"error":   err.Error(),
		})
	}

	var talent entity.Talent
	err = database.DB.Where("email = ?", claims["email"]).First(&talent).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting talent",
			"error":   err.Error(),
		})
	}

	var projects []entity.Project
	results := database.DB.Find(&projects)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	// Filter project yang belum pernah diambil oleh talent
	var availableProjects []entity.Project
	for _, project := range projects {
		var talents []entity.Talent
		database.DB.Model(&project).Association("Talents").Find(&talents)

		var isTaken bool = false
		for _, talent := range talents {
			if talent.ID == talent.ID {
				isTaken = true
				break
			}
		}

		if !isTaken {
			availableProjects = append(availableProjects, project)
		}
	}

	responseDTO := make([]dto.ProjectGetResponseDTO, len(availableProjects))

	for i, project := range availableProjects {
		var mentors []entity.Mentor
		var talents []entity.Talent
		var applications []entity.ProjectApply

		database.DB.Model(&project).Association("Mentors").Find(&mentors)
		database.DB.Model(&project).Association("Applications").Find(&applications)
		database.DB.Model(&project).Association("Talents").Find(&talents)

		responseDTO[i] = dto.ProjectGetResponseDTO{
			ID:           project.ID,
			Name:         project.Name,
			Description:  project.Description,
			Deadline:     project.Deadline,
			PartnerID:    project.PartnerID,
			Mentors:      mentors,
			Talents:      talents,
			Applications: applications,
		}
	}

	return c.Status(200).JSON(responseDTO)
}

func ProjectRegister(c *fiber.Ctx) error {
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

	var partner entity.Partner
	err = database.DB.Where("id = ?", project.PartnerID).First(&partner).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting partner",
			"error":   err.Error(),
		})
	}

	newProject := entity.Project{
		Name:        project.Name,
		Description: project.Description,
		Deadline:    deadline,
		PartnerID:   partner.ID,
	}

	newProjectRes := database.DB.Create(&newProject)

	if newProjectRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new Project",
			"error":   newProjectRes.Error.Error(),
		})
	}

	responseDTO := dto.ProjectRegisterResponseDTO{
		ID:           newProject.ID,
		Message:      "New Project created successfully",
		Name:         newProject.Name,
		Description:  newProject.Description,
		Deadline:     newProject.Deadline,
		PartnerID:    newProject.PartnerID,
		Mentors:      newProject.Mentors,
		Talents:      newProject.Talents,
		Applications: newProject.Applications,
	}

	return c.Status(201).JSON(responseDTO)
}

func ProjectApplyRegister(c *fiber.Ctx) error {
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

	if user.Role != "Talent" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Only talents can apply to projects",
		})
	}

	var talent entity.Talent
	err = database.DB.Where("email = ?", claims["email"]).First(&talent).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting talent",
			"error":   err.Error(),
		})
	}

	var project entity.Project
	err = database.DB.Where("id = ?", c.Params("id")).First(&project).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting project",
			"error":   err.Error(),
		})
	}

	var existingApplication entity.ProjectApply
	database.DB.Where("project_id = ? AND talent_id = ?", project.ID, talent.ID).First(&existingApplication)
	if existingApplication != (entity.ProjectApply{}) {
		return c.Status(400).JSON(fiber.Map{
			"message": "An application already exists for this project",
		})
	}

	newApplication := entity.ProjectApply{
		ProjectID: project.ID,
		TalentID:  talent.ID,
		Status:    "Pending",
	}

	newApplicationRes := database.DB.Create(&newApplication)

	if newApplicationRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new application",
			"error":   newApplicationRes.Error.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Application created successfully",
	})
}

// belom disaring pake middleware
func ProjectApplyAccept(c *fiber.Ctx) error {
	var application entity.ProjectApply
	err := database.DB.Where("id = ?", c.Params("id")).First(&application).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting application",
			"error":   err.Error(),
		})
	}

	var project entity.Project
	err = database.DB.Where("id = ?", application.ProjectID).First(&project).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting project",
			"error":   err.Error(),
		})
	}

	var talent entity.Talent
	err = database.DB.Where("id = ?", application.TalentID).First(&talent).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting talent",
			"error":   err.Error(),
		})
	}

	var existingTalent entity.Talent
	database.DB.Model(&project).Association("Talents").Find(&existingTalent)
	if existingTalent != (entity.Talent{}) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Talent already applied to project",
		})
	}

	project.Talents = append(project.Talents, talent)
	database.DB.Save(&project)

	application.Status = "Accepted"
	database.DB.Save(&application)

	return c.Status(201).JSON(fiber.Map{
		"message": "Talent accepted to project successfully",
	})
}

func ProjectApplyGetAllForProject(c *fiber.Ctx) error {
	var project entity.Project
	err := database.DB.Where("id = ?", c.Params("idProject")).First(&project).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting project",
			"error":   err.Error(),
		})
	}

	var applications []entity.ProjectApply
	database.DB.Model(&project).Association("Applications").Find(&applications)

	responseDTO := make([]dto.ProjectApplyGetResponseDTO, len(applications))

	for i, application := range applications {
		var talent entity.Talent
		database.DB.Where("id = ?", application.TalentID).First(&talent)

		responseDTO[i] = dto.ProjectApplyGetResponseDTO{
			ID:        application.ID,
			ProjectID: application.ProjectID,
			Talent:    talent,
			Status:    string(application.Status),
		}
	}

	return c.Status(200).JSON(responseDTO)
}
