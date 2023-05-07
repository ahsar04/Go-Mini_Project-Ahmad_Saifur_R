package models

import "github.com/jinzhu/gorm"

type Registration struct {
    gorm.Model
    Exam_id         uint        `json:"exam_id" from:"exam_id"`
    Participant_id  uint        `json:"participant_id" from:"participant_id"`
    Exam            Exam        `gorm:"foreignKey:Exam_id"`
    Participant     Participant `gorm:"foreignKey:Participant_id"`
}
type RegistrationResponse struct {
	ID    int    `json:"id" form:"id"`
	Exam_id   		int    `json:"exam_id" from:"exam_id"`
	Participant_id  int    `json:"participant_id" from:"participant_id"`
}
