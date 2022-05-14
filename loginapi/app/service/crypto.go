package service

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordEncrypt ﾊﾟｽﾜｰﾄﾞをhash化
func PasswordEncrypt(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

// const (
// 	MinCost     int = 4
// 	MaxCost     int = 31
// 	DefaultCost int = 10
// )

// CompareHashAndPassword hashと非hashﾊﾟｽﾜｰﾄﾞド比較
func CompareHashAndPassword(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}