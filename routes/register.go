package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)



func registerForEvent(context *gin.Context) {

userID := context.GetInt64("userID")
eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID.",
			"error":   err.Error(),
		})
		return
	}
event, err := models.GetEventByID(eventId)

if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{
		"message": "Failed to retrieve event.",
		"error":   err.Error(),
	})
	return
}
err = event.Register(userID)

if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{
		"message": "Failed to register for event.",
		"error":   err.Error(),
	})
	return
}

context.JSON(http.StatusOK, gin.H{
	"message": "Successfully registered for event.",
})


}





func cancelRegistration(context *gin.Context){}