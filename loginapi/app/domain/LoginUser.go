package domain

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTToken struct {
	ID       string
	claims   *Claims
	jwttoken *jwt.Token
	cookie   string
}

type Claims struct {
	jwt.StandardClaims
}

// create JWT token
func (j *JWTToken) CreateToken() (token string, err error) {
	Claims := jwt.StandardClaims{Issuer: j.ID /*stringに型変換*/, ExpiresAt: time.Now().Add(time.Hour * 24).Unix() /*有効期限*/}
	j.jwttoken = jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	token, err = j.jwttoken.SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("Token has not generated: %s\n", err)
		return token, err
	}
	return token, err
}
