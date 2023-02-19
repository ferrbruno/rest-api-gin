package main

import (
	"github.com/ferrbruno/rest-api-gin/database"
	"github.com/ferrbruno/rest-api-gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")
	router.LoadRoutes(r)

	r.Run()
}
