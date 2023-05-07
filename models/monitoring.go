package models

import "github.com/jinzhu/gorm"

type Monitoring struct {
    gorm.Model
    Exam_reg    	uint    `json:"exam_reg" form:"exam_reg"`
    Registration 	Registration `gorm:"foreignKey:Exam_reg"`
    Screenshot  	string  `json:"screenshot" form:"screenshot"`
    Look_at     	string  `json:"look_at" form:"look_at"`    
    Time        	string  `json:"time" form:"time"`
}

type MonitoringResponse struct {
	ID    		int    `json:"id" form:"id"`
	Exam_reg   	uint    `json:"exam_reg" from:"exam_reg"`
	Screenshot  string    `json:"screenshot" from:"screenshot"`
	Look_at    	string    `json:"look_at" from:"look_at"`	
	Time    	string    `json:"time" from:"time"`
}
