package models

import (
	"github.com/jinzhu/gorm"
)

type Participant struct {
	gorm.Model	
    Name     string `json:"judul" form:"judul"`
    Gender   string `json:"gender" form:"gender"`
    Email    string `json:"email" form:"email"`
    Phone    string `json:"phone" form:"phone"`
}