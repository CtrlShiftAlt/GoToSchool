package controllers

import "github.com/gin-gonic/gin"

// Ping .
func Ping(C *gin.Context) {
	C.JSON(200, gin.H{
		"message":  "pong",
	})
}

