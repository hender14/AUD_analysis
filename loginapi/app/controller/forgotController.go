package controller

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"app/model"
	"app/service"
	"app/service/gcp"
)

func Forgot(ctx *gin.Context) {
	var data map[string]string
	// parse request data
	if err := ctx.Bind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}
	token := RandStringRunes(12)
	passwordReset := model.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	// check hasn not registered
	email := data["email"]
	// query the entity
	qrfield, err := gcp.Fsquery(&model.Fsqparam{Collection: "User", Key: "email", Condition: "==", Value: email})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "query has not completed"})
		return
	}
	if len(qrfield) == 0 || len(qrfield) == 2{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email had not registered"})
		return
	}

	// save for DB
	gcp.Fscreate_rst(&passwordReset)

	// send Sendgrid mail by Web API
	baseurl := os.Getenv("CORS_ADDRESS")
	url := baseurl + "/reset/" + token
	mailconfig := model.Mails{
		From: model.Sendaddress{ Express: "AUD Support Team", Address: os.Getenv("SENDGRID_FROM_EMAIL")},
		To: model.Sendaddress{ Express: "User", Address: data["email"]},
		Subject: "Password Reset",
		Text: model.Content{ Plantext: "Click here to reset password!", Htmltext: url },
  }
	_, err = service.Sendmail(mailconfig, url)
	if err != nil {
		// errmessage := "firebase auth has problem: " + err
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err })
		return
	}

	ctx.JSON(http.StatusOK, gin.H {"message": "SUCCESS"})
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

func Reset(ctx *gin.Context) {
	var data map[string]string
	// parse request data
	if err := ctx.Bind(&data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Tt has problem for input data"})
		return
	}
	// validate password
	if data["password"] != data["password_confirm"] {
    ctx.JSON(http.StatusBadRequest, gin.H {"message": "Passwords do not match!",})
	}

	// get token data from JWT Token
	token := data["token"]
	// query the entity
	qrfield, err := gcp.Fsquery(&model.Fsqparam{Collection: "PasswordReset", Key: "token", Condition: "==", Value: token})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "query has not completed"})
		return
	}
	if len(qrfield) == 0 || len(qrfield) == 2{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "reset had not registered"})
		return
	}

	email := qrfield[0].Email
	// query the entity
	qrfield, err = gcp.Fsquery(&model.Fsqparam{Collection: "User", Key: "email", Condition: "==", Value: email})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "query has not completed"})
		return
	}
	if len(qrfield) == 0 || len(qrfield) == 2{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "reset had not registered"})
		return
	}
	user := qrfield[0]
  // encode password
	password, _ :=service.PasswordEncrypt([]byte(data["password"]))
	user.Password = password
	// Update the entity
	// timeも更新する?
	_, err = gcp.Fsupdate(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "update has not completed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {"message": "SUCCESS"})
}