package main

import "time"

// MyClaims .
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// TokenExpireDuration .
const TokenExpireDuration = time.Hour * 2

// GenToken .
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		"username",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMetiodHS256, c)
	return token.SignedString(MySecret)
}

func main() {

}
