package controllers

import (
	"net/http"
	"strconv"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

// UPDATE USER DETAILS
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Notification int
		Role         int
	}
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	updateBody := &models.User{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&user, id)
	initializers.DB.Model(&user).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, user)

}

// DELETE USER
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.User{}, id)
	c.Status(http.StatusOK)
}

// GET USER DETAILS
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var body models.User
	initializers.DB.Where(&models.User{Reference: id}).Find(&body)
	user := &UserResponseBody{}
	deepcopier.Copy(body).To(user)
	c.IndentedJSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// GET ALL USERS
func GetUsers(c *gin.Context) {
	var user []models.User
	initializers.DB.Find(&user)
	var users []UserResponseBody
	for i := range user {
		_user := &UserResponseBody{}
		deepcopier.Copy(user[i]).To(_user)
		users = append(users, *_user)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"users": users,
	})
}
