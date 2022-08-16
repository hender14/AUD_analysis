package controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/hender14/app/model"
	"github.com/hender14/app/service/gcp"
)

type inLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginUser struct {
	input   inLoginUser
	account model.SignUser
}
type JWTToken struct {
	ID       string
	claims   *Claims
	token    string
	jwttoken *jwt.Token
	cookie   string
}

type Claims struct {
	jwt.StandardClaims
}

func (l *loginUser) QueryEmail() (err error) {
	qparam := &model.Fsqparam{Collection: "SignUser", Key: "email", Condition: "==", Value: l.input.Email}
	qrfield, err := gcp.Fsquery(qparam)
	if err != nil {
		return err
	}
	if len(qrfield) == 0 || len(qrfield) == 2 {
		err = errors.New("email had not registered")
		fmt.Printf("email query error: %s/n", err)
		return err
	}
	l.account = qrfield[0]

	if l.account.ID == "" {
		err = errors.New("id had not registered")
		fmt.Printf("id has not generated: %s/n", err)
		return err
	}
	return err
}

func (j *JWTToken) CreateToken() (err error) {
	Claims := jwt.StandardClaims{Issuer: j.ID /*stringに型変換*/, ExpiresAt: time.Now().Add(time.Hour * 24).Unix() /*有効期限*/}
	j.jwttoken = jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	j.token, err = j.jwttoken.SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("Token has not generated: %s\n", err)
		return err
	}
	return err
}

// user login
func Login(ctx *gin.Context) {
	l := new(loginUser)
	if err := ctx.BindJSON(&l.input); err != nil {
		fmt.Printf("Tt has problem for input data: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	// query the entity
	err := l.QueryEmail()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email address has not registered"})
		return
	}

	// compare user password
	err = bcrypt.CompareHashAndPassword(l.account.Password, []byte(l.input.Password))
	if err != nil {
		fmt.Printf("Tt has problem for comparing password: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect password"})
		return
	}

	// create JWT token
	j := JWTToken{ID: l.account.ID}
	err = j.CreateToken()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Token has not generated"})
		return
	}

	// config Cookie
	ctx.SetSameSite(http.SameSiteNoneMode) // samesiteをnonemodeにする
	// 第7引数:http requestのみcookieを利用できるようにするか指定する。trueの場合JavaScriptからcookieを利用できない。
	// trueにすることで、XSS攻撃を緩和することが可能。
	ctx.SetCookie("jwt", j.token, 3600 /*[s]*/, "/app", os.Getenv("CORS_ADDRESS"), true, false)

	ctx.JSON(http.StatusOK, gin.H{"jwt": j.token})
}

// user logout
func Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteNoneMode) // samesiteをnonemodeにする
	ctx.SetCookie("jwt", "", 3600 /*[s]*/, "/app", os.Getenv("CORS_ADDRESS"), true, false)
	cookie, err := ctx.Cookie("jwt") // info when login
	// if cookie != "" {
	// 	fmt.Printf("cookie data has not deleted: %s\n", cookie)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cookie has not deleted"})
	// 	return
	// }
	if err != nil {
		fmt.Printf("cookie reading error: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cookie operation has not done"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"jwt": cookie})
}

// get user info /* add coockie info */
func User(ctx *gin.Context) {
	var err error
	j := new(JWTToken)

	// get JWT token from Cookie
	j.cookie, err = ctx.Cookie("jwt") // info when login
	if err != nil {
		fmt.Printf("cookie reading error: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "token info has not"})
		return
	}

	// get token
	j.jwttoken, err = jwt.ParseWithClaims(j.cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !j.jwttoken.Valid {
		fmt.Printf("jwt token parse error: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unauthenticated"})
	}

	// get User ID
	// claims := j.jwttoken.Claims.(*Claims)
	j.claims = j.jwttoken.Claims.(*Claims)
	j.ID = j.claims.Issuer

	// read the entity
	user, err := gcp.Fsread(j.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "read has not completed"})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H {"message": "success",})
	ctx.JSON(http.StatusOK, user)
}
