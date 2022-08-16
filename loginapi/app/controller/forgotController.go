package controller

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hender14/app/model"
	"github.com/hender14/app/service"
	"github.com/hender14/app/service/gcp"
)

type inforgotUser struct {
	Email            string `json:"email"`
}

type inresetUser struct {
	Password         string `json:"password"`
	Password_confirm string `json:"password_confirm"`
	Token            string `json:"token"`
}

type forgotUser struct {
	input   inforgotUser
	account model.PasswordReset
}

type resetUser struct {
	input   inresetUser
	account_rst model.PasswordReset
	account model.SignUser
}

func (f *forgotUser) QueryEmail() (err error) {
	qparam := &model.Fsqparam{Collection: "PasswordReset", Key: "email", Condition: "==", Value: f.input.Email}
	qrfield, err := gcp.Fsquery(qparam)
	if err != nil {
		return err
	}
	if len(qrfield) != 0 {
		err = errors.New("email had already registered")
		fmt.Printf("email query error: %s query result: %+v/n", err, qrfield)
		return err
	}
	return err
}

func (r *resetUser) QueryEmail() (err error) {
	qparam := &model.Fsqparam{Collection: "SignUser", Key: "email", Condition: "==", Value: r.account_rst.Email}
	qrfield, err := gcp.Fsquery(qparam)
	if err != nil {
		return err
	}
	if len(qrfield) == 0 || len(qrfield) == 2{
		err = errors.New("reset had not registered")
		fmt.Printf("email query error: %s/n", err)
		return err
	}
	r.account = qrfield[0]
	return err
}

func (r *resetUser) QueryToken() (err error) {
	qparam := &model.Fsqparam{Collection: "PasswordReset", Key: "token", Condition: "==", Value: r.input.Token}
	qrfield, err := gcp.Fsquery_rst(qparam)
	if err != nil {
		return err
	}
	if len(qrfield) == 0 || len(qrfield) == 2{
		err = errors.New("reset had not registered")
		fmt.Printf("token query error: %s/n", err)
		return err
	}
	r.account_rst = qrfield[0]
	return err
}

// password check
func (r *resetUser) CheckPassword() (err error) {
	if r.input.Password != r.input.Password_confirm {
		err = errors.New("password is different")
		fmt.Printf("password confirm error: %s/n", err)
		return err
	}
	return err
}

func (f *forgotUser) CreateAccoount(token string) (err error) {
	f.account.Email = f.input.Email
	f.account.Token = token
	f.account.Year = time.Now()

	err = gcp.Fscreate_rst(&f.account)
	if err != nil {
		fmt.Printf("create result: %s/n", f.account)
		fmt.Printf("err result: %s/n", err)
		return err
	}

	return err
}

func (r *resetUser) UpdateAccoount() (err error) {
	// password encode
	password, err :=service.PasswordEncrypt([]byte(r.input.Password))
	if err != nil {
		fmt.Printf("password encode err: %s\n", err)
		return err
	}

	r.account.Password = password
	r.account.Year = time.Now()

	err = gcp.Fsupdate(&r.account)
	if err != nil {
		fmt.Printf("update result: %s/n", r.account)
		fmt.Printf("err result: %s/n", err)
		return err
	}
	return err
}

func (r *resetUser) DeleteAccoount() (err error) {
	err = gcp.Fsdelete_rst(&r.account_rst)
	if err != nil {
		return err
	}
	return err
}

// forgot user info
func Forgot(ctx *gin.Context) {
	f := new(forgotUser)
	// parse request data
	if err := ctx.BindJSON(&f.input); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	// check hasn not registered
	// query the entity
	err := f.QueryEmail()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email address has not registered"})
		return
	}

	// create random token
	token := RandStringRunes(12)

	// save for DB
	err = f.CreateAccoount(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "account create error"})
		return
	}

	// mail config
	mailconfig := Resetmail(f.account.Email)
	// param := map[string]string{"reseturl": os.Getenv("CORS_ADDRESS") + "/reset/" + token, "email": os.Getenv("SENDGRID_FROM_EMAIL")}
	param := model.RstmailPara{Reseturl: os.Getenv("CORS_ADDRESS") + "/reset/" + token, Email: os.Getenv("SENDGRID_FROM_EMAIL")}
	// param["Reseturl"] = os.Getenv("CORS_ADDRESS") + "/reset/" + token
	// fmt.Printf("%+v\n", param)

	// send Sendgrid mail by Web API
	_, err = service.Sendmail(mailconfig, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err })
		return
	}

	ctx.JSON(http.StatusOK, gin.H {"token": token,})
}

func Resetmail(email string) ( model.Mails ) {

	mailconfig := model.Mails{
		From: model.Sendaddress{ Express: "AUD Support Team", Address: os.Getenv("SENDGRID_FROM_EMAIL")},
		To: model.Sendaddress{ Express: "User", Address: email},
		Subject: "Password Reset",
		// Text: model.Content{ Plantext: "Click here to reset password!", Htmltext: url },
  }
	return mailconfig
}

// ﾗﾝﾀﾞﾑ文字列を返す関数
func RandStringRunes(n int) string {
	var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return string(b)
}

// reset user info
func Reset(ctx *gin.Context) {
	r := new(resetUser)
	// parse request data
	if err := ctx.BindJSON(&r.input); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	// validate password
	err := r.CheckPassword()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Password is different"})
		return
	}

	// get token data from JWT Token
	// query the entity
	err = r.QueryToken()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email address has already registered"})
		return
	}

	// query the entity
	err = r.QueryEmail()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "reset has not completed"})
		return
	}

	// Update the entity
	// timeも更新する?
	err = r.UpdateAccoount()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "update has not completed"})
		return
	}

	// user_rst := qrfield_rst[0]
	// err = gcp.Fsdelete_rst( &user_rst )

	err = r.DeleteAccoount()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "delete has not completed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {"message": "SUCCESS"})
}