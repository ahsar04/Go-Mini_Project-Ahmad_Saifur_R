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
