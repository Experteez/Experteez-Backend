package dto

import (
	"github.com/google/uuid"
	"time"
)

type ProjectRegisterRequestDTO struct {
	Name        string		`json:"name" validate:"required"`
	Description string 		`json:"description" validate:"required"`
	Deadline    string  	`json:"deadline" validate:"required"`
	Mentors		[]string	`json:"mentors"`
}

type ProjectRegisterResponseDTO struct {
	ID        	uuid.UUID 	`json:"id"`
	Message   	string    	`json:"message"`
	Name  	  	string    	`json:"name"`
	Description string    	`json:"description"`
	Deadline 	time.Time 	`json:"deadline"`
	PartnerID 	uuid.UUID	`json:"partner_id"`
}