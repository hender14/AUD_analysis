package domain

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type InUser struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Password_confirm string `json:"password_confirm"`
}

type SignUser struct {
	ID        string    `dynamo:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `dynamo:"email"`
	Password  []byte    `json:"-"` // -を指定すると非表示にできる
	Year      time.Time `json:"year"`
}

type deleteUser struct {
	ID      string `json:"id"`
	account *SignUser
}

// password check
func CheckPassword(password string, password_confirm string) (err error) {
	if password != password_confirm {
		err = errors.New("password is different")
		fmt.Printf("password confirm err: %s/n", err)
		return err
	}
	return err
}

func (s *SignUser) CreateAccoount(i *InUser) error {
	password, err := EncodePassword(i.Password)
	if err != nil {
		return err
	}

	s.FirstName = i.FirstName
	s.LastName = i.LastName
	s.Email = i.Email
	s.Password = password
	s.Year = time.Now()

	return err
}

// password encode
func EncodePassword(password string) ([]byte, error) {
	epassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Printf("password encode err: %s\n", err)
		return nil, err
	}
	return epassword, err
}
