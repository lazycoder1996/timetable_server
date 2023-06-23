package main

import (
	"os"
	"strings"
	"time"
	"timetable_server/controllers"
	"timetable_server/initializers"
	"timetable_server/migrate"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func clearData() {
	initializers.DB.Exec("delete from schedules where recursive = false")
	initializers.DB.Exec("delete from bookings")
	initializers.DB.Exec("update schedules set status = true")
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.SyncDB()
}

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	c := cron.New()
	c.AddFunc("00 21 * * 5", clearData)
	c.AddFunc("00 18 * * 1-5", func() {
		initializers.DB.Exec("delete from bookings where lower(day) = ?", strings.ToLower(time.Now().Weekday().String()))
	})
	c.Start()
	// r.LoadHTMLGlob("templates/*.tmpl.html")
	// r.Static("/static", "static")

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })
	api := r.Group("api")
	auth := api.Group("auth")
	{
		auth.POST("/login", controllers.LoginUser)
		auth.POST("/register", controllers.CreateUser)
	}

	rooms := api.Group("rooms")
	{
		rooms.POST("/", controllers.CreateRoom)
		rooms.GET("/", controllers.GetRooms)
		rooms.GET("/:name", controllers.GetRoom)
		rooms.PUT("/:name", controllers.UpdateRoom)
		rooms.DELETE("/:name", controllers.DeleteRoom)
		rooms.GET("/live", controllers.LiveRooms)
		rooms.GET("/vacant", controllers.AvailableRooms)
		rooms.GET("/available_times", controllers.RoomAvailability)
	}

	courses := api.Group("courses")
	{
		courses.POST("/", controllers.CreateCourse)
		courses.GET("/", controllers.GetCourses)
		courses.GET("/:code", controllers.GetCourse)
		courses.PUT("/:code", controllers.UpdateCourse)
		courses.DELETE("/:code", controllers.DeleteCourse)
	}

	classes := api.Group("programmes")
	{
		classes.POST("/", controllers.CreateProgramme)
		classes.GET("/", controllers.GetProgrammes)
		classes.GET("/:name", controllers.GetProgramme)
		classes.PUT("/:name", controllers.UpdateProgramme)
		classes.DELETE("/:name", controllers.DeleteProgramme)
	}

	bookings := api.Group("bookings")
	{
		bookings.POST("/", controllers.CreateBooking)
		bookings.GET("/", controllers.GetBookings)
		bookings.PUT("/:id", controllers.UpdateBooking)
		bookings.DELETE("/:id", controllers.DeleteBooking)
	}

	users := api.Group("users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	timetable := api.Group("timetable")
	{
		timetable.GET("", controllers.GetTimeTable)
	}

	schedules := api.Group("schedules")
	{
		schedules.POST("/", controllers.CreateSchedule)
		schedules.GET("/", controllers.GetSchedules)
		schedules.GET("/:day", controllers.GetSchedule)
		schedules.PUT("/:id", controllers.UpdateSchedule)
		schedules.DELETE("/:id", controllers.DeleteSchedule)

	}

	r.Run(":" + port)
}
