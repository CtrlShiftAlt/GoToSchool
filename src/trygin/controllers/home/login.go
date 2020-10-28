package home

import (
	"fmt"
	"net/http"
	"time"
	"trygin/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// SignUp 注册会员
func SignUp(C *gin.Context) {
	username := C.PostForm("username")
	password := C.PostForm("password")

	userModel := models.User{}
	userInfo := userModel.FindUser(username)
	if userInfo.UID != 0 {
		C.JSON(http.StatusOK, gin.H{
			"message": "user already exists",
		})
		return
	}
	userModel.AddUser(username, password)
	C.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    "sign up",
		"uid":     userInfo.UID,
	})
}

// SignIn 登录
func SignIn(C *gin.Context) {
	username := C.PostForm("username")
	password := C.PostForm("password")
	userModel := models.User{}
	userInfo := userModel.FindUser(username)
	if userInfo.UID == 0 {
		C.JSON(http.StatusOK, gin.H{
			"message": "user not exists",
		})
		return
	}
	if password != userInfo.Password {
		C.JSON(http.StatusOK, gin.H{
			"message": "password err",
		})
		return
	}
	// 记录session
	session, err := store.Get(C.Request, "SID")
	if err != nil {
		C.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("store Get SID failed: %v", err),
		})
		return
	}
	// session uid
	session.Values["uid"] = userInfo.UID

	var tokenStr string
	// JWT
	{
		mySigningKey := []byte("AllYourBase")
		TokenExpireDuration := time.Hour * time.Duration(1)
		// MyClaims .
		type MyClaims struct {
			jwt.StandardClaims
			UID uint `json:"uid"`
		}
		claims := MyClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
				Issuer:    "websocket",
			},
			userInfo.UID,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err = token.SignedString(mySigningKey)
		if err != nil {
			C.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("get token failed: %v", err),
			})
			return
		}
	}
	// session jwt
	session.Values["socket_jwt"] = tokenStr
	session.Save(C.Request, C.Writer)

	data := make(map[string]interface{})

	data["token"] = tokenStr
	C.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    data,
	})
}
