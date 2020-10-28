package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims .
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func getToken() (string, error) {
	mySigningKey := []byte("AllYourBase")
	const TokenExpireDuration = time.Hour * 2
	claims := MyClaims{
		"username22",
		jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// Issuer:    "my-project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func parseToken(tokenString string) (*MyClaims, error) {
	mySigningKey := []byte("AllYourBase")
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Printf("ParseWithClaims failed, err: %v\n", err)
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func main() {
	// 加密
	tokenStr, err := getToken()
	if err != nil {
		fmt.Printf("get Token failed, err: %v\n", err)
	}
	fmt.Printf("%v\n", tokenStr)

	// 解密
	// tokenStr := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1ODgwNTc3NjQsImV4cCI6MTU4ODA2NDk2NCwidXNlcm5hbWUiOiJnb29nIn0.kxxeyk42GGo8ibPK_ZkxXGqlqYIfufPWX8kGxamoDPI"
	claims, err := parseToken(tokenStr)
	if err != nil {
		fmt.Printf("parse Token failed, err: %v\n", err)
	}
	fmt.Printf("%#v\n", claims)
}
