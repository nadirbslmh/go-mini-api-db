package routes

import (
	"go-mini-api-db/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/posts", controllers.GetAllPosts)
	e.POST("/posts", controllers.CreatePost)
}
