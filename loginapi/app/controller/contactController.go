package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/hender14/app/model"
	"github.com/hender14/app/service"
)

func Contact(ctx *gin.Context) {
	var input model.CntmailPara
	// parse request data
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	// send Sendgrid mail by Web API
	mailconfig := Contactmail()
	// fmt.Printf("mailconfig: %+v\n", mailconfig)
	// fmt.Printf("data: %+v\n", data)
	_, err := service.Cnt_Sendmail(mailconfig, input)
	if err != nil {
		// errmessage := "firebase auth has problem: " + err
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "contact has completed"})
}

func Contactmail() model.Mails {
	mailconfig := model.Mails{
		From:    model.Sendaddress{Express: "AUD Contact", Address: os.Getenv("SENDGRID_FROM_EMAIL")},
		To:      model.Sendaddress{Express: "Owner", Address: os.Getenv("SENDGRID_FROM_EMAIL")},
		Subject: "Contact Owner",
	}
	return mailconfig
}
