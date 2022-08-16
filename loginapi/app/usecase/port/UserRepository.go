package port

import (
	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
)

type UserRepository interface {
	QueryEmail(*domain.InUser) (err error)
	RegisterAccoount(*domain.SignUser) (err error)
	ReadID(*domain.SignUser) error
	DeleteAccount(*domain.SignUser) error
}

type UserInputPort interface {
	Sign(*domain.InUser)
	Delete(*gin.Context, *domain.SignUser)
}

type UserOutputPort interface {
	Render(interface{}, error)
	RenderError(interface{}, error)
}
