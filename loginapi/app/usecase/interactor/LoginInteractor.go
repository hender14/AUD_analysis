package interactor

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hender14/app/domain"
	"golang.org/x/crypto/bcrypt"
)

// login
func (interactor *UserInteractor) Login(c *gin.Context, input *domain.LoginUser) {
	// query the entity
	account, err := interactor.UserRepo.QueryEmail(input.Email)
	if err != nil {
		interactor.OutputPort.RenderError(input, err)
		return
	}

	// compare user password
	err = bcrypt.CompareHashAndPassword(account.Password, []byte(input.Password))
	if err != nil {
		fmt.Printf("Tt has problem for comparing password: %s\n", err)
		interactor.OutputPort.RenderError(input, err)
		return
	}

	// create JWT token
	j := domain.JWTToken{ID: account.ID}
	token, err := j.CreateToken()
	if err != nil {
		interactor.OutputPort.RenderError(j, err)
		return
	}

	// config Cookie
	c.SetSameSite(http.SameSiteNoneMode) // samesiteをnonemodeにする
	// 第7引数:http requestのみcookieを利用できるようにするか指定する。trueの場合JavaScriptからcookieを利用できない。
	// trueにすることで、XSS攻撃を緩和することが可能。
	c.SetCookie("jwt", token, 3600 /*[s]*/, "/app", os.Getenv("CORS_ADDRESS"), true, false)

	interactor.OutputPort.Render(gin.H{"jwt": token}, err)
	return
}

// logout
func (interactor *UserInteractor) Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode) // samesiteをnonemodeにする
	c.SetCookie("jwt", "", 3600 /*[s]*/, "/app", os.Getenv("CORS_ADDRESS"), true, false)
	cookie, err := c.Cookie("jwt") // info when login
	// if cookie != "" {
	// 	fmt.Printf("cookie data has not deleted: %s\n", cookie)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cookie has not deleted"})
	// 	return
	// }
	if err != nil {
		fmt.Printf("cookie reading error: %s\n", err)
		interactor.OutputPort.RenderError(cookie, err)
		return
	}
	interactor.OutputPort.Render(gin.H{"jwt": cookie}, err)
	return
}

func (interactor *UserInteractor) User(c *gin.Context, cookie string) {
	// get token
	jwttoken, err := jwt.ParseWithClaims(cookie, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !jwttoken.Valid {
		fmt.Printf("jwt token parse error: %s\n", err)
		interactor.OutputPort.RenderError(jwttoken, err)
		return
	}

	// get User ID
	// claims := j.jwttoken.Claims.(*Claims)
	claims := jwttoken.Claims.(*domain.Claims)
	ID := claims.Issuer

	// read the entity
	user, err := interactor.UserRepo.ReadID(ID)
	if err != nil {
		interactor.OutputPort.RenderError(user, err)
		return
	}

	interactor.OutputPort.Render(user, err)
	return
}
