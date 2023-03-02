package routes

import (
	"waysbeans/handler"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.ConnDB)
	h := handler.HandlerProduct(productRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/products/:id", h.GetProducts)
}
