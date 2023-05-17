package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.Run("localhost:8080")
}
