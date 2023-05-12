package services

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/config"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"
	"github.com/labstack/echo"
)

func GetMonitorings() (models.MonitoringResponse, error) {
	monitorings := models.Monitoring{}
	err := config.DB.Preload("Registration").Preload("Registration.Exam").Preload("Registration.Participant").Find(&monitorings).Error
	if err != nil {
		return models.MonitoringResponse{}, err
	}
	MonitoringResponse := models.MonitoringResponse{
		ID:         int(monitorings.ID),
		Exam_reg:   monitorings.Exam_reg,
		Screenshot: monitorings.Screenshot,
		Look_at:    monitorings.Look_at,
		Time:       monitorings.Time,
	}
	return MonitoringResponse, nil
}
func CreateMonitoringService(c echo.Context) (models.MonitoringResponse, error) {
	monitoring := models.Monitoring{}
	c.Bind(&monitoring)

	if err := config.DB.Preload("Registration").Save(&monitoring).Error; err != nil {
		return models.MonitoringResponse{}, err
	}
	
	MonitoringResponse := models.MonitoringResponse{
		ID:         int(monitoring.ID),
		Exam_reg:   monitoring.Exam_reg,
		Screenshot: monitoring.Screenshot,
		Look_at:    monitoring.Look_at,
		Time:       monitoring.Time,
	}
	
	return MonitoringResponse, nil
}

func CreateMonitoring(c echo.Context) (*models.MonitoringResponse, error) {
	monitoring := models.Monitoring{}
	err := c.Bind(&monitoring)
	if err != nil {
		return nil, errors.New("failed to bind monitoring data")
	}

	// Get the uploaded file
	file, err := c.FormFile("screenshot")
	if err != nil {
		return nil, errors.New("failed to get uploaded file")
	}

	// Save the file to a local directory with a unique name
	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d-%d%s", time.Now().Unix(), rand.Intn(1000), ext)
	filePath := filepath.Join("assets/monitoring", fileName)
	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open uploaded file")
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New("failed to create file")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, errors.New("failed to save uploaded file")
	}

	// Set the file path to the monitoring struct
	monitoring.Screenshot = filePath

	// Save the monitoring data to the database
	if err := config.DB.Preload("Registration").Save(&monitoring).Error; err != nil {
		return nil, errors.New("failed to save monitoring data to database")
	}

	monitoringResponse := models.MonitoringResponse{
		ID:          int(monitoring.ID),
		Exam_reg:    monitoring.Exam_reg,
		Screenshot:  monitoring.Screenshot,
		Look_at:     monitoring.Look_at,
		Time:        monitoring.Time,
	}

	return &monitoringResponse, nil
}
