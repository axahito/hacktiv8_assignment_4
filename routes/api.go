package routes

import (
	"hacktiv8_assignment_4/handlers"

	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	route := gin.Default()
	route.LoadHTMLFiles("index.html")

	route.GET("/", handlers.PageView)
	route.GET("/update-status", handlers.Update)

	return route
}
