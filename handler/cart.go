package handler

import (
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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

func (h *cartHandler) GetCart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: cart})
}

func (h *cartHandler) CreateCart(c echo.Context) error {
	request := new(dto.CartRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	productId, _ := strconv.Atoi(c.Param("product_id"))

	cart := models.Cart{
		ProductID:     productId,
		OrderQuantity: request.OrderQuantity,
		UserID:        int(userId),
	}

	data, err := h.CartRepository.CreateCart(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convertResponseCart(data)})
}

func (h *cartHandler) UpdateCart(c echo.Context) error {
	request := new(dto.CartRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.ProductID != 0 {
		cart.ProductID = request.ProductID
	}

	if request.OrderQuantity != 0 {
		cart.OrderQuantity = request.OrderQuantity
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	if request.UserID != 0 {
		cart.UserID = int(userId)
	}

	data, err := h.CartRepository.UpdateCart(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convertResponseCart(data)})
}

func (h *cartHandler) DeleteCart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CartRepository.DeleteCart(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convertResponseCart(data)})
}

func convertResponseCart(u models.Cart) dto.CartResponse {
	return dto.CartResponse{
		ProductID:     u.ProductID,
		OrderQuantity: u.OrderQuantity,
		UserID:        u.UserID,
	}
}
