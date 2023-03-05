package handler

import (
	"net/http"
	"waysbeans/dto/result"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

// SETUP HANDLER STRUCT
type cartHandler struct {
	CartRepository repositories.CartRepository
}

// SETUP HANDLER FUNCTION
func HandlerCart(CartRepository repositories.CartRepository) *cartHandler {
	return &cartHandler{CartRepository}
}

// FUNCTION FIND CARTS
func (h *cartHandler) FindCarts(c echo.Context) error {
	cart, err := h.CartRepository.FindCarts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: cart})
}
