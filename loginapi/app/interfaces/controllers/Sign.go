package controllers

import (
	"github.com/hender14/app/domain"
	"github.com/hender14/app/interfaces/gateway"
	"github.com/hender14/app/usecase/port"

	"fmt"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	OutputFactory func(ctx *gin.Context) port.UserOutputPort
	// -> presenter.NewUserOutputPort
	InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	// -> interactor.NewUserInputPort
	RepoFactory func(c gateway.CRUD) port.UserRepository
	// -> gateway.NewUserRepository
	Conn gateway.CRUD
}

type deleteUser struct {
	ID string `json:"id"`
}

// user register
func (controller *UsersController) Sign(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn)
	inputPort := controller.InputFactory(outputPort, repository)

	user := new(domain.InUser)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		outputPort.RenderError(user, err)
		return
	}

	inputPort.Sign(user)
	return
}

func (controller *UsersController) Delete(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn)
	inputPort := controller.InputFactory(outputPort, repository)

	user := new(deleteUser)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		outputPort.RenderError(user.ID, err)
		return
	}
	inputPort.Delete(c, user.ID)
	return
}
