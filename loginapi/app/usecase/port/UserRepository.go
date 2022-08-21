package port

import (
	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
	"github.com/sendgrid/rest"
)

type UserRepository interface {
	QueryEmail_none(string) (*domain.SignUser, error)
	QueryEmail(string) (*domain.SignUser, error)
	RegisterAccoount(*domain.SignUser) error
	ReadID(string) (*domain.SignUser, error)
	DeleteAccount(*domain.SignUser) error
	ResetAccount(string, string) (*domain.ForgotUser, error)
	Sendmail(config *domain.Mails, param *domain.RstmailPara) (*rest.Response, error)
	Sendmail_Cnt(config *domain.Mails, param *domain.CntmailPara) (*rest.Response, error)
	QueryToken(string) (*domain.ForgotUser, error)
	UpdateAccoount(*domain.SignUser, string) (*domain.SignUser, error)
	DeleteAccoount(*domain.ForgotUser) error
}

type UserInputPort interface {
	Sign(*domain.InUser)
	Delete(*gin.Context, string)
	Login(*gin.Context, *domain.LoginUser)
	Logout(*gin.Context)
	User(*gin.Context, string)
	Forgot(*gin.Context, string)
	Reset(*gin.Context, *domain.ResetUser)
	Contact(*gin.Context, *domain.CntmailPara)
}

type UserOutputPort interface {
	Render(interface{}, error)
	RenderError(interface{}, error)
}
