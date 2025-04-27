package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Render!")
	})

	fmt.Println("Server running on port", port)
	r.Run("0.0.0.0:" + port)
}
