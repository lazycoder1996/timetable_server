package controllers

import (
	"net/http"
	"strconv"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
)

// GET TIMETABLE
func GetTimeTable(c *gin.Context) {
	programme := c.Query("programme")
	year := c.Query("year")
	recursive := c.Query("recursive")
	var schedules []models.Schedule
	body := &models.Schedule{}
	if programme != "" {
		body.Programme = programme
	}
	if year != "" {
		body.Year, _ = strconv.Atoi(year)
	}
	if recursive != "" {
		body.Recursive, _ = strconv.Atoi(recursive)
	}
	initializers.DB.Where(body).Find(&schedules)
	c.IndentedJSON(http.StatusOK, gin.H{
		"timetable": schedules,
	})
}
