// tests/task_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task-manager/router" // Adjust import path as necessary

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestCreateTask tests task creation
func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()
	taskJSON := `{"title": "New Task", "description": "Task description", "user_id": 1, "status": "pending"}`
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(taskJSON)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var response map[string]string
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Task created successfully", response["message"])
}

// TestGetTask tests fetching a task by ID
func TestGetTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1.0, response["id"])
	assert.Equal(t, "Sample Task", response["title"])
}

// TestGetTask_NotFound tests fetching a non-existent task
func TestGetTask_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/tasks/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	var response map[string]string
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Task not found", response["error"])
}

// TestGetAllTasks tests fetching all tasks
func TestGetAllTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, []interface{}{}, response["tasks"])
	assert.Equal(t, 2, len(response["tasks"].([]interface{})))
}

// TestUpdateTask tests updating a task
func TestUpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()
	taskJSON := `{"title": "Updated Task"}`

	req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer([]byte(taskJSON)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]string
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Task updated successfully", response["message"])
}

// TestDeleteTask tests deleting a task
func TestDeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("DELETE", "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}
