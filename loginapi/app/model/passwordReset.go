package model

import "time"

type PasswordReset struct {
	ID    string
	Email string
	Token string
	Year  time.Time
}
