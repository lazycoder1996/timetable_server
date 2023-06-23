package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateRoom(c *gin.Context) {
	var body models.Room
	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(&body)
	res := initializers.DB.Create(&body)
	if res.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"room": body,
	})

}
func UpdateRoom(c *gin.Context) {
	name := c.Param("name")
	var body models.Room
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// var room models.Room
	updateBody := &models.Room{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&body, name)
	initializers.DB.Model(&body).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, body)

}
func DeleteRoom(c *gin.Context) {
	name := c.Param("name")

	initializers.DB.Delete(&models.Room{}, name)
	c.Status(http.StatusOK)
}

func GetRoom(c *gin.Context) {
	name := c.Param("name")
	var room models.Room
	initializers.DB.Where(&models.Room{Name: name}).Find(&room)

	c.IndentedJSON(http.StatusOK, gin.H{
		"room": room,
	})
}

func GetRooms(c *gin.Context) {
	var room []models.Room
	initializers.DB.Where("lower(type) = 'class' and lower(location) = 'engineering'").Find(&room)
	c.IndentedJSON(http.StatusOK, gin.H{
		"rooms": room,
	})
}

// GET ROOMS LIVE ROOMS

// ROOMS IN USE AT A GIVEN TIME IN A GIVE DAY OR DATE
func LiveRooms(c *gin.Context) {
	// date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	day := c.Query("day")

	// query := fmt.Sprintf("drop function if exists open; create function open() returns setof schedules as $$ begin return query select * from schedules where start_time <= %d and end_time > %d and status=true and day = '%s'; end; $$ language 'plpgsql';",
	// 	time, time, strings.ToLower(day))
	// initializers.DB.Exec(query)

	var rooms []models.Schedule
	initializers.DB.Preload("Room").Preload("Course").Raw("select * from open(? ,?)", time, strings.ToLower(day)).Find(&rooms)
	liveRooms := make([]models.RoomStatusResponse, 0, 10)
	for i := range rooms {
		liveRoom := &models.RoomStatusResponse{}
		deepcopier.Copy(rooms[i]).To(liveRoom)
		liveRooms = append(liveRooms, *liveRoom)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"rooms": liveRooms,
	})
}

// EMPTY ROOMS AT A GIVEN TIME IN A GIVEN DAY
func AvailableRooms(c *gin.Context) {
	// date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	day := c.Query("day")
	// initializers.DB.Where(&models.Schedule{Date: date}).Where("? < start_time or ? > end_time", time, time).Where("and status = ?", false).Where("day = ? ", day).Find(&rooms)

	var rooms []models.Schedule

	// query := fmt.Sprintf("drop function if exists empty; create function empty() returns setof schedules as $$ begin return query select * from schedules where ((start_time <= %d and %d < end_time and status = false) or start_time > %d) and day = '%s'; end; $$ language 'plpgsql';",
	// 	time, time, time, strings.ToLower(day))
	// cols := "rooms.name, ok.programme, ok.year, ok.course_code, ok.day, ok.start_time, ok.end_time, ok.recursive, ok.date, ok.status, ok.booked_by, ok.booking_id"
	// initializers.DB.Exec(query)
	// query = fmt.Sprintf("drop function if exists available_now; create function available_now() returns setof schedules as $$ begin return query select %s from (select * from empty() where room_name not in (select room_name from open())) as ok right join rooms on rooms.name = ok.room_name; end; $$ language 'plpgsql';", cols)
	// initializers.DB.Exec(query)

	// initializers.DB.Where(&models.Schedule{Date: date}).Raw("select * from schedules where ((start_time <= ? and ? < end_time and status = false) or start_time > ?) and day = ?", time, time, time, strings.ToLower(day)).Scan(&rooms)

	query := fmt.Sprintf("select distinct on (room_name) room_name, programme, year, course_code, lower(day), start_time, end_time, recursive, date, status, booked_by, booking_id from available_now(%d, '%s') order by room_name, start_time", time, strings.ToLower(day))
	initializers.DB.Preload("Room").Preload("Course").Raw(query).Find(&rooms)
	vacantRooms := make([]models.RoomStatusResponse, 0, 10)
	for i := range rooms {
		vacantRoom := &models.RoomStatusResponse{}
		deepcopier.Copy(rooms[i]).To(vacantRoom)
		vacantRooms = append(vacantRooms, *vacantRoom)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"rooms": vacantRooms,
	})

}

// AVAILABLE TIMES WHICH A ROOM BECOMES EMTPY IN A GIVEN DAY/ DATE
func RoomAvailability(c *gin.Context) {
	room := c.Query("room")
	day := c.Query("day")
	time, _ := strconv.Atoi(c.Query("time"))

	var schedules []models.Schedule
	initializers.DB.Preload("Room").Preload("Course").Raw("select * from schedules where lower(day) = ? and start_time >= ? and room_name = ? and status = true", strings.ToLower(day), time, room).Find(&schedules)
	busyTimes := make([]models.RoomAvailableTimes, 0, 10)
	for i := range schedules {
		availableTime := &models.RoomAvailableTimes{}
		deepcopier.Copy(schedules[i]).To(availableTime)
		busyTimes = append(busyTimes, *availableTime)
	}

	// TODO: PERFORM OPERATION HERE
	c.IndentedJSON(http.StatusOK, gin.H{
		"usage_times": busyTimes,
	})

}
