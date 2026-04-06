package main

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents) //Get,Post,Put,Delete,Patch,Head,Options,Any
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse Data.",
			"error":   err.Error(),
		})
		// Handle error
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully.",
		"event":   event,
	})

}
