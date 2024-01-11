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
	user := new(dto.UserRegisterRequestDTO)

	if err := c.BodyParser(user); err != nil {
		return err
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

	var existingTalent entity.Talent
	res := database.DB.Where("email = ?", user.Email).First(&existingTalent)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	newTalentRes := database.DB.Create(&newTalent)

	if newTalentRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new Talent user",
			"error":   newTalentRes.Error.Error(),
		})
	}

	responseDTO := dto.UserTalentRegisterResponseDTO{
		Message:   "New Talent user created successfully",
		Role:      string(newTalent.User.Role),
		FullName:  newTalent.User.FullName,
		Email:     newTalent.User.Email,
		BirthDate: newTalent.User.BirthDate,
		Bio:       &newTalent.Bio,
		Points:    &newTalent.Points,
		Photo:     &newTalent.Photo,
	}

	return c.Status(201).JSON(responseDTO)
}
