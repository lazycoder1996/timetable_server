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
	rooms := api.Group("rooms")
	{
		rooms.POST("/", controllers.CreateRoom)
		rooms.GET("/", controllers.GetRooms)
		rooms.GET("/:name", controllers.GetRoom)
		rooms.PUT("/:name", controllers.UpdateRoom)
		rooms.DELETE("/:name", controllers.DeleteRoom)

	}

	courses := api.Group("courses")
	{
		courses.POST("/", controllers.CreateCourse)
		courses.GET("/", controllers.GetCourses)
		courses.GET("/:code", controllers.GetCourse)
		courses.PUT("/:code", controllers.UpdateCourse)
		courses.DELETE("/:code", controllers.DeleteCourse)

	}

	classes := api.Group("classes")
	{
		classes.POST("/", controllers.CreateClass)
		classes.GET("/", controllers.GetClasses)
		classes.GET("/:name", controllers.GetClass)
		classes.PUT("/:name", controllers.UpdateClass)
		classes.DELETE("/:name", controllers.DeleteClass)

	}

	bookings := api.Group("bookings")
	{
		bookings.POST("/", controllers.CreateBooking)
		bookings.GET("/", controllers.GetBookings)
		bookings.GET("/:id", controllers.GetBooking)
		bookings.PUT("/:id", controllers.UpdateBooking)
		bookings.DELETE("/:id", controllers.DeleteBooking)

	}

	users := api.Group("users")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)

	}

	schedules := api.Group("schedules")
	{
		schedules.POST("/", controllers.CreateSchedule)
		schedules.GET("/", controllers.GetSchedules)
		schedules.GET("/:id", controllers.GetSchedule)
		schedules.PUT("/:id", controllers.UpdateSchedule)
		schedules.DELETE("/:id", controllers.DeleteSchedule)

	}

	r.Run(":" + port)
}
