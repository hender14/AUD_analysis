package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/hender14/app/interfaces/controllers"
	"github.com/hender14/app/interfaces/gateway"
	"github.com/hender14/app/interfaces/presenter"
	"github.com/hender14/app/usecase/interactor"
)

type Routing struct {
	Fsc  *Fscontext
	Conf *Sgconfig
	Gin  *gin.Engine
	Port string
}

func NewRouting(f *Fscontext) *Routing {
	r := &Routing{
		Fsc:  f,
		Conf: rstNewRequest(),
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
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          r.Fsc,
		Config:        r.Conf,
	}
	// user registration
	r.Gin.POST("/register", func(c *gin.Context) { usersController.Sign(c) })
	// login
	r.Gin.POST("/app/login", func(c *gin.Context) { usersController.Login(c) })
	// logout
	r.Gin.GET("/app/logout", func(c *gin.Context) { usersController.Logout(c) })
	// get user info
	r.Gin.GET("/app/user", func(c *gin.Context) { usersController.User(c) })
	// reset user info
	r.Gin.POST("/app/forgot", func(c *gin.Context) { usersController.Forgot(c) })
	r.Gin.POST("/app/reset", func(c *gin.Context) { usersController.Reset(c) })
	// delete user info
	r.Gin.GET("/app/delete", func(c *gin.Context) { usersController.Delete(c) })
	// contact from user
	r.Gin.POST("/app/contact", func(c *gin.Context) { usersController.Contact(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(":" + r.Port)
}
