package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateSchedule(c *gin.Context) {
	var body models.Schedule
	if err := c.ShouldBindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// CHECK IF PROGRAMME HAS A CLASS ONGOING
	var count int
	classes := initializers.DB.Model(&models.Schedule{}).Preload("Course").Where("start_time <= ? and ? < end_time and lower(day) = ? and programme = ? and year = ? and status=true", body.StartTime, body.StartTime, strings.ToLower(body.Day), body.Programme, body.Year)
	classes.Count(&count)
	if count != 0 {
		var class models.Schedule
		classes.Find(&class)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("You already have %s class during the given time in %s", class.Course.Name, class.RoomName),
		})
		return
	}
	classes = initializers.DB.Model(&models.Schedule{}).Where("start_time <= ? and ? < end_time and lower(day) = ? and room_name = ? and status=true", body.StartTime, body.StartTime, strings.ToLower(body.Day), body.RoomName)
	classes.Count(&count)
	if count != 0 {
		var class models.Schedule
		classes.Find(&class)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%s %d already have a class during the given time", class.Programme, class.Year),
		})
		return
	}
	body.Day = strings.ToLower(body.Day)
	// INSERTING INTO BOOKINGS
	if body.BookedBy != 0 {
		var booking models.Booking
		deepcopier.Copy(&body).To(&booking)
		booking.Reference = body.BookedBy
		booking.Room = body.RoomName
		fmt.Println(booking)
		res := initializers.DB.Create(&booking)
		if res.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": res.Error.Error(),
			})
			return
		}
		body.BookingID = int(booking.ID)
		body.Recursive = false
	}
	res := initializers.DB.Create(&body)
	if res.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
		return
	}
	initializers.DB.Preload("Course").Preload("Room").Find(&body)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": body,
	})

}
func UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
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
	var schedule models.Schedule
	updateBody := &models.Schedule{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&schedule, id)
	initializers.DB.Model(&schedule).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, schedule)

}
func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Schedule{}, id)
	c.Status(http.StatusOK)
}

func GetSchedule(c *gin.Context) {
	day := c.Param("day")
	var schedule []models.Schedule
	initializers.DB.Preload("Course").Preload("Room").Where("lower(day) = ? and status= true", strings.ToLower(day)).Find(&schedule)

	c.IndentedJSON(http.StatusOK, gin.H{
		"schedules": schedule,
	})
}

func GetSchedules(c *gin.Context) {
	var schedule []models.Schedule
	initializers.DB.Preload("Course").Preload("Room").Find(&schedule)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": schedule,
	})
}
