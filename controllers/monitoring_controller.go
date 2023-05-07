package controllers

import (
	"net/http"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/config"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"

	"github.com/labstack/echo"
)

// get all Monitorings
func GetMonitoringsController(c echo.Context) error {
	monitorings := []models.Monitoring{}
	err := config.DB.Preload("Registration").Preload("Registration.Exam").Preload("Registration.Participant").Find(&monitorings).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     http.StatusOK,
		"message":    "success get all data",
		"monitorings": monitorings,
	})
}

// // // get Monitoring by id
func GetMonitoringController(c echo.Context) error {
	monitorings := models.Monitoring{}
	Exam_reg := c.Param("exam_reg")
	err := config.DB.Preload("Registration").Find(&monitorings, "exam_reg = ?", Exam_reg).Error;
	if  err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get Monitoring by exam_reg",
		"data":    monitorings,
	})
}
// create new Monitoring
func CreateMonitoringController(c echo.Context) error {
	monitoring := models.Monitoring{}
	c.Bind(&monitoring)
	

	if err := config.DB.Preload("Registration").Save(&monitoring).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	MonitoringResponse:=models.MonitoringResponse{int(monitoring.ID),monitoring.Exam_reg,monitoring.Screenshot,monitoring.Look_at,monitoring.Time}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new Monitoring",
		"data":    MonitoringResponse,
	})
}
// // delete Monitoring by id
// func DeleteMonitoringController(c echo.Context) error {
// 	monitoring := models.Monitoring{}
// 	MonitoringID := c.Param("id")

// 	if err := config.DB.Preload("Exam").Preload("Participant").First(&monitoring, MonitoringID).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	if err := config.DB.Delete(&monitoring).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": http.StatusOK,
// 		"message": "success delete Monitoring by id",
// 	})
// }
// // // update Monitoring by id
// func UpdateMonitoringController(c echo.Context) error {
// 	monitoring := models.Monitoring{}
// 	MonitoringID := c.Param("id")

// 	if err := config.DB.First(&monitoring, MonitoringID).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	if err := c.Bind(&monitoring); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	if err := config.DB.Save(&monitoring).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	MonitoringResponse:=models.MonitoringResponse{int(monitoring.ID),int(monitoring.Exam_id),int(monitoring.Participant_id)}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": http.StatusOK,
// 		"message": "success update Monitoring by id",
// 		"data":    MonitoringResponse,
// 	})
// }