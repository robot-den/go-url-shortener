package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		// gin.H is a shortcut for map[string]interface{}
		context.JSON(http.StatusOK, gin.H{
			"message": "Use shortener!",
		})
	})
	err := router.Run(":9909")
	if err != nil {
		panic(fmt.Sprintf("failed to start web server: %s", err.Error()))
	}
}
