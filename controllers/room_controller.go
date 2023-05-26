package controllers

import (
	"fmt"
	"net/http"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
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
		"room": body,
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
		"room": room,
	})
}

func GetRooms(c *gin.Context) {
	var room []models.Rooms
	initializers.DB.Find(&room)
	c.IndentedJSON(http.StatusOK, gin.H{
		"room": room,
	})
}
