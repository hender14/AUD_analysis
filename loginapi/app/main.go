package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"app/controller"
	"app/middleware"
	"app/service/gcp"
)

// main関数
func main() {
	router := setupRouter()

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  router.Run(":" + port)
}

func setupRouter() *gin.Engine {
  if err := gcp.NoSqlconnect(); err != nil {
    fmt.Printf("gcp connect error: %s\n", err)
  }
	router := gin.Default()
  //CORSはﾙｰﾃｨﾝｸﾞの前に宣言する
  router.Use(middleware.CORS())

  // user registration
  router.POST("/app/register", controller.Sign)
  // login
  router.POST("/app/login", controller.Login)
  // logout
  router.GET("/app/logout", controller.Logout)
  // get user info
  router.GET("/app/user", controller.User)
  // reset user info
  router.POST("/app/forgot", controller.Forgot)
  router.POST("/app/reset", controller.Reset)
  // delete user info
  router.GET("/app/delete", controller.Delete)
  router.POST("/app/contact", controller.Contact)
  
	return router
}