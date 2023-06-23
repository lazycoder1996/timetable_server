package migrate

import (
	"fmt"
	"timetable_server/initializers"
	"timetable_server/models"
)

func SyncDB() {
	openFunctionQuery := "drop function if exists open; create function open(a integer, b text) returns setof schedules as $$ begin return query select * from schedules where start_time <= a and end_time > a and status=true and day = b; end; $$ language 'plpgsql';"
	initializers.DB.Exec(openFunctionQuery)

	emptyFunctionQuery := "drop function if exists empty; create function empty(a integer, b text) returns setof schedules as $$ begin return query select * from schedules where ((start_time <= a and a < end_time and status = false) or start_time > a) and lower(day) = b; end; $$ language 'plpgsql';"
	initializers.DB.Exec(emptyFunctionQuery)

	cols := "rooms.name, ok.programme, ok.year, ok.course_code, lower(ok.day), ok.start_time, ok.end_time, ok.recursive, ok.date, ok.status, ok.booked_by, ok.booking_id"
	availableFunctionQuery := fmt.Sprintf("drop function if exists available_now; create function available_now(a integer, b text) returns setof schedules as $$ begin return query select %s from (select * from empty(a , b) where room_name not in (select room_name from open(a, b))) as ok right join rooms on rooms.name = ok.room_name; end; $$ language 'plpgsql';", cols)
	initializers.DB.Exec(availableFunctionQuery)

	// initializers.DB.Debug().DropTableIfExists(models.Schedule{})
	// initializers.DB.Debug().DropTableIfExists(models.User{}, models.Booking{}, models.Schedule{}, models.Room{}, models.Course{}, models.Programme{})
	// initializers.DB.Debug().SetJoinTableHandler(&models.Schedule{}, "room_name", &gorm.JoinTableHandler{})
	initializers.DB.Debug().AutoMigrate(&models.User{}, &models.Programme{}, &models.Room{}, &models.Schedule{}, &models.Course{}, &models.Booking{})

	// ADDING USERS FORIEGN KEYS => PROGRAMME
	initializers.DB.Model(&models.User{}).AddForeignKey("programme", "programmes(programme)", "cascade", "cascade")

	// ADDING SCHEDULES FOREIGN KEYS => ROOM, COURSE, PROGRAMME, BOOKED BY,
	initializers.DB.Model(&models.Schedule{}).AddForeignKey("room_name", "rooms(name)", "cascade", "cascade")
	initializers.DB.Model(&models.Schedule{}).AddForeignKey("course_code", "courses(code)", "cascade", "cascade")
	initializers.DB.Model(&models.Schedule{}).AddForeignKey("programme", "programmes(programme)", "cascade", "cascade")


	initializers.DB.Model(&models.Booking{}).AddForeignKey("reference", "users(reference)", "cascade", "cascade")
	initializers.DB.Model(&models.Booking{}).AddForeignKey("room", "rooms(name)", "cascade", "cascade")
	initializers.DB.Model(&models.Booking{}).AddForeignKey("course_code", "courses(code)", "cascade", "cascade")
	

}
