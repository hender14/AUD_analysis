package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/hender14/app/model"
	"github.com/hender14/app/service/gcp"
)

type inUser struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Password_confirm string `json:"password_confirm"`
}

type signUser struct {
	input   inUser
	account model.SignUser
}

type indeleteUser struct {
	ID string `json:"id"`
}

type deleteUser struct {
	input   indeleteUser
	account *model.SignUser
}

func (s *signUser) QueryEmail() (err error) {
	qparam := &model.Fsqparam{Collection: "SignUser", Key: "email", Condition: "==", Value: s.input.Email}
	qrfield, err := gcp.Fsquery(qparam)
	if err != nil {
		return err
	}
	if len(qrfield) != 0 {
		err = errors.New("email had already registered")
		fmt.Printf("email query err: %s/n", err)
		return err
	}
	return err
}

func (d *deleteUser) ReadID() (err error) {
	d.account, err = gcp.Fsread(d.input.ID)
	if err != nil {
		fmt.Printf("query result: %s/n", d.account)
		fmt.Printf("query err result: %s/n", err)
		return err
	}
	if d.account.ID == "" {
		err = errors.New("user should have an auto generated ID")
		fmt.Printf("no ID err: %s/n", err)
		return err
	}
	return err
}

// password check
func (s *signUser) CheckPassword() (err error) {
	if s.input.Password != s.input.Password_confirm {
		err = errors.New("password is different")
		fmt.Printf("password confirm err: %s/n", err)
		return err
	}
	return err
}

func (s *signUser) CreateAccoount() (err error) {
	// timestmp := time.Now().Format(time.RFC850)

	// password encode
	password, err := bcrypt.GenerateFromPassword([]byte(s.input.Password), 14)
	if err != nil {
		fmt.Printf("password encode err: %s\n", err)
		return err
	}

	s.account.FirstName = s.input.FirstName
	s.account.LastName = s.input.LastName
	s.account.Email = s.input.Email
	s.account.Password = password
	s.account.Year = time.Now()

	err = gcp.Fscreate(&s.account)
	if err != nil {
		fmt.Printf("create result: %s/n", s.account)
		fmt.Printf("err result: %s/n", err)
		return err
	}

	return err
}

func (d *deleteUser) DeleteAccount() (err error) {
	err = gcp.Fsdelete(d.account)
	if err != nil {
		return err
	}
	return err
}

// user register
func Sign(ctx *gin.Context) {
	s := new(signUser)
	// parse request data
	if err := ctx.BindJSON(&s.input); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	// password check
	err := s.CheckPassword()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Password is different"})
		return
	}

	// query the entity
	err = s.QueryEmail()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email address has already registered"})
		return
	}

	// user register
	err = s.CreateAccoount()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "account create error"})
		return
	}

	ctx.JSON(http.StatusOK, s.account)
}

// user delete
func Delete(ctx *gin.Context) {
	d := new(deleteUser)

	if err := ctx.BindJSON(&d.input); err != nil {
		log.Print(err)
		fmt.Printf("parse request err: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	// query the entity
	err := d.ReadID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "query has not completed"})
		return
	}
	err = gcp.Fsdelete(d.account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Delete has not completed"})
		return
	}

	// samesiteをnonemodeにする
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("jwt", "", 3600 /* time.Now().Add(time.Hour * 24) */, "/app", os.Getenv("CORS_ADDRESS"), true, false)
	ctx.JSON(http.StatusOK, d.account)
}
