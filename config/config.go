package config

import (
	"fmt"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
  	DB *gorm.DB
)


func init() {
	InitDB()
	InitialMigration()
}


type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}


func InitDB() {


	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "mini_project",
	}


	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)


	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
  	DB.AutoMigrate(&models.User{})
  	DB.AutoMigrate(&models.Exam{})
  	DB.AutoMigrate(&models.Participant{})
  	DB.AutoMigrate(&models.Registration{})
  	DB.AutoMigrate(&models.Monitoring{})
}
