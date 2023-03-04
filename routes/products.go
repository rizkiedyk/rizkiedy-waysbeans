package routes

import (
	"waysbeans/handler"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.ConnDB)
	h := handler.HandlerProduct(productRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProducts)
	e.POST("/product", middleware.Auth(middleware.UploadFile(h.CreateProducts)))
	e.PATCH("/product/:id", middleware.Auth(middleware.UploadFile(h.UpdateProducts)))
	e.DELETE("/product/:id", h.DeleteProduct)
}
