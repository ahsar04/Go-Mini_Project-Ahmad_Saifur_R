package controllers

import (
	"code_structure/config"
	"code_structure/models"
	"net/http"

	"github.com/labstack/echo"
)

// get all users
func GetExamsController(c echo.Context) error {
	var exams []models.Exam


	if err := config.DB.Find(&exams).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all exams",
		"Data":   exams,
	})
}
// get exam by id
func GetExamController(c echo.Context) error {
	exam := models.Exam{}
	examID := c.Param("id")

	if err := config.DB.First(&exam, examID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get exam by id",
		"data":    exam,
	})
}
// create new exam
func CreateExamController(c echo.Context) error {
	exam := models.Exam{}
	c.Bind(&exam)


	if err := config.DB.Save(&exam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new exam",
		"data":    exam,
	})
}
// delete exam by id
func DeleteExamController(c echo.Context) error {
	exam := models.Exam{}
	examID := c.Param("id")

	if err := config.DB.First(&exam, examID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&exam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success delete exam by id",
	})
}
// update exam by id
func UpdateExamController(c echo.Context) error {
	exam := models.Exam{}
	examID := c.Param("id")

	if err := config.DB.First(&exam, examID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&exam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&exam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update exam by id",
		"data":    exam,
	})
}