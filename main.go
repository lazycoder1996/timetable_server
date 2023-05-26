package main

import (
	"os"
	"timetable_server/controllers"
	"timetable_server/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// r.LoadHTMLGlob("templates/*.tmpl.html")
	// r.Static("/static", "static")

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })
	api := r.Group("api")
	api.POST("/room", controllers.CreateRoom)
	api.GET("/room", controllers.GetRooms)
	api.GET("/room/:name", controllers.GetRoom)
	api.PUT("/room/:name", controllers.UpdateRoom)
	api.DELETE("/room/:name", controllers.DeleteRoom)
	r.Run(":" + port)
}
