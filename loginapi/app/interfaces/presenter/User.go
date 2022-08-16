package presenter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hender14/app/usecase/port"
)

type UserRepository struct {
	context *gin.Context
}

// NewUserRepository はUserRepositoryを返します．
func NewUserOutputPort(c *gin.Context) port.UserOutputPort {
	return &UserRepository{
		context: c,
	}
}

// usecase.UserOutputPortを実装している
// Render はNameを出力します．
func (u *UserRepository) Render(user interface{}, err error) {
	u.context.JSON(http.StatusOK, user)
}

// RenderError はErrorを出力します．
func (u *UserRepository) RenderError(user interface{}, err error) {
	u.context.JSON(http.StatusInternalServerError, gin.H{"message": "Tt has problem for input data"})
	fmt.Printf("Tt has problem: %s\n", err)
	// fmt.Printf(user, err)
}
