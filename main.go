package main

/**
 gin run main.go -> live reload
 */
import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 會寫 gin.log 也會輸出
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// global use middleware
	r.Use(DummaryMiddleware())

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

	// url: /v1/welcome
	v1 := r.Group("/v1")
	{
		v1.GET("welcome", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "message",
			})
		})
	}

	// 當發生 panic 的時候避免程式 crash
	//r.Use(gin.Recovery())

	// 設定 login struct
	type Login struct {
		User string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	r.POST("/loginJson", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "test" && json.Password !=" test" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	})

	r.POST("/loginForm", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBind(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "test" && json.Password != "test" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	})

	r.Run()
}

// custom middleware
func DummaryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
