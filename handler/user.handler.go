package handler

import (
	"Experteez-Backend/database"
	"Experteez-Backend/model/dto"
	"Experteez-Backend/model/entity"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserTalentRegister(c *fiber.Ctx) error {
	user := new(dto.TalentRegisterRequestDTO)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new Talent user",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new Talent user",
			"error":   errValidate.Error(),
		})
	}

	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing BirthDate",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		Role:      "Talent",
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  string(hashedPassword),
		BirthDate: birthDate,
	}

	newTalent := entity.Talent{
		User: newUser,
	}

	var existingUser entity.User
	res := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	newUserRes := database.DB.Create(&newUser)
	newTalentRes := database.DB.Create(&newTalent)

	if newTalentRes.Error != nil || newUserRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new Talent user",
			"error":   newTalentRes.Error.Error(),
		})
	}

	responseDTO := dto.TalentRegisterResponseDTO{
		Message:   "New Talent user created successfully",
		ID:        newTalent.ID,
		Role:      string(newTalent.User.Role),
		FullName:  newTalent.User.FullName,
		Email:     newTalent.User.Email,
		BirthDate: newTalent.User.BirthDate.Format("2006-01-02"),
		Bio:       newTalent.Bio,
		Points:    newTalent.Points,
		Photo:     newTalent.Photo,
	}

	return c.Status(201).JSON(responseDTO)
}

func UserPartnerRegister(c *fiber.Ctx) error {
	user := new(dto.PartnerRegisterRequestDTO)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new Partner user",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new Partner user",
			"error":   errValidate.Error(),
		})
	}

	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing BirthDate",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		Role:      "Partner",
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  string(hashedPassword),
		BirthDate: birthDate,
	}

	newPartner := entity.Partner{
		User:        newUser,
		Description: user.Description,
		Photo:       user.Photo,
		Projects:    user.Projects,
	}

	var existingUser entity.User
	res := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	newUserRes := database.DB.Create(&newUser)
	newPartnerRes := database.DB.Create(&newPartner)

	if newPartnerRes.Error != nil || newUserRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new Partner user",
			"error":   newPartnerRes.Error.Error(),
		})
	}

	responseDTO := dto.PartnerRegisterResponseDTO{
		Message:     "New Partner user created successfully",
		ID:          newPartner.ID,
		Role:        string(newPartner.User.Role),
		FullName:    newPartner.User.FullName,
		Email:       newPartner.User.Email,
		BirthDate:   newPartner.User.BirthDate.Format("2006-01-02"),
		Description: newPartner.Description,
		Photo:       newPartner.Photo,
		Projects:    newPartner.Projects,
	}

	return c.Status(201).JSON(responseDTO)
}

func UserMentorRegister(c *fiber.Ctx) error {
	user := new(dto.MentorRegisterRequestDTO)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new Mentor user",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new Mentor user",
			"error":   errValidate.Error(),
		})
	}

	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing BirthDate",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		Role:      "Mentor",
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  string(hashedPassword),
		BirthDate: birthDate,
	}

	newMentor := entity.Mentor{
		User:      newUser,
		Company:   user.Company,
		Specialty: user.Specialty,
		Bio:       user.Bio,
		Photo:     user.Photo,
		Talents:   user.Talents,
	}

	var existingUser entity.User
	res := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	newUserRes := database.DB.Create(&newUser)
	newMentorRes := database.DB.Create(&newMentor)

	if newMentorRes.Error != nil || newUserRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new Mentor user",
			"error":   newMentorRes.Error.Error(),
		})
	}

	responseDTO := dto.MentorRegisterResponseDTO{
		Message:   "New Mentor user created successfully",
		ID:        newMentor.ID,
		Role:      string(newMentor.User.Role),
		FullName:  newMentor.User.FullName,
		Email:     newMentor.User.Email,
		BirthDate: newMentor.User.BirthDate.Format("2006-01-02"),
		Company:   newMentor.Company,
		Specialty: newMentor.Specialty,
		Bio:       newMentor.Bio,
		Photo:     newMentor.Photo,
		Talents:   newMentor.Talents,
	}

	return c.Status(201).JSON(responseDTO)
}
