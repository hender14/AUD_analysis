package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/hender14/app/interfaces/controllers"
	"github.com/hender14/app/interfaces/gateway"
	"github.com/hender14/app/interfaces/presenter"
	"github.com/hender14/app/usecase/interactor"
)

type Routing struct {
	Awsc *Awscontext
	Conf *Sgconfig
	Gin  *gin.Engine
	Port string
}

func NewRouting(a *Awscontext) *Routing {
	r := &Routing{
		Awsc: a,
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
		Conn:          r.Awsc,
		Config:        r.Conf,
	}
	// user registration
	r.Gin.POST("/register", func(c *gin.Context) { usersController.Sign(c) })
	// login
	r.Gin.POST("/login", func(c *gin.Context) { usersController.Login(c) })
	// logout
	r.Gin.GET("/logout", func(c *gin.Context) { usersController.Logout(c) })
	// get user info
	r.Gin.GET("/user", func(c *gin.Context) { usersController.User(c) })
	// reset user info
	r.Gin.POST("/forgot", func(c *gin.Context) { usersController.Forgot(c) })
	r.Gin.POST("/reset", func(c *gin.Context) { usersController.Reset(c) })
	// delete user info
	r.Gin.GET("/delete", func(c *gin.Context) { usersController.Delete(c) })
	// contact from user
	r.Gin.POST("/contact", func(c *gin.Context) { usersController.Contact(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(":" + r.Port)
}
