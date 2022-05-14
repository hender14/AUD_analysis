package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

type Testuser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// main関数
func TestMain(t *testing.T) {
	router := setupRouter()
  // router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}

func TestLogin(t *testing.T) {
	router := setupRouter()
	input := Testuser{os.Getenv("TESTUSER_EMAIL"), os.Getenv("TESTUSER_PASSWORD")}
	input_json, _ := json.Marshal(input)
  body := strings.NewReader(string(input_json))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/app/login", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}