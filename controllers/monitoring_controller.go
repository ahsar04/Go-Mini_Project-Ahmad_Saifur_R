package controllers

import (
	"net/http"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/services"

	"github.com/labstack/echo"
)

// get all Monitorings
func GetMonitoringsController(c echo.Context) error {
	monitorings, err := services.GetMonitorings()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     http.StatusOK,
		"message":    "success get all data",
		"monitorings": monitorings,
	})
}


// create new Monitoring
func CreateMonitoringController(c echo.Context) error {
    result, err := services.CreateMonitoringService(c)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success create new Monitoring",
        "data":    result,
    })
}

func CreateMonitoringController2(c echo.Context) error {
	monitoringResponse, err := services.CreateMonitoring(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "success create new Monitoring",
		"data":    monitoringResponse,
	})
}
