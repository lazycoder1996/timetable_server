package controllers

import (
	"fmt"
	"net/http"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateCourse(c *gin.Context) {
	var body models.Courses
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
func UpdateCourse(c *gin.Context) {
	code := c.Param("code")
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
	var course models.Courses
	updateBody := &models.Courses{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&course, code)
	initializers.DB.Model(&course).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, course)

}
func DeleteCourse(c *gin.Context) {
	code := c.Param("code")

	initializers.DB.Delete(&models.Courses{}, code)
	c.Status(http.StatusOK)
}

func GetCourse(c *gin.Context) {
	code := c.Param("code")
	var course models.Courses
	initializers.DB.Where(&models.Courses{CourseID: code}).Find(&course)

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": course,
	})
}

func GetCourses(c *gin.Context) {
	var course []models.Courses
	initializers.DB.Find(&course)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": course,
	})
}
