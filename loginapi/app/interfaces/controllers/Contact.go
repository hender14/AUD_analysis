package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
)

func (controller *UsersController) Contact(c *gin.Context) {
	outputPort := controller.OutputFactory(c)
	repository := controller.RepoFactory(controller.Conn, controller.Config)
	inputPort := controller.InputFactory(outputPort, repository)

	user := new(domain.CntmailPara)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		outputPort.RenderError(user, err)
		return
	}
	inputPort.Contact(c, user)
	return
}
