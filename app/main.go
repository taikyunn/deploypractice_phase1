package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLFiles("./templates/index.html")
		c.HTML(200, "index.html", gin.H{})
	})
	r.Run(":3000")
}
