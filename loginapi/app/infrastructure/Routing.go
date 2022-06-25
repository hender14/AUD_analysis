package infrastructure

import (
	"github.com/gin-gonic/gin"

	"app/controller"
	"app/interfaces/controllers"
	"app/interfaces/gateway"
	"app/usecase/interactor"
)

type Routing struct {
	Fsc  *Fscontext
	Gin  *gin.Engine
	Port string
}

func NewRouting(f *Fscontext) *Routing {
	r := &Routing{
		Fsc:  f,
		Gin:  gin.Default(),
		Port: setPort(),
	}
	//CORSはﾙｰﾃｨﾝｸﾞの前に宣言する
	r.setMiddleware()
	r.setRouting()
	return r
}

func (r *Routing) setMiddleware() {
	r.Gin.Use(CORS())
}

func (r *Routing) setRouting() {
	// usersController := controllers.NewUsersController(r.Fsc)
	usersController := controllers.UsersController{
		InputFactory: interactor.NewUserInputPort,
		RepoFactory:  gateway.NewUserRepository,
		Conn:         r.Fsc,
	}
	// user registration
	// r.Gin.POST("/app/register", controller.Sign)
	r.Gin.POST("/register", func(c *gin.Context) { usersController.Sign(c) })
	// login
	r.Gin.POST("/app/login", controller.Login)
	// logout
	r.Gin.GET("/app/logout", controller.Logout)
	// get user info
	r.Gin.GET("/app/user", controller.User)
	// reset user info
	r.Gin.POST("/app/forgot", controller.Forgot)
	r.Gin.POST("/app/reset", controller.Reset)
	// delete user info
	r.Gin.GET("/app/delete", controller.Delete)
}

func (r *Routing) Run() {
	r.Gin.Run(":" + r.Port)
}
