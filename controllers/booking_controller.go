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
	var body models.Bookings
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
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": body,
	})

}
func UpdateBooking(c *gin.Context) {
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
	var booking models.Bookings
	updateBody := &models.Bookings{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&booking, id)
	initializers.DB.Model(&booking).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, booking)

}
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Bookings{}, id)
	c.Status(http.StatusOK)
}

func GetBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var booking models.Bookings
	initializers.DB.Where(&models.Bookings{ID: id}).Find(&booking)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": booking,
	})
}

func GetBookings(c *gin.Context) {
	var booking []models.Bookings
	initializers.DB.Find(&booking)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": booking,
	})
}
