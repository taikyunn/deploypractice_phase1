package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLFiles("./templates/index.html")
		c.HTML(200, "index.html", gin.H{})
	})
	r.Run(":" + port)
	// r.Run(":3000")
}
