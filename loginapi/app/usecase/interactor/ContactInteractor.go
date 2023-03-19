package interactor

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
)

func (interactor *UserInteractor) Contact(c *gin.Context, input *domain.CntmailPara) {
	// send Sendgrid mail by Web API
	mailconfig := Contactmail()
	// fmt.Printf("mailconfig: %+v\n", mailconfig)
	// fmt.Printf("data: %+v\n", data)
	_, err := interactor.UserRepo.Sendmail_Cnt(mailconfig, input)
	if err != nil {
		interactor.OutputPort.RenderError(mailconfig, err)
		return
	}
	interactor.OutputPort.Render(gin.H{"message": "contact has completed"}, err)
}

func Contactmail() *domain.Mails {
	mailconfig := &domain.Mails{
		From:    domain.Sendaddress{Express: "AUD Contact", Address: os.Getenv("SENDGRID_FROM_EMAIL")},
		To:      domain.Sendaddress{Express: "Owner", Address: os.Getenv("SENDGRID_FROM_EMAIL")},
		Subject: "Contact Owner",
	}
	return mailconfig
}
