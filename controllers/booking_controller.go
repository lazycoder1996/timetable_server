package controllers

import (
	"net/http"
	"strconv"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateBooking(c *gin.Context) {
	var body models.Booking
	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res := initializers.DB.Create(&body)
	if res.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
		return
	}
	// var user models.User
	// initializers.DB.Where("reference = ?", body.Reference).Find(&user)
	// fmt.Println(user.Programme)
	// schedule := models.Schedule{
	// 	RoomName:   body.Room,
	// 	Programme:  user.Programme,
	// 	Year:       user.Year,
	// 	CourseCode: body.CourseCode,
	// 	Day:        body.Day,
	// 	StartTime:  body.StartTime,
	// 	EndTime:    body.EndTime,
	// 	Recursive:  false,
	// 	Status:     true,
	// 	BookedBy:   body.Reference,
	// 	BookingID:  int(body.ID),
	// }
	// res = initializers.DB.Create(&schedule)
	// if res.Error != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{
	// 		"error": res.Error.Error(),
	// 	})
	// 	return
	// }
	c.IndentedJSON(http.StatusOK, gin.H{
		"booking": body,
	})
}

func UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Course string
	}
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var booking models.Booking
	updateBody := &models.Booking{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&booking, id)
	initializers.DB.Model(&booking).UpdateColumns(&updateBody)

	var schedule models.Schedule
	scheduleUpdateBody := &models.Schedule{}
	deepcopier.Copy(body).To(scheduleUpdateBody)

	initializers.DB.First(&schedule, id)
	initializers.DB.Model(&schedule).UpdateColumns(scheduleUpdateBody)

	c.IndentedJSON(http.StatusOK, gin.H{
		"booking": booking,
	})

}
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Booking{}, id)
	initializers.DB.Where("booking_id=?", id).Delete(&models.Schedule{})
	c.Status(http.StatusOK)
}

func GetBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var booking models.Booking
	initializers.DB.Where(&models.Booking{}, id).Find(&booking)
	c.IndentedJSON(http.StatusOK, gin.H{
		"booking": booking,
	})
}

func GetBookings(c *gin.Context) {
	programme := c.Query("programme")
	year, _ := strconv.Atoi(c.Query("year"))
	var booking []models.Booking
	initializers.DB.Preload("Course").Model(&models.Schedule{
		Programme: programme, Year: year, Recursive: false,
	}).Find(&booking)
	c.IndentedJSON(http.StatusOK, gin.H{
		"bookings": booking,
	})
}
