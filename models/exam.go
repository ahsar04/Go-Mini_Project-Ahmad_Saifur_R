package models

import "github.com/jinzhu/gorm"

type Exam struct {
	gorm.Model
	Exam_name    string `json:"exam_name" form:"exam_name"`
	Exam_code    string `json:"exam_code" form:"exam_code"`
	Exam_date    string `json:"exam_date" form:"exam_date"`
}
