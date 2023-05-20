package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetExamsController(t *testing.T) {
	// create a new Echo instance
	e := echo.New()

	// create a request and recorder
	req := httptest.NewRequest(http.MethodGet, "/exams", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call the handler function
	err := GetExamsController(c)

	// assertions on the response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Status  int           `json:"status"`
		Message string        `json:"message"`
		Data    []models.Exam `json:"Data"`
	}

	err = json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, "success get all exams", response.Message)
	assert.NotEmpty(t, response.Data)
}

func TestGetExamController(t *testing.T) {
	// create a new Echo instance
	e := echo.New()

	// create a request and recorder with an exam id
	examID := "1"
	req := httptest.NewRequest(http.MethodGet, "/exams/"+examID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(examID)

	// call the handler function
	err := GetExamController(c)

	// assertions on the response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    models.Exam `json:"data"`
	}

	err = json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, "success get exam by id", response.Message)
	assert.NotEmpty(t, response.Data)
}
func TestCreateExamController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"exam_name":"Math Exam", "exam_code":"MATH101", "exam_date":"2022-01-01"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	if err := CreateExamController(c); err != nil {
		t.Errorf("error creating exam: %v", err)
	}

	// Assertions
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %v but got %v", http.StatusOK, rec.Code)
	}

	var resp struct {
		Status  int                    `json:"status"`
		Message string                 `json:"message"`
		Data    models.ExamResponse    `json:"data"`
	}

	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	if resp.Status != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, resp.Status)
	}

	if resp.Data.Exam_code != "MATH101" {
		t.Errorf("expected exam code MATH101 but got %v", resp.Data.Exam_code)
	}
}
func TestDeleteExamController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/exams/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Test
	if err := DeleteExamController(c); err != nil {
		t.Errorf("error deleting exam: %v", err)
	}

	// Assertions
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %v but got %v", http.StatusOK, rec.Code)
	}

	var resp struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	if resp.Status != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, resp.Status)
	}
}

