package models

import (
	"github.com/jinzhu/gorm"
)

type Participant struct {
	gorm.Model	
    Name     string `json:"name" form:"name"`
    Gender   string `json:"gender" form:"gender"`
    Email    string `json:"email" form:"email"`
    Phone    string `json:"phone" form:"phone"`
}
type ParticipantResponse struct {
	ID    	 int    `json:"id" form:"id"`	
    Name     string `json:"name" form:"name"`
    Gender   string `json:"gender" form:"gender"`
    Email    string `json:"email" form:"email"`
    Phone    string `json:"phone" form:"phone"`
}