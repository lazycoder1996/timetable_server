package migrate

import (
	"timetable_server/initializers"
	"timetable_server/models"
)

func SyncDB() {
	initializers.DB.AutoMigrate(&models.Room{}, &models.Booking{}, &models.Programme{}, &models.User{}, &models.Schedule{}, &models.Course{})
}
