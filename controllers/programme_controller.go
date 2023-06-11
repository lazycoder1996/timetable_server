package controllers

import (
	"fmt"
	"net/http"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateProgramme(c *gin.Context) {
	var body models.Programme
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
		"programme": body,
	})

}
func UpdateProgramme(c *gin.Context) {
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
	var programme models.Programme
	updateBody := &models.Programme{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&programme, name)
	initializers.DB.Model(&programme).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, programme)

}
func DeleteProgramme(c *gin.Context) {
	name := c.Param("name")

	initializers.DB.Delete(&models.Programme{}, name)
	c.Status(http.StatusOK)
}

func GetProgramme(c *gin.Context) {
	name := c.Param("name")
	var programme models.Programme
	initializers.DB.Where(&models.Programme{}, name).Find(&programme)

	c.IndentedJSON(http.StatusOK, gin.H{
		"programme": programme,
	})
}

func GetProgrammes(c *gin.Context) {
	var programmes []models.Programme
	initializers.DB.Find(&programmes)
	c.IndentedJSON(http.StatusOK, gin.H{
		"programmes": programmes,
	})
}
