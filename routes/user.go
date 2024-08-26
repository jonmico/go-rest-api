package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
}
