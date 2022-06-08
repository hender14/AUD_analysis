package model

import (
	"time"
)

// 構造体宣言
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"` // -を指定すると非表示にできる
}

type SignUser struct {
	ID string `json:"id"`
	User
	Year time.Time
}
