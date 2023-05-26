package main

import (
	"timetable_server/initializers"
	"timetable_server/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Rooms{}, &models.Bookings{}, &models.Classes{}, &models.Users{}, &models.Schedules{}, &models.Courses{})
}
