package migrate

import (
	"timetable_server/initializers"
	"timetable_server/models"
)

func SyncDB() {
	// initializers.DB.Debug().DropTableIfExists( models.Schedule{})
	// initializers.DB.Debug().DropTableIfExists(models.User{}, models.Booking{}, models.Schedule{}, models.Room{}, models.Course{}, models.Programme{})
	// initializers.DB.Debug().SetJoinTableHandler(&models.Schedule{}, "room_name", &gorm.JoinTableHandler{})
	initializers.DB.Debug().AutoMigrate(&models.User{}, &models.Programme{}, &models.Room{}, &models.Schedule{}, &models.Course{})

	// ADDING USERS FORIEGN KEYS => PROGRAMME
	initializers.DB.Model(&models.User{}).AddForeignKey("programme", "programmes(programme)", "cascade", "cascade")

	// ADDING SCHEDULES FOREIGN KEYS => ROOM, COURSE, PROGRAMME, BOOKED BY,
	initializers.DB.Model(&models.Schedule{}).AddForeignKey("room_name", "rooms(name)", "cascade", "cascade")
	initializers.DB.Model(&models.Schedule{}).AddForeignKey("course_code", "courses(code)", "cascade", "cascade")
	initializers.DB.Model(&models.Schedule{}).AddForeignKey("programme", "programmes(programme)", "cascade", "cascade")

}
