package controllers

import (
	"net/http"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/config"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"

	"github.com/labstack/echo"
)

// get all participants
func GetParticipantsController(c echo.Context) error {
	participants:=[]models.Participant{}


	if err := config.DB.Find(&participants).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	participantsResponse := make([]models.ParticipantResponse, len(participants))
	for i, participants := range participants {
	participantsResponse[i]=models.ParticipantResponse{int(participants.ID),participants.Name,participants.Gender,participants.Email,participants.Phone}
    }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all participants",
		"Data":   participantsResponse,
	})
}
// get participant by id
func GetParticipantController(c echo.Context) error {
	participant := models.Participant{}
	participantID := c.Param("id")

	if err := config.DB.First(&participant, participantID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	participantsResponse:=models.ParticipantResponse{int(participant.ID),participant.Name,participant.Gender,participant.Email,participant.Phone}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get participant by id",
		"data":    participantsResponse,
	})
}
// create new participant
func CreateParticipantController(c echo.Context) error {
	participant := models.Participant{}
	c.Bind(&participant)
	if participant.Email=="" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusBadRequest,
			"message": "email tidak boleh kosong",
		})
	}
	if err := config.DB.Find(&participant, "email =?", participant.Email).Error; err == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusBadRequest,
			"message": "email sudah terdaftar",
		})
	}
	if err := config.DB.Save(&participant).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	participantsResponse:=models.ParticipantResponse{int(participant.ID),participant.Name,participant.Gender,participant.Email,participant.Phone}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new participant",
		"data":    participantsResponse,
	})
}
// delete participant by id
func DeleteParticipantController(c echo.Context) error {
	participant := models.Participant{}
	participantID := c.Param("id")

	if err := config.DB.First(&participant, participantID).Error; err != nil {
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
	participantsResponse:=models.ParticipantResponse{int(participant.ID),participant.Name,participant.Gender,participant.Email,participant.Phone}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update participant by id",
		"data":    participantsResponse,
	})
}