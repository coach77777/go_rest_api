package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) //Get,Post,Put,Delete,Patch,Head,Options,Any

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
