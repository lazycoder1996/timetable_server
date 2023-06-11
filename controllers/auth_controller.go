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
	Reference      string
	FirstName      string
	MiddleName     string
	Surname        string
	Programme      string
	Year           int
	ProfilePicture string
	Notification   int
	Role           int
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
	if user.ID == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Reference ID",
		})
		return
	}
	response := &UserResponseBody{}
	deepcopier.Copy(body).To(response)

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
