package routes

import (
	"waysbeans/handler"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.ConnDB)
	h := handler.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login) // add this code
}
