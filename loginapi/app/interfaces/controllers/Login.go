package controllers

import (
	"github.com/hender14/app/domain"

	"fmt"

	"github.com/gin-gonic/gin"
)

// login
func (controller *UsersController) Login(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn, controller.Config)
	inputPort := controller.InputFactory(outputPort, repository)

	user := new(domain.LoginUser)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		outputPort.RenderError(user, err)
		return
	}
	inputPort.Login(c, user)
	return
}

// logout
func (controller *UsersController) Logout(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn, controller.Config)
	inputPort := controller.InputFactory(outputPort, repository)

	inputPort.Logout(c)
	return
}

// get user info
func (controller *UsersController) User(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn, controller.Config)
	inputPort := controller.InputFactory(outputPort, repository)

	// get JWT token from Cookie
	cookie, err := c.Cookie("jwt") // info when login
	if err != nil {
		fmt.Printf("cookie reading error: %s\n", err)
		outputPort.RenderError(cookie, err)
		return
	}
	inputPort.User(c, cookie)
	return
}
