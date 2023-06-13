package controllers

import (
	"fmt"
	"net/http"
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
	res := initializers.DB.Preload("Course").Create(&body)
	fmt.Println(&body)
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
	id := c.Param("id")
	var schedule models.Schedule
	initializers.DB.Preload("Course").Find(&schedule, id)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": schedule,
	})
}

func GetSchedules(c *gin.Context) {
	var schedule []models.Schedule
	initializers.DB.Preload("Course").Find(&schedule)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": schedule,
	})
}
