package controllers

import (
	"code_structure/config"
	"code_structure/models"
	"net/http"

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