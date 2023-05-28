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

func CreateUser(c *gin.Context) {
	var body models.Users
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
func UpdateUser(c *gin.Context) {
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
	var user models.Users
	updateBody := &models.Users{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&user, id)
	initializers.DB.Model(&user).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, user)

}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Users{}, id)
	c.Status(http.StatusOK)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.Users
	initializers.DB.Where(&models.Users{UserID: id}).Find(&user)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func GetUsers(c *gin.Context) {
	var user []models.Users
	initializers.DB.Find(&user)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": user,
	})
}
