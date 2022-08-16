package controllers

import (
	"os"

	"github.com/hender14/app/domain"
	"github.com/hender14/app/interfaces/gateway"
	"github.com/hender14/app/usecase/port"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	// OutputFactory interactor.UserInteractor
	// -> presenter.NewUserOutputPort
	InputFactory func(u port.UserRepository) port.UserInputPort
	// -> interactor.NewUserInputPort
	RepoFactory func(c gateway.CRUD) port.UserRepository
	// -> gateway.NewUserRepository
	Conn gateway.CRUD
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

	repository := controller.RepoFactory(controller.Conn)
	inputPort := controller.InputFactory(repository)
	account, err := inputPort.Sign(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for account register"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (controller *UsersController) Delete(c *gin.Context) {
	user := new(domain.SignUser)
	// parse request data
	if err := c.BindJSON(&user); err != nil {
		fmt.Printf("parse request err: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	repository := controller.RepoFactory(controller.Conn)
	inputPort := controller.InputFactory(repository)
	account, err := inputPort.Delete(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for account register"})
		return
	}

	// samesiteをnonemodeにする
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", "", 3600 /* time.Now().Add(time.Hour * 24) */, "/app", os.Getenv("CORS_ADDRESS"), true, false)
	c.JSON(http.StatusOK, account)
}
