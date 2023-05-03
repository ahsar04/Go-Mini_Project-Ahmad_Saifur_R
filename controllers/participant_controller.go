package controllers

import (
	"code_structure/config"
	"code_structure/models"
	"net/http"

	"github.com/labstack/echo"
)

// get all participants
func GetParticipantsController(c echo.Context) error {
	var participants []models.Participant


	if err := config.DB.Preload("User").Find(&participants).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all participants",
		"Data":   participants,
	})
}
// get participant by id
func GetParticipantController(c echo.Context) error {
	participant := models.Participant{}
	participantID := c.Param("id")

	if err := config.DB.Preload("User").First(&participant, participantID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get participant by id",
		"data":    participant,
	})
}
// create new participant
func CreateParticipantController(c echo.Context) error {
	participant := models.Participant{}
	c.Bind(&participant)


	if err := config.DB.Preload("User").Save(&participant).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new participant",
		"data":    participant,
	})
}
// delete participant by id
func DeleteParticipantController(c echo.Context) error {
	participant := models.Participant{}
	participantID := c.Param("id")

	if err := config.DB.Preload("User").First(&participant, participantID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&participant).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success delete participant by id",
	})
}
// update participant by id
func UpdateParticipantController(c echo.Context) error {
	participant := models.Participant{}
	participantID := c.Param("id")

	if err := config.DB.First(&participant, participantID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&participant); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&participant).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update participant by id",
		"data":    participant,
	})
}