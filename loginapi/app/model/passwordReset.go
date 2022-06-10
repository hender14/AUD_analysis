package model

import "time"

type PasswordReset struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Year  time.Time `json:"year"`
}
