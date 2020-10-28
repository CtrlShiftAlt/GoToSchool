package admin

import (
	"net/http"
	"trygin/models"

	"github.com/gin-gonic/gin"
)

// Index .
func Index(C *gin.Context) {
	book := models.Book{}
	book.AddBook()
	C.JSON(200, gin.H{
		"status": http.StatusOK,
	})
}
