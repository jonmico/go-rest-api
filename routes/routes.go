package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Event routes and handlers
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// User routes and handlers
	server.POST("/signup", signup)
	server.POST("/login", login)
}
