package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

type TimetableModel struct {
	CourseCode string        `json:"course_code"`
	Course     models.Course `json:"course_details"`
	StartTime  int           `json:"start_time"`
	EndTime    int           `json:"end_time"`
	Day        string        `json:"day"`
	Room       models.Room   `json:"room"`
	Recursive  bool          `json:"recursive"`
}

// GET TIMETABLE
func GetTimeTable(c *gin.Context) {
	programme := c.Query("programme")
	year := c.Query("year")
	day := c.Query("day")
	var schedules []models.Schedule
	body := &models.Schedule{}
	if programme != "" {
		body.Programme = programme
	}
	if year != "" {
		body.Year, _ = strconv.Atoi(year)
	}
	if day != "" {
		body.Day = strings.ToLower(day)
	}
	initializers.DB.Preload("Room").Preload("Course").Where(&body).Find(&schedules)
	timetables := make([]TimetableModel, 0, 10)
	for i := range schedules {
		timetable := &TimetableModel{}
		deepcopier.Copy(schedules[i]).To(timetable)
		timetables = append(timetables, *timetable)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"timetable": timetables,
	})
}
