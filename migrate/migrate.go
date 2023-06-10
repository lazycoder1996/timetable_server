package migrate

import (
	"timetable_server/initializers"
	"timetable_server/models"
)

func SyncDB() {
	initializers.DB.AutoMigrate(&models.Rooms{}, &models.Bookings{}, &models.Programme{}, &models.User{}, &models.Schedules{}, &models.Courses{})
}
