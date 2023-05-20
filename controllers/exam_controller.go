package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/config"
	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"

	"github.com/labstack/echo"
)

// get all users
func GetExamsController(c echo.Context) error {
	exams :=[]models.Exam{}


	if err := config.DB.Find(&exams).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
    examResponse := make([]models.ExamResponse, len(exams))
	for i, exams := range exams {
		examResponse[i]=models.ExamResponse{int(exams.ID), exams.Exam_name, exams.Exam_code,exams.Exam_date}
        
    }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all exams",
		"Data":   examResponse,
	})
}
// get exam by id
func GetExamController(c echo.Context) error {
	exam := models.Exam{}
	examID := c.Param("id")

	if err := config.DB.First(&exam, examID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	examResponse:=models.ExamResponse{int(exam.ID), exam.Exam_name, exam.Exam_code,exam.Exam_date}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get exam by id",
		"data":    examResponse,
	})
}
// create new exam
func CreateExamController(c echo.Context) error {
	exam := models.Exam{}
	c.Bind(&exam)
	if exam.Exam_code=="" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusBadRequest,
			"message": "exam code tidak boleh kosong",
		})
	}
	if err := config.DB.Find(&exam, "exam_code =?", exam.Exam_code).Error; err == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusBadRequest,
			"message": "exam code sudah terdaftar",
		})
	}
	if err := config.DB.Save(&exam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	examResponse:=models.ExamResponse{int(exam.ID), exam.Exam_name, exam.Exam_code,exam.Exam_date}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new exam",
		"data":    examResponse,
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
	examResponse:=models.ExamResponse{int(exam.ID), exam.Exam_name, exam.Exam_code,exam.Exam_date}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update exam by id",
		"data":    examResponse,
	})
}
func TestUpdateExamController(t *testing.T) {
	// Initialize a new Echo instance
	e := echo.New()

	// Create a new request
	req := httptest.NewRequest(http.MethodPut, "/exams/1", strings.NewReader(`{"exam_name":"New Exam Name","exam_code":"NE","exam_date":"2023-06-01"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the controller function
	if err := UpdateExamController(c); err != nil {
		t.Fatalf("failed to update exam: %v", err)
	}

	// Assert the response
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	expectedBody := `{"data":{"exam_code":"NE","exam_date":"2023-06-01","exam_id":1,"exam_name":"New Exam Name"},"message":"success update exam by id","status":200}`
	if rec.Body.String() != expectedBody {
		t.Errorf("unexpected response body:\n expected: %s\n got: %s", expectedBody, rec.Body.String())
	}
}
