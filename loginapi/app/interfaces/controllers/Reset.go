package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
)

type forgotUser struct {
	Email string `json:"email"`
}

func (controller *UsersController) Forgot(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn, controller.Config)
	inputPort := controller.InputFactory(outputPort, repository)

	user := new(forgotUser)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		outputPort.RenderError(user.Email, err)
		return
	}
	inputPort.Forgot(c, user.Email)
	return
}

// reset
func (controller *UsersController) Reset(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn, controller.Config)
	inputPort := controller.InputFactory(outputPort, repository)

	user := new(domain.ResetUser)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		outputPort.RenderError(user, err)
		return
	}
	inputPort.Reset(c, user)
	return
}
