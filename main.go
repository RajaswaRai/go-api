package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GO API is Online!",
		})
	})

	router.GET("/about", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This GO API made by RajaswaRai",
		})
	})

	return router
}

func main() {
	r := setupRouter()
	r.Run(":8000")

	fmt.Println("API Running on port:8000")
}
