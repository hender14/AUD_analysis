package domain

import (
	"time"
)

type ForgotUser struct {
	ID    string    `dynamo:"id"`
	Email string    `dynamo:"email"`
	Token string    `dynamo:"token"`
	Year  time.Time `json:"year"`
}

type ResetUser struct {
	Password         string `json:"password"`
	Password_confirm string `json:"password_confirm"`
	Token            string `json:"token"`
}

type RstmailPara struct {
	Reseturl string `json:"reseturl"`
	Email    string `json:"email"`
}

type Mails struct {
	From    Sendaddress
	To      Sendaddress
	Subject string
	// Text    Content
}

type Sendaddress struct {
	Express string
	Address string
}
