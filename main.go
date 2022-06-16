package main

import (
	"github.com/gin-gonic/gin"
	"github.com/korawichtawan/eightPuzzleWebApp/controllers"
)

func main() {
	r := gin.Default();

	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", controllers.ViewHandler)
	r.POST("/", controllers.SolverHandler)

	r.Run();

}