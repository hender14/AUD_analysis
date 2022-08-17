package port

import (
	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
)

type UserRepository interface {
	QueryEmail_none(string) (domain.SignUser, error)
	QueryEmail(string) (domain.SignUser, error)
	RegisterAccoount(*domain.SignUser) error
	ReadID(string) (*domain.SignUser, error)
	DeleteAccount(*domain.SignUser) error
}

type UserInputPort interface {
	Sign(*domain.InUser)
	Delete(*gin.Context, string)
	Login(*gin.Context, *domain.LoginUser)
	Logout(*gin.Context)
	User(*gin.Context, string)
}

type UserOutputPort interface {
	Render(interface{}, error)
	RenderError(interface{}, error)
}
