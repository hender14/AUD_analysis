package interactor

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
	"github.com/hender14/app/usecase/port"
)

type UserInteractor struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &UserInteractor{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

// user register
func (interactor *UserInteractor) Sign(input *domain.InUser) {
	s := new(domain.SignUser)

	// password check
	err := domain.CheckPassword(input)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		return
	}

	// query the entity
	_, err = interactor.UserRepo.QueryEmail_none(input.Email)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		return
	}

	// create account info
	err = s.CreateAccoount(input)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		return
	}

	// user register
	err = interactor.UserRepo.RegisterAccoount(s)
	if err != nil {
		interactor.OutputPort.RenderError(s, err)
		return
	}
	interactor.OutputPort.Render(s, err)
	return
}

func (interactor *UserInteractor) Delete(c *gin.Context, input string) {
	// query the entity
	account, err := interactor.UserRepo.ReadID(input)
	if err != nil {
		interactor.OutputPort.RenderError(account, err)
		return
	}

	err = interactor.UserRepo.DeleteAccount(account)
	if err != nil {
		interactor.OutputPort.RenderError(account, err)
		return
	}

	// samesiteをnonemodeにする
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", "", 3600 /* time.Now().Add(time.Hour * 24) */, "/app", os.Getenv("CORS_ADDRESS"), true, false)
	interactor.OutputPort.Render(account, err)
	return
}
