package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"servicemanager/model"
	"testing"
)

var router *gin.Engine

func init() {
	router = InitEngine()
}

func performRequestBody(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetList(t *testing.T) {

	//router := InitEngine()
	w := performRequest(router, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	var response []model.Service
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(response))
}

func TestStartService(t *testing.T) {
	cmd := model.Command{
		Command: "cmd",
		Args:    []string{"/c", "dir"},
	}
	body, _ := json.Marshal(cmd)
	w := performRequestBody(router, "POST", "/", body)
	assert.Equal(t, http.StatusCreated, w.Code)
	var response model.Service

	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, response.PID)

	// wrong command

	body, _ = json.Marshal(model.Command{
		Command: "fakecmn",
		Args:    nil,
	})
	w = performRequestBody(router, "POST", "/", body)
	assert.Equal(t, http.StatusNotAcceptable, w.Code)

	// wrong param
	wp := []byte("Fake param")
	w = performRequestBody(router, "POST", "/", wp)
	assert.Equal(t, http.StatusBadRequest, w.Code)

}
