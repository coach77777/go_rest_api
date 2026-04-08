package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/rest-api/models"
)

func signup(context *gin.Context) {
	var user models.User

err := context.ShouldBindJSON(&user)

if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind user data"})
	return
}
err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
