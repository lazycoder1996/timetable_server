package controllers

import (
	"fmt"
	"net/http"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateClass(c *gin.Context) {
	var body models.Classes
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
func UpdateClass(c *gin.Context) {
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
	var class models.Classes
	updateBody := &models.Classes{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&class, name)
	initializers.DB.Model(&class).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, class)

}
func DeleteClass(c *gin.Context) {
	name := c.Param("name")

	initializers.DB.Delete(&models.Classes{}, name)
	c.Status(http.StatusOK)
}

func GetClass(c *gin.Context) {
	name := c.Param("name")
	var class models.Classes
	initializers.DB.Where(&models.Classes{ClassName: name}).Find(&class)

	c.IndentedJSON(http.StatusOK, gin.H{
		"class": class,
	})
}

func GetClasses(c *gin.Context) {
	var class []models.Classes
	initializers.DB.Find(&class)
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": class,
	})
}
