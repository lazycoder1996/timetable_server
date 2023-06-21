package controllers

import (
	"net/http"
	"os"
	"time"
	"timetable_server/initializers"
	"timetable_server/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ulule/deepcopier"
	"golang.org/x/crypto/bcrypt"
)

type UserResponseBody struct {
	Reference      int    `json:"reference"`
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	Surname        string `json:"surname"`
	Programme      string `json:"programme_name"`
	Year           int    `json:"year"`
	ProfilePicture string `json:"profile_picture"`
	Notification   int    `json:"notification"`
	Role           int    `json:"role"`
}

// CREATE A USER IN DATABASE
func CreateUser(c *gin.Context) {
	var body models.User
	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to has password",
		})
	}
	body.Password = string(hash)
	res := initializers.DB.Create(&body)
	if res.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
		return
	}
	response := &UserResponseBody{}
	deepcopier.Copy(body).To(response)
	c.IndentedJSON(http.StatusOK, gin.H{
		"user": response,
	})

}

func LoginUser(c *gin.Context) {
	var body models.LoginBody
	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "reference = ?", body.Reference)
	if user.Reference == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Reference ID",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	// Generate a jwt token TODO: learn about paseto
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Reference,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	response := &UserResponseBody{}
	deepcopier.Copy(user).To(response)

	// Sign and get the complete encode token as a string using the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  response,
	})
}
