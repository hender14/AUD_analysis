package interactor

import (
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
)

func (interactor *UserInteractor) Forgot(c *gin.Context, email string) {
	// check hasn not registered
	// query the entity
	user, err := interactor.UserRepo.QueryEmail(email)
	if err != nil {
		interactor.OutputPort.RenderError(user, err)
		return
	}

	// create random token
	token := RandStringRunes(12)

	// save for DB
	forgotuser, err := interactor.UserRepo.ResetAccount(email, token)
	// var forgotuser *domain.ForgotUser
	// err = f.CreateAccoount(token)
	if err != nil {
		interactor.OutputPort.RenderError(forgotuser, err)
		return
	}

	// mail config
	mailconfig := Resetmail(forgotuser.Email)
	// param := map[string]string{"reseturl": os.Getenv("CORS_ADDRESS") + "/reset/" + token, "email": os.Getenv("SENDGRID_FROM_EMAIL")}
	param := &domain.RstmailPara{Reseturl: os.Getenv("CORS_ADDRESS") + "/reset/" + token, Email: os.Getenv("SENDGRID_FROM_EMAIL")}
	// param["Reseturl"] = os.Getenv("CORS_ADDRESS") + "/reset/" + token
	// fmt.Printf("%+v\n", param)

	// send Sendgrid mail by Web API
	_, err = interactor.UserRepo.Sendmail(mailconfig, param)
	if err != nil {
		interactor.OutputPort.RenderError(mailconfig, err)
		return
	}
	interactor.OutputPort.Render(gin.H{"token": token}, err)
}

func Resetmail(email string) *domain.Mails {
	mailconfig := &domain.Mails{
		From:    domain.Sendaddress{Express: "AUD Support Team", Address: os.Getenv("SENDGRID_FROM_EMAIL")},
		To:      domain.Sendaddress{Express: "User", Address: email},
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

// reset
func (interactor *UserInteractor) Reset(c *gin.Context, input *domain.ResetUser) {
	// validate password
	err := domain.CheckPassword(input.Password, input.Password_confirm)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		return
	}

	// get token data from JWT Token
	// query the entity
	user_rst, err := interactor.UserRepo.QueryToken(input.Token)
	if err != nil {
		interactor.OutputPort.RenderError(user_rst, err)
		return
	}

	// query the entity
	user, err := interactor.UserRepo.QueryEmail(user_rst.Email)
	if err != nil {
		interactor.OutputPort.RenderError(user, err)
		return
	}

	// Update the entity
	// timeも更新する?
	user, err = interactor.UserRepo.UpdateAccoount(user, input.Password)
	if err != nil {
		interactor.OutputPort.RenderError(user, err)
		return
	}

	err = interactor.UserRepo.DeleteAccoount(user_rst)
	if err != nil {
		interactor.OutputPort.RenderError(user, err)
		return
	}
	interactor.OutputPort.Render(gin.H{"message": "SUCCESS"}, err)
	return
}
