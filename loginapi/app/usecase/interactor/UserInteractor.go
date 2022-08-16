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
	User       port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &UserInteractor{
		OutputPort: outputPort,
		User:       userRepository,
	}
}

// user register
func (interactor *UserInteractor) Sign(input *domain.InUser) {
	s := new(domain.SignUser)

	// password check
	err := domain.CheckPassword(input)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		// return nil, err
	}

	// query the entity
	err = interactor.User.QueryEmail(input)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		// return nil, err
	}

	// create account info
	err = s.CreateAccoount(input)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		// return s, err
	}

	// user register
	err = interactor.User.RegisterAccoount(s)
	if err != nil {
		interactor.OutputPort.RenderError(s, err)
		// return s, err
	}
	interactor.OutputPort.Render(s, err)

	return
}

func (interactor *UserInteractor) Delete(c *gin.Context, input *domain.SignUser) {
	d := new(domain.SignUser)
	// query the entity
	err := interactor.User.ReadID(d)
	if err != nil {
		interactor.OutputPort.RenderError(d, err)
		// return d, err
	}
	err = interactor.User.DeleteAccount(input)
	if err != nil {
		interactor.OutputPort.RenderError(d, err)
		// return d, err
	}

	// samesiteをnonemodeにする
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", "", 3600 /* time.Now().Add(time.Hour * 24) */, "/app", os.Getenv("CORS_ADDRESS"), true, false)
	// c.JSON(http.StatusOK, d)
	interactor.OutputPort.Render(d, err)
	// return d, err
	return
}
