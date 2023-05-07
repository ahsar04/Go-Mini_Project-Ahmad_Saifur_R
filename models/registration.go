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
type ClienLoginResponse struct {
	ID              uint    `json:"id" form:"id"`
    Exam_id         uint    `json:"exam_id" from:"exam_id"`
    Participant_id  uint    `json:"participant_id" from:"participant_id"`
	Token           string  `json:"token" form:"token"`
	Ujian           string  `json:"ujian" form:"ujian"`
	Participant           string  `json:"participant" form:"participant"`

}
