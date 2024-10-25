package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// LoginHandler handles the login requests
func LoginHandler(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dummy check for the sake of example
	if user.Username == "newuser" && user.Password == "securePassword123" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
}

// LogoutHandler handles the logout requests
func LogoutHandler(c *gin.Context) {
	// Assuming we clear session or token here
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// TestLogin tests the LoginHandler function
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/login", LoginHandler)

	userJSON := `{"username": "newuser", "password": "securePassword123"}`

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(userJSON)))
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

	assert.Equal(t, "Login successful", response["message"])
}

// TestLogout tests the LogoutHandler function
func TestLogout(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/logout", LogoutHandler) // Using POST for logout, can also be GET

	req, err := http.NewRequest("POST", "/logout", nil) // No body needed
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]string
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Logout successful", response["message"])
}
