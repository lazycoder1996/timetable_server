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
	var body models.Rooms
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
	var room models.Rooms
	updateBody := &models.Rooms{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&room, name)
	initializers.DB.Model(&room).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, room)

}
func DeleteRoom(c *gin.Context) {
	name := c.Param("name")

	initializers.DB.Delete(&models.Rooms{}, name)
	c.Status(http.StatusOK)
}

func GetRoom(c *gin.Context) {
	name := c.Param("name")
	var room models.Rooms
	initializers.DB.Where(&models.Rooms{RoomName: name}).Find(&room)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": room,
	})
}

func GetRooms(c *gin.Context) {
	var room []models.Rooms
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

func LiveRooms(c *gin.Context) {
	date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	day := c.Query("day")

	var rooms []models.Schedules
	initializers.DB.Where(&models.Schedules{Date: date}).Where("start_time <= ? and end_time >= ?", time, time).Where("day = ?", day).Find(&rooms)
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

func AvailableRooms(c *gin.Context) {
	date := c.Query("date")
	time, _ := strconv.Atoi(c.Query("time"))
	day := c.Query("day")
	var rooms []models.Schedules
	initializers.DB.Where(&models.Schedules{Date: date}).Where("start_time >= ? and end_time <= ?", time, time).Where("day = ?", day).Find(&rooms)
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
