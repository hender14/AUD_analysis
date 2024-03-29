package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/hender14/app/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type Tsignuser struct {
	SFirstname        string `json:"first_name"`
	SLastname         string `json:"last_name"`
	SEmail            string `json:"email"`
	SPassword         string `json:"password"`
	SPassword_confirm string `json:"password_confirm"`
}

type Tloginuser struct {
	LEmail    string `json:"email"`
	LPassword string `json:"password"`
}

type TForgotuser struct {
	TEmail string `json:"email"`
}

type LJwttoken struct {
	Jwttoken string `json:"jwt"`
}

type RToken struct {
	Token string `json:"token"`
}

type Tdeleteuser struct {
	ID string `json:"id"`
}

type TResetuser struct {
	SToken            string `json:"token"`
	SPassword         string `json:"password"`
	SPassword_confirm string `json:"password_confirm"`
}

type Cntmail struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Email   string `json:"email"`
}

// global変数
var jwttoken LJwttoken
var token RToken
var user Tdeleteuser

func testSetting() *gin.Engine {
	fs, err := infrastructure.NewDB()
	if err != nil {
		log.Fatalf("Listen and serve failed. %s\n", err)
	}
	r := infrastructure.NewRouting(fs)
	return r.Gin
}

func TestSign(t *testing.T) {
	router := testSetting()
	input := Tsignuser{os.Getenv("TESTUSER_FIRSTNAME"), os.Getenv("TESTUSER_LASTNAME"), os.Getenv("TESTUSER_EMAIL"), os.Getenv("TESTUSER_PASSWORD"), os.Getenv("TESTUSER_PASSWORD_CONFIRM")}
	input_json, _ := json.Marshal(input)
	body := strings.NewReader(string(input_json))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	// println(w.Body.String())
}

func TestLogin(t *testing.T) {
	router := testSetting()
	input := Tloginuser{os.Getenv("TESTUSER_EMAIL"), os.Getenv("TESTUSER_PASSWORD")}
	input_json, _ := json.Marshal(input)
	body := strings.NewReader(string(input_json))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &jwttoken)
	assert.Equal(t, nil, err)
	// println(w.Body.String())
}

func TestUser(t *testing.T) {
	router := testSetting()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", nil)
	// Cookie
	req.AddCookie(&http.Cookie{
		Name: "jwt", Value: jwttoken.Jwttoken, Path: "/app", Domain: os.Getenv("CORS_ADDRESS"),
		MaxAge: 3600 /* seconds */, Secure: true, HttpOnly: false,
	})
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, nil, err)
	// println(w.Body.String())
}

func TestLogout(t *testing.T) {
	router := testSetting()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/logout", nil)
	// Cookie
	req.AddCookie(&http.Cookie{
		Name: "jwt", Value: jwttoken.Jwttoken, Path: "/app", Domain: os.Getenv("CORS_ADDRESS"),
		MaxAge: 3600 /* seconds */, Secure: true, HttpOnly: false,
	})
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &jwttoken)
	assert.Equal(t, nil, err)
	// println(w.Body.String())
}

func TestForgot(t *testing.T) {
	router := testSetting()
	input := TForgotuser{os.Getenv("TESTUSER_EMAIL")}
	input_json, _ := json.Marshal(input)
	body := strings.NewReader(string(input_json))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/forgot", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &token)
	assert.Equal(t, nil, err)
	// println(token.Token)
}

func TestReset(t *testing.T) {
	router := testSetting()
	input := TResetuser{token.Token, os.Getenv("TESTUSER_PASSWORD"), os.Getenv("TESTUSER_PASSWORD_CONFIRM")}
	input_json, _ := json.Marshal(input)
	body := strings.NewReader(string(input_json))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/reset", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	// println(w.Body.String())
}

func TestDelete(t *testing.T) {
	router := testSetting()
	input := Tdeleteuser{user.ID}
	input_json, _ := json.Marshal(input)
	body := strings.NewReader(string(input_json))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/delete", body)
	// Cookie
	req.AddCookie(&http.Cookie{
		Name: "jwt", Value: jwttoken.Jwttoken, Path: "/app", Domain: os.Getenv("CORS_ADDRESS"),
		MaxAge: 3600 /* seconds */, Secure: true, HttpOnly: false,
	})
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	// err := json.Unmarshal(w.Body.Bytes(), &jwttoken)
	// assert.Equal(t, nil, err)
	// println(w.Body.String())
}

func TestContact(t *testing.T) {
	router := testSetting()
	input := Cntmail{os.Getenv("TESTMAIL_ID"), os.Getenv("TESTMAIL_TITLE"), os.Getenv("TESTMAIL_CONTENT"), os.Getenv("TESTMAIL_EMAIL")}
	input_json, _ := json.Marshal(input)
	body := strings.NewReader(string(input_json))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/contact", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, nil, w.Body.String())
	// println(w.Body.String())
}
