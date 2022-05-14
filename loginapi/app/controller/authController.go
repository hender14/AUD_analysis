package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"app/model"
	"app/service/gcp"
)

type Claims struct {
	jwt.StandardClaims
}

func Test(ctx *gin.Context) {
  // ctx.AbortWithStatus(http.StatusNoContent)
		ctx.JSON(200, gin.H{
			"msg": "hello world",
		})
}

func Test2(ctx *gin.Context) {
	fmt.Println("test")
	// var data map[string]string
  // ctx.AbortWithStatus(http.StatusNoContent)
		ctx.JSON(200, gin.H{
			"msg": "hello world",
		})
}

// user register
func Sign(ctx *gin.Context) {
	var data map[string]string
	// parse request data
		if err := ctx.Bind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	email := data["email"]
	  // query the entity
	qrfield, err := gcp.Fsquery(&model.Fsqparam{Collection: "User", Key: "email", Condition: "==", Value: email})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "query has not completed"})
		return
	}
	if len(qrfield) != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email had already registered"})
		return
  }

	// password check
	if data["password"] != data["password_confirm"] {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Password is different"})
		return
	}

	// password encode
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	timestp, _ := time.Parse(time.RFC3339, "2001-01-01T00:00:00.000Z")
	user := model.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
		Year:      timestp,
	}

	// user register
	gcp.Fscreate(&user)

	ctx.JSON(http.StatusOK, user)
}

// user login
func Login(ctx *gin.Context) {
	var data map[string]string
	// if err := ctx.Bind(&data); err != nil {
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Tt has problem for input data"})
		return
	}

	email := data["email"]
	// query the entity
	qrfield, err := gcp.Fsquery(&model.Fsqparam{Collection: "User", Key: "email", Condition: "==", Value: email})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "query has not completed"})
		return
	}
	if len(qrfield) == 0 || len(qrfield) == 2{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email had not registered"})
		return
	}
	user := qrfield[0]

	if user.ID == "" {
		ctx.JSON(404, gin.H {"message": "User not found"})
		return
	}
	// compare user password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {"message": "Incorrect password" })
		return
		// ctx.Abort()
	}
	// JWT
	claims := jwt.StandardClaims {
		Issuer: user.ID,            // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Token has not generated"})
		return
	}
  // samesiteをnonemodeにする
	ctx.SetSameSite(http.SameSiteNoneMode)
	// Cookie
	ctx.SetCookie ("jwt",	token, 3600 /* seconds */,   "/app", os.Getenv("CORS_ADDRESS"), true ,false,
	)
// 第7引数・・・httpリクエストのみcookieを利用できるようにするか指定します。trueの場合JavaScriptからcookieを利用することができません。trueにすることで、クロスサイトスクリプティング（XSS）攻撃を緩和することができます。

ctx.JSON(200, gin.H {"jwt": token,})
}

// user logout
func Logout(ctx *gin.Context) {
	// samesiteをnonemodeにする
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie ("jwt",	"", 3600/* time.Now().Add(time.Hour * 24) */, "/app", os.Getenv("CORS_ADDRESS"), true ,false,
)

	ctx.JSON(http.StatusOK, gin.H {"message": "success",})
}

	// get user info *add coockie info
func User(ctx *gin.Context) {
	// get JWT from Cookie
	cookie, err := ctx.Cookie("jwt") // info when login
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "token info has not"})
		return
	}
	// get token
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unauthenticated",})
	}

	claims := token.Claims.(*Claims)
	// get User ID
	id := claims.Issuer

	// read the entity
	rdfield, err := gcp.Fsread(&model.User{ID:id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "read has not completed"})
		return
	}
	user := rdfield

	// ctx.JSON(http.StatusOK, gin.H {"message": "success",})
	ctx.JSON(200, user)
}