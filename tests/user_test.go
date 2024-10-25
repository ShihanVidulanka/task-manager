// tests/user_test.go
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

// TestCreateUser tests user registration
func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()
	userJSON := `{"username": "newuser", "password": "securePassword123"}`

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(userJSON)))
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

	assert.Equal(t, "User registered successfully", response["message"])
}

// TestGetUser tests fetching a user by ID
func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/users/1", nil)
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
	assert.Equal(t, "newuser", response["username"])
}

// TestGetUser_NotFound tests fetching a non-existent user
func TestGetUser_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/users/999", nil)
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

	assert.Equal(t, "User not found", response["error"])
}

// TestGetAllUsers tests fetching all users
func TestGetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/users", nil)
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

	assert.IsType(t, []interface{}{}, response["users"])
	assert.Equal(t, 2, len(response["users"].([]interface{})))
}

// TestUpdateUser tests updating a user
func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()
	userJSON := `{"username": "updateduser", "password": "newpassword"}`

	req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer([]byte(userJSON)))
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

	assert.Equal(t, "User updated successfully", response["message"])
}

// TestDeleteUser tests deleting a user
func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := router.SetupRouter()

	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}
