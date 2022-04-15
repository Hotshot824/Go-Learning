package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

    return router
}

func main() {
    router := setupRouter()
    router.Run(":8080") // default localhost:8000
}