package model

import (
  // "github.com/jinzhu/gorm"
	"time"
  // _ "github.com/go-sql-driver/mysql"
)

// 構造体宣言
// type User struct {
//   gorm.Model
//   Username string //`form:"username" binding:"required" gorm:"unique;not null"`
//   // Email string
//   Password string //`form:"password" binding:"required"`
// }

type User struct {
	ID        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"` // -を指定すると非表示にできる
	Year      time.Time
}

// type User struct {
// 	gorm.Model
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Email     string `json:"email" gorm:"unique"`
// 	Password  []byte `json:"-"` // -を指定すると非表示にできる
// }

type Info struct {
  Name   string
  Id   string
  Test   string
  Status int
	// Status string `json:"statusCode"`
	// Name   string `json:"name"`
}
