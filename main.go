package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", handler)

	router.Run()
}

func handler(ctx *gin.Context) {

	user := User{"User", 20}

	ctx.HTML(200, "index.html", gin.H{
		"user": user,
	})
}
