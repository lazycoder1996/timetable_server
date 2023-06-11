package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"gorm.io/gorm"
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
		"data": body,
	})

}
func UpdateRoom(c *gin.Context) {
	name := c.Param("name")
	var body struct {
		Name string
		Size int
	}
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var room models.Room
	updateBody := &models.Room{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&room, name)
	initializers.DB.Model(&room).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, room)

}
func DeleteRoom(c *gin.Context) {
	name := c.Param("name")

	initializers.DB.Delete(&models.Room{}, name)
	c.Status(http.StatusOK)
}

func GetRoom(c *gin.Context) {
	name := c.Param("name")
	var room models.Room
	initializers.DB.Where(&models.Room{RoomName: name}).Find(&room)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": room,
	})
}

func GetRooms(c *gin.Context) {
	var room []models.Room
	initializers.DB.Find(&room)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": room,
	})
}

// GET ROOMS LIVE ROOMS
type RoomStatusResponse struct {
	gorm.Model
	Room      string
	Programme string
	Year      int
	Course    string
	StartTime int
	EndTime   int
}

// ROOMS IN USE AT A GIVEN TIME IN A GIVE DAY OR DATE
func LiveRooms(c *gin.Context) {
	date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	day := c.Query("day")

	var rooms []models.Schedule
	initializers.DB.Where(&models.Schedule{Date: date}).Where("start_time <= ? and end_time >= ?", time, time).Where("day = ?", day).Find(&rooms)
	liveRooms := make([]RoomStatusResponse, 0, 10)
	for i := range rooms {
		liveRoom := &RoomStatusResponse{}
		deepcopier.Copy(rooms[i]).To(liveRoom)
		liveRooms = append(liveRooms, *liveRoom)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"rooms": liveRooms,
	})
}

// EMPTY ROOMS AT A GIVEN TIME IN A GIVEN DAY
func AvailableRooms(c *gin.Context) {
	date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	day := c.Query("day")
	var rooms []models.Schedule
	initializers.DB.Where(&models.Schedule{Date: date}).Where("start_time >= ? and end_time <= ?", time, time).Where("day = ?", day).Find(&rooms)
	vacantRooms := make([]RoomStatusResponse, 0, 10)
	for i := range rooms {
		vacantRoom := &RoomStatusResponse{}
		deepcopier.Copy(rooms[i]).To(vacantRoom)
		vacantRooms = append(vacantRooms, *vacantRoom)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"rooms": vacantRooms,
	})

}

// AVAILABLE TIMES WHICH A ROOM BECOMES EMTPY IN A GIVEN DAY/ DATE
func RoomAvailability(c *gin.Context) {
	room := c.Param("room")
	var schedules []models.Schedule
	initializers.DB.Where(&models.Schedule{Room: room}).Find(&schedules)
	// TODO: PERFORM OPERATION HERE
	c.IndentedJSON(http.StatusOK, gin.H{
		"available_times": schedules,
	})

}
