package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateSchedule(c *gin.Context) {
	var body models.Schedules
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
	var schedule models.Schedules
	updateBody := &models.Schedules{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&schedule, id)
	initializers.DB.Model(&schedule).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, schedule)

}
func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Schedules{}, id)
	c.Status(http.StatusOK)
}

func GetSchedule(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var schedule models.Schedules
	initializers.DB.Where(&models.Schedules{ScheduleID: id}).Find(&schedule)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": schedule,
	})
}

func GetSchedules(c *gin.Context) {
	var schedule []models.Schedules
	initializers.DB.Find(&schedule)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": schedule,
	})
}
