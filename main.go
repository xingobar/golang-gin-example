package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// get parameter
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(200, "Hello %s ", name)
	})

	// query string
	r.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Gary")
		lastname := context.Query("lastname")
		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	
	// post
	// url: http://localhost:3000/form_post
	r.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymouns")
		context.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	
	r.Run()
}
