package controllers

import (
	"app/domain"
	"app/interfaces/gateway"
	"app/usecase/interactor"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	Interactor interactor.UserInteractor
}

func NewUsersController(f gateway.Fsc) *UsersController {
	return &UsersController{
		Interactor: interactor.UserInteractor{
			User: gateway.UserRepository{Fsc: f},
		},
	}
}

// user register
func (controller *UsersController) Sign(c *gin.Context) {
	user := new(domain.InUser)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	account, err := controller.Interactor.Sign(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for account register"})
		return
	}

	c.JSON(http.StatusOK, account)
}
