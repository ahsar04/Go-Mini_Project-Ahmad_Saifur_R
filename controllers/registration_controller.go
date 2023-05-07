package controllers

import (
	"net/http"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/config"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/middlewares"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"

	"github.com/labstack/echo"
)

// // get all Registrations
func GetRegistrationsController(c echo.Context) error {
	var registrations []models.Registration

	if err := config.DB.Preload("Exam").Preload("Participant").Find(&registrations).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all Registrations",
		"Data":   registrations,
	})
}
// // get Registration by id
func GetRegistrationController(c echo.Context) error {
	registration := models.Registration{}
	RegistrationID := c.Param("id")

	if err := config.DB.Preload("Exam").Preload("Participant").First(&registration, RegistrationID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get Registration by id",
		"data":    registration,
	})
}
// create new Registration
func CreateRegistrationController(c echo.Context) error {
	registration := models.Registration{}
	c.Bind(&registration)
	

	if err := config.DB.Preload("Exam").Preload("Participant").Save(&registration).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	RegistrationResponse:=models.RegistrationResponse{int(registration.ID),int(registration.Exam_id),int(registration.Participant_id)}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new Registration",
		"data":    RegistrationResponse,
	})
}
// delete Registration by id
func DeleteRegistrationController(c echo.Context) error {
	registration := models.Registration{}
	RegistrationID := c.Param("id")

	if err := config.DB.Preload("Exam").Preload("Participant").First(&registration, RegistrationID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&registration).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success delete Registration by id",
	})
}
// // update Registration by id
func UpdateRegistrationController(c echo.Context) error {
	registration := models.Registration{}
	RegistrationID := c.Param("id")

	if err := config.DB.First(&registration, RegistrationID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&registration); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&registration).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	RegistrationResponse:=models.RegistrationResponse{int(registration.ID),int(registration.Exam_id),int(registration.Participant_id)}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update Registration by id",
		"data":    RegistrationResponse,
	})
}

func LoginClientController(c echo.Context) error {
	req := struct {
        Exam_code    string `json:"exam_code"`
        Email string `json:"email"`
    }{}
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }
    registration := models.Registration{}
    // Get user from database by email
    if err := config.DB.Preload("Exam", "exam_code = ?",req.Exam_code).Preload("Participant", "email = ?",req.Email).First(&registration).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "exam not found",
            "error":   err.Error(),
        })
    }
    // Generate JWT token
    token, err := middlewares.CreateClientToken(int(registration.ID), registration.Participant.Name)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "fail login",
            "error":   err.Error(),
        })
    }

    // Return user data and JWT token
    dataResponse := models.ClienLoginResponse{registration.ID,registration.Exam_id,registration.Participant_id,token,registration.Exam.Exam_name,registration.Participant.Name}
    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success login",
        "data":    dataResponse,
    })
}

