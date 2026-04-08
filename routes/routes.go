package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)  //Get, post, put, delete events
	server.GET("/events/:id", getEvent) //examples of getting a single event by ID
	server.POST("/events", createEvent) //example of creating a new event
	server.PUT("/events/:id", updateEvent) //example of updating an existing event
	server.DELETE("/events/:id", deleteEvent) //example of deleting an existing event
	//************************************
	server.POST("/signup", signup)
	server.POST("/login", login)
}
