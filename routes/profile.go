package routes

import (
	"waysbeans/handler"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	ProfileRepository := repositories.RepositoryProfile(mysql.ConnDB)
	h := handler.HandlerProfile(ProfileRepository)

	e.GET("/profiles", h.FindProfile)
	e.GET("/profile/:id", h.GetProfile)
	e.POST("/profile", middleware.Auth(h.CreateProfile))
}
